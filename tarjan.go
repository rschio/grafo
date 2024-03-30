package grafo

func Tarjan[T any](g Graph[T]) [][]int {
	n := g.Order()
	pre := make([]int, n)
	low := make([]int, n)
	for i := range n {
		pre[i], low[i] = -1, -1
	}

	s := &scc{
		stk: new(stack[int]),
		pre: pre,
		low: low,
	}
	for v := range n {
		if s.pre[v] == -1 {
			strongConnected(g, s, v)
		}
	}

	return s.components
}

type scc struct {
	stk        *stack[int]
	pre        []int
	low        []int
	cnt        int
	components [][]int
}

func strongConnected[T any](g Graph[T], s *scc, v int) {
	// TODO: My brain melted. I really don't know why or how this works.
	// It's a mix of Sedgewick with https://pure.tue.nl/ws/portalfiles/portal/167977703/Schols_W..pdf.
	// Try to understand and simplify.
	work := new(stack[[2]int])
	work.Push([2]int{v, 0})

	var minV int
	for work.Len() > 0 {
		ww := work.Pop()
		v, j := ww[0], ww[1]
		if j == 0 {
			s.pre[v] = s.cnt
			s.cnt++
			s.low[v] = s.pre[v]
			minV = s.low[v]
			s.stk.Push(v)
		}
		recurse := false
		i := j
		for w, _ := range g.EdgesFrom(v) {
			if s.pre[w] == -1 {
				work.Push([2]int{v, i + 1})
				work.Push([2]int{w, 0})
				recurse = true
				break
			}
			minV = min(minV, s.low[w])
		}
		if !recurse {
			if minV < s.low[v] {
				s.low[v] = minV
				continue
			}

			var comp []int
			for {
				u := s.stk.Pop()
				s.low[u] = g.Order()
				comp = append(comp, u)
				if u == v {
					break
				}
			}
			s.components = append(s.components, comp)
		}
	}
}

//func SCdfsR[T any](g Graph[T], s *scc, v int) {
//	s.pre[v] = s.cnt
//	s.cnt++
//	s.low[v] = s.pre[v]
//	minV := s.low[v]
//	s.stk.Push(v)
//
//	for w, _ := range g.EdgesFrom(v) {
//		if s.pre[w] == -1 {
//			SCdfsR(g, s, w)
//		}
//		minV = min(minV, s.low[w])
//	}
//
//	if minV < s.low[v] {
//		s.low[v] = minV
//		return
//	}
//
//	var comp []int
//	for {
//		u := s.stk.Pop()
//		s.low[u] = g.Order()
//		comp = append(comp, u)
//		if u == v {
//			break
//		}
//	}
//	s.components = append(s.components, comp)
//}
