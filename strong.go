package grafo

import (
	"iter"
	"math"
)

func StrongComponents[T any](g Graph[T]) [][]int {
	// StrongComponents use the algorithmT with non-recursive
	// depth-first search as described in
	// "Finding strong components using depth-fist search".
	//
	// TODO: What to choose?
	// The non-recursive algorithmT needs a pull iterator to
	// work, but the Graph interface is a push iterator.
	// To solve the problem I could copy the graph and store it in
	// a favorable struct or change the push iterator to a pull one.
	// To copy the graph we need |E| space (in dense graphs ~|V²|) or
	// changing the iterator type slows down by a factor of 10 or 20x.
	// I choosed to slow down and keep the memory in a safe amount.
	n := g.Order()
	stk := new(stack[int])
	low := make([]uint, n+1)

	// Dummy.
	stk.Push(n)
	low[n] = 0

	s := &scc[T]{
		g:          g,
		stk:        stk,
		P:          new(stack[int]),
		iters:      make([]func() (int, T, bool), n),
		low:        low,
		components: [][]int{},
	}
	for v := range g.Order() {
		if s.unvisited(v) {
			s.dfsI(v)
		}
	}

	return s.components
}

type scc[T any] struct {
	g          Graph[T]
	stk        *stack[int]
	P          *stack[int]
	iters      []func() (int, T, bool)
	low        []uint
	components [][]int
	time       uint
}

func (s *scc[T]) dfsI(start int) {
	v := start
	s.previsit(v)
	next, stop := iter.Pull2(s.g.EdgesFrom(v))
	defer stop()
	w, _, ok := next()

	for {
		if ok {
			if s.unvisited(w) {
				s.previsit(w)
				// FORWARD
				s.iters[v] = next
				s.P.Push(v)
				v = w
				next, stop = iter.Pull2(s.g.EdgesFrom(v))
				defer stop()
				// FORWARD END
				continue
			}
		} else {
			s.postvisit(v)
			if v == start {
				return
			}
			// BACKWARD
			w = v
			v = s.P.Pop()
			next = s.iters[v]
			// BACKWARD END
		}
		s.retreat(v, w)
		w, _, ok = next()
	}
}

func (s *scc[_]) unvisited(v int) bool {
	return s.low[v] == 0
}

func (s *scc[_]) leader(v int) bool {
	// Encoding trick. The leader bit is encoded
	// as the negation of least significative bit.
	return s.low[v]&1 == 0
}

func (s *scc[_]) previsit(v int) {
	s.time = s.time + 2
	s.low[v] = s.time
}

func (s *scc[_]) retreat(v, w int) {
	if s.low[w] < s.low[v] {
		s.low[v] = s.low[w] | 1
	}
}

func (s *scc[_]) postvisit(v int) {
	if !s.leader(v) {
		s.stk.Push(v)
		return
	}

	var comp []int
	for s.low[s.stk.Top()] >= s.low[v] {
		x := s.stk.Pop()
		comp = append(comp, x)
		s.low[x] = math.MaxUint
	}
	comp = append(comp, v)
	s.low[v] = math.MaxUint
	s.components = append(s.components, comp)
}
