package main

import "container/heap"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	// lc：稳定排序

	// 个人
	//arr := make(hp, 0, len(nums))
	//s := 0
	//for i, v := range nums {
	//	arr = append(arr, [2]int{v, i})
	//	s += v
	//}
	//h := &arr
	//heap.Init(h)
	//
	//m := len(queries)
	//ans := make([]int64, m)
	//for i, q := range queries {
	//	idx, k := q[0], q[1]
	//	if nums[idx] > 0 {
	//		s -= nums[idx]
	//		nums[idx] = 0
	//	}
	//	for k > 0 && h.Len() > 0 {
	//		p := heap.Pop(h).([2]int)
	//		v, j := p[0], p[1]
	//		if nums[j] == 0 {
	//			continue
	//		}
	//		s -= v
	//		nums[j] = 0
	//		k--
	//	}
	//	ans[i] = int64(s)
	//	if h.Len() == 0 {
	//		break
	//	}
	//}
	//return ans
}

type hp [][2]int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i][0] < h[j][0] || h[i][0] == h[j][0] && h[i][1] < h[j][1] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *hp) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

var _ heap.Interface = (*hp)(nil)

//leetcode submit region end(Prohibit modification and deletion)
