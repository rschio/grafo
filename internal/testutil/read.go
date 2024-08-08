package testutil

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/rschio/grafo/internal/encoding/simple"
	"github.com/rschio/grafo/internal/multigraph"
)

func ReadFile[T any](
	file string,
	parseWeight func(string) (T, error),
) (*multigraph.Multigraph[T], error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readGraph(f, parseWeight)
}

func ReadGraph[T any](
	t testing.TB,
	r io.Reader,
	parseWeight func(string) (T, error),
) *multigraph.Multigraph[T] {
	t.Helper()

	g, err := readGraph(r, parseWeight)
	if err != nil {
		t.Fatal(err)
	}

	return g
}

func readGraph[T any](r io.Reader, parseWeight func(string) (T, error),
) (*multigraph.Multigraph[T], error) {
	g, err := simple.NewDecoder(r, parseWeight).Decode()
	if err != nil {
		return nil, err
	}
	mg, ok := g.(*multigraph.Multigraph[T])
	if !ok {
		return nil, fmt.Errorf("failed to case graph decoded from simple to *Multigraph")
	}
	return mg, nil
}
