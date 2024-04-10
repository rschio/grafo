package testutil

import (
	"math"
	"strings"
	"testing"
)

func TestParseGraph(t *testing.T) {
	var file = `
	v 3
	0 2 inf`

	t.Run("int8", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[int8](r)
		if err != nil {
			t.Fatalf("ParseGraph[int](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		inf := InfFor[int8]()
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != inf {
				t.Errorf("got %d %d want %d %d", w, weight, 2, inf)
			}
		}
	})
	t.Run("int16", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[int16](r)
		if err != nil {
			t.Fatalf("ParseGraph[int](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		inf := InfFor[int16]()
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != inf {
				t.Errorf("got %d %d want %d %d", w, weight, 2, inf)
			}
		}
	})
	t.Run("int32", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[int32](r)
		if err != nil {
			t.Fatalf("ParseGraph[int](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		inf := InfFor[int32]()
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != inf {
				t.Errorf("got %d %d want %d %d", w, weight, 2, inf)
			}
		}
	})
	t.Run("int64", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[int64](r)
		if err != nil {
			t.Fatalf("ParseGraph[int](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		inf := InfFor[int64]()
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != inf {
				t.Errorf("got %d %d want %d %d", w, weight, 2, inf)
			}
		}
	})
	t.Run("int", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[int](r)
		if err != nil {
			t.Fatalf("ParseGraph[int](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		inf := InfFor[int]()
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != inf {
				t.Errorf("got %d %d want %d %d", w, weight, 2, inf)
			}
		}
	})

	t.Run("uint8", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[uint8](r)
		if err != nil {
			t.Fatalf("ParseGraph[uint](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		inf := InfFor[uint8]()
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != inf {
				t.Errorf("got %d %d want %d %d", w, weight, 2, inf)
			}
		}
	})
	t.Run("uint16", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[uint16](r)
		if err != nil {
			t.Fatalf("ParseGraph[uint16](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		inf := InfFor[uint16]()
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != inf {
				t.Errorf("got %d %d want %d %d", w, weight, 2, inf)
			}
		}
	})
	t.Run("uint32", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[uint32](r)
		if err != nil {
			t.Fatalf("ParseGraph[uint32](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		inf := InfFor[uint32]()
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != inf {
				t.Errorf("got %d %d want %d %d", w, weight, 2, inf)
			}
		}
	})
	t.Run("uint64", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[uint64](r)
		if err != nil {
			t.Fatalf("ParseGraph[uint64](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		inf := InfFor[uint64]()
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != inf {
				t.Errorf("got %d %d want %d %d", w, weight, 2, inf)
			}
		}
	})
	t.Run("uint", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[uint](r)
		if err != nil {
			t.Fatalf("ParseGraph[uint](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		inf := InfFor[uint]()
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != inf {
				t.Errorf("got %d %d want %d %d", w, weight, 2, inf)
			}
		}
	})
	t.Run("uintptr", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[uintptr](r)
		if err != nil {
			t.Fatalf("ParseGraph[uintptr](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		inf := InfFor[uintptr]()
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != inf {
				t.Errorf("got %d %d want %d %d", w, weight, 2, inf)
			}
		}
	})

	t.Run("float32", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[float32](r)
		if err != nil {
			t.Fatalf("ParseGraph[float32](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		inf := InfFor[float32]()
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != inf {
				t.Errorf("got %d %f want %d %f", w, weight, 2, inf)
			}
		}
	})

	t.Run("float64", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[float64](r)
		if err != nil {
			t.Fatalf("ParseGraph[float64](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		inf := InfFor[float64]()
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != inf {
				t.Errorf("got %d %f want %d %f", w, weight, 2, inf)
			}
		}
	})

	t.Run("struct{}", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[struct{}](r)
		if err != nil {
			t.Fatalf("ParseGraph[struct{}](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != struct{}{} {
				t.Errorf("got %d %v want %d %v", w, weight, 2, struct{}{})
			}
		}
	})
}

func TestParseGraphNegativeInf(t *testing.T) {
	var file = `
	v 3
	0 2 -inf`

	t.Run("int32", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[int32](r)
		if err != nil {
			t.Fatalf("ParseGraph[int32](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		negInf := int32(math.MinInt32)
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != negInf {
				t.Errorf("got %d %d want %d %d", w, weight, 2, negInf)
			}
		}
	})
	t.Run("float64", func(t *testing.T) {
		r := strings.NewReader(file)
		g, err := ParseGraph[float64](r)
		if err != nil {
			t.Fatalf("ParseGraph[float64](r): %v", err)
		}

		if g.Order() != 3 {
			t.Errorf("g.Order() = %d, want %d", g.Order(), 3)
		}

		negInf := math.Inf(-1)
		for w, weight := range g.EdgesFrom(0) {
			if w != 2 || weight != negInf {
				t.Errorf("got %d %f want %d %f", w, weight, 2, negInf)
			}
		}
	})
}
