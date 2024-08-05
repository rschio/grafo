package grafo

import (
	"bytes"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rschio/grafo/internal/testutil"
	"golang.org/x/tools/txtar"
)

func TestMaxFlow(t *testing.T) {
	archive, err := txtar.ParseFile(filepath.Join("testdata", "maxflow.txtar"))
	if err != nil {
		t.Fatal(err)
	}
	files := archive.Files

	for i := 0; i+1 < len(files); i += 2 {
		t.Run(files[i].Name, func(t *testing.T) {
			g := testutil.ReadGraph(t, bytes.NewReader(files[i].Data), strconv.Atoi)
			source, target, wantFlow, wantGraph := readTxtarAnswer(t, files[i+1])

			flow, res := MaxFlow(g, source, target)
			if flow != wantFlow {
				t.Errorf("got %v flow want %v", flow, wantFlow)
			}

			if wantGraph == "-" { // Skip graph string test.
				return
			}
			if diff := cmp.Diff(String(res), wantGraph); diff != "" {
				t.Errorf("MaxFlow(%d, %d) = %s", source, target, diff)
			}
		})
	}
}

func readTxtarAnswer(t testing.TB, f txtar.File) (source, target, flow int, graph string) {
	str := string(bytes.TrimSpace(f.Data))
	lines := strings.Split(str, "\n")
	if len(lines) != 2 {
		t.Fatalf("invalid answer file, expected 2 lines got %d", len(lines))
	}

	vals := strings.Split(lines[0], " ")
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
		if vals[2] == "max" {
			v2 = InfFor[int]()
		} else {
			t.Fatalf("failed to read answer: %v", err)
		}
	}

	return v0, v1, v2, lines[1]
}
