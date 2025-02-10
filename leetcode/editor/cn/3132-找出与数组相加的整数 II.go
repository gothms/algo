package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums1 := []int{3, 5, 5, 3}
	nums2 := []int{7, 7}
	integer := minimumAddedInteger(nums1, nums2)
	fmt.Println(integer)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumAddedInteger(nums1 []int, nums2 []int) int {
	ans, minV := math.MaxInt32, nums2[len(nums2)-1]
	arr := [3]int(nums1[:3])
	sort.Ints(arr[:])
	for i, v := range nums1[3:] {
		if arr[2] > v {
			arr[2] = v
			sort.Ints(arr[:])
		}
		minV = min(minV, nums2[i])
	}
	for _, val := range arr {
		memo := make(map[int]int)
		cnt, d := 0, minV-val
		for _, v := range nums2 {
			memo[v]++
		}
		for _, v := range nums1 {
			if memo[v+d] == 0 {
				cnt++
			} else {
				memo[v+d]--
			}
		}
		if cnt == 2 {
			ans = min(ans, d)
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
//func minimumAddedInteger(nums1 []int, nums2 []int) int {
//	// hash
//	// 其他方法：排序后，取 nums1[2] nums1[1] nums1[0] 尝试
//	hp, m := &hp3132{}, math.MaxInt32
//	for _, v := range nums1 {
//		if hp.Len() < 3 {
//			heap.Push(hp, v)
//		} else if v < hp.IntSlice[0] {
//			hp.IntSlice[0] = v
//			heap.Fix(hp, 0)
//		}
//	}
//
//	ans := math.MaxInt32
//	for _, mv := range hp.IntSlice {
//		memo := make(map[int]int, len(nums2))
//		for _, v := range nums2 {
//			m = min(m, v)
//			memo[v]++
//		}
//		d, cnt := m-mv, 0
//		for _, v := range nums1 {
//			if memo[v+d] <= 0 {
//				cnt++
//			}
//			memo[v+d]--
//		}
//		if cnt == 2 {
//			ans = min(ans, d)
//		}
//	}
//	return ans
//}
//
//type hp3132 struct {
//	sort.IntSlice
//}
//
//func (h hp3132) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
//func (h *hp3132) Push(x any)        { h.IntSlice = append(h.IntSlice, x.(int)) }
//func (h *hp3132) Pop() any {
//	v := h.IntSlice[len(h.IntSlice)-1]
//	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
//	return v
//}
