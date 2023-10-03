package main

import "fmt"

func main() {
	edges := []int{1, 2, 0, 0, 3}
	edges = []int{1, 2, 3, 4, 0}
	edges = []int{7, 0, 7, 0, 5, 3, 3, 0}
	edges = []int{6, 3, 6, 1, 0, 8, 0, 6, 6}
	nodes := countVisitedNodes(edges)
	fmt.Println(nodes)
}

/*
重点：
	题意：有向图，包含 n 个节点，节点编号从 0 到 n - 1
	edges：并且每个顶点都有下一个顶点
	结论：每个连通的结尾，都是一个环，且环中的顶点都没有环以外的下一个顶点
*/

func countVisitedNodes(edges []int) []int {
	// 拓扑排序
	// 1.入度为 0 的顶点为当前起始顶点，最后剩下的顶点都在环中（可能是不同的环中）
	// 2.求出环中的顶点的值
	// 3.反向求环以外的顶点的值

	// 图：dfs
	n := len(edges)
	cvn, visited := make([]int, n), make([]bool, n)
	var dfs func(int) (int, int)
	dfs = func(t int) (int, int) {
		if visited[t] { // 找到环 / 已计算值
			return t, cvn[t]
		}
		visited[t] = true
		i, d := dfs(edges[t]) // 有向图
		d++                   // 距离 +1
		cvn[t] = d            // 环中的顶点的值会被修改
		if i == t {
			for down := edges[t]; down != i; down = edges[down] {
				cvn[down] = d // 修正环中顶点的值
			}
		}
		return i, d
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
		}
	}
	return cvn
}
