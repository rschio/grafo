package grafo

import "testing"

func BenchmarkBFS(b *testing.B) {
	g := Sort(completeGraph(1000))

	b.ResetTimer()
	b.ReportAllocs()
	for range b.N {
		for e := range BFS(g, 0) {
			_ = e
		}
	}
}

func BenchmarkDFSIter(b *testing.B) {
	//g := Sort(GenerateRandomEdges(1_000, 5_000, 20))
	g := Sort(completeGraph(1000))

	b.ResetTimer()
	b.ReportAllocs()
	for range b.N {
		for e := range DFS(g, 0) {
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
