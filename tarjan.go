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

	s := &tarjanS[T]{
		g:          g,
		stk:        stk,
		low:        low,
		components: [][]int{},
	}
	for v := range g.Order() {
		if s.unvisited(v) {
			dfsR(s, v)
		}
	}

	return s.components
}

func dfsR[T any](s *tarjanS[T], v int) {
	s.previsit(v)
	scan(s, v)
	s.postvisit(v)
}

func scan[T any](s *tarjanS[T], v int) {
	for w, _ := range s.g.EdgesFrom(v) {
		if s.unvisited(w) {
			dfsR(s, w)
		}
		s.retreat(v, w)
	}
}

type tarjanS[T any] struct {
	g          Graph[T]
	stk        *stack[int]
	low        []int
	components [][]int
	time       int
}

func (s *tarjanS[T]) unvisited(v int) bool {
	return s.low[v] == 0
}

func (s *tarjanS[T]) leader(v int) bool {
	// Encoding trick. The leader bit is encoded
	// as the negation of least significative bit.
	return s.low[v]&1 == 0
}

func (s *tarjanS[T]) previsit(v int) {
	s.time = s.time + 2
	s.low[v] = s.time
}

func (s *tarjanS[T]) retreat(v, w int) {
	if s.low[w] < s.low[v] {
		s.low[v] = s.low[w] | 1
	}
}

func (s *tarjanS[T]) postvisit(v int) {
	if s.leader(v) {
		s.extractComponent(v)
	} else {
		s.stk.Push(v)
	}
}

func (s *tarjanS[T]) extractComponent(v int) {
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
