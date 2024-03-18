package grafo

// ShortestPath computes a shortest path from v to w.
// Only edges with non-negative costs are included.
// The number dist is the length of the path, or inf if w cannot be reached.
// (inf is +inf for floats and the maximum value for integers).
//
// The time complexity is O((|E| + |V|)⋅log|V|), where |E| is the number of edges
// and |V| the number of vertices in the graph.
func ShortestPath[T IntegerOrFloat](g Graph[T], v, w int) (path []int, dist T) {
	parent, distances := shortestPath(g, v, w)
	path, dist = []int{}, distances[w]
	if dist == InfFor[T]() {
		return
	}
	for v := w; v != -1; v = parent[v] {
		path = append(path, v)
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return
}

// ShortestPaths computes the shortest paths from v to all other vertices.
// Only edges with non-negative costs are included.
// The number parent[w] is the predecessor of w on a shortest path from v to w,
// or -1 if none exists.
// The number dist[w] equals the length of a shortest path from v to w,
// or is inf if w cannot be reached.
// (inf is +inf for floats and the maximum value for integers).
//
// The time complexity is O((|E| + |V|)⋅log|V|), where |E| is the number of edges
// and |V| the number of vertices in the graph.
func ShortestPaths[T IntegerOrFloat](g Graph[T], v int) (parent []int, dist []T) {
	// Use -1 to search for a vertex that doesn't exists so it will
	// search for all the shortest paths from v.
	return shortestPath(g, v, -1)
}

func shortestPath[T IntegerOrFloat](g Graph[T], v, w int) (parent []int, dist []T) {
	n := g.Order()
	dist = make([]T, n)
	parent = make([]int, n)
	inf := InfFor[T]()
	for i := range dist {
		dist[i], parent[i] = inf, -1
	}
	dist[v] = 0

	// Dijkstra's algorithm
	Q := emptyPrioQueue(dist)
	Q.Push(v)

	dst := w
	for Q.Len() > 0 {
		v = Q.Pop()
		if v == dst {
			return
		}
		for w, weight := range g.EdgesFrom(v) {
			// Skip negative edges.
			if weight < 0 {
				continue
			}
			alt := dist[v] + weight
			switch {
			case dist[w] == inf:
				dist[w], parent[w] = alt, v
				Q.Push(w)
			case alt < dist[w]:
				dist[w], parent[w] = alt, v
				Q.Fix(w)
			}
		}
	}

	return
}
