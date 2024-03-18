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

import "cmp"

type prioQueue[S ~[]E, E cmp.Ordered] struct {
	heap  []int // vertices in heap order
	index []int // index of each vertex in the heap
	cost  S
}

func emptyPrioQueue[S ~[]E, E cmp.Ordered](cost S) *prioQueue[S, E] {
	return &prioQueue[S, E]{
		index: make([]int, len(cost)),
		cost:  cost,
	}
}

func newPrioQueue[S ~[]E, E cmp.Ordered](cost S) *prioQueue[S, E] {
	n := len(cost)
	q := &prioQueue[S, E]{
		heap:  make([]int, n),
		index: make([]int, n),
		cost:  cost,
	}
	for i := range q.heap {
		q.heap[i] = i
		q.index[i] = i
	}
	return q
}

// Len returns the number of elements in the queue.
func (q *prioQueue[S, E]) Len() int {
	return len(q.heap)
}

// Push pushes v onto the queue.
// The time complexity is O(log n) where n = q.Len().
func (q *prioQueue[S, E]) Push(v int) {
	n := q.Len()
	q.heap = append(q.heap, v)
	q.index[v] = n
	q.up(n)
}

// Pop removes the minimum element from the queue and returns it.
// The time complexity is O(log n) where n = q.Len().
func (q *prioQueue[S, E]) Pop() int {
	n := q.Len() - 1
	q.swap(0, n)
	q.down(0, n)

	v := q.heap[n]
	q.index[v] = -1
	q.heap = q.heap[:n]
	return v
}

// Contains tells whether v is in the queue.
func (q *prioQueue[S, E]) Contains(v int) bool {
	return q.index[v] >= 0
}

// Fix re-establishes the ordering after the cost for v has changed.
// The time complexity is O(log n) where n = q.Len().
func (q *prioQueue[S, E]) Fix(v int) {
	if i := q.index[v]; !q.down(i, q.Len()) {
		q.up(i)
	}
}

func (q *prioQueue[S, E]) less(i, j int) bool {
	return q.cost[q.heap[i]] < q.cost[q.heap[j]]
}

func (q *prioQueue[S, E]) swap(i, j int) {
	q.heap[i], q.heap[j] = q.heap[j], q.heap[i]
	q.index[q.heap[i]] = i
	q.index[q.heap[j]] = j
}

func (q *prioQueue[S, E]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !q.less(j, i) {
			break
		}
		q.swap(i, j)
		j = i
	}
}

func (q *prioQueue[S, E]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && q.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !q.less(j, i) {
			break
		}
		q.swap(i, j)
		i = j
	}
	return i > i0
}
