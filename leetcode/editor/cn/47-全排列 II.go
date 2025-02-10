package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	nums = []int{1, 2, 3, 4}
	nums = []int{1, 1, 1, 2}
	ret := permuteUnique(nums)
	fmt.Println(ret)
}

// leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	n := len(nums)
	ret := make([][]int, 0)
	check := func(i, j int) bool {
		for ; i < j; i++ {
			if nums[i] == nums[j] { // 存在重复元素
				return false
			}
		}
		return true
	}
	var dfs func(int)
	dfs = func(i int) {
		if i == n-1 {
			ret = append(ret, append([]int(nil), nums...))
			return
		}
		//memo := make(map[int]bool)	// 也可以每次使用 map 记录元素：第一次出现时操作
		dfs(i + 1) // 分支：n*(n-1)*(n-2)*...*1
		for j := i + 1; j < n; j++ {
			if check(i, j) { // [i,j] 存在重复元素：最后一次出现时才操作
				nums[i], nums[j] = nums[j], nums[i]
				dfs(i + 1)
				nums[i], nums[j] = nums[j], nums[i] // 回溯
			}
		}
	}
	dfs(0)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
