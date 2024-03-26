package grafo

import (
	"bufio"
	"bytes"
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
			_, err := fmt.Fprintf(w, "%d,%d,%v\n", v, ww, weight)
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
