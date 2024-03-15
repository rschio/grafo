package grafo

import "testing"

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
