package grafo

import (
	"fmt"
	"iter"
	"log"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rschio/grafo/internal/multigraph"
	"github.com/rschio/grafo/internal/testutil"
	"golang.org/x/tools/txtar"
)

func TestDFSPostvisit(t *testing.T) {
	g := multigraph.New[struct{}](6)
	wt := struct{}{}
	g.Add(0, 1, wt)
	g.Add(0, 2, wt)
	g.Add(1, 3, wt)
	g.Add(1, 4, wt)
	g.Add(2, 5, wt)

	want := []int{3, 4, 1, 5, 2, 0}
	got := []int{}
	for v := range DFSPostvisit(g, 0) {
		got = append(got, v)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("posvisit diff %s", diff)
	}

	// Test if iterator respects the break.
	for _ = range DFSPostvisit(g, 0) {
		break
	}
}

func TestDFSPrevisit(t *testing.T) {
	g := multigraph.New[struct{}](6)
	wt := struct{}{}
	g.Add(0, 1, wt)
	g.Add(0, 2, wt)
	g.Add(1, 3, wt)
	g.Add(1, 4, wt)
	g.Add(2, 5, wt)

	want := []int{0, 1, 3, 4, 2, 5}
	got := []int{}
	for v := range DFSPrevisit(g, 0) {
		got = append(got, v)
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("posvisit diff %s", diff)
	}

	// Test if iterator respects the break.
	for _ = range DFSPrevisit(g, 0) {
		break
	}

	t.Run("vertex out of range +", func(t *testing.T) {
		defer func() {
			recover()
		}()

		g := multigraph.New[struct{}](2)
		wt := struct{}{}
		g.Add(0, 1, wt)

		for _ = range DFSPrevisit(g, 2) {
			t.Fatal("should panic before this")
		}
	})
	t.Run("vertex out of range -", func(t *testing.T) {
		defer func() {
			recover()
		}()

		g := multigraph.New[struct{}](2)
		wt := struct{}{}
		g.Add(0, 1, wt)

		for _ = range DFSPrevisit(g, -1) {
			t.Fatal("should panic before this")
		}
	})
}

func TestDFS(t *testing.T) {
	t.Run("break", func(t *testing.T) {
		g := NewMutable[struct{}](6)
		wt := struct{}{}
		g.Add(0, 1, wt)
		g.Add(0, 2, wt)

		// Test if iterator respects the break.
		for _ = range DFS(g, 0) {
			break
		}
	})

	t.Run("test1", func(t *testing.T) {
		g := NewMutable[struct{}](6)
		wt := struct{}{}
		g.Add(0, 1, wt)
		g.Add(0, 2, wt)
		g.Add(2, 3, wt)
		g.Add(3, 4, wt)
		g.Add(4, 5, wt)

		visited := make([]int, 0)

		for e := range DFS(g, 0) {
			visited = append(visited, e.W)
		}

		path1 := []int{1}
		path2 := []int{2, 3, 4, 5}
		// The algorithm is not deterministic, but one
		// of the options should occur.
		err1 := visitedOrder(visited, path1, path2)
		err2 := visitedOrder(visited, path2, path1)
		if err1 != nil && err2 != nil {
			t.Error(err1)
			t.Error(err2)
		}

		// Visited all vertices, execept the starting one.
		sort.Ints(visited)
		want := []int{1, 2, 3, 4, 5}
		if slices.Compare(visited, want) != 0 {
			t.Error("not all vertices were visited")
		}
	})
}

func TestDFSPossibleTrees(t *testing.T) {
	file := filepath.Join("testdata", "7_dfs_graph")
	g, err := testutil.ReadFile(file, strconv.Atoi)
	if err != nil {
		t.Fatal(err)
	}

	possiblePaths := [][]int{
		{0, 1, 4, 5, 3, 2, 6},
		{0, 1, 4, 5, 2, 6, 3},
		{0, 1, 3, 4, 5, 2, 6},
		{0, 1, 3, 2, 6, 4, 5},
		{0, 1, 2, 6, 3, 4, 5},
		{0, 1, 2, 6, 4, 5, 3},
	}

	path := make([]int, 0)
	path = append(path, 0)
	for e := range DFS(g, 0) {
		path = append(path, e.W)
	}

	allDiferent := true
	for _, p := range possiblePaths {
		if slices.Equal(path, p) {
			allDiferent = false
			break
		}
	}

	if allDiferent {
		t.Errorf("the path is diferent from all possible paths.\ngot %v\nwant%v",
			path, possiblePaths)
	}
}

func TestDFSExhaustion(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	archive, err := txtar.ParseFile(filepath.Join("testdata", "exhaust5.txtar"))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range archive.Files {
		for v := range 5 {
			func() {
				g, err := graphFromFile(f.Data)
				if err != nil {
					log.Fatal(err)
				}

				next1, stop1 := iter.Pull(DFS(g, v))
				defer stop1()

				next2, stop2 := iter.Pull(dfsRec(g, v))
				defer stop2()

				path := make([]Edge[int], 0)
				for {
					e1, ok1 := next1()
					e2, ok2 := next2()

					path = append(path, e1)
					if diff := cmp.Diff(e1, e2); diff != "" || ok1 != ok2 {
						t.Errorf("%s_%d: ok1 %v ok2 %v diff: %s\npath: %v", f.Name, v, ok1, ok2, diff, path)
						return
					}
					if ok1 == false {
						return
					}
				}
			}()
		}
	}
}

func dfsRec[T any](g Graph[T], v int) iter.Seq[Edge[T]] {
	return func(yield func(e Edge[T]) bool) {
		visited := make([]bool, g.Order())
		visited[v] = true
		dfsR(g, visited, yield, v)
	}
}

func dfsR[T any](g Graph[T], visited []bool, yield func(e Edge[T]) bool, v int) {
	for w, wt := range g.EdgesFrom(v) {
		if visited[w] {
			continue
		}
		visited[w] = true
		if !yield(Edge[T]{v, w, wt}) {
			return
		}
		dfsR(g, visited, yield, w)
	}
}

func graphFromFile(data []byte) (*multigraph.Multigraph[int], error) {
	lines := strings.Split(string(data), "\n")
	first := true
	allNums := make([]int, 0)
	V := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if first {
			nums, err := parseLine(line)
			if err != nil {
				return nil, err
			}
			if len(nums) != 2 {
				return nil, fmt.Errorf("should return V and E got %d values", len(nums))
			}
			V = nums[0]
			first = false
			continue
		}

		nums, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		allNums = append(allNums, nums...)
	}
	if len(allNums)%2 != 0 {
		return nil, fmt.Errorf("malformed file, odd number of vertices doesn't form edges")
	}

	g := multigraph.New[int](V)
	for i := 0; i < len(allNums); i = i + 2 {
		g.Add(allNums[i], allNums[i+1], 0)
	}

	return g, nil
}

func parseLine(line string) ([]int, error) {
	fields := strings.Fields(line)
	numbers := make([]int, len(fields))
	for i, field := range fields {
		number, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		numbers[i] = number
	}
	return numbers, nil
}
