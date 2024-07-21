package grafo

import (
	"bytes"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/rschio/grafo/internal/testutil"
	"golang.org/x/tools/txtar"
)

func TestMaxFlowTxtar(t *testing.T) {
	archive, err := txtar.ParseFile(filepath.Join("testdata", "maxflow.txtar"))
	if err != nil {
		t.Fatal(err)
	}
	files := archive.Files

	for i := 0; i+1 < len(files); i += 2 {
		t.Run(files[i].Name, func(t *testing.T) {
			g := testutil.ReadGraph(t, bytes.NewReader(files[i].Data), strconv.Atoi)
			source, target, answer := readTxtarAnswer(t, files[i+1])

			flow := MaxFlow(g, source, target)
			if flow != answer {
				t.Fatalf("got %v flow want %v", flow, answer)
			}
		})
	}
}

func readTxtarAnswer(t testing.TB, f txtar.File) (source, target, answer int) {
	str := string(bytes.TrimSpace(f.Data))
	vals := strings.Split(str, " ")
	if len(vals) != 3 {
		t.Fatal("got invalid answer file, expected \"s, t, answer\" format")
	}

	v0, err := strconv.Atoi(vals[0])
	if err != nil {
		t.Fatalf("failed to read s: %v", err)
	}
	v1, err := strconv.Atoi(vals[1])
	if err != nil {
		t.Fatalf("failed to read t: %v", err)
	}
	v2, err := strconv.Atoi(vals[2])
	if err != nil {
		t.Fatalf("failed to read answer: %v", err)
	}

	return v0, v1, v2
}
