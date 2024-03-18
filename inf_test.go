package grafo

import (
	"math"
	"testing"
)

func Test_infFor(t *testing.T) {
	if got := InfFor[uintptr](); got != math.MaxUint {
		t.Errorf("uintptr: got %v want %v", got, uint(math.MaxUint))
	}
	if got := InfFor[uint](); got != math.MaxUint {
		t.Errorf("uint: got %v want %v", got, uint(math.MaxUint))
	}
	if got := InfFor[uint64](); got != math.MaxUint64 {
		t.Errorf("uint64: got %v want %v", got, uint64(math.MaxUint64))
	}
	if got := InfFor[uint32](); got != math.MaxUint32 {
		t.Errorf("uint32: got %v want %v", got, uint32(math.MaxUint32))
	}
	if got := InfFor[uint16](); got != math.MaxUint16 {
		t.Errorf("uint16: got %v want %v", got, math.MaxUint16)
	}
	if got := InfFor[uint8](); got != math.MaxUint8 {
		t.Errorf("uint8: got %v want %v", got, math.MaxUint8)
	}
	if got := InfFor[int](); got != math.MaxInt {
		t.Errorf("int: got %v want %v", got, math.MaxInt)
	}
	if got := InfFor[int64](); got != math.MaxInt64 {
		t.Errorf("int64: got %v want %v", got, int64(math.MaxInt64))
	}
	if got := InfFor[int8](); got != math.MaxInt8 {
		t.Errorf("int8: got %v want %v", got, math.MaxInt8)
	}
	if got := InfFor[float64](); got != math.Inf(1) {
		t.Errorf("float64: got %v want %v", got, math.Inf(1))
	}
	if got := InfFor[float32](); float64(got) != math.Inf(1) {
		t.Errorf("float32: got %v want %v", got, math.Inf(1))
	}

	type (
		myUint   uint
		myUint64 uint64
		myUint32 uint32
		myUint16 uint16
		myUint8  uint8

		myInt   int
		myInt64 int64
		myInt32 int32
		myInt16 int16
		myInt8  int8

		myFloat32 float32
		myFloat64 float64
	)

	if got := InfFor[myUint](); got != math.MaxUint {
		t.Errorf("myUint64: got %v want %v", got, uint(math.MaxUint))
	}
	if got := InfFor[myUint64](); got != math.MaxUint64 {
		t.Errorf("myUint64: got %v want %v", got, uint64(math.MaxUint64))
	}
	if got := InfFor[myUint32](); got != math.MaxUint32 {
		t.Errorf("myUint32: got %v want %v", got, uint32(math.MaxUint32))
	}
	if got := InfFor[myUint16](); got != math.MaxUint16 {
		t.Errorf("myUint16: got %v want %v", got, math.MaxUint16)
	}
	if got := InfFor[myUint8](); got != math.MaxUint8 {
		t.Errorf("myUint8: got %v want %v", got, math.MaxUint8)
	}

	if got := InfFor[myInt](); got != math.MaxInt {
		t.Errorf("myInt64: got %v want %v", got, math.MaxInt)
	}
	if got := InfFor[myInt64](); got != math.MaxInt64 {
		t.Errorf("myInt64: got %v want %v", got, int64(math.MaxInt64))
	}
	if got := InfFor[myInt32](); got != math.MaxInt32 {
		t.Errorf("myInt32: got %v want %v", got, math.MaxInt32)
	}
	if got := InfFor[myInt16](); got != math.MaxInt16 {
		t.Errorf("myInt16: got %v want %v", got, math.MaxInt16)
	}
	if got := InfFor[myInt8](); got != math.MaxInt8 {
		t.Errorf("myInt8: got %v want %v", got, math.MaxInt8)
	}

	if got := InfFor[myFloat32](); !math.IsInf(float64(got), 1) {
		t.Errorf("myFloat32: got %v want %v", got, math.Inf(1))
	}
	if got := InfFor[myFloat64](); !math.IsInf(float64(got), 1) {
		t.Errorf("myFloat64: got %v want %v", got, math.Inf(1))
	}
}
