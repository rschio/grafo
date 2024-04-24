package grafo_test

import (
	"bytes"
	"fmt"
	"iter"

	"github.com/rschio/grafo"
)

// This example is inspired on section 5.7 Graph Reductions: Flood Fill
// of the book https://jeffe.cs.illinois.edu/teaching/algorithms/.
func ExampleBFS() {
	fmt.Println(image)

	v := 78
	// Find all the vertices of the connected region of v.
	// This problem can be reduced to find all vertices reachable from v.
	vertices := []int{v}
	for e := range grafo.BFS(image, v) {
		vertices = append(vertices, e.W)
	}

	// Change the connected region color to green.
	for _, v := range vertices {
		row, col := image.rowCol(v)
		image[row][col] = '🟩'
	}

	fmt.Println(image)
	// Output:
	//
	//⬜⬛⬜⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜
	//⬛⬜⬛⬜⬛⬛⬜⬜⬜⬜⬜⬜⬜
	//⬜⬛⬜⬛⬜⬛⬛⬜⬜⬜⬜⬜⬜
	//⬜⬜⬛⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜
	//⬛⬛⬜⬜⬜⬛⬛⬛⬜⬜⬜⬜⬜
	//⬜⬛⬛⬜⬛⬜⬛⬜⬜⬜⬜⬜⬜
	//⬜⬜⬛⬛⬛⬛⬜⬜⬛⬜⬜⬜⬜
	//⬜⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜
	//⬜⬜⬜⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜
	//⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬛⬛
	//⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬛⬜
	//⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬛⬜⬛
	//⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬛⬜
	//
	//⬜⬛⬜⬜⬛🟩🟩🟩🟩🟩🟩🟩🟩
	//⬛⬜⬛⬜⬛⬛🟩🟩🟩🟩🟩🟩🟩
	//⬜⬛⬜⬛⬜⬛⬛🟩🟩🟩🟩🟩🟩
	//⬜⬜⬛⬜⬜⬜⬛🟩🟩🟩🟩🟩🟩
	//⬛⬛⬜⬜⬜⬛⬛⬛🟩🟩🟩🟩🟩
	//🟩⬛⬛⬜⬛⬜⬛🟩🟩🟩🟩🟩🟩
	//🟩🟩⬛⬛⬛⬛🟩🟩⬛🟩🟩🟩🟩
	//🟩🟩🟩🟩⬛🟩🟩🟩🟩🟩🟩🟩🟩
	//🟩🟩🟩🟩🟩🟩⬛🟩🟩🟩🟩🟩🟩
	//🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩⬛⬛⬛
	//🟩🟩🟩🟩🟩🟩🟩🟩🟩⬛⬜⬛⬜
	//🟩🟩🟩🟩🟩🟩🟩🟩🟩⬛⬛⬜⬛
	//🟩🟩🟩🟩🟩🟩🟩🟩🟩⬛⬜⬛⬜
}

var image graph = [][]rune{
	[]rune("⬜⬛⬜⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜"),
	[]rune("⬛⬜⬛⬜⬛⬛⬜⬜⬜⬜⬜⬜⬜"),
	[]rune("⬜⬛⬜⬛⬜⬛⬛⬜⬜⬜⬜⬜⬜"),
	[]rune("⬜⬜⬛⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜"),
	[]rune("⬛⬛⬜⬜⬜⬛⬛⬛⬜⬜⬜⬜⬜"),
	[]rune("⬜⬛⬛⬜⬛⬜⬛⬜⬜⬜⬜⬜⬜"),
	[]rune("⬜⬜⬛⬛⬛⬛⬜⬜⬛⬜⬜⬜⬜"),
	[]rune("⬜⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜"),
	[]rune("⬜⬜⬜⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜"),
	[]rune("⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬛⬛"),
	[]rune("⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬛⬜"),
	[]rune("⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬛⬜⬛"),
	[]rune("⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬛⬜"),
}

type graph [][]rune

func (g graph) Order() int { return len(g) * len(g[0]) }

// EdgesFrom return the edges from v.
// To express a connected region we define that v can possibly
// have 4 neighbors that are top, bottom, left and right.
// The edge only exists if the vertices are of the same color of v.
func (g graph) EdgesFrom(v int) iter.Seq2[int, struct{}] {
	n := len(g[0])

	neighbors := make([]int, 0, 4)
	top := v - n
	bottom := v + n
	left := v - 1
	right := v + 1

	if top >= 0 {
		neighbors = append(neighbors, top)
	}
	if bottom < g.Order() {
		neighbors = append(neighbors, bottom)
	}

	row := v / n
	if left >= 0 && left/n == row {
		neighbors = append(neighbors, left)
	}
	if right/n == row {
		neighbors = append(neighbors, right)
	}

	return func(yield func(w int, weight struct{}) bool) {
		color := g.colorOf(v)
		for _, neighbor := range neighbors {
			if g.colorOf(neighbor) != color {
				continue
			}
			if !yield(neighbor, struct{}{}) {
				return
			}
		}
	}
}

func (g graph) String() string {
	buf := new(bytes.Buffer)
	for _, line := range g {
		fmt.Fprintln(buf, string(line[:]))
	}
	return buf.String()
}

func (g graph) rowCol(v int) (int, int) {
	n := len(g[0])
	row := v / n
	col := v % n
	return row, col
}

func (g graph) colorOf(v int) rune {
	row, col := g.rowCol(v)
	return g[row][col]
}
