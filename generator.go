package grafo

import (
	"math/rand/v2"
)

// GenerateRandomEdges genereate a random graph with V vertices and E edges,
// without self-loops or parallel edges.
// It panics if E > V * (V - 1).
//
// This function is generally not suitable for generating huge dense graphs.
func GenerateRandomEdges(V, E, maxWeight int) *Mutable[int] {
	if E > V*(V-1) {
		panic("GenerateRandomEdges does not generate self-loops or parallel edges, but got: E > V * (V - 1)")
	}

	g := NewMutable[int](V)

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
