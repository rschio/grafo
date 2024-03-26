package grafo

// BellmanFord calculate the shortest path from v to all other vertices.
// The function returns the a slice of parent, a slice of dist containing the dists
// between v and the vertex and return ok if there isn't a negative cycle.
//
// The algorithm skip NaN weighted edges.
func BellmanFord[T IntegerOrFloat](g Graph[T], v int) (parent []int, dist []T, ok bool) {
	n := g.Order()
	inf := InfFor[T]()
	parent = make([]int, n)
	dist = make([]T, n)
	onQueue := make([]bool, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
		dist[i] = inf
		onQueue[i] = false
	}

	// TODO: parent[v] should be v or -1?
	parent[v] = -1
	dist[v] = 0
	Q := newBfsQueue(n)
	Q.Insert(v)
	onQueue[v] = true

	sentinel := n
	Q.Insert(sentinel)

	k := 0
	for {
		v = Q.Remove()
		if v < sentinel {
			for w, weight := range g.EdgesFrom(v) {
				if isNaN(weight) {
					continue
				}
				alt := add(dist[v], weight)
				if alt < dist[w] {
					dist[w] = alt
					parent[w] = v
					if !onQueue[w] {
						Q.Insert(w)
						onQueue[w] = true
					}
				}
			}
		} else {
			if Q.Len() == 0 {
				ok = true
				break
			}

			k++
			if k >= n {
				ok = false // Negative cycle.
				break
			}

			Q.Insert(sentinel)
			for i := range n {
				onQueue[i] = false
			}
		}
	}

	return parent, dist, ok
}

// add add two numbers and check for overflow and
// underflow, if overflow occurs add return positive
// inf, if underflows it returns negative inf.
func add[T IntegerOrFloat](a, b T) T {
	c := a + b
	if a > 0 && b > 0 && c < a {
		return InfFor[T]()
	}
	if a < 0 && b < 0 && c > a {
		return InfFor[T]() + 1 // Negative inf.
	}
	return c
}
