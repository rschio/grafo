package grafo

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func DOT[T any](g Graph[T], w io.Writer) error {
	buf := new(bytes.Buffer)

	fmt.Fprintln(buf, "digraph {")
	for v := range g.Order() {
		for ww, wt := range g.EdgesFrom(v) {
			fmt.Fprintf(buf, "\t%d -> %d [weight=%v]\n", v, ww, wt)
		}
	}
	fmt.Fprintln(buf, "}")

	_, err := io.Copy(w, buf)
	return err
}

func Write[T any](g Graph[T], w io.Writer) error {
	n := g.Order()

	_, err := fmt.Fprintf(w, "%d\n", n)
	if err != nil {
		return err
	}

	for v := range n {
		for ww, weight := range g.EdgesFrom(v) {
			_, err := fmt.Fprintf(w, "%d %d %v\n", v, ww, weight)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Read(path string) (*Mutable[int], error) {
	numVertices, _, _ := strings.Cut(filepath.Base(path), "_")
	n, err := strconv.Atoi(numVertices)
	if err != nil {
		return nil, fmt.Errorf("failed to get number of vertices: %w", err)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return read(n, f)
}

func read(n int, r io.Reader) (*Mutable[int], error) {
	g := NewMutable[int](n)

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		parts := bytes.Split(sc.Bytes(), []byte{','})
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
		weight, err := strconv.Atoi(wt)
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

func readGrFile(path string) (*Immutable[int], error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readGr(f)
}

func readGr(r io.Reader) (*Immutable[int], error) {
	var V int
	sc := bufio.NewScanner(r)
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

	g := newMultigraph[int](V)
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
		weight, err := strconv.Atoi(string(fields[2]))
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

	return Sort(g), nil
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
