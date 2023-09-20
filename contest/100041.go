package main

import "fmt"

func main() {
	edges := [][]int{{2, 0}, {2, 1}, {1, 3}, {0, 4}}
	n := 5
	reversals := minEdgeReversals(n, edges)
	fmt.Println(reversals)
}
func minEdgeReversals(n int, edges [][]int) []int {
	adj, inverseAdj := make([][]int, n), make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		inverseAdj[e[1]] = append(inverseAdj[e[1]], e[0])
	}
	var dfs func(int, int) int
	dfs = func(i, f int) int {
		cnt := 0
		for _, t := range adj[i] {
			if f != t {
				cnt += dfs(t, i) // 邻接表的边：不需要反转
			}
		}
		for _, t := range inverseAdj[i] {
			if f != t {
				cnt += dfs(t, i) + 1 // 逆邻接表的边：需要反转
			}
		}
		return cnt
	}
	mr := make([]int, n)
	mr[0] = dfs(0, -1) // 计算顶点 0 的反转次数
	var dfsRange func(int, int)
	dfsRange = func(i, f int) {
		for _, t := range adj[i] {
			if f != t {
				mr[t] = mr[i] + 1 // 邻接表的边：反转次数 +1
				dfsRange(t, i)
			}
		}
		for _, t := range inverseAdj[i] {
			if f != t {
				mr[t] = mr[i] - 1 // 逆邻接表的边：反转次数 -1
				dfsRange(t, i)
			}
		}
	}
	dfsRange(0, -1)
	return mr
}
