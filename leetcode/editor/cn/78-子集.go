//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
//
// 示例 2：
//
//
//输入：nums = [0]
//输出：[[],[0]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// nums 中的所有元素 互不相同
//
//
// Related Topics 位运算 数组 回溯 👍 2099 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	nums = []int{1, 2, 3, 4}
	i := subsets(nums)
	for _, v := range i {
		fmt.Println(v)
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	// 迭代
	//n := len(nums)
	//m := 1 << n
	//ret := make([][]int, 0, m)
	//for i := 0; i < m; i++ {
	//	temp := make([]int, 0, n)
	//	for j := 0; j < n; j++ {
	//		if i>>j&1 > 0 {
	//			temp = append(temp, nums[j])
	//		}
	//	}
	//	ret = append(ret, temp)
	//}
	//return ret

	// 递归
	m, n := 1<<len(nums), len(nums)
	ret, temp := make([][]int, 0, m), make([]int, 0, n)
	var dfs func(int)
	dfs = func(i int) {
		if i < 0 {
			ret = append(ret, append([]int(nil), temp...))
			return
		}
		temp = append(temp, nums[i])
		dfs(i - 1)
		temp = temp[:len(temp)-1] // 复用内存
		dfs(i - 1)
	}
	dfs(n - 1)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
