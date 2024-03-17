// Modified from github.com/yourbasic/graph.

// BSD 2-Clause License
//
// Copyright (c) 2017, Stefan Nilsson
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package grafo

import (
	"iter"
	"strconv"
)

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
