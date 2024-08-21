package simple

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/rschio/grafo/internal"
	"github.com/rschio/grafo/internal/multigraph"
)

type Encoder[T any] struct {
	w         io.Writer
	fmtWeight func(T) string
}

func NewEncoder[T any](w io.Writer, fmtWeight func(T) string) *Encoder[T] {
	return &Encoder[T]{
		w:         w,
		fmtWeight: fmtWeight,
	}
}

func (e *Encoder[T]) Encode(g internal.Graph[T]) error {
	_, err := fmt.Fprintln(e.w, g.Order())
	if err != nil {
		return err
	}
	for v := range g.Order() {
		for w, wt := range g.EdgesFrom(v) {
			_, err := fmt.Fprintf(e.w, "%d %d %s\n", v, w, e.fmtWeight(wt))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

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
	sc := bufio.NewScanner(d.r)

	n, err := readNumberOfVertices(sc)
	if err != nil {
		return nil, err
	}

	g := multigraph.New[T](n)
	sep := []byte(" ")
	for sc.Scan() {
		line := bytes.TrimSpace(sc.Bytes())
		if len(line) == 0 {
			continue
		}
		parts := bytes.Split(line, sep)
		if len(parts) != 3 {
			return nil, fmt.Errorf("got %d elements in one line, want 3", len(parts))
		}
		vv := strings.TrimSpace(string(parts[0]))
		ww := strings.TrimSpace(string(parts[1]))
		wt := strings.TrimSpace(string(parts[2]))
		v, err := strconv.Atoi(vv)
		if err != nil {
			return nil, err
		}
		w, err := strconv.Atoi(ww)
		if err != nil {
			return nil, err
		}
		weight, err := d.parseWeight(wt)
		if err != nil {
			return nil, err
		}
		g.Add(v, w, weight)
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return g, nil
}

func readNumberOfVertices(sc *bufio.Scanner) (int, error) {
	if ok := sc.Scan(); !ok {
		return 0, fmt.Errorf("failed to read number of vertices: %w", sc.Err())
	}
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		return 0, fmt.Errorf("failed to read number of vertices: %w", err)
	}
	return n, nil
}
