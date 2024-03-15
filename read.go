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
