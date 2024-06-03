package main

import "fmt"

func main() {
	edges := [][]int{{4, 3},
		{1, 4},
		{4, 8},
		{1, 7},
		{6, 4},
		{4, 2},
		{7, 4},
		{4, 0},
		{0, 9},
		{5, 4}}
	n, source, destination := 10, 5, 9
	path := validPath(n, edges, source, destination)
	fmt.Println(path)
}

// leetcode submit region begin(Prohibit modification and deletion)
func validPath(n int, edges [][]int, source int, destination int) bool {
	// 并查集

	// dfs
	adj, vis := make([][]int, n), make([]bool, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		adj[x], adj[y] = append(adj[x], y), append(adj[y], x)
	}
	vis[source] = true
	var dfs func(int) bool
	dfs = func(f int) bool {
		if f == destination {
			return true
		}
		for _, t := range adj[f] {
			if !vis[t] {
				vis[t] = true
				if dfs(t) {
					return true
				}
			}
		}
		return false
	}
	return dfs(source)

	//bfs
}

//leetcode submit region end(Prohibit modification and deletion)
