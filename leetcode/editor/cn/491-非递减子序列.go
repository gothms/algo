//给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
//
// 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,6,7,7]
//输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
//
//
// 示例 2：
//
//
//输入：nums = [4,4,3,2,1]
//输出：[[4,4]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 15
// -100 <= nums[i] <= 100
//
//
// Related Topics 位运算 数组 哈希表 回溯 👍 783 👎 0

package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	nums := []int{4, 6, 7, 7}   // 8
	nums = []int{4, 5, 6, 7, 7} // 19
	//nums = []int{9, 2, 5, 10, 3, 7, 11, 12} // 63
	//nums = []int{9, 2, 5, 10}               // 5
	//nums = []int{1, 2, 3, 4, 5, 5}          // 42
	nums = []int{4, 4, 3, 2, 1} // 1
	subsequences := findSubsequences(nums)
	fmt.Println(subsequences, len(subsequences))
}

// leetcode submit region begin(Prohibit modification and deletion)
func findSubsequences(nums []int) [][]int {
	// lc：剪枝
	n := len(nums)
	ans, path := make([][]int, 0), make([]int, n+1)
	path[0] = math.MinInt32
	var dfs func(int, int)
	dfs = func(i, idx int) {
		if i == n {
			if idx > 2 {
				ans = append(ans, slices.Clone(path[1:idx]))
			}
			return
		}
		if path[idx-1] <= nums[i] {
			path[idx] = nums[i]
			dfs(i+1, idx+1)
		}
		if path[idx-1] != nums[i] { // 剪枝：path[idx-1] == arr [i]：已经被上面选择了
			dfs(i+1, idx)
		}
	}
	dfs(0, 1)
	return ans

	// hash
	//const prime, val = 16777619, 101 // 100+1：为了区分 path 为空时，hash==0
	//n := len(nums)
	//ans, path := make([][]int, 0), make([]int, n+1)
	//memo := make(map[uint32]struct{})
	//path[0] = math.MinInt32 // 哨兵
	//var dfs func(int, int, uint32)
	//dfs = func(i, idx int, hash uint32) { // idx：path 的下一个填充位置
	//	if i == n {
	//		if _, ok := memo[hash]; !ok && idx > 2 { // hash 去重
	//			ans = append(ans, slices.Clone(path[1:idx]))
	//			memo[hash] = struct{}{}
	//		}
	//		return
	//	}
	//	dfs(i+1, idx, hash) // 不选
	//	hash *= prime
	//	if path[idx-1] <= nums[i] { // 非递减
	//		cur := hash + uint32(nums[i]+val)
	//		path[idx] = nums[i]
	//		dfs(i+1, idx+1, cur)
	//	}
	//}
	//dfs(0, 1, 0)
	//return ans
}

//func findSubsequences(nums []int) [][]int {
//	// 错误：比如 [1,2,2]，hash 记录了  [1,2] 是第二个 2，则无法再记录第一个 2 的 [1,2]，则没有 [1,2,2]
//	const prime, val = 16777619, 100
//	n := len(nums)
//	ans, path := make([][]int, 0), make([]int, n+1)
//	memo := make(map[uint32]struct{})
//	path[0] = math.MinInt32 // 哨兵
//	var dfs func(int, int, uint32)
//	dfs = func(i, idx int, hash uint32) { // idx：path 的下一个填充位置
//		if i == n {
//			if _, ok := memo[hash]; !ok && idx > 2 {
//				fmt.Println(path[1:idx], hash)
//				ans = append(ans, slices.Clone(path[1:idx]))
//				memo[hash] = struct{}{}
//			}
//			return
//		}
//		//if _, ok := memo[hash]; !ok {
//		//memo[hash] = struct{}{}
//		dfs(i+1, idx, hash) // 不选
//		//}
//		hash *= prime
//		if path[idx-1] <= nums[i] {
//			//fmt.Println(i, nums[i], hash, hash+uint32(nums[i]+val))
//			cur := hash + uint32(nums[i]+val)
//			//if _, ok := memo[cur]; !ok {
//			//memo[cur] = struct{}{}
//			path[idx] = nums[i]
//			//fmt.Println(i, cur, path[1:idx+1])
//			dfs(i+1, idx+1, cur)
//		}
//	}
//	//for i := range nums {
//	//	dfs(i, 1, 0)
//	//}
//	dfs(0, 1, 0)
//	fmt.Println(memo)
//	return ans
//}

//leetcode submit region end(Prohibit modification and deletion)
