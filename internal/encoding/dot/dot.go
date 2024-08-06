package dot

import (
	"fmt"
	"io"

	"github.com/rschio/grafo/internal"
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
	if _, err := fmt.Fprintln(e.w, "digraph {"); err != nil {
		return err
	}
	for v := range g.Order() {
		for ww, wt := range g.EdgesFrom(v) {
			_, err := fmt.Fprintf(e.w, "\t%d -> %d [weight=%s]\n", v, ww, e.fmtWeight(wt))
			if err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fprintln(e.w, "}"); err != nil {
		return err
	}
	return nil
}
