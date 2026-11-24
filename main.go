package main

import "fmt"

// HasCycle reports whether a DIRECTED graph contains a cycle, via DFS with
// three vertex states: unvisited, on the current recursion path, done.
// An edge into an on-path vertex is a back edge — a cycle.
func HasCycle(n int, edges [][2]int) bool {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
	}
	const (
		white = 0 // untouched
		gray  = 1 // on the current DFS path
		black = 2 // fully explored
	)
	state := make([]int, n)
	var walk func(v int) bool
	walk = func(v int) bool {
		state[v] = gray
		for _, w := range adj[v] {
			switch state[w] {
			case gray:
				return true // back edge into our own path
			case white:
				if walk(w) {
					return true
				}
			}
		}
		state[v] = black
		return false
	}
	for s := 0; s < n; s++ {
		if state[s] == white && walk(s) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(HasCycle(3, [][2]int{{0, 1}, {1, 2}, {2, 0}}))
	fmt.Println(HasCycle(3, [][2]int{{0, 1}, {0, 2}, {1, 2}}))
}
