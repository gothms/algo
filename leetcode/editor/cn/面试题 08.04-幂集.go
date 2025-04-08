package main

import (
	"fmt"
	"slices"
)

func main() {
	n := 6
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	i := subsets(nums)
	fmt.Println(len(i), 1<<n, i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0, 1<<n)
	path := make([]int, 0, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		dfs(i + 1)
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
	}
	dfs(0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
