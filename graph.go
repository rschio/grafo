package grafo

import (
	"iter"
	"math"
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
