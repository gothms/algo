package main

import "slices"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0, func(n int) int {
		v := 1
		for i := 2; i <= n; i++ {
			v *= i
		}
		return v
	}(len(nums)))
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(nums))
			return
		}
		dfs(i + 1)
		for j := i + 1; j < n; j++ {
			nums[i], nums[j] = nums[j], nums[i]
			dfs(i + 1)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	dfs(0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
