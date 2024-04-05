// Modified from github.com/yourbasic/graph.

// BSD 2-Clause License
//
// Copyright (c) 2017, Stefan Nilsson
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package grafo

// TopSort returns a topological ordering of the vertices in
// a directed acyclic graph; if the graph is not acyclic,
// no such ordering exists and ok is set to false.
//
// In a topological order v comes before w for every directed edge from v to w.
func TopSort[T any](g Graph[T]) (order []int, ok bool) {
	return topsort(g, true)
}

// Acyclic tells if g has no cycles.
func Acyclic[T any](g Graph[T]) bool {
	_, acyclic := topsort(g, false)
	return acyclic
}

// Kahn's algorithm.
func topsort[T any](g Graph[T], output bool) (order []int, acyclic bool) {
	n := g.Order()
	indegree := make([]int, n)
	for v := range indegree {
		for w, _ := range g.EdgesFrom(v) {
			indegree[w]++
		}
	}

	// Invariant: this queue holds all vertices with indegree 0.
	queue := newBfsQueue(10)
	for v, degree := range indegree {
		if degree == 0 {
			queue.Insert(v)
		}
	}

	order = []int{}
	vertexCount := 0
	for queue.Len() > 0 {
		v := queue.Remove()
		if output {
			order = append(order, v)
		}
		vertexCount++
		for w, _ := range g.EdgesFrom(v) {
			indegree[w]--
			if indegree[w] == 0 {
				queue.Insert(w)
			}
		}
	}

	if vertexCount == n {
		acyclic = true
	}

	return order, acyclic
}
