package multigraph

import (
	"iter"
	"strconv"
)

type Multigraph[T any] struct {
	edges [][]neighbor[T]
}

type neighbor[T any] struct {
	vertex int
	weight T
}

func New[T any](n int) *Multigraph[T] {
	return &Multigraph[T]{edges: make([][]neighbor[T], n)}
}

func (g *Multigraph[T]) Add(v, w int, weight T) {
	if w < 0 || w >= g.Order() {
		panic("vertex out of range: " + strconv.Itoa(w))
	}
	g.edges[v] = append(g.edges[v], neighbor[T]{vertex: w, weight: weight})
}

func (g *Multigraph[T]) AddBoth(v, w int, weight T) {
	g.Add(v, w, weight)
	g.Add(w, v, weight)
}

func (g *Multigraph[T]) Order() int { return len(g.edges) }

func (g *Multigraph[T]) EdgesFrom(v int) iter.Seq2[int, T] {
	return func(yield func(w int, weight T) bool) {
		for _, e := range g.edges[v] {
			if !yield(e.vertex, e.weight) {
				return
			}
		}
	}
}

func (g *Multigraph[T]) Visit(v int, do func(w int, c int64) bool) bool {
	for _, e := range g.edges[v] {
		wt, ok := any(e.weight).(int64)
		if !ok {
			wtInt, ok := any(e.weight).(int)
			wt = int64(wtInt)
			if !ok {
				wt = 1
			}
		}
		if do(e.vertex, wt) {
			return true
		}
	}
	return false
}
