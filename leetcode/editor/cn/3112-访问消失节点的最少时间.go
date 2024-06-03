package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	edges := [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}}
	disappear := []int{1, 3, 5}
	n := 3
	time := minimumTime(n, edges, disappear)
	fmt.Println(time)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumTime(n int, edges [][]int, disappear []int) []int {
	adj, dis := make([][][2]int, n), make([]int, n)
	for _, e := range edges {
		x, y, t := e[0], e[1], e[2]
		adj[x], adj[y] = append(adj[x], [2]int{y, t}), append(adj[y], [2]int{x, t})
	}
	ans, h := make([]int, n), &hp3112{{}}
	for i := range ans {
		ans[i], dis[i] = -1, math.MaxInt32
	}
	dis[0] = 0
	for h.Len() > 0 {
		cur := heap.Pop(h).([2]int)
		i, t := cur[0], cur[1]
		if ans[i] > 0 {
			continue
		}
		ans[i] = t
		for _, e := range adj[i] {
			if d := t + e[1]; d < disappear[e[0]] && d < dis[e[0]] {
				dis[e[0]] = d
				heap.Push(h, [2]int{e[0], d})
			}
		}
	}
	return ans
}

type hp3112 [][2]int

func (h hp3112) Len() int           { return len(h) }
func (h hp3112) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h hp3112) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp3112) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *hp3112) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
