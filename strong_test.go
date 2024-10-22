package grafo

import (
	"cmp"
	"iter"
	"log"
	"path/filepath"
	"slices"
	"testing"

	gcmp "github.com/google/go-cmp/cmp"
	"github.com/rschio/graph"
	"golang.org/x/tools/txtar"
)

func TestStrongComponents(t *testing.T) {
	g := NewMutable[int64](6)
	g.Add(0, 1, 1)
	g.Add(1, 2, 1)
	g.Add(1, 3, 1)
	g.Add(1, 4, 1)
	g.Add(2, 0, 1)
	g.Add(2, 4, 1)
	g.Add(3, 5, 1)
	g.Add(4, 5, 1)
	g.Add(5, 4, 1)

	comps1 := StrongComponents(g)
	comps2 := graph.StrongComponents(toIterator(g))
	sortComponents(comps1)
	sortComponents(comps2)
	if !gcmp.Equal(comps1, comps2) {
		t.Errorf("\ngot %v\nwant%v\n", comps1, comps2)
	}
}

func TestStrongComponentsStackOverflow(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	var g line = 3_000_000
	StrongComponents(g)
}

func TestStrongComponentsExhaust(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	archive, err := txtar.ParseFile(filepath.Join("testdata", "exhaust5.txtar"))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range archive.Files {
		g, err := graphFromFile(f.Data)
		if err != nil {
			log.Fatal(err)
		}
		comps1 := StrongComponents(g)
		comps2 := graph.StrongComponents(g)
		sortComponents(comps1)
		sortComponents(comps2)
		if !gcmp.Equal(comps1, comps2) {
			t.Errorf("%s: got %v\nwant%v\n", f.Name, comps1, comps2)
		}
	}
}

func sortComponents(comps [][]int) {
	for i := range comps {
		slices.Sort(comps[i])
	}
	slices.SortFunc(comps, func(a, b []int) int {
		// a and b have at least one element.
		return cmp.Compare(a[0], b[0])
	})
}

func toIterator[T ~int | ~int64](g Graph[T]) *graph.Mutable {
	h := graph.New(g.Order())
	for v := range g.Order() {
		for w, weight := range g.EdgesFrom(v) {
			h.AddCost(v, w, int64(weight))
		}
	}
	return h
}

type line int

func (g line) Order() int { return int(g) }

func (g line) EdgesFrom(v int) iter.Seq2[int, int] {
	return func(yield func(w, weight int) bool) {
		if v+1 >= g.Order() {
			return
		}
		yield(v+1, 1)
	}
}
