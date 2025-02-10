package main

import (
	"fmt"
	"math/bits"
)

func main() {
	nums := []int{8, 4, 2, 30, 15} // true
	//nums = []int{75, 34, 30}       // false
	//nums = []int{3, 16, 8, 4, 2}   // false
	//nums = []int{20, 16}           // false
	//for _, v := range nums {
	//	fmt.Printf("%b\n", v)
	//	fmt.Println(bits.OnesCount(uint(v)))
	//}
	array := canSortArray(nums)
	fmt.Println(array)
}

// leetcode submit region begin(Prohibit modification and deletion)
func canSortArray(nums []int) bool {
	// lc：记录每一段的最小值和最大值
	preMax := 0
	for i, n := 0, len(nums); i < n; {
		mx := 0
		ones := bits.OnesCount(uint(nums[i]))
		for ; i < n && bits.OnesCount(uint(nums[i])) == ones; i++ {
			if nums[i] < preMax { // 无法排成有序的
				return false
			}
			mx = max(mx, nums[i]) // 更新本组最大值
		}
		preMax = mx
	}
	return true

	// 相邻
	//lastCnt, cur, maxV := -1, -1, -1
	//for _, v := range nums {
	//	if v < maxV {
	//		return false
	//	}
	//	i := bits.OnesCount(uint(v))
	//	if i != lastCnt {
	//		lastCnt, cur, maxV = i, v, cur
	//		if v < maxV {
	//			return false
	//		}
	//	} else {
	//		cur = max(cur, v)
	//	}
	//}
	//return true

	// 不相邻
	//memo := [8][]int{}
	//for _, v := range nums {
	//	i := bits.OnesCount(uint(v))
	//	memo[i] = append(memo[i], v)
	//}
	//for _, vs := range memo {
	//	sort.Ints(vs)
	//}
	//last := -1
	//for _, v := range nums {
	//	i := bits.OnesCount(uint(v))
	//	cur := memo[i][0]
	//	if cur < last {
	//		return false
	//	}
	//	last = cur
	//	memo[i] = memo[i][1:]
	//}
	//return true
}

//leetcode submit region end(Prohibit modification and deletion)
