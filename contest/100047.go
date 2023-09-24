package main

import "fmt"

func main() {

	//for i, p := range primes {
	//	if !p {
	//		fmt.Print(i, " ")
	//	}
	//}
	edges := [][]int{{1, 5},
		{5, 4},
		{4, 6},
		{5, 8},
		{8, 9},
		{9, 10},
		{5, 2},
		{2, 3},
		{3, 7}}
	n := 10 // [[1,5],[5,4],[4,6],[5,8],[8,9],[9,10],[5,2],[2,3],[3,7]]	// 17
	//edges = [][]int{{1, 2}, {4, 1}, {3, 4}} // 4
	//n = 4
	paths := countPaths(n, edges)
	fmt.Println(paths)
}

/*
示例：[[1,5],[5,4],[4,6],[5,8],[8,9],[9,10],[5,2],[2,3],[3,7]]	// 17
	与 5 相连的子树：[1],[4,6],[8,9,10]
	数学：f(5) = 1+2+3 + 1*2+1*3+2*3
		数学方式数据计算有误，如 edges = [][]int{{1, 2}, {4, 1}, {3, 4}}	// 4
	前缀和：f(5) = 0*1 + 1*2 + (1+2)*3 + 1+2+3
*/

func countPaths(n int, edges [][]int) int64 {
	// 前缀和：优化 dfs

	// 个人写法用的数学，lc 写法用的前缀和
	var cnt int64
	cp, adj := make([]int, n+1), make([][]int, n+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		adj[x], adj[y] = append(adj[x], y), append(adj[y], x)
	}
	var vs []int
	var dfs func(int, int)
	dfs = func(f, t int) {
		vs = append(vs, t)
		for _, i := range adj[t] {
			if i != f && primes[i] { // 每个区域只能有一个质数
				dfs(t, i)
			}
		}
	}
	for i := 1; i <= n; i++ {
		if primes[i] { // 只统计质数
			continue
		}
		preSum := 0 // 前缀和
		for _, next := range adj[i] {
			if !primes[next] { // 每个质数为一个连通区域
				continue
			}
			if cp[next] == 0 { // 未计算的区域
				vs = make([]int, 0) // 连通区域内的节点
				dfs(i, next)        // 计算每个连通区域的节点总数
				for _, v := range vs {
					cp[v] = len(vs) // 更新连通区域内的节点数
				}
			}
			cp[i] += preSum * cp[next] // 前缀和 * cp[next]
			preSum += cp[next]         // 前缀和
		}
		cp[i] += preSum     // 单条路径
		cnt += int64(cp[i]) // 累加质数的路径
	}
	return cnt
}

const N = 1e5

var primes [N + 1]bool

func init() {
	primes[1] = true // false 为质数
	for i := 2; i <= N; i++ {
		if !primes[i] { // 质数
			for j := i << 1; j <= N; j += i {
				primes[j] = true
			}
		}
	}
}
