//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
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
//
//
// Related Topics 位运算 数组 回溯 👍 1126 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	dup := subsetsWithDup(nums)
	fmt.Println(dup)
}

/*
思路：迭代
	1.在不考虑元素重复的情况，求子集
		类似 LeetCode-78
		所以我们只需要再加上重复元素的判断即可
	2.在排序的情况下，二进制位标记 nums[i] 是否被选择，原则就是
		当 nums[i] 标记被选择，而 nums[j] == nums[j-1] 且 nums[j-1] 被标记不被选择时
		放弃本轮标记位的子集
*/
// leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	// 迭代
	//	ret, n, m := make([][]int, 1), len(nums), 1<<len(nums)-1
	//	sort.Ints(nums) // 有序则更好去重
	//out:
	//	for i := 1; i < m; i++ {
	//		next := make([]int, 0)
	//		for j := 0; j < n; j++ {
	//			if i>>j&1 > 0 { // nums[j] 被选择
	//				if j > 0 && nums[j] == nums[j-1] && i>>(j-1)&1 == 0 { // 连续相同的数，但前一个数没被选择
	//					continue out // 去重，放弃本轮标记位的子集
	//				}
	//				next = append(next, nums[j])
	//			}
	//		}
	//		if len(next) == 3 {
	//			ret = append(ret, next) // 找到一个子集
	//		}
	//		//ret = append(ret, next) // 找到一个子集
	//	}
	//	ret = append(ret, append([]int(nil), nums...)) // 加尾
	//	return ret

	// 递归
	//	sort.Ints(nums)
	//	n := len(nums)
	//	N := 1 << n
	//	ret := make([][]int, 1)
	//	temp := make([]int, 0, n)
	//out:
	//	for i := 1; i < N; i++ {
	//		for j := 0; j < n; j++ {
	//			if i>>j&1 != 0 {
	//				if j > 0 && nums[j] == nums[j-1] && i>>(j-1)&1 == 0 {
	//					temp = temp[:0]
	//					continue out
	//				}
	//				temp = append(temp, nums[j])
	//			}
	//		}
	//		ret = append(ret, append([]int(nil), temp...))
	//		temp = temp[:0]
	//	}
	//	return ret

	// dfs
	sort.Ints(nums) // 有序
	n := len(nums)
	ret := make([][]int, 0)
	temp := make([]int, 0, n) // 复用 path
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ret = append(ret, append([]int(nil), temp...))
			return
		}
		temp = append(temp, nums[i])
		dfs(i + 1)
		temp = temp[:len(temp)-1]             // 回溯
		for i+1 < n && nums[i] == nums[i+1] { // 重复元素：上一个 dfs 已选择，此处全跳过
			i++
		}
		dfs(i + 1)
	}
	dfs(0)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
