package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{0, 0, 1, 2}
	unique := permuteUnique(nums)
	fmt.Println(unique)
}

// leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	//sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(nums))
			return
		}
		dfs(i + 1)
	out:
		for j := i + 1; j < n; j++ {
			for k := i; k < j; k++ {
				if nums[k] == nums[j] {
					continue out
				}
			}
			nums[i], nums[j] = nums[j], nums[i]
			dfs(i + 1)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	dfs(0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
