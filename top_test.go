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

func TestTopSort(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		g := NewMutable[struct{}](0)
		order, ok := TopSort(g)
		if diff := cmp.Diff(order, []int{}); diff != "" {
			t.Errorf("TopSort %s", diff)
		}
		if !ok {
			t.Errorf("TopSort should return OK")
		}
	})

	t.Run("1", func(t *testing.T) {
		g := NewMutable[struct{}](4)
		g.Add(3, 2, struct{}{})
		g.Add(2, 1, struct{}{})
		g.Add(1, 0, struct{}{})
		expOrder := []int{3, 2, 1, 0}
		order, ok := TopSort(g)
		if diff := cmp.Diff(order, expOrder); diff != "" {
			t.Errorf("TopSort %s", diff)
		}
		if !ok {
			t.Errorf("TopSort should return OK")
		}
	})

	t.Run("2", func(t *testing.T) {
		g := NewMutable[struct{}](5)
		g.Add(0, 1, struct{}{})
		g.Add(1, 2, struct{}{})
		g.Add(1, 3, struct{}{})
		g.Add(2, 4, struct{}{})
		g.Add(3, 4, struct{}{})
		order, ok := TopSort(g)
		expOrder1 := []int{0, 1, 2, 3, 4}
		expOrder2 := []int{0, 1, 3, 2, 4}
		if !cmp.Equal(order, expOrder1) && !cmp.Equal(order, expOrder2) {
			t.Errorf("TopSort %s", cmp.Diff(order, expOrder1))
		}
		if !ok {
			t.Errorf("TopSort should return OK")
		}
	})
}

func TestAcyclic(t *testing.T) {
	g := NewMutable[struct{}](0)
	if diff := cmp.Diff(Acyclic(g), true); diff != "" {
		t.Errorf("Acyclic %s", diff)
	}

	g = NewMutable[struct{}](1)
	if diff := cmp.Diff(Acyclic(g), true); diff != "" {
		t.Errorf("Acyclic %s", diff)
	}
	g.Add(0, 0, struct{}{})
	if diff := cmp.Diff(Acyclic(g), false); diff != "" {
		t.Errorf("Acyclic %s", diff)
	}

	g = NewMutable[struct{}](4)
	g.Add(0, 1, struct{}{})
	g.Add(2, 1, struct{}{})
	if diff := cmp.Diff(Acyclic(g), true); diff != "" {
		t.Errorf("Acyclic %s", diff)
	}
	g.Add(0, 2, struct{}{})
	if diff := cmp.Diff(Acyclic(g), true); diff != "" {
		t.Errorf("Acyclic %s", diff)
	}
	g.Add(3, 0, struct{}{})
	if diff := cmp.Diff(Acyclic(g), true); diff != "" {
		t.Errorf("Acyclic %s", diff)
	}
	g.Add(1, 3, struct{}{})
	if diff := cmp.Diff(Acyclic(g), false); diff != "" {
		t.Errorf("Acyclic %s", diff)
	}
}

func BenchmarkAcyclic(b *testing.B) {
	n := 1000
	b.StopTimer()
	g := NewMutable[struct{}](n)
	for i := 0; i < 2*n; i++ {
		v, w := rand.IntN(n), rand.IntN(n)
		if v < w {
			g.AddBoth(v, w, struct{}{})
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = Acyclic(g)
	}
}

func BenchmarkTopSort(b *testing.B) {
	n := 1000
	b.StopTimer()
	g := NewMutable[struct{}](n)
	for i := 0; i < 2*n; i++ {
		v, w := rand.IntN(n), rand.IntN(n)
		if v < w {
			g.AddBoth(v, w, struct{}{})
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TopSort(g)
	}
}
