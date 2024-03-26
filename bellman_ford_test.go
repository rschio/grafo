package grafo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBellmanFord(t *testing.T) {
	gC := NewMutable[int](5)
	gC.Add(0, 1, 10)
	gC.Add(1, 2, 10)
	gC.Add(2, 3, 10)
	gC.Add(3, 4, 10)

	gD := NewMutable[int](6)
	gD.Add(0, 1, 20)
	gD.Add(1, 2, 30)
	gD.Add(0, 3, 9)
	gD.Add(3, 1, 9)
	gD.Add(0, 4, 5)
	gD.Add(4, 5, 5)
	gD.Add(5, 1, 5)

	gE := NewMutable[int](5)
	gE.Add(0, 1, -1)
	gE.Add(1, 0, -1)
	gE.Add(1, 2, -1)
	gE.Add(2, 1, -1)
	gE.Add(2, 3, -1)
	gE.Add(3, 2, -1)
	gE.Add(3, 4, -1)
	gE.Add(4, 3, -1)

	gG := NewMutable[int](6)
	gG.Add(0, 1, 41)
	gG.Add(0, 5, 29)
	gG.Add(1, 2, 51)
	gG.Add(1, 4, 32)
	gG.Add(2, 3, 50)
	gG.Add(3, 0, 45)
	gG.Add(3, 5, -38)
	gG.Add(4, 2, 32)
	gG.Add(4, 3, 36)
	gG.Add(5, 1, -29)
	gG.Add(5, 4, 21)

	tests := []struct {
		name       string
		v          int
		g          Graph[int]
		wantParent []int
		wantDist   []int
		wantOK     bool
	}{
		{"Exemplo C", 0, gC, []int{-1, 0, 1, 2, 3}, []int{0, 10, 20, 30, 40}, true},
		{"Exemplo D", 0, gD, []int{-1, 5, 1, 0, 0, 4}, []int{0, 15, 45, 9, 5, 10}, true},
		{"Exemplo E", 0, gE, []int{1, 0, 1, 2, 3}, []int{-4, -5, -4, -5, -4}, false},
		{"Exemplo G", 4, gG, []int{3, 5, 1, 4, -1, 3}, []int{81, -31, 20, 36, 0, -2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parent, dist, ok := BellmanFord(tt.g, tt.v)
			if ok != tt.wantOK {
				t.Errorf("should return ok = %v", tt.wantOK)
			}
			if diff := cmp.Diff(parent, tt.wantParent); diff != "" {
				t.Errorf("wrong parent: %v", diff)
			}
			if diff := cmp.Diff(dist, tt.wantDist); diff != "" {
				t.Errorf("wrong dist: %v", diff)
			}
		})
	}
}
