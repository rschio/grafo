package grafo

import (
	"iter"
	"math"
)

// StrongComponents returns a slice of g's strong connected components.
// Each slice contains the vertices of each component.
//
// A component is strongly connected if all its vertices are reachable
// from every other vertex in the component.
func StrongComponents[T any](g Graph[T]) [][]int {
	// StrongComponents use the algorithmT with non-recursive
	// depth-first search as described in
	// "Finding strong components using depth-fist search".
	n := g.Order()
	s := &scc[T]{
		g:          g,
		stk:        new(stack[int]),
		P:          new(stack[int]),
		iters:      make([]func() (int, T, bool), n),
		low:        make([]uint, n+1),
		components: [][]int{},
	}

	// Dummy.
	s.stk.Push(n)
	s.low[n] = 0

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
