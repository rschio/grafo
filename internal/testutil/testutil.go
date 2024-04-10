package testutil

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func ParseGraph[T any](r io.Reader) (*Multigraph[T], error) {
	sc := bufio.NewScanner(r)

	n, err := parseVerticeCount(sc)
	if err != nil {
		return nil, err
	}

	parseWeight := weightFunc[T]()

	g := NewMultigraph[T](n)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		v, w, wt, err := parseEdge(line, parseWeight)
		if err != nil {
			return nil, err
		}
		g.Add(v, w, wt)
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return g, nil
}

func weightFunc[T any]() func(string) (T, error) {
	switch reflect.TypeFor[T]().Kind() {
	case reflect.Int8:
		intInf := int64(InfFor[int8]())
		return parseInt[T](intInf, 8)
	case reflect.Int16:
		intInf := int64(InfFor[int16]())
		return parseInt[T](intInf, 16)
	case reflect.Int32:
		intInf := int64(InfFor[int32]())
		return parseInt[T](intInf, 32)
	case reflect.Int64:
		intInf := InfFor[int64]()
		return parseInt[T](intInf, 64)
	case reflect.Int:
		intInf := int64(InfFor[int]())
		return parseInt[T](intInf, 0)

	case reflect.Uint8:
		uintInf := uint64(InfFor[uint8]())
		return parseUint[T](uintInf, 8)
	case reflect.Uint16:
		uintInf := uint64(InfFor[uint16]())
		return parseUint[T](uintInf, 16)
	case reflect.Uint32:
		uintInf := uint64(InfFor[uint32]())
		return parseUint[T](uintInf, 32)
	case reflect.Uint64:
		uintInf := uint64(InfFor[uint64]())
		return parseUint[T](uintInf, 64)
	case reflect.Uint:
		uintInf := uint64(InfFor[uint]())
		return parseUint[T](uintInf, 0)
	case reflect.Uintptr:
		uintInf := uint64(InfFor[uintptr]())
		return parseUint[T](uintInf, 0)

	case reflect.Float32:
		return parseFloat[T](32)
	case reflect.Float64:
		return parseFloat[T](64)

	case reflect.Struct:
		if reflect.ValueOf(*new(T)).NumField() != 0 {
			break
		}

		return func(_ string) (T, error) {
			v := struct{}{}
			typ := reflect.TypeFor[T]()
			val := reflect.ValueOf(v).Convert(typ)
			return val.Interface().(T), nil
		}
	}

	return func(string) (T, error) {
		return *new(T), fmt.Errorf("parseWeight not implemented")
	}

}

func parseFloat[T any](bitSize int) func(s string) (T, error) {
	return func(s string) (T, error) {
		var v float64
		if s == "inf" || s == "+inf" {
			v = math.Inf(1)
		} else if s == "-inf" {
			v = math.Inf(-1)
		} else {
			var err error
			v, err = strconv.ParseFloat(s, bitSize)
			if err != nil {
				return *new(T), err
			}
		}
		typ := reflect.TypeFor[T]()
		val := reflect.ValueOf(v).Convert(typ)
		return val.Interface().(T), nil
	}
}

func parseUint[T any](uintInf uint64, bitSize int) func(s string) (T, error) {
	return func(s string) (T, error) {
		var v uint64
		if s == "inf" {
			v = uintInf
		} else {
			var err error
			v, err = strconv.ParseUint(s, 10, bitSize)
			if err != nil {
				return *new(T), err
			}
			v = min(v, uintInf)
		}
		typ := reflect.TypeFor[T]()
		val := reflect.ValueOf(v).Convert(typ)
		return val.Interface().(T), nil
	}
}

func parseInt[T any](intInf int64, bitSize int) func(s string) (T, error) {
	return func(s string) (T, error) {
		var v int64
		if s == "inf" || s == "+inf" {
			v = intInf
		} else if s == "-inf" {
			v = -intInf - 1
		} else {
			var err error
			v, err = strconv.ParseInt(s, 10, bitSize)
			if err != nil {
				return *new(T), err
			}

			if v >= 0 {
				v = min(v, intInf)
			} else {
				v = max(v, -intInf-1)
			}
		}
		typ := reflect.TypeFor[T]()
		val := reflect.ValueOf(v).Convert(typ)
		return val.Interface().(T), nil
	}
}

func parseEdge[T any](line string, parseWeight func(string) (T, error)) (int, int, T, error) {
	fields := strings.SplitN(line, " ", 3)
	if len(fields) != 3 {
		return 0, 0, *new(T), fmt.Errorf("failed to parse edge %q", line)
	}
	v, err := strconv.Atoi(fields[0])
	if err != nil {
		return 0, 0, *new(T), fmt.Errorf("failed to parse edge %q", line)
	}
	w, err := strconv.Atoi(fields[1])
	if err != nil {
		return 0, 0, *new(T), fmt.Errorf("failed to parse edge %q", line)
	}

	weight, err := parseWeight(fields[2])
	if err != nil {
		return 0, 0, *new(T), fmt.Errorf("failed to parse weight %q: %v", line, err)
	}

	return v, w, weight, nil
}

func parseVerticeCount(sc *bufio.Scanner) (int, error) {
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 2 || fields[0] != "v" {
			return 0, fmt.Errorf("wrong graph format for vertice count, got %q", line)
		}
		n, err := strconv.Atoi(fields[1])
		if err != nil {
			return 0, fmt.Errorf("failed to parse graph vertice count: %v", err)
		}

		return n, nil
	}

	return 0, fmt.Errorf("wrong graph format, not found vertice count")
}
