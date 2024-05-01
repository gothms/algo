package main

import (
	"algo/algo"
	"fmt"
)

func main() {
	// 拓扑排序
	n := 6
	edges := [][]int{{0, 3}, {3, 2}, {0, 1}, {1, 3}, {4, 1}, {4, 2}, {5, 1}}
	//edges = [][]int{{0, 3}, {3, 2}, {0, 1}, {1, 3}, {4, 1}, {4, 2}, {2, 1}} // 存在环
	algo.TopoSortByKahn(n, edges)
	fmt.Println()
	algo.TopoSortByDFS(n, edges)
	fmt.Println()
	circle := algo.TopoSortByDFSCircle(n, edges)
	fmt.Println(circle)

	// bfs
	//inDegree, adj := make([]int, n), make([][]int, n)
	//for _, edge := range edges {
	//	s, t := edge[0], edge[1] // s->t：s先于t，t依赖于s
	//	adj[s] = append(adj[s], t)
	//	inDegree[t]++ // 统计每个顶点的入度
	//}
	//kahn := make([]int, 0)
	//bfs := make([]int, 0, n)
	//for i, v := range inDegree {
	//	if v == 0 {
	//		kahn = append(kahn, i)
	//		bfs = append(bfs, i)
	//	}
	//}
	//for l := len(kahn); l > 0; l = len(kahn) {
	//	for i := 0; i < l; i++ {
	//		for _, v := range adj[kahn[i]] {
	//			inDegree[v]--
	//			if inDegree[v] == 0 {
	//				kahn = append(kahn, v)
	//				bfs = append(bfs, v)
	//			}
	//		}
	//	}
	//	kahn = kahn[l:]
	//}
	//for _, v := range bfs {
	//	fmt.Print("->", v)
	//}
}
