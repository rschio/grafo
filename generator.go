package grafo

import (
	"iter"
	"math/rand/v2"
	"strconv"
)

// GenerateRandomEdges genereate a random graph with V vertices and E edges,
// without self-loops or parallel edges.
// It panics if E > V * (V - 1).
//
// This function is generally not suitable for generating huge dense graphs.
func GenerateRandomEdges[T int | int64](V, E int, maxWeight T) *Mutable[T] {
	if E > V*(V-1) {
		panic("GenerateRandomEdges does not generate self-loops or parallel edges, but got: E > V * (V - 1)")
	}

	g := NewMutable[T](V)

	inserted := make(map[[2]int]struct{})
	for len(inserted) < E {
		v, w, weight := rand.N(V), rand.N(V), rand.N(maxWeight)

		edge := [2]int{v, w}
		// Avoid self-loops and parallel edges.
		if _, ok := inserted[edge]; v == w || ok {
			continue
		}
		inserted[edge] = struct{}{}

		g.Add(v, w, weight)
	}

	return g
}

func GenerateRandom(V, E, maxWeight int) *Mutable[int] {
	if E > V*(V-1) {
		panic("GenerateRandom does not generate self-loops or parallel edges, but got: E > V * (V - 1)")
	}

	p := float64(E) / float64(V) / float64(V-1)
	g := NewMutable[int](V)

	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			if i == j {
				continue
			}
			if rand.Float64() < p {
				g.Add(i, j, rand.N(maxWeight))
			}
		}
	}

	return g
}

func generateRandomWithRand[T any](V, E int, weightFn func() T, rnd *rand.Rand) *multigraph[T] {
	g := newMultigraph[T](V)
	for range E {
		g.Add(rnd.IntN(V), rnd.IntN(V), weightFn())
	}
	return g
}

type multigraph[T any] struct {
	edges [][]neighbor[T]
}

func newMultigraph[T any](n int) *multigraph[T] {
	return &multigraph[T]{edges: make([][]neighbor[T], n)}
}

func (g *multigraph[T]) Add(v, w int, weight T) {
	if w < 0 || w >= g.Order() {
		panic("vertex out of range: " + strconv.Itoa(w))
	}
	g.edges[v] = append(g.edges[v], neighbor[T]{vertex: w, weight: weight})
}

func (g *multigraph[T]) AddBoth(v, w int, weight T) {
	g.Add(v, w, weight)
	g.Add(w, v, weight)
}

func (g *multigraph[T]) Order() int { return len(g.edges) }

func (g *multigraph[T]) EdgesFrom(v int) iter.Seq2[int, T] {
	return func(yield func(w int, weight T) bool) {
		for _, e := range g.edges[v] {
			if !yield(e.vertex, e.weight) {
				return
			}
		}
	}
}

func (g *multigraph[T]) Visit(v int, do func(w int, c int64) bool) bool {
	for _, e := range g.edges[v] {
		wt, ok := any(e.weight).(int64)
		if !ok {
			wtInt, ok := any(e.weight).(int)
			wt = int64(wtInt)
			if !ok {
				wt = 1
			}
		}
		if do(e.vertex, wt) {
			return true
		}
	}
	return false
}
