package gr

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/rschio/grafo/internal"
	"github.com/rschio/grafo/internal/multigraph"
)

type Decoder[T any] struct {
	r           io.Reader
	parseWeight func(string) (T, error)
}

func NewDecoder[T any](r io.Reader, parseWeight func(string) (T, error)) *Decoder[T] {
	return &Decoder[T]{
		r:           r,
		parseWeight: parseWeight,
	}
}

func (d *Decoder[T]) Decode() (internal.Graph[T], error) {
	var V int
	sc := bufio.NewScanner(d.r)
	for sc.Scan() {
		line := sc.Bytes()
		// Skip empty lines and comments.
		if len(line) == 0 || line[0] == 'c' {
			continue
		}
		if len(line) >= 5 && string(line[:5]) == "p sp " {
			var err error
			V, _, err = grVerticesAndEdges(line)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	g := multigraph.New[T](V)
	for sc.Scan() {
		line := sc.Bytes()
		// Skip empty lines and comments.
		if len(line) == 0 || line[0] == 'c' {
			continue
		}
		if line[0] != 'a' {
			return nil, fmt.Errorf("wrong format: %s", line)
		}
		fields := bytes.Fields(line[1:])
		if len(fields) != 3 {
			return nil, fmt.Errorf("got %d elements in one line, want 3", len(fields))
		}

		v, err := strconv.Atoi(string(fields[0]))
		if err != nil {
			return nil, err
		}
		w, err := strconv.Atoi(string(fields[1]))
		if err != nil {
			return nil, err
		}
		weight, err := d.parseWeight(string(fields[2]))
		if err != nil {
			return nil, err
		}

		if v > g.Order() {
			return nil, fmt.Errorf("vertex %d out of valid range", v)
		}
		if w > g.Order() {
			return nil, fmt.Errorf("vertex %d out of valid range", w)
		}

		// The gr format use 1 idexed vertices.
		// We use 0 indexed.
		v = v - 1
		w = w - 1

		g.Add(v, w, weight)
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return g, nil
}

func grVerticesAndEdges(line []byte) (int, int, error) {
	line = line[len("p sp "):]
	fields := bytes.Fields(line)
	if len(fields) != 2 {
		return 0, 0, errors.New("failed to get vertices and edges")
	}
	v, err := strconv.Atoi(string(fields[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert vertices: %w", err)
	}
	e, err := strconv.Atoi(string(fields[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert edges: %w", err)
	}
	return v, e, nil
}
