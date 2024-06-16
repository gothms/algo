package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 6, 1, 2}
	k := 2
	beauty := maximumBeauty(nums, k)
	fmt.Println(beauty)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumBeauty(nums []int, k int) int {
	// 排序+滑动窗体
	//sort.Ints(nums)
	//ans, j := 0, 0
	//k <<= 1
	//for i, v := range nums {
	//	for nums[j] < v-k {
	//		j++
	//	}
	//	ans = max(ans, i-j+1)
	//}
	//return ans

	// 差分数组：优化
	maxV := 0
	for _, v := range nums {
		maxV = max(maxV, v)
	}
	maxV++
	diff := make([]int, maxV+1)
	for _, v := range nums {
		diff[max(0, v-k)]++
		diff[min(maxV, v+k+1)]--
	}
	ans, curV := 0, 0
	for _, v := range diff {
		curV += v
		ans = max(ans, curV)
	}
	return ans

	// 差分数组
	//minV, maxV := math.MaxInt32, math.MinInt32
	//for _, v := range nums {
	//	minV = min(minV, v)
	//	maxV = max(maxV, v)
	//}
	//minV -= k
	//maxV += k
	//temp := make([]int, maxV-minV+2) // 差分数组：实质元素总个数为 maxV-minV+1
	//for _, v := range nums {
	//	temp[v-k-minV]++
	//	temp[v+k-minV+1]--
	//}
	//ans, curV := 0, 0
	//for _, v := range temp {
	//	curV += v
	//	ans = max(ans, curV)
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
