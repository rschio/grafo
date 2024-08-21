package grafo

import "testing"

func TestGenerator(t *testing.T) {
	V := 10
	E := 90
	maxWeight := 50
	g := generateRandomEdges(V, E, maxWeight)

	if g.Order() != V {
		t.Errorf("got %d want %d vertices", g.Order(), V)
	}

	count := 0
	for v := range g.Order() {
		for _, wt := range g.EdgesFrom(v) {
			count++
			if wt < 0 || wt > maxWeight {
				t.Errorf("invalid weight %d, max is %d", wt, maxWeight)
			}
		}
	}

	if count != E {
		t.Errorf("got %d want %d edges", count, E)
	}
}

func TestGeneratorRandom(t *testing.T) {
	V := 50
	E := 500
	maxWeight := 50
	g := generateRandom(V, E, maxWeight)

	if g.Order() != V {
		t.Errorf("got %d want %d vertices", g.Order(), V)
	}

	count := 0
	for v := range g.Order() {
		for _, wt := range g.EdgesFrom(v) {
			count++
			if wt < 0 || wt > maxWeight {
				t.Errorf("invalid weight %d, max is %d", wt, maxWeight)
			}
		}
	}

	// The edges generation is probabilistic, so it will generate
	// aproximatly E edges. Consider +-20% of error.
	pErr := 0.2
	smallest := int(float64(E) * (1 - pErr))
	largest := int(float64(E) * (1 + pErr))
	if count < smallest || count > largest {
		t.Errorf("got %d want between [%d, %d] edges", count, smallest, largest)
	}
}
