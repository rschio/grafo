package grafo

import (
	"bytes"
	"cmp"
	"fmt"
	"maps"
	"slices"
)

// String returns a description of g with two elements:
// the number of vertices, followed by a sorted list of all edges.
func String[T IntegerOrFloat](g Graph[T]) string {
	n := g.Order()
	// This may be a multigraph, so we look for duplicates by counting.
	count := make(map[Edge[T]]int)
	for v := range n {
		for w, weight := range g.EdgesFrom(v) {
			count[Edge[T]{V: v, W: w, Weight: weight}]++
		}
	}
	// Sort lexicographically on (v, w, c).
	edges := slices.SortedFunc(maps.Keys(count), func(a, b Edge[T]) int {
		v := a.V == b.V
		w := a.W == b.W
		switch {
		case v && w:
			return cmp.Compare(a.Weight, b.Weight)
		case v:
			return cmp.Compare(a.W, b.W)
		default:
			return cmp.Compare(a.V, b.V)
		}
	})
	// Build the string.
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%d [", n)
	for _, e := range edges {
		c := count[e]
		if e.V < e.W {
			// Collect edges in opposite directions into an undirected edge.
			back := Edge[T]{e.W, e.V, e.Weight}
			m := min(c, count[back])
			count[back] -= m
			appendEdge(buf, e, m, true)
			appendEdge(buf, e, c-m, false)
		} else {
			appendEdge(buf, e, c, false)
		}
	}
	if len(edges) > 0 {
		buf.Truncate(buf.Len() - 1) // Remove trailing ' '.
	}
	buf.WriteByte(']')
	return buf.String()
}

func appendEdge[T IntegerOrFloat](buf *bytes.Buffer, e Edge[T], count int, bi bool) {
	if count <= 0 {
		return
	}
	if count > 1 {
		fmt.Fprintf(buf, "%dx", count)
	}

	start, end := '(', ')'
	if bi {
		start, end = '{', '}'
	}
	fmt.Fprintf(buf, "%c%d %d%c", start, e.V, e.W, end)

	if e.Weight != 0 {
		fmt.Fprint(buf, ":")
		switch e.Weight {
		case InfFor[T]():
			fmt.Fprintf(buf, "max")
		// TODO: Switch Min?
		default:
			fmt.Fprint(buf, e.Weight)
		}
	}
	fmt.Fprint(buf, " ")
}
