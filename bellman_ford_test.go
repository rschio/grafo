package grafo

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rschio/grafo/internal/testutil"
	"golang.org/x/tools/txtar"
)

func TestBellmanFord(t *testing.T) {
	for name, tt := range readBFTestcases(t, "bellman_ford.txtar") {
		t.Run(name, func(t *testing.T) {
			parent, dist, ok := BellmanFord(tt.g, tt.v)
			if diff := cmp.Diff(ok, tt.wantOK); diff != "" {
				t.Errorf("wrong negative cycle detection %v", diff)
			}
			if diff := cmp.Diff(parent, tt.wantParent); diff != "" {
				t.Errorf("wrong parent: %v", diff)
			}
			if diff := cmp.Diff(dist, tt.wantDist); diff != "" {
				t.Errorf("wrong dist: %v", diff)
			}

		})
	}
}

type bfTestCase struct {
	g          Graph[int]
	v          int
	wantParent []int
	wantDist   []int
	wantOK     bool
}

func readBFTestcases(t *testing.T, fname string) map[string]bfTestCase {
	archive, err := txtar.ParseFile(filepath.Join("testdata", fname))
	if err != nil {
		t.Fatal(err)
	}

	tcs := make(map[string]bfTestCase)

	for _, f := range archive.Files {
		ext := filepath.Ext(f.Name)
		name := f.Name[:len(f.Name)-len(ext)]

		switch ext {
		case ".graph":
			g, err := testutil.ParseGraph[int](bytes.NewReader(f.Data))
			if err != nil {
				t.Fatal(err)
			}
			tc := tcs[name]
			tc.g = g
			tcs[name] = tc

		case ".result":
			lines := bytes.SplitN(f.Data, []byte("\n"), 4)
			if len(lines) != 4 {
				t.Fatalf("failed to parse file %s", f.Name)
			}
			for i := range lines {
				lines[i] = bytes.TrimSpace(lines[i])
			}

			v, err := testutil.ParseValue[int](string(lines[0]))
			if err != nil {
				t.Fatalf("failed to parse file %s: %v", f.Name, err)
			}
			parent, err := testutil.ParseSlice[int](string(lines[1]))
			if err != nil {
				t.Fatalf("failed to parse file %s: %v", f.Name, err)
			}

			dist, err := testutil.ParseSlice[int](string(lines[2]))
			if err != nil {
				t.Fatalf("failed to parse file %s: %v", f.Name, err)
			}

			ok, err := testutil.ParseValue[bool](string(lines[3]))
			if err != nil {
				t.Fatalf("failed to parse file %s: %v", f.Name, err)
			}

			tc := tcs[name]
			tc.v = v
			tc.wantParent = parent
			tc.wantDist = dist
			tc.wantOK = ok
			tcs[name] = tc

		default:
			t.Fatalf("file with invalid ext: %v", f.Name)
		}
	}

	return tcs
}
