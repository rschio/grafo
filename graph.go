package grafo

import (
	"iter"
	"math"
	"reflect"
)

// Graph describes a directed weighted graph.
//
// The Graph interface can be used to describe both oridinary graphs and
// multigraphs.
type Graph[T any] interface {
	// Order returns the number of vertices in a graph.
	Order() int

	// EdgesFrom returns an iterator that iterates over the outgoing edges of v.
	// Each iteration returns w and weight of an edge v-[weight]->w, weight is
	// the weight of the edge.
	//
	// The iteration may occur in any order, and the order may vary.
	EdgesFrom(v int) iter.Seq2[int, T]
}

type IntegerOrFloat interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

func InfFor[T IntegerOrFloat]() T {
	switch reflect.TypeFor[T]().Kind() {
	case reflect.Int8:
		return math.MaxInt8
	case reflect.Int16:
		v := uint64(math.MaxInt16)
		return T(v)
	case reflect.Int32:
		v := uint64(math.MaxInt32)
		return T(v)
	case reflect.Int64:
		v := uint64(math.MaxInt64)
		return T(v)
	case reflect.Int:
		v := uint64(math.MaxInt)
		return T(v)
	case reflect.Uint8:
		v := uint64(math.MaxUint8)
		return T(v)
	case reflect.Uint16:
		v := uint64(math.MaxUint16)
		return T(v)
	case reflect.Uint32:
		v := uint64(math.MaxUint32)
		return T(v)
	case reflect.Uint64:
		v := uint64(math.MaxUint64)
		return T(v)
	case reflect.Uint, reflect.Uintptr:
		v := uint64(math.MaxUint)
		return T(v)
	case reflect.Float32, reflect.Float64:
		return T(math.Inf(1))
	default:
		return *new(T)
	}
}
