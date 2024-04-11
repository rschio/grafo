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
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBipartition(t *testing.T) {
	g1 := NewMutable[int](1)
	g1.Add(0, 0, 0)

	g2 := NewMutable[int](2)
	g2.AddBoth(0, 1, 0)

	g3 := NewMutable[int](2)
	g3.AddBoth(0, 1, 0)
	g3.Add(0, 0, 0)

	g4 := NewMutable[int](5)
	g4.Add(0, 1, 0)
	g4.Add(0, 2, 0)
	g4.Add(0, 3, 0)

	g5 := NewMutable[int](5)
	g5.Add(0, 1, 0)
	g5.Add(0, 2, 0)
	g5.Add(0, 3, 0)
	g5.Add(2, 3, 0)

	g6 := NewMutable[int](5)
	g6.AddBoth(0, 1, 0)
	g6.AddBoth(1, 2, 0)
	g6.AddBoth(2, 3, 0)
	g6.AddBoth(3, 0, 0)

	tests := []struct {
		name     string
		g        Graph[int]
		wantPart []int
		wantOK   bool
	}{
		{"0", NewMutable[int](0), []int{}, true},
		{"1", NewMutable[int](1), []int{0}, true},
		{"2", g1, []int{}, false},
		{"3", NewMutable[int](2), []int{0, 1}, true},
		{"4", g2, []int{0}, true},
		{"5", g3, []int{}, false},
		{"6", g4, []int{0, 4}, true},
		{"7", g5, []int{}, false},
		{"8", g6, []int{0, 2, 4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			part, ok := Bipartition(tt.g)
			if diff := cmp.Diff(part, tt.wantPart); diff != "" {
				t.Errorf("Bipartition: %s", diff)
			}
			if diff := cmp.Diff(ok, tt.wantOK); diff != "" {
				t.Errorf("Bipartition: %s", diff)
			}
		})
	}
}
