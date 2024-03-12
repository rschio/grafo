package grafo

import "testing"

func TestGenerator(t *testing.T) {
	V := 10
	E := 90
	maxWeight := 50
	g := GenerateRandomEdges(V, E, maxWeight)
	t.Logf("%+v", g)
}

func TestGeneratorRandom(t *testing.T) {
	V := 10
	E := 30
	maxWeight := 50
	g := GenerateRandom(V, E, maxWeight)
	t.Logf("%+v", g)
}
