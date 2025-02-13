package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{9}
	maxOperations := 2
	nums = []int{2, 4, 8, 2}
	maxOperations = 4
	nums = []int{7, 17}
	maxOperations = 2
	nums = []int{1, 1, 1}
	maxOperations = 2
	nums = []int{431, 922, 158, 60, 192, 14, 788, 146, 788, 775, 772, 792, 68, 143, 376, 375, 877, 516, 595, 82, 56, 704, 160, 403, 713, 504, 67, 332, 26}
	maxOperations = 80 // 129
	//sort.Ints(nums)
	//cnt := 0
	//for _, v := range nums[7:] {
	//	//cnt += (v - 1) / 128 // 81
	//	cnt += v / (128 + 1) //
	//}
	//fmt.Println(cnt)

	//nums = []int{99, 99, 99, 99}
	//maxOperations = 3
	size := minimumSize(nums, maxOperations)
	fmt.Println(size)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumSize(nums []int, maxOperations int) int {
	// 二分查找：个人
	//sort.Ints(nums) // 为了二分：2
	//s, maxV, n := 0, 0, len(nums)
	//if nums[0] == nums[n-1] { // fast path
	//	return max(1, (nums[0]-1)/(maxOperations/n+1)+1)
	//}
	//for _, v := range nums {
	//	s += v              // 和
	//	maxV = max(maxV, v) // 最大值
	//}
	//minV := max(1, (s+1)/(n+maxOperations))
	//v := sort.Search(maxV-minV, func(i int) bool { // 二分：1
	//	op, t := 0, i+minV                     // t 可能是 0，防止 /0，则初始化 minV >= 1
	//	j := sort.Search(n, func(j int) bool { // 二分：2
	//		return nums[j] > t
	//	})
	//	for _, v := range nums[j:] {
	//		op += (v - 1) / t
	//		if op > maxOperations {
	//			return false
	//		}
	//	}
	//	//return op <= maxOperations  // 操作次数 <= maxOperations 为合法
	//	return true
	//}) + minV
	//return v

	// 二分
	s, maxV, n := 0, 0, len(nums)
	for _, v := range nums {
		s += v
		maxV = max(maxV, v)
	}
	minV := max(1, (s+1)/(n+maxOperations))
	v := sort.Search(maxV-minV, func(i int) bool {
		op, t := 0, i+minV
		for _, v := range nums {
			if op += (v - 1) / t; op > maxOperations {
				return false
			}
		}
		return true
	}) + minV
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
