package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 0, 1}
	goal := 1
	//nums = []int{ 0, 0, 0, 0, 0}
	//goal = 0
	withSum := numSubarraysWithSum(nums, goal)
	fmt.Println(withSum)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numSubarraysWithSum(nums []int, goal int) int {
	// lc
	memo := make(map[int]int)
	ans, sum := 0, 0
	for _, v := range nums {
		memo[sum]++
		sum += v
		ans += memo[sum-goal]
	}
	return ans

	// 个人
	//n := len(nums)
	//memo := make([]int, 1, n+2)
	//memo[0] = -1
	//for i, v := range nums {
	//	if v == 1 {
	//		memo = append(memo, i)
	//	}
	//}
	//memo = append(memo, n)
	//ans := 0
	//if goal == 0 {
	//	for l, r := 0, 1; r < len(memo); r++ {
	//		c := memo[r] - memo[l] - 1
	//		ans += (c + 1) * c >> 1
	//		l++
	//	}
	//} else {
	//	for l, r := 1, goal; r < len(memo)-1; r++ {
	//		ans += (memo[l] - memo[l-1]) * (memo[r+1] - memo[r])
	//		l++
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
