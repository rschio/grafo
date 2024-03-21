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
	"math/rand/v2"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMST(t *testing.T) {
	g := NewMutable[uint8](0)
	if diff := cmp.Diff(MST(g), []int{}); diff != "" {
		t.Errorf("MST: %s", diff)
	}

	g = NewMutable[uint8](10)
	g.AddBoth(0, 1, 4)
	g.AddBoth(0, 7, 8)
	g.AddBoth(1, 2, 8)
	g.AddBoth(1, 7, 11)
	g.AddBoth(2, 3, 7)
	g.AddBoth(2, 8, 2)
	g.AddBoth(2, 5, 4)
	g.AddBoth(3, 4, 9)
	g.AddBoth(3, 5, 14)
	g.AddBoth(4, 5, 10)
	g.AddBoth(5, 6, 2)
	g.AddBoth(6, 7, 1)
	g.AddBoth(6, 8, 6)
	g.AddBoth(7, 8, 7)
	exp := []int{-1, 0, 5, 2, 3, 6, 7, 0, 2, -1}
	if diff := cmp.Diff(MST(g), exp); diff != "" {
		t.Errorf("MST: %s", diff)
	}
}

func BenchmarkMST(b *testing.B) {
	n := 1000
	b.StopTimer()
	g := NewMutable[int64](n)
	for i := 0; i < 2*n; i++ {
		g.Add(rand.IntN(n), rand.IntN(n), int64(rand.Int()))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = MST(g)
	}
}
