package grafo

import (
	"iter"
	"math"
)

type Graph[T any] interface {
	Order() int
	EdgesFrom(v int) iter.Seq2[int, T]
}

type IntegerOrFloat interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

var maxes = [...]uint64{
	math.MaxInt8,
	math.MaxUint8,
	math.MaxInt16,
	math.MaxUint16,
	math.MaxInt32,
	math.MaxUint32,
	math.MaxInt64,
	math.MaxUint64,
}

func infFor[T IntegerOrFloat]() T {
	// Check if T is a float.
	var f float64 = 1.5
	if float64(T(f)) == f {
		return T(math.Inf(1))
	}

	// Check when v overflows.
	var v T
	for i := 0; v+1 > 0; i++ {
		v = T(maxes[i])
	}

	return v
}
