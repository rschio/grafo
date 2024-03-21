package grafo

import "iter"

// DFS returns an iterator of edges that traverse the graph
// in Depth First Search way, starting from vertex v.
func DFS[T any](g Graph[T], v int) iter.Seq[Edge[T]] {
	return func(yield func(e Edge[T]) bool) {
		visited := make([]bool, g.Order())
		visited[v] = true
		stk := new(stack[T])

		for w, weight := range g.EdgesFrom(v) {
			stk.Push(Edge[T]{V: v, W: w, Weight: weight})
		}

		for stk.Len() > 0 {
			edge := stk.Pop()
			if visited[edge.W] {
				continue
			}

			if !yield(edge) {
				return
			}

			v = edge.W
			visited[v] = true

			for w, weight := range g.EdgesFrom(v) {
				if visited[w] {
					continue
				}
				stk.Push(Edge[T]{V: v, W: w, Weight: weight})
			}
		}
	}
}

type stack[T any] struct {
	s []Edge[T]
}

func (s *stack[T]) Len() int { return len(s.s) }

func (s *stack[T]) Push(v Edge[T]) {
	s.s = append(s.s, v)
}

func (s *stack[T]) Pop() Edge[T] {
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v
}
