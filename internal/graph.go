package internal

import (
	"iter"
)

// Graph is a copy of the grafo.Graph interface used only to avoid cyclic imports.
type Graph[T any] interface {
	Order() int
	EdgesFrom(v int) iter.Seq2[int, T]
}
