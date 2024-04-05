package grafo

type queue struct {
	q     []int
	first int
}

func newQueue(cap int) *queue {
	return &queue{q: make([]int, 0, cap)}
}

func (q *queue) Len() int { return len(q.q) - q.first }

func (q *queue) Pop() int {
	v := q.q[q.first]
	q.first++
	return v
}

func (q *queue) Push(v int) {
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
