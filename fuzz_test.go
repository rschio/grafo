package grafo

import (
	"bytes"
	"math"
	"math/rand/v2"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func FuzzShortestPaths(f *testing.F) {
	f.Add(uint(5), int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, VV uint, maxValue int64) {
		V := int(VV%300 + 1) // Use a small V to test.
		E := rand.IntN(V*(V-1) + 1)
		if maxValue < 0 {
			maxValue = -(maxValue + 1)
		}
		if maxValue == 0 {
			maxValue = 1
		}
		g := GenerateRandomEdges(V, E, maxValue)

		v := rand.IntN(V)

		_, dist1 := ShortestPaths(g, v)
		_, dist2, _ := BellmanFord(g, v)
		//	_, dist3 := graph.ShortestPaths(toIterator(g), v)
		//	for i, val := range dist3 {
		//		if val == -1 {
		//			dist3[i] = InfFor[int64]()
		//		}
		//	}

		if diff := cmp.Diff(dist1, dist2); diff != "" {
			var buf bytes.Buffer
			if err := DOT(g, &buf); err != nil {
				t.Errorf("failed to DOT: %v", err)
			}
			t.Errorf("V=%d E=%d maxValue=%d v=%d\nGraph=[%s]\ndiff=%v", V, E, maxValue, v, buf.String(), diff)
		}
	})
}

//func toIterator(g Graph[int64]) graph.Iterator {
//	it := graph.New(g.Order())
//	for v := range g.Order() {
//		for w, weight := range g.EdgesFrom(v) {
//			it.AddCost(v, w, weight)
//		}
//	}
//	return it
//}
