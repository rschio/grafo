package grafo

import (
	"iter"
	"math"
	"strconv"
)

type Graph[T any] interface {
	Order() int
	EdgesFrom(v int) iter.Seq2[int, T]
}

func infFor[T ~int64 | ~float64]() T {
	var v T
	switch any(v).(type) {
	case int64:
		v = math.MaxInt64
	case float64:
		v = T(math.Inf(1))
	}
	return v
}

type Mutable[T any] struct {
	edges []map[int]T
}

func NewMutable[T any](n int) *Mutable[T] {
	return &Mutable[T]{edges: make([]map[int]T, n)}
}

func (g *Mutable[T]) Add(v, w int, weight T) {
	if w < 0 || w >= g.Order() {
		panic("vertex out of range: " + strconv.Itoa(w))
	}
	if g.edges[v] == nil {
		g.edges[v] = make(map[int]T)
	}
	g.edges[v][w] = weight
}

func (g *Mutable[T]) AddBoth(v, w int, weight T) {
	g.Add(v, w, weight)
	g.Add(w, v, weight)
}

func (g *Mutable[T]) Order() int { return len(g.edges) }

func (g *Mutable[T]) EdgesFrom(i int) iter.Seq2[int, T] {
	return func(yield func(w int, weight T) bool) {
		for w, weight := range g.edges[i] {
			if !yield(w, weight) {
				return
			}
		}
	}
}
