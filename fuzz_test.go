package grafo

import (
	"bytes"
	"math/rand/v2"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func FuzzShortestPaths(f *testing.F) {
	f.Fuzz(func(t *testing.T, VV uint, maxValue int) {
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

		if diff := cmp.Diff(dist1, dist2); diff != "" {
			var buf bytes.Buffer
			if err := DOT(g, &buf); err != nil {
				t.Errorf("failed to DOT: %v", err)
			}
			t.Errorf("V=%d E=%d maxValue=%d\nGraph=[%s]\ndiff=%v", V, E, maxValue, buf.String(), diff)
		}
	})
}
