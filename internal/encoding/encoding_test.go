package encoding_test

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/rschio/grafo/internal/encoding"
	"github.com/rschio/grafo/internal/encoding/simple"
	"github.com/rschio/grafo/internal/multigraph"
)

// TODO: Test and document that Transform do not need to generate an equal file.
// It needs to generate a file that can be decoded as a equivalent Graph.
// The order of the edges im Simple format file for example don't need to be equal.
func TestTransform(t *testing.T) {
	g := multigraph.New[int](5)
	g.Add(0, 1, 1)
	g.Add(0, 1, 3)
	g.Add(1, 2, 2)
	g.Add(2, 1, 500)
	g.Add(3, 0, -10)

	w := new(bytes.Buffer)
	enc := simple.NewEncoder[int](w, strconv.Itoa)
	if err := enc.Encode(g); err != nil {
		t.Fatalf("enc.Encode(g): %v", err)
	}
	encoded := w.Bytes()
	w.Reset()

	dec := simple.NewDecoder[int](bytes.NewReader(encoded), strconv.Atoi)

	if err := encoding.Transform(enc, dec); err != nil {
		t.Fatalf("Transform(enc, dec): %v", err)
	}

	if !bytes.Equal(w.Bytes(), encoded) {
		t.Fatalf("got diferent encodes from simple encode and transform:\n%s\n%s", w.Bytes(), encoded)
	}
}
