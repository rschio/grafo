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

func TestShortestPath(t *testing.T) {
	t.Run("float64", func(t *testing.T) {
		g := NewMutable[float64](6)
		g.Add(0, 1, 1.0)
		g.Add(0, 2, 1.0)
		g.Add(0, 3, 3.0)
		g.Add(1, 3, 0.0)
		g.Add(2, 3, 1.0)
		g.Add(2, 5, 8.0)
		g.Add(3, 5, 7.0)
		g.Add(1, 5, -1)
		parent, dist := ShortestPaths(g, 0)
		expParent := []int{-1, 0, 0, 1, -1, 3}
		expDist := []float64{0, 1, 1, 1, InfFor[float64](), 8}
		if diff := cmp.Diff(parent, expParent); diff != "" {
			t.Errorf("ShortestPaths->parent %s", diff)
		}
		if diff := cmp.Diff(dist, expDist); diff != "" {
			t.Errorf("ShortestPaths->dist %s", diff)
		}
	})

	t.Run("int64", func(t *testing.T) {
		g := NewMutable[int64](6)
		g.Add(0, 1, 1.0)
		g.Add(0, 2, 1.0)
		g.Add(0, 3, 3.0)
		g.Add(1, 3, 0.0)
		g.Add(2, 3, 1.0)
		g.Add(2, 5, 8.0)
		g.Add(3, 5, 7.0)
		g.Add(1, 5, -1)
		parent, dist := ShortestPaths(g, 0)
		expParent := []int{-1, 0, 0, 1, -1, 3}
		expDist := []int64{0, 1, 1, 1, InfFor[int64](), 8}
		if diff := cmp.Diff(parent, expParent); diff != "" {
			t.Errorf("ShortestPaths->parent %s", diff)
		}
		if diff := cmp.Diff(dist, expDist); diff != "" {
			t.Errorf("ShortestPaths->dist %s", diff)
		}
	})

	type myUint8 uint8
	t.Run("myUint8", func(t *testing.T) {
		g := NewMutable[myUint8](6)
		g.Add(0, 1, 1)
		g.Add(0, 2, 1)
		g.Add(0, 3, 3)
		g.Add(1, 3, 0)
		g.Add(2, 3, 1)
		g.Add(2, 5, 8)
		g.Add(3, 5, 7)
		parent, dist := ShortestPaths(g, 0)
		expParent := []int{-1, 0, 0, 1, -1, 3}
		expDist := []myUint8{0, 1, 1, 1, InfFor[myUint8](), 8}
		if diff := cmp.Diff(parent, expParent); diff != "" {
			t.Errorf("ShortestPaths->parent %s", diff)
		}
		if diff := cmp.Diff(dist, expDist); diff != "" {
			t.Errorf("ShortestPaths->dist %s", diff)
		}
	})

	t.Run("uint32 shortest path", func(t *testing.T) {
		g := NewMutable[uint32](6)
		g.Add(0, 1, 1)
		g.Add(0, 2, 1)
		g.Add(0, 3, 3)
		g.Add(1, 3, 0)
		g.Add(2, 3, 1)
		g.Add(2, 5, 8)
		g.Add(3, 5, 7)

		path, d := ShortestPath(g, 0, 5)
		if diff := cmp.Diff(path, []int{0, 1, 3, 5}); diff != "" {
			t.Errorf("ShortestPath->path %s", diff)
		}
		if diff := cmp.Diff(d, uint32(8)); diff != "" {
			t.Errorf("ShortestPath->dist %s", diff)
		}

		path, d = ShortestPath(g, 0, 0)
		if diff := cmp.Diff(path, []int{0}); diff != "" {
			t.Errorf("ShortestPath->path %s", diff)
		}
		if diff := cmp.Diff(d, uint32(0)); diff != "" {
			t.Errorf("ShortestPath->dist %s", diff)
		}

		path, d = ShortestPath(g, 0, 4)
		if diff := cmp.Diff(path, []int{}); diff != "" {
			t.Errorf("ShortestPath->path %s", diff)
		}
		if diff := cmp.Diff(d, InfFor[uint32]()); diff != "" {
			t.Errorf("ShortestPath->dist %s", diff)
		}

	})
}