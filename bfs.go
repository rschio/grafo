package grafo

import "iter"

// Edge is a directed graph edge V -[Weight]-> W.
type Edge[T any] struct {
	V, W   int
	Weight T
}

// BFS returns an iterator of edges that traverse the graph
// in Breadth First Search way, starting from vertex v.
func BFS[T any](g Graph[T], v int) iter.Seq[Edge[T]] {
	return func(yield func(e Edge[T]) bool) {
		visited := make([]bool, g.Order())
		visited[v] = true
		queue := newQueue(10)
		queue.Push(v)

		for queue.Len() > 0 {
			v = queue.Pop()
			for w, weight := range g.EdgesFrom(v) {
				if visited[w] {
					continue
				}

				if !yield(Edge[T]{V: v, W: w, Weight: weight}) {
					return
				}

				visited[w] = true
				queue.Push(w)
			}
		}
	}
}
