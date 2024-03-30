package grafo

import (
	"math/rand/v2"
	"path/filepath"
	"testing"

	"github.com/rschio/graph"
)

var complete = Sort(completeGraph(1000))
var pathG = Sort(pathGraph(1000))
var dimacsG = readDIMACS()

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

func BenchmarkShortestPaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ShortestPaths(pathG, 0)
	}
}

func BenchmarkShortestPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ShortestPath(pathG, 0, pathG.Order()-1)
	}
}

func BenchmarkBellmanFord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = BellmanFord(pathG, 0)
	}
}

func BenchmarkBigShortestPaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ShortestPaths(dimacsG, 0)
	}
}

func BenchmarkBigShortestPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ShortestPath(dimacsG, 0, pathG.Order()-1)
	}
}

func BenchmarkBigBellmanFord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = BellmanFord(dimacsG, 0)
	}
}

func BenchmarkTarjan(b *testing.B) {
	for range b.N {
		_ = StrongComponents(pathG)
	}
}

func BenchmarkStrongComponent(b *testing.B) {
	pathGIterator := graph.Sort(toIterator(pathG))
	b.ResetTimer()
	for range b.N {
		_ = graph.StrongComponents(pathGIterator)
	}
}

func pathGraph(n int) *Mutable[int64] {
	g := NewMutable[int64](n)
	for i := 0; i < 2*n; i++ {
		g.Add(rand.IntN(n), rand.IntN(n), int64(rand.Int()))
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

func readDIMACS() *Immutable[int] {
	g, err := Read(filepath.Join("testdata", "264346_DIMACS"))
	if err != nil {
		panic(err)
	}
	return Sort(g)
}
