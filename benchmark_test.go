package grafo

import (
	"math/rand/v2"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/rschio/grafo/internal/testutil"
)

var complete = Sort(completeGraph(1000))
var complete2 = Sort(completeGraph2(1000))
var pathG = Sort(pathGraph(1000))
var dimacsG = readDIMACS("usa-road-d.ny")

func BenchmarkBellmanFord(b *testing.B) {
	benchmarks := []struct {
		name string
		g    *Immutable[int]
	}{
		{"Normal", pathG},
		{"Big", dimacsG},
	}
	for _, bb := range benchmarks {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, _ = BellmanFord(bb.g, 0)
			}
		})
	}
}

func BenchmarkShortestPaths(b *testing.B) {
	benchmarks := []struct {
		name string
		g    *Immutable[int]
	}{
		{"Normal", pathG},
		{"Big", dimacsG},
	}

	for _, bb := range benchmarks {
		b.Run(bb.name, func(b *testing.B) {
			for range b.N {
				_, _ = ShortestPaths(bb.g, 0)
			}
		})
	}
}

func BenchmarkShortestPath(b *testing.B) {
	benchmarks := []struct {
		name string
		g    *Immutable[int]
	}{
		{"Normal", pathG},
		{"Big", dimacsG},
	}

	for _, bb := range benchmarks {
		b.Run(bb.name, func(b *testing.B) {
			for range b.N {
				_, _ = ShortestPath(bb.g, 0, bb.g.Order()-1)
			}
		})
	}
}

func BenchmarkStrongComponents(b *testing.B) {
	benchmarks := []struct {
		name string
		g    *Immutable[int]
	}{
		{"Complete2", complete2},
		{"Big", dimacsG},
	}

	for _, bb := range benchmarks {
		b.Run(bb.name, func(b *testing.B) {
			for range b.N {
				_ = StrongComponents(bb.g)
			}
		})
	}
}

func BenchmarkBFS(b *testing.B) {
	for range b.N {
		for e := range BFS(complete, 0) {
			_ = e
		}
	}
}

func BenchmarkDFSIter(b *testing.B) {
	for range b.N {
		for e := range DFS(complete, 0) {
			_ = e
		}
	}
}

func Benchmark_infFor(b *testing.B) {
	for range b.N {
		v := InfFor[uint64]()
		_ = v
	}
}

func BenchmarkRange(b *testing.B) {
	var g Graph[struct{}] = complete

	for range b.N {
		for v := range g.Order() {
			for w, wt := range g.EdgesFrom(v) {
				_, _ = w, wt
			}
		}
	}
}

func pathGraph(n int) *Mutable[int] {
	g := NewMutable[int](n)
	for i := 0; i < 2*n; i++ {
		g.Add(rand.IntN(n), rand.IntN(n), rand.Int())
	}
	return g
}

func completeGraph2(n int) *Mutable[int] {
	g := NewMutable[int](n)
	for v := range n {
		for w := range n {
			if v == w {
				continue
			}
			g.Add(v, w, 1)
		}
	}
	return g
}

func completeGraph(n int) *Mutable[struct{}] {
	g := NewMutable[struct{}](n)
	for v := range n {
		for w := range n {
			if v == w {
				continue
			}
			g.Add(v, w, struct{}{})
		}
	}
	return g
}

func readDIMACS(fname string) *Immutable[int] {
	g, err := testutil.ReadFile(filepath.Join("testdata", fname), strconv.Atoi)
	if err != nil {
		panic(err)
	}
	return Sort(g)
}
