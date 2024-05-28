package main

import "fmt"

func main() {
	n := 3
	dislikes := [][]int{{1, 2},
		{1, 3},
		{2, 3}} // false
	n = 10
	dislikes = [][]int{{4, 7},
		{4, 8},
		{5, 6},
		{1, 6},
		{3, 7},
		{2, 5},
		{5, 8},
		{1, 2},
		{4, 9},
		{6, 10},
		{8, 10},
		{3, 6},
		{2, 10},
		{9, 10},
		{3, 9},
		{2, 3},
		{1, 9},
		{4, 6},
		{5, 7},
		{3, 8},
		{1, 8},
		{1, 7},
		{2, 4}} // true
	bipartition := possibleBipartition(n, dislikes)
	fmt.Println(bipartition)
}

// leetcode submit region begin(Prohibit modification and deletion)
func possibleBipartition(n int, dislikes [][]int) bool {
	// 并查集
	//p := make([]int, (n+1)<<1)
	//for i := range p {
	//	p[i] = i
	//}
	//var find func(int) int
	//find = func(i int) int {
	//	if p[i] != i {
	//		p[i] = find(p[i])
	//	}
	//	return p[i]
	//}
	//union := func(i, j int) {
	//	p[find(i)] = p[find(j)]
	//}
	//for _, d := range dislikes {
	//	i, j := d[0], d[1]
	//	if find(i) == find(j) {
	//		return false
	//	}
	//	union(i, j+n)
	//	union(j, i+n)
	//}
	//return true

	// bfs

	// dfs：染色法，判断二分图的模板题
	adj, color := make([][]int, n), make([]int, n)
	for _, d := range dislikes {
		x, y := d[0]-1, d[1]-1
		adj[x], adj[y] = append(adj[x], y), append(adj[y], x) // 因为要标记，所以需要一次遍历完有关系的人，所以要逆关系
	}
	var dfs func(int, int) bool
	dfs = func(i, col int) bool {
		color[i] = col
		for _, j := range adj[i] {
			if color[j] == col || color[j] == 0 && !dfs(j, -col) { // 染色为 -col
				return false
			}
		}
		return true
	}
	for i, c := range color {
		if c == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
