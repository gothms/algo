package main

import (
	"container/heap"
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func smallestK(arr []int, k int) []int {
	// 快排

	// lc
	if k == 0 {
		return nil
	}
	h := &hp{arr[:k]}
	heap.Init(h)
	for _, v := range arr[k:] {
		if h.IntSlice[0] > v {
			h.IntSlice[0] = v
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice

	// 堆
	//if k == 0 {
	//	return nil
	//}
	//h := &hp{}
	//for _, v := range arr {
	//	if h.Len() < k {
	//		heap.Push(h, v)
	//	} else if h.IntSlice[0] > v {
	//		h.IntSlice[0] = v
	//		heap.Fix(h, 0)
	//	}
	//}
	//return h.IntSlice
}

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(x any)        { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hp) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

var _ heap.Interface = (*hp)(nil)

//leetcode submit region end(Prohibit modification and deletion)
