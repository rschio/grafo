// Modified from github.com/yourbasic/graph.

//BSD 2-Clause License
//
//Copyright (c) 2017, Stefan Nilsson
//All rights reserved.
//
//Redistribution and use in source and binary forms, with or without
//modification, are permitted provided that the following conditions are met:
//
//* Redistributions of source code must retain the above copyright notice, this
//  list of conditions and the following disclaimer.
//
//* Redistributions in binary form must reproduce the above copyright notice,
//  this list of conditions and the following disclaimer in the documentation
//  and/or other materials provided with the distribution.
//
//THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
//AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
//IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
//DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
//FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
//DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
//SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
//CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
//OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
//OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package grafo

import (
	"cmp"
	"iter"
	"sort"
	"strconv"
)

// Immutable is a compact representation of an immutable graph.
// The implementation uses lists to associate each vertex in the graph
// with its adjacent vertices. This makes for fast and predictable
// iteration: the Visit method produces its elements by reading
// from a fixed sorted precomputed list. This type supports multigraphs.
type Immutable[T any] struct {
	// edges[v] is a sorted list of v's neighbors.
	edges [][]neighbor[T]
	stats Stats
}

type neighbor[T any] struct {
	vertex int
	weight T
}

// Stats holds basic data about an Iterator.
type Stats struct {
	Size  int // Number of unique edges.
	Multi int // Number of duplicate edges.
	//Weighted int // Number of edges with non-zero cost.
	Loops    int // Number of self-loops.
	Isolated int // Number of vertices with outdegree zero.
}

// Sort returns an immutable copy of g with a Visit method
// that returns its neighbors in increasing numerical order.
func Sort[T any](g Graph[T]) *Immutable[T] {
	if g, ok := g.(*Immutable[T]); ok {
		return g
	}
	return build(g, false)
}

// Transpose returns the transpose graph of g.
// The transpose graph has the same set of vertices as g,
// but all of the edges are reversed compared to the orientation
// of the corresponding edges in g.
func Transpose[T any](g Graph[T]) *Immutable[T] {
	return build(g, true)
}

func build[T any](g Graph[T], transpose bool) *Immutable[T] {
	n := g.Order()
	h := &Immutable[T]{edges: make([][]neighbor[T], n)}
	for v := range h.edges {
		for w, wt := range g.EdgesFrom(v) {
			if w < 0 || w >= n {
				panic("vertex out of range: " + strconv.Itoa(w))
			}
			if transpose {
				h.edges[w] = append(h.edges[w], neighbor[T]{v, wt})
			} else {
				h.edges[v] = append(h.edges[v], neighbor[T]{w, wt})
			}
		}
		sort.Slice(h.edges[v], func(i, j int) bool {
			e := h.edges[v]
			// Parallel edges don't have a defined order based on its weight.
			return cmp.Less(e[i].vertex, e[j].vertex)
		})
	}

	for v, neighbors := range h.edges {
		if len(neighbors) == 0 {
			h.stats.Isolated++
		}
		prev := -1
		for _, e := range neighbors {
			w, _ := e.vertex, e.weight
			if v == w {
				h.stats.Loops++
			}
			//	if wt != zero {
			//		h.stats.Weighted++
			//	}
			if w == prev {
				h.stats.Multi++
			} else {
				h.stats.Size++
				prev = w
			}
		}
	}
	return h
}

func (g *Immutable[T]) EdgesFrom(v int) iter.Seq2[int, T] {
	return func(yield func(w int, weight T) bool) {
		for _, e := range g.edges[v] {
			if !yield(e.vertex, e.weight) {
				return
			}
		}
	}
}

//// VisitFrom calls the do function starting from the first neighbor w
//// for which w ≥ a, with c equal to the cost of the edge from v to w.
//// The neighbors are then visited in increasing numerical order.
//// If do returns true, VisitFrom returns immediately,
//// skipping any remaining neighbors, and returns true.
//func (g *Immutable) VisitFrom(v int, a int, do func(w int, c int64) bool) bool {
//	neighbors := g.edges[v]
//	n := len(neighbors)
//	i := sort.Search(n, func(i int) bool { return a <= neighbors[i].vertex })
//	for ; i < n; i++ {
//		e := neighbors[i]
//		if do(e.vertex, e.weight) {
//			return true
//		}
//	}
//	return false
//}

//// String returns a string representation of the graph.
//func (g *Immutable) String() string {
//	return String(g)
//}

// Order returns the number of vertices in the graph.
func (g *Immutable[T]) Order() int {
	return len(g.edges)
}

// Edge tells if there is an edge from v to w.
func (g *Immutable[T]) Edge(v, w int) bool {
	if v < 0 || v >= len(g.edges) {
		return false
	}
	edges := g.edges[v]
	n := len(edges)
	i := sort.Search(n, func(i int) bool { return w <= edges[i].vertex })
	return i < n && w == edges[i].vertex
}

// Degree returns the number of outward directed edges from v.
func (g *Immutable[T]) Degree(v int) int {
	return len(g.edges[v])
}
