package main

import (
	"fmt"
	"math"
)

func main() {
	edges := [][]int{{0, 2, 2},
		{0, 5, 6},
		{1, 0, 3},
		{1, 4, 5},
		{2, 1, 1},
		{2, 3, 3},
		{2, 3, 4},
		{3, 4, 2},
		{4, 5, 1}}
	n := 6
	src1, src2, dest := 0, 1, 5
	//edges = [][]int{{4, 2, 20},
	//	{4, 3, 46},
	//	{0, 1, 15},
	//	{0, 1, 43},
	//	{0, 1, 32},
	//	{3, 1, 13}}
	//n = 5
	//src1, src2, dest = 0, 4, 1 // 74
	//edges = [][]int{{0, 1, 1}, {2, 1, 1}}
	//n = 3
	//src1, src2, dest = 0, 1, 2 // -1
	weight := minimumWeight(n, edges, src1, src2, dest)
	fmt.Println(weight)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	// Dijkstra
	//const inf = math.MaxInt64 / 3
	//dijkstra := func(s int, adj [][][2]int) []int {
	//	dis, h := make([]int, n), &hp2203{{s, 0}}
	//	for i := range dis {
	//		dis[i] = inf
	//	}
	//	dis[s] = 0
	//	for h.Len() > 0 {
	//		cur := heap.Pop(h).([2]int)
	//		f, d := cur[0], cur[1]
	//		if d > dis[f] { // 重要：否则超时
	//			continue
	//		}
	//		for _, e := range adj[f] {
	//			t, w := e[0], e[1]
	//			if dist := d + w; dist < dis[t] {
	//				dis[t] = dist
	//				heap.Push(h, [2]int{t, dist})
	//			}
	//		}
	//	}
	//	return dis
	//}
	//adj, ivAdj := make([][][2]int, n), make([][][2]int, n)
	//for _, e := range edges {
	//	f, t, w := e[0], e[1], e[2]
	//	adj[f] = append(adj[f], [2]int{t, w})
	//	ivAdj[t] = append(ivAdj[t], [2]int{f, w}) // 逆邻接表
	//}
	//d1 := dijkstra(src1, adj)
	//d2 := dijkstra(src2, adj)
	//d3 := dijkstra(dest, ivAdj)
	//
	//ans := inf
	//for i := 0; i < n; i++ { // 三个节点到 i 的距离之和
	//	ans = min(ans, d1[i]+d2[i]+d3[i])
	//}
	//if ans >= inf {
	//	return -1
	//}
	//return int64(ans)

	// floyd：超时
	const inf = math.MaxInt32 >> 1
	floyd, adj := make([][]int, n), make([][][2]int, n)
	for i := range floyd {
		floyd[i] = make([]int, n)
		for j := range floyd[i] {
			floyd[i][j] = inf
		}
		floyd[i][i] = 0
	}
	for _, e := range edges {
		f, t, w := e[0], e[1], e[2]
		adj[f] = append(adj[f], [2]int{t, w})
		floyd[f][t] = min(floyd[f][t], w)
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j {
					floyd[i][j] = min(floyd[i][j], floyd[i][k]+floyd[k][j])
				}
			}
		}
	}
	for i, f := range floyd {
		fmt.Println(i, f)
	}
	ans := inf
	for k := 0; k < n; k++ {
		ans = min(ans, floyd[src1][k]+floyd[src2][k]+floyd[k][dest])
	}
	if ans >= inf {
		return -1
	}
	return int64(ans)
}

type hp2203 [][2]int

func (h hp2203) Len() int           { return len(h) }
func (h hp2203) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h hp2203) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp2203) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *hp2203) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
