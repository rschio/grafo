package grafo

import "iter"

// BFS returns an iterator of edges that traverse the graph
// in Breadth First Search way, starting from vertex v.
func BFS[T any](g Graph[T], v int) iter.Seq[Edge[T]] {
	return func(yield func(e Edge[T]) bool) {
		visited := make([]bool, g.Order())
		visited[v] = true
		queue := newBfsQueue(10)
		queue.Insert(v)

		for queue.Len() > 0 {
			v = queue.Remove()
			for w, weight := range g.EdgesFrom(v) {
				if visited[w] {
					continue
				}

				if !yield(Edge[T]{V: v, W: w, Weight: weight}) {
					return
				}

				visited[w] = true
				queue.Insert(w)
			}
		}
	}
}

type bfsQueue struct {
	q     []int
	first int
}

func newBfsQueue(cap int) *bfsQueue {
	return &bfsQueue{q: make([]int, 0, cap)}
}

func (q *bfsQueue) Len() int { return len(q.q) - q.first }

func (q *bfsQueue) Remove() int {
	v := q.q[q.first]
	q.first++
	return v
}

func (q *bfsQueue) Insert(v int) {
	if len(q.q) == cap(q.q) {
		if q.first > len(q.q)/4 {
			l := q.Len()
			copy(q.q[:], q.q[q.first:])
			q.q = q.q[:l]
			q.first = 0
		}
	}
	q.q = append(q.q, v)
}
