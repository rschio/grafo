package grafo

import (
	"cmp"
	"path/filepath"
	"slices"
	"testing"
)

func TestRead(t *testing.T) {
	fname := "5_edge_list"
	g, err := Read(filepath.Join("testdata", fname))
	if err != nil {
		t.Fatalf("failed to read graph: %v", err)
	}

	if g.Order() != 5 {
		t.Errorf("got %d want %d vertices", g.Order(), 5)
	}

	edgeCount := 0
	for _, edges := range g.edges {
		edgeCount += len(edges)
	}

	if edgeCount != 9 {
		t.Errorf("got %d want %d edges", edgeCount, 9)
	}

	wantEdges := []Edge[int]{
		{V: 0, W: 1, Weight: 6},
		{V: 1, W: 2, Weight: 15},
		{V: 1, W: 3, Weight: 7},
		{V: 1, W: 4, Weight: 7},
		{V: 2, W: 4, Weight: 20},
		{V: 2, W: 1, Weight: 13},
		{V: 3, W: 0, Weight: 10},
		{V: 4, W: 0, Weight: 1},
		{V: 4, W: 3, Weight: 14},
	}

	gotEdges := make([]Edge[int], 0)
	for v := range g.Order() {
		for w, weight := range g.EdgesFrom(v) {
			gotEdges = append(gotEdges, Edge[int]{V: v, W: w, Weight: weight})
		}
	}

	cmpFn := func(a, b Edge[int]) int {
		switch {
		case a.V != b.V:
			return cmp.Compare(a.V, b.V)
		case a.W != b.W:
			return cmp.Compare(a.W, b.W)
		default:
			return cmp.Compare(a.Weight, b.Weight)
		}
	}

	slices.SortFunc(wantEdges, cmpFn)
	slices.SortFunc(gotEdges, cmpFn)

	if !slices.EqualFunc(gotEdges, wantEdges, func(a, b Edge[int]) bool {
		return cmpFn(a, b) == 0
	}) {
		t.Errorf("got diferent edges:\n got %v\nwant %v", gotEdges, wantEdges)
	}
}
