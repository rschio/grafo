package grafo

import (
	"bytes"
	"iter"
	"math/rand/v2"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rschio/graph"
)

func FuzzStrongComponents(f *testing.F) {
	f.Fuzz(func(t *testing.T, VV uint) {
		V := int(VV%100_000 + 1)
		E := rand.IntN(V*(V-1) + 1)
		g := GenerateRandomEdges(V, E, 1)

		comps1 := StrongComponents(g)
		comps2 := graph.StrongComponents(toIterator(g))
		sortComponents(comps1)
		sortComponents(comps2)

		if !cmp.Equal(comps1, comps2) {
			var buf bytes.Buffer
			if err := DOT(g, &buf); err != nil {
				t.Errorf("failed to DOT: %v", err)
			}
			t.Errorf("V=%d E=%d\nGraph=[%s]\ncomps1=%v\ncomps2=%v", V, E, buf.String(), comps1, comps2)
		}
	})
}

func FuzzShortestPaths(f *testing.F) {
	f.Add(uint(10), uint(20), InfFor[int64](), uint64(0), uint64(1))
	f.Fuzz(func(t *testing.T, VV, EE uint, maxValue int64, seed1, seed2 uint64) {
		V := int(VV%500 + 1) // Use a small V to test.
		E := int(EE % uint(V*V))
		if maxValue <= 0 {
			maxValue = -maxValue
			if maxValue == 0 {
				maxValue = 1
			}
		}
		rnd := rand.New(rand.NewPCG(seed1, seed2))
		weightFn := func() int64 {
			return rnd.Int64N(maxValue)
		}
		g := generateRandomWithRand(V, E, weightFn, rnd)
		v := rand.IntN(V)

		_, dist1 := ShortestPaths(g, v)
		_, dist2, _ := BellmanFord(g, v)

		if diff := cmp.Diff(dist1, dist2); diff != "" {
			var buf bytes.Buffer
			if err := DOT(g, &buf); err != nil {
				t.Errorf("failed to DOT: %v", err)
			}
			t.Errorf("V=%d E=%d maxValue=%d v=%d\nGraph=[%s]\ndiff=%v", V, E, maxValue, v, buf.String(), diff)
		}
	})
}

func FuzzDFS(f *testing.F) {
	f.Fuzz(func(t *testing.T, VV, EE uint, seed1, seed2 uint64) {
		V := int(VV%1000) + 1
		E := int(EE % uint((V * V)))
		rnd := rand.New(rand.NewPCG(seed1, seed2))
		g := generateRandomWithRand(V, E, func() int64 { return 1 }, rnd)

		next1, stop1 := iter.Pull(DFS(g, 0))
		defer stop1()

		next2, stop2 := iter.Pull(dfsRec(g, 0))
		defer stop2()

		path := make([]Edge[int64], 0)
		for {
			e1, ok1 := next1()
			e2, ok2 := next2()

			path = append(path, e1)
			if diff := cmp.Diff(e1, e2); diff != "" || ok1 != ok2 {
				var buf bytes.Buffer
				DOT(g, &buf)
				t.Fatalf("ok1 %v ok2 %v diff: %s\npath[%v]\n%s", ok1, ok2, diff, path, buf.String())
			}
			if ok1 == false {
				break
			}
		}
	})
}

func dfsRec[T any](g Graph[T], v int) iter.Seq[Edge[T]] {
	return func(yield func(e Edge[T]) bool) {
		visited := make([]bool, g.Order())
		dfsR(g, visited, yield, v)
	}
}

func dfsR[T any](g Graph[T], visited []bool, yield func(e Edge[T]) bool, v int) {
	for w, wt := range g.EdgesFrom(v) {
		if visited[w] {
			continue
		}
		visited[w] = true
		if !yield(Edge[T]{v, w, wt}) {
			return
		}
		dfsR(g, visited, yield, w)
	}
}
