package grafo

import (
	"iter"
	"strconv"
)

// DFS returns an iterator of edges that traverse the graph
// in Depth First Search way, starting from vertex v.
func DFS[T any](g Graph[T], v int) iter.Seq[Edge[T]] {
	return func(yield func(e Edge[T]) bool) {
		noop := func(int) bool { return true }
		dfs(g, v, yield, noop)
	}
}

func DFSPrevisit[T any](g Graph[T], v int) iter.Seq[int] {
	return func(yield func(v int) bool) {
		if v < 0 || v >= g.Order() {
			panic("vertex out of range: " + strconv.Itoa(v))
		}

		if !yield(v) {
			return
		}

		previsit := func(e Edge[T]) bool {
			return yield(e.W)
		}
		noop := func(int) bool { return true }
		dfs(g, v, previsit, noop)
	}
}

func DFSPostvisit[T any](g Graph[T], v int) iter.Seq[int] {
	return func(yield func(v int) bool) {
		noop := func(Edge[T]) bool { return true }
		dfs(g, v, noop, yield)
	}
}

func dfs[T any](g Graph[T], v int, previsit func(e Edge[T]) bool, postvisit func(v int) bool) {
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
			if !postvisit(v) {
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
