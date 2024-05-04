//有 n 个网络节点，标记为 1 到 n。
//
// 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点，
//wi 是一个信号从源节点传递到目标节点的时间。
//
// 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
//输出：2
//
//
// 示例 2：
//
//
//输入：times = [[1,2,1]], n = 2, k = 1
//输出：1
//
//
// 示例 3：
//
//
//输入：times = [[1,2,1]], n = 2, k = 2
//输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= k <= n <= 100
// 1 <= times.length <= 6000
// times[i].length == 3
// 1 <= ui, vi <= n
// ui != vi
// 0 <= wi <= 100
// 所有 (ui, vi) 对都 互不相同（即，不含重复边）
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 最短路 堆（优先队列） 👍 667 👎 0

package main

import (
	"container/heap"
	"math"
	"slices"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func networkDelayTime(times [][]int, n int, k int) int {
	// 优化
	adj := make([][][2]int, n+1)
	for _, time := range times {
		adj[time[0]] = append(adj[time[0]], [2]int{time[1], time[2]})
	}
	h, dijkstra := &hp743{{0, k}}, make([]int, n+1)
	for i := 1; i <= n; i++ {
		dijkstra[i] = math.MaxInt32
	}
	dijkstra[k] = 0
	for h.Len() > 0 {
		cur := heap.Pop(h).([2]int)
		if cur[0] > dijkstra[cur[1]] { // cur[1] 已出堆，则已找到最短路径
			continue
		}
		for _, next := range adj[cur[1]] {
			if d := cur[0] + next[1]; d < dijkstra[next[0]] {
				dijkstra[next[0]] = d
				heap.Push(h, [2]int{d, next[0]})
			}
		}
	}
	ans := slices.Max(dijkstra)
	if ans < math.MaxInt32 {
		return ans
	}
	return -1

	// Dijkstra：个人
	//adj := make([][][2]int, n+1)
	//for _, time := range times {
	//	adj[time[0]] = append(adj[time[0]], [2]int{time[1], time[2]})
	//}
	//h, dijkstra, vis := &hp743{{0, k}}, make([]int, n+1), make([]bool, n+1)
	//for i := range dijkstra {
	//	dijkstra[i] = math.MaxInt32
	//}
	//dijkstra[k] = 0
	//for h.Len() > 0 {
	//	cur := heap.Pop(h).([2]int)
	//	if vis[cur[1]] {
	//		continue
	//	}
	//	vis[cur[1]] = true
	//	for _, next := range adj[cur[1]] {
	//		if d := cur[0] + next[1]; d < dijkstra[next[0]] {
	//			dijkstra[next[0]] = d
	//			heap.Push(h, [2]int{d, next[0]})
	//		}
	//	}
	//}
	//var ans int
	//for i := 1; i <= n; i++ {
	//	if !vis[i] {
	//		return -1
	//	}
	//	ans = max(ans, dijkstra[i])
	//}
	//return ans
}

type hp743 [][2]int

func (h hp743) Len() int {
	return len(h)
}
func (h hp743) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}
func (h hp743) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp743) Push(x any) {
	*h = append(*h, x.([2]int))
}
func (h *hp743) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
