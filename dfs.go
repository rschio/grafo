package grafo

import (
	"iter"
	"strconv"
)

// DFS returns an iterator of edges that traverse the graph
// in Depth First Search way, starting from vertex v.
func DFS[T any](g Graph[T], v int) iter.Seq[Edge[T]] {
	return func(yield func(e Edge[T]) bool) {
		dfs(g, v, yield, func(int) bool { return true })
	}
}

func DFSPrevisit[T any](g Graph[T], v int) iter.Seq[int] {
	return func(yield func(v int) bool) {
		if v >= g.Order() {
			panic("vertex out of range: " + strconv.Itoa(v))
		}

		if !yield(v) {
			return
		}

		previsit := func(e Edge[T]) bool {
			return yield(e.W)
		}
		dfs(g, v, previsit, func(int) bool { return true })
	}
}

func DFSPostvisit[T any](g Graph[T], v int) iter.Seq[int] {
	return func(yield func(v int) bool) {
		dfs(g, v, func(Edge[T]) bool { return true }, yield)
	}
}

func dfs[T any](g Graph[T], v int, previsit func(e Edge[T]) bool, posvisit func(v int) bool) {
	visited := make([]bool, g.Order())
	visited[v] = true
	path := new(stack[vIter[T]])

	next, stop := iter.Pull2(g.EdgesFrom(v))
	defer stop()
	w, weight, ok := next()

	for {
		switch {
		case ok && visited[w]:
			w, weight, ok = next()

		case ok && !visited[w]:
			if !previsit(Edge[T]{v, w, weight}) {
				return
			}
			visited[w] = true

			path.Push(vIter[T]{v, next})
			v = w
			next, stop = iter.Pull2(g.EdgesFrom(v))
			defer stop()

		case !ok:
			if !posvisit(v) {
				return
			}
			if path.Len() == 0 {
				return
			}
			vi := path.Pop()
			v, next = vi.v, vi.iter
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
