package grafo

import "iter"

// DFS returns an iterator of edges that traverse the graph
// in Depth First Search way, starting from vertex v.
func DFS[T any](g Graph[T], v int) iter.Seq[Edge[T]] {
	return func(yield func(e Edge[T]) bool) {
		visited := make([]bool, g.Order())
		P := new(stack[vIter[T]])

		start := v
		next, stop := iter.Pull2(g.EdgesFrom(v))
		defer stop()
		w, weight, ok := next()

		for {
			if ok {
				if !visited[w] {
					visited[w] = true
					if !yield(Edge[T]{v, w, weight}) {
						return
					}

					P.Push(vIter[T]{v, next})
					v = w
					next, stop = iter.Pull2(g.EdgesFrom(v))
					defer stop()
					continue
				}
			} else {
				if v == start && P.Len() == 0 {
					return
				}

				w = v
				vi := P.Pop()
				v, next = vi.v, vi.iter
			}
			w, weight, ok = next()
		}
	}
}

type vIter[T any] struct {
	v    int
	iter func() (int, T, bool)
}

type stack[T any] struct {
	s []T
}

func (s *stack[T]) Len() int { return len(s.s) }

func (s *stack[T]) Push(v T) {
	s.s = append(s.s, v)
}

func (s *stack[T]) Top() T {
	return s.s[len(s.s)-1]
}

func (s *stack[T]) Pop() T {
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v
}
