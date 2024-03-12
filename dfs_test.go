package grafo

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func TestDFS(t *testing.T) {
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
		fmt.Println(visited)

		// Visited all vertices, execept the starting one.
		sort.Ints(visited)
		want := []int{1, 2, 3, 4, 5}
		if slices.Compare(visited, want) != 0 {
			t.Error("not all vertices were visited")
		}
	})
}
