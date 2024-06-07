package maths

import (
	"math"
	"slices"
)

/*
lc
	491
*/

// ====================子序列====================

// SubsequenceHash 整数数组 arr 中可能包含重复元素，返回该数组所有可能的非递减子序列
// 元素可能是负值，要求子序列长度不能小于 k
func SubsequenceHash(arr []int, k int) [][]int {
	const prime, val = 16777619, 101 // 假设最小的负值为 -100，则 100+1 为了区分 path 为空时，hash==0
	n := len(arr)
	ans, path := make([][]int, 0), make([]int, n+1)
	memo := make(map[uint32]struct{})
	path[0] = math.MinInt32 // 哨兵
	var dfs func(int, int, uint32)
	dfs = func(i, idx int, hash uint32) { // idx：path 的下一个填充位置
		if i == n {
			if _, ok := memo[hash]; !ok && idx > k { // hash 去重
				ans = append(ans, slices.Clone(path[1:idx]))
				memo[hash] = struct{}{}
			}
			return
		}
		dfs(i+1, idx, hash)        // 不选
		if path[idx-1] <= arr[i] { // 非递减
			path[idx] = arr[i]
			dfs(i+1, idx+1, hash*prime+uint32(arr[i]+val))
		}
	}
	dfs(0, 1, 0)
	return ans
}

// Subsequence 整数数组 arr 中可能包含重复元素，返回该数组所有可能的非递减子序列
// 元素可能是负值，要求子序列长度不能小于 k
func Subsequence(arr []int, k int) [][]int {
	n := len(arr)
	ans, path := make([][]int, 0), make([]int, n+1)
	path[0] = math.MinInt32
	var dfs func(int, int)
	dfs = func(i, idx int) {
		if i == n {
			if idx > k {
				ans = append(ans, slices.Clone(path[1:idx]))
			}
			return
		}
		if path[idx-1] <= arr[i] {
			path[idx] = arr[i]
			dfs(i+1, idx+1)
		}
		if path[idx-1] != arr[i] { // 剪枝：path[idx-1] == arr [i]：已经被上面选择了
			dfs(i+1, idx)
		}
	}
	dfs(0, 1)
	return ans
}
