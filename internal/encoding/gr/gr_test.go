package gr

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func TestDecode(t *testing.T) {
	const fname = "USA-road-d.NY.gr"
	f, err := os.Open(filepath.Join("testdata", fname))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	g, err := NewDecoder(f, strconv.Atoi).Decode()
	if err != nil {
		t.Fatal(err)
	}

	wantV := 264346
	wantE := 733846
	if g.Order() != wantV {
		t.Errorf("got wrong number of vertices: %d, want %d", g.Order(), wantV)
	}

	edges := 0
	for v := range g.Order() {
		for _, _ = range g.EdgesFrom(v) {
			edges++
		}
	}
	if edges != wantE {
		t.Errorf("got wrong number of edges: %d, want %d", edges, wantE)
	}
}
