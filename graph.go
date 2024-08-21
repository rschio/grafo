// Package grafo contains generic implementations of basic graph algorithms.
//
// Most of the functions of this package operates on the Graph[T] interface,
// a small but powerfull interface that represents a directed weighted graph.
package grafo

import (
	"iter"
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
