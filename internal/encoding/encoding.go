package encoding

import (
	"fmt"

	"github.com/rschio/grafo/internal"
)

type Encoder[T any] interface {
	Encode(internal.Graph[T]) error
}

type Decoder[T any] interface {
	Decode() (internal.Graph[T], error)
}

func Transform[T any](to Encoder[T], from Decoder[T]) error {
	g, err := from.Decode()
	if err != nil {
		return fmt.Errorf("decoding: %w", err)
	}
	if err := to.Encode(g); err != nil {
		return fmt.Errorf("encoding: %w", err)
	}
	return nil
}
