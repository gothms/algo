package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1}
	nums = []int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720} // [9,18,90,180,360,720]
	subset := largestDivisibleSubset(nums)
	fmt.Println(subset)
}

// leetcode submit region begin(Prohibit modification and deletion)
func largestDivisibleSubset(nums []int) []int {
	// 个人：dp
	sort.Ints(nums)
	n := len(nums)
	idx, mx := 0, 0
	dp := make([][2]int, n)
	for i, v := range nums {
		dp[i] = [2]int{-1, 1} // 上一个元素的索引，最大子集长度
		for j := 0; j < i; j++ {
			if v%nums[j] == 0 && dp[j][1] >= dp[i][1] { // 长度相同的子集，选其一即可
				dp[i] = [2]int{j, dp[j][1] + 1}
			}
		}
		if dp[i][1] > mx { // 更新最大子集
			idx, mx = i, dp[i][1]
		}
	}
	ans := make([]int, 0, mx)
	for i := idx; i >= 0; i = dp[i][0] {
		ans = append(ans, nums[i])
	}
	return ans

	// 超时
	//sort.Ints(nums)
	//i := 0
	//if nums[0] == 1 {
	//	i = 1
	//}
	//memo := make([][]int, 0)
	//for ; i < len(nums); i++ {
	//	v, add := nums[i], false
	//	for _, m := range memo {
	//		if v%m[len(m)-1] == 0 {
	//			memo = append(memo, append(slices.Clone(m), v))
	//			add = true
	//		}
	//	}
	//	if !add {
	//		memo = append(memo, []int{v})
	//	}
	//}
	////fmt.Println(memo)
	//mx := 0
	//for k, m := range memo {
	//	if len(m) > len(memo[mx]) {
	//		mx = k
	//	}
	//}
	//if nums[0] == 1 {
	//	if len(memo) > 0 {
	//		return append([]int{1}, memo[mx]...)
	//	}
	//	return []int{1}
	//}
	//return memo[mx]
}

//leetcode submit region end(Prohibit modification and deletion)
