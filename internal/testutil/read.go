package testutil

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/rschio/grafo/internal/multigraph"
)

func ReadGraph(t testing.TB, r io.Reader) *multigraph.Multigraph[int] {
	t.Helper()

	g, err := readGraph(r)
	if err != nil {
		t.Fatal(err)
	}

	return g
}

func readGraph(r io.Reader) (*multigraph.Multigraph[int], error) {
	sc := bufio.NewScanner(r)

	n, err := readNumberOfVertices(sc)
	if err != nil {
		return nil, err
	}

	g := multigraph.New[int](n)
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
