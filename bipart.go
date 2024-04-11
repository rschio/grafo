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

// Bipartition returns a subset U of g's vertices with the property
// that every edge of g connects a vertex in U to one outside of U.
// If g isn't bipartite, it returns an empty slice and sets ok to false.
func Bipartition[T any](g Graph[T]) (part []int, ok bool) {
	type color byte
	const (
		none color = iota
		white
		black
	)
	colors := make([]color, g.Order())
	whiteCount := 0

	for v := range colors {
		if colors[v] != none {
			continue
		}
		colors[v] = white
		whiteCount++
		queue := newQueue(1)
		queue.Push(v)
		for queue.Len() > 0 {
			v := queue.Pop()
			for w, _ := range g.EdgesFrom(v) {
				switch {
				case colors[w] != none:
					if colors[v] == colors[w] {
						return []int{}, false
					}
					continue
				case colors[v] == white:
					colors[w] = black
				default:
					colors[w] = white
					whiteCount++
				}
				queue.Push(w)
			}
		}
	}

	part = make([]int, 0, whiteCount)
	for v, color := range colors {
		if color == white {
			part = append(part, v)
		}
	}
	return part, true
}
