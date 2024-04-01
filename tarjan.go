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
		time:       0,
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

//func StrongComponents[T any](g Graph[T]) [][]int {
//	n := g.Order()
//
//	s := &scc{
//		stk:     new(stack[int]),
//		visited: make([]bool, n),
//		low:     make([]int, n),
//	}
//
//	for v := range n {
//		if !s.visited[v] {
//			strongConnected(g, s, v)
//		}
//	}
//
//	return s.components
//}

//// Tarjan algorithm.
//type scc struct {
//	stk        *stack[int]
//	visited    []bool
//	low        []int
//	cnt        int
//	components [][]int
//}

//func strongConnected[T any](g Graph[T], s *scc, v int) {
//	// TODO: My brain melted. I really don't know why or how this works.
//	// It's a mix of Sedgewick with https://pure.tue.nl/ws/portalfiles/portal/167977703/Schols_W..pdf.
//	// Try to understand and simplify.
//	work := new(stack[[2]int])
//	work.Push([2]int{v, 0})
//
//	var minV int
//	for work.Len() > 0 {
//		ww := work.Pop()
//		v, j := ww[0], ww[1]
//		if j == 0 {
//			s.visited[v] = true
//			s.low[v] = s.cnt
//			s.cnt++
//			minV = s.low[v]
//			s.stk.Push(v)
//		}
//		recurse := false
//		for w, _ := range g.EdgesFrom(v) {
//			if !s.visited[w] {
//				work.Push([2]int{v, j + 1})
//				work.Push([2]int{w, 0})
//				recurse = true
//				break
//			}
//			minV = min(minV, s.low[w])
//		}
//		if !recurse {
//			if minV < s.low[v] {
//				s.low[v] = minV
//				continue
//			}
//			extractComponent(s, v)
//		}
//	}
//}

//func strongConnectedRecursive[T any](g Graph[T], s *scc, v int) {
//	s.visited[v] = true
//	s.low[v] = s.cnt
//	s.cnt++
//	minV := s.low[v]
//	s.stk.Push(v)
//
//	for w, _ := range g.EdgesFrom(v) {
//		if !s.visited[w] {
//			strongConnectedRecursive(g, s, w)
//		}
//		minV = min(minV, s.low[w])
//	}
//
//	if minV < s.low[v] {
//		s.low[v] = minV
//		return
//	}
//
//	// We are in the head.
//	extractComponent(s, v)
//}

//func extractComponent(s *scc, v int) {
//	var comp []int
//	for {
//		u := s.stk.Pop()
//		s.low[u] = math.MaxInt
//		comp = append(comp, u)
//		if u == v {
//			break
//		}
//	}
//	s.components = append(s.components, comp)
//}
