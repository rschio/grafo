package grafo

import (
	"testing"
)

var complete = Sort(completeGraph(1000))

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
