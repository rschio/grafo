package grafo

import (
	"fmt"
	"testing"
)

func TestBFS(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		g := NewMutable[struct{}](5)
		wt := struct{}{}
		g.Add(0, 1, wt)
		g.Add(0, 2, wt)
		g.Add(0, 3, wt)
		g.Add(2, 4, wt)

		visited := make([]int, 0)
		for e := range BFS(g, 0) {
			visited = append(visited, e.W)
		}

		before := []int{1, 2, 3}
		after := []int{4}
		if err := visitedOrder(visited, before, after); err != nil {
			t.Error(err)
		}
	})
}

func visitedOrder(visited []int, before []int, after []int) error {
	a := make(map[int]struct{})
	b := make(map[int]struct{})

	for _, v := range before {
		b[v] = struct{}{}
	}
	for _, v := range after {
		a[v] = struct{}{}
	}

	for _, v := range visited {
		if len(b) == 0 {
			break
		}
		if _, ok := a[v]; ok {
			return fmt.Errorf("visited in wrong order:\nvisited %v\nwant before %v\nwant after%v",
				visited, before, after)
		}
		delete(b, v)
	}

	return nil
}
