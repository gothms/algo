package main

import (
	"fmt"
	"math"
)

func main() {
	edges := [][]int{{4, 1},
		{5, 1},
		{3, 2},
		{5, 0},
		{4, 0},
		{3, 0},
		{2, 1}}
	n := 6 // 4
	edges = [][]int{{4, 2},
		{5, 1},
		{5, 0},
		{0, 3},
		{5, 2},
		{1, 4},
		{1, 3},
		{3, 4}}
	n = 6 // 3
	cycle := findShortestCycle(n, edges)
	fmt.Println(cycle)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findShortestCycle(n int, edges [][]int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func findShortestCycle_(n int, edges [][]int) int {
	// Dijkstra：比 bfs 慢
	//adj := make([][]int, n)
	//for _, e := range edges {
	//	x, y := e[0], e[1]
	//	adj[x], adj[y] = append(adj[x], y), append(adj[y], x)
	//}
	//ans := math.MaxInt32
	//for i := 0; i < n; i++ { // 枚举每个点作为起点
	//	dis := make([]int, n)
	//	dis[i] = 1
	//	h := &hp2608{{-1, i, 1}}
	//out:
	//	for h.Len() > 0 {
	//		cur := heap.Pop(h).([3]int)
	//		f, t, d := cur[0], cur[1], cur[2]
	//		for _, next := range adj[t] {
	//			if next == f {
	//				continue
	//			}
	//			if dis[next] > 0 {
	//				ans = min(ans, d+dis[next]-1)
	//				break out
	//			} else {
	//				dis[next] = d + 1
	//				heap.Push(h, [3]int{t, next, d + 1})
	//			}
	//		}
	//	}
	//}
	//if ans == math.MaxInt32 {
	//	return -1
	//}
	//return ans

	// bfs
	adj := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		adj[x], adj[y] = append(adj[x], y), append(adj[y], x)
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ { // 枚举每个点作为起点
		dis := make([]int, n) // 重点：标记深度
		dis[i] = 1
		tag := true // 标记：遍历完同一深度的顶点
		q := [][2]int{{-1, i}}
		for l := len(q); l > 0 && tag; l = len(q) {
			for j := 0; j < l; j++ {
				f, t := q[j][0], q[j][1]
				for _, next := range adj[t] {
					if next == f {
						continue
					}
					if dis[next] > 0 { // 找到环了
						ans = min(ans, dis[t]+dis[next]-1)
						tag = false // 需要把同深度的环都访问完，因为有可能先找到一个 x 的环，但本深度下还有 x-1 的环
					} else {
						dis[next] = dis[t] + 1
						q = append(q, [2]int{t, next})
					}
				}
			}
			q = q[l:]
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

//type hp2608 [][3]int
//
//func (h hp2608) Len() int           { return len(h) }
//func (h hp2608) Less(i, j int) bool { return h[i][2] < h[j][2] }
//func (h hp2608) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//func (h *hp2608) Push(x any)        { *h = append(*h, x.([3]int)) }
//func (h *hp2608) Pop() any {
//	v := (*h)[len(*h)-1]
//	*h = (*h)[:len(*h)-1]
//	return v
//}
