package algo

import "fmt"

/*
课程表系列
	lc-207
	lc-210
	lc-630
	lc-1462
*/

// TopoSortByKahn Kahn 算法实现拓扑排序
// n 为顶点数，edges 为所有边
func TopoSortByKahn(n int, edges [][]int) {
	inDegree, adj := make([]int, n), make([][]int, n)
	for _, edge := range edges {
		s, t := edge[0], edge[1] // s->t：s先于t，t依赖于s
		adj[s] = append(adj[s], t)
		inDegree[t]++ // 统计每个顶点的入度
	}
	fmt.Println(adj)
	fmt.Println(inDegree)
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i) // 初始化：入度为 0 的顶点
		}
	}
	for len(queue) > 0 { // 遍历次数 < n，说明有环
		s := queue[0]
		fmt.Print("->", s) // 拓扑排序
		queue = queue[1:]
		for _, t := range adj[s] {
			inDegree[t]-- // 入度 -1
			if inDegree[t] == 0 {
				queue = append(queue, t)
			}
		}
	}
}

// TopoSortByDFS DFS 算法实现拓扑排序
// n 为顶点数，edges 为所有边
func TopoSortByDFS(n int, edges [][]int) {
	inverseAdj := make([][]int, n)
	for _, edge := range edges {
		s, t := edge[0], edge[1]                 // s->t：s先于t，t依赖于s
		inverseAdj[t] = append(inverseAdj[t], s) // dfs 方案适合逆邻接表：inverseAdj
		//inverseAdj[s] = append(inverseAdj[s], t) // 邻接表 dfs 的拓扑排序是反的
	}
	visited := make([]bool, n)
	for i := 0; i < n; i++ { // 遍历每个顶点
		if !visited[i] { // 未访问的顶点
			visited[i] = true
			topoDfs(inverseAdj, visited, i)
		}
	}
}
func topoDfs(inverseAdj [][]int, visited []bool, t int) {
	for _, s := range inverseAdj[t] {
		if !visited[s] {
			visited[s] = true
			topoDfs(inverseAdj, visited, s)
		}
	}
	fmt.Print("->", t) // 拓扑排序
}

// TopoSortByDFSCircle DFS 算法实现拓扑排序，检测是否有环
// n 为顶点数，edges 为所有边
func TopoSortByDFSCircle(n int, edges [][]int) bool {
	inverseAdj := make([][]int, n)
	for _, edge := range edges {
		s, t := edge[0], edge[1]                 // s->t：s先于t，t依赖于s
		inverseAdj[t] = append(inverseAdj[t], s) // dfs 方案适合逆邻接表：inverseAdj
	}
	visited := make([]int8, n)
	for i := 0; i < n; i++ { // 遍历每个顶点
		if visited[i] == 0 && !topoDfsCircle(inverseAdj, visited, i) { // 未访问的顶点
			return false
		}
	}
	return true
}
func topoDfsCircle(inverseAdj [][]int, visited []int8, t int) bool {
	if visited[t] != 0 { // 已访问的顶点
		return visited[t] == 1 // -1 则有环
	}
	visited[t] = -1 // 检测中
	for _, s := range inverseAdj[t] {
		//if visited[s] <= 0 && !topoDfsCircle(inverseAdj, visited, s) {
		if !topoDfsCircle(inverseAdj, visited, s) {
			return false // 存在环
		}
	}
	visited[t] = 1 // 已检测
	return true
}
