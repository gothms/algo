package main

import (
	"container/heap"
	"fmt"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func kClosest(points [][]int, k int) [][]int {
	arr := make(hp, 0, k)
	h := &arr
	for i, p := range points {
		dis := p[0]*p[0] + p[1]*p[1]
		if h.Len() < k || (*h)[0][0] > dis {
			if h.Len() < k {
				heap.Push(h, [2]int{dis, i})
				continue
			}
			(*h)[0] = [2]int{dis, i}
			heap.Fix(h, 0)
		}
	}
	ans := make([][]int, 0, k)
	for _, p := range *h {
		ans = append(ans, points[p[1]])
	}
	return ans
}

type hp [][2]int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *hp) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

var _ heap.Interface = (*hp)(nil)

//leetcode submit region end(Prohibit modification and deletion)
