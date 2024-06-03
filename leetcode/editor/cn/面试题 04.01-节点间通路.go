package main

import "fmt"

func main() {
	graph := [][]int{{0, 1},
		{1, 2},
		{1, 3},
		{1, 10},
		{1, 11},
		{1, 4},
		{2, 4},
		{2, 6},
		{2, 9},
		{2, 10},
		{2, 4},
		{2, 5},
		{2, 10},
		{3, 7},
		{3, 7},
		{4, 5},
		{4, 11},
		{4, 11},
		{4, 10},
		{5, 7},
		{5, 10},
		{6, 8},
		{7, 11},
		{8, 10}}
	n, start, target := 12, 2, 3
	path := findWhetherExistsPath(n, graph, start, target)
	fmt.Println(path)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	// 并查集：错误，因为是有向图
	//uni := make([]int, n)
	//for i := range uni {
	//	uni[i] = i
	//}
	//var find func(int) int
	//find = func(i int) int {
	//	for uni[i] != i {
	//		uni[i], i = uni[uni[i]], uni[i]
	//	}
	//	return uni[i]
	//}
	//union := func(p, q int) {
	//	if rootP, rootQ := find(p), find(q); rootP != rootQ {
	//		uni[rootQ] = rootP
	//	}
	//}
	//for _, g := range graph {
	//	union(g[0], g[1])
	//}
	//return find(start) == find(target)

	// dfs

	// bfs
	adj := make([][]int, n)
	for _, g := range graph {
		x, y := g[0], g[1]
		adj[x] = append(adj[x], y)
	}
	q, vis := []int{start}, make([]bool, n)
	vis[start] = true
	for len(q) > 0 {
		i := q[0]
		if i == target {
			return true
		}
		for _, j := range adj[i] {
			if !vis[j] {
				vis[j] = true
				q = append(q, j)
			}
		}
		q = q[1:]
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func findWhetherExistsPath_(n int, graph [][]int, start int, target int) bool {
	adj, visited := make([][]int, n), make([]bool, n)
	for _, g := range graph {
		adj[g[0]] = append(adj[g[0]], g[1])
	}
	visited[start] = true
	var dfs func(int, int) bool
	dfs = func(f, t int) bool {
		for _, i := range adj[t] {
			if i == target {
				return true
			}
			if !visited[i] {
				visited[i] = true
				if dfs(t, i) {
					return true
				}
			}
		}
		return false
	}
	return dfs(-1, start)
}
