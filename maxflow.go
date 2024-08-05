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

import "golang.org/x/exp/constraints"

func MaxFlow[T constraints.Integer](g Graph[T], s, t int) (flow T, graph Graph[T]) {
	// Edmonds-Karp's algorithm
	inf := InfFor[T]()
	n := g.Order()
	prev := make([]int, n)
	residual := Copy(g)
	for residualFlow(residual, s, t, prev) && flow < inf {
		pathFlow := inf
		for v := t; v != s; {
			u := prev[v]
			if c := residual.Weight(u, v); c < pathFlow {
				pathFlow = c
			}
			v = u
		}
		flow += pathFlow
		for v := t; v != s; {
			u := prev[v]
			residual.Add(u, v, residual.Weight(u, v)-pathFlow)
			residual.Add(v, u, residual.Weight(v, u)+pathFlow)
			v = u
		}
	}
	res := NewMutable[T](n)
	for v := 0; v < n; v++ {
		for w, weight := range g.EdgesFrom(v) {
			if flow := weight - residual.Weight(v, w); flow > 0 {
				res.Add(v, w, flow)
			}
		}
	}
	return flow, Sort(res)
}

func residualFlow[T constraints.Integer](g *Mutable[T], s, t int, prev []int) bool {
	visited := make([]bool, g.Order())
	prev[s], visited[s] = -1, true
	queue := newQueue(10)
	queue.Push(s)

	for queue.Len() > 0 {
		v := queue.Pop()
		for w, weight := range g.EdgesFrom(v) {
			if !visited[w] && weight > 0 {
				prev[w] = v
				visited[w] = true
				queue.Push(w)
			}
		}
	}
	return visited[t]
}
