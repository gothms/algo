package main

import (
	"container/heap"
	"fmt"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSum(grid [][]int, limits []int, k int) int64 {
	h, hr := &hp{}, &hp{}
	for i, g := range grid {
		for _, v := range g {
			heap.Push(hr, v)
			if hr.Len() > limits[i] {
				heap.Pop(hr)
			}
		}
		for hr.Len() > 0 {
			heap.Push(h, heap.Pop(hr))
			if h.Len() > k {
				heap.Pop(h)
			}
		}
	}
	ans := 0
	for _, v := range *h {
		ans += v
	}
	return int64(ans)
}

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] < h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.(int)) }
func (h *hp) Pop() any {
	V := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return V
}

var _ heap.Interface = (*hp)(nil)

//leetcode submit region end(Prohibit modification and deletion)
