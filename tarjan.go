package grafo

import (
	"math"
)

func StrongComponents[T any](g Graph[T]) [][]int {
	n := g.Order()
	stk := new(stack[int])
	low := make([]int, n+1)

	// Dummy.
	stk.Push(n)
	low[n] = 0

	s := &scc[T]{
		g:          g,
		stk:        stk,
		low:        low,
		components: [][]int{},
	}
	for v := range g.Order() {
		if s.unvisited(v) {
			s.dfsR(v)
		}
	}

	return s.components
}

type scc[T any] struct {
	g          Graph[T]
	stk        *stack[int]
	low        []int
	components [][]int
	time       int
}

func (s *scc[_]) dfsR(v int) {
	s.previsit(v)
	s.scan(v)
	s.postvisit(v)
}

func (s *scc[_]) scan(v int) {
	for w, _ := range s.g.EdgesFrom(v) {
		if s.unvisited(w) {
			s.dfsR(w)
		}
		s.retreat(v, w)
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
		s.low[x] = math.MaxInt
	}
	comp = append(comp, v)
	s.low[v] = math.MaxInt
	s.components = append(s.components, comp)
}
