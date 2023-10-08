package math

import "sort"

/*
子集
	78
	90
*/

// ====================子集 I====================

// Subsets 枚举：迭代
// 整数数组 nums 中的元素 互不相同，返回该数组所有可能的子集（幂集）
func Subsets(nums []int) [][]int {
	n := len(nums)
	N := 1 << n // 2^n
	ret := make([][]int, N)
	for i := 1; i < N; i++ {
		temp := make([]int, 0, n) // 与其复用切片时每次拷贝，不如直接先创建切片
		for v, idx := i, 0; v > 0; v >>= 1 {
			if v&1 != 0 {
				temp = append(temp, nums[idx])
			}
			idx++
		}
		ret[i] = temp // 通过 i 标识索引
	}
	return ret
}

// SubsetsSlow slow 的原因是复用 arr
func SubsetsSlow(nums []int) [][]int {
	n := len(nums)
	N := 1 << n
	ret := make([][]int, N)
	arr := make([]int, 0, n)
	for i := 1; i < N; i++ {
		//for v := i; v > 0; v &= v - 1 { // 写法一
		//	idx := bits.TrailingZeros(uint(v & -v))
		//	arr = append(arr, nums[idx])
		//}
		for v, idx := i, 0; v > 0; v >>= 1 { // 写法二：由于 nums 长度很小，所以更快
			if v&1 != 0 {
				arr = append(arr, nums[idx])
			}
			idx++
		}
		ret[i] = append([]int(nil), arr...) // O(n)：每次都要拷贝数组
		arr = arr[:0]
	}
	return ret
}

// SubsetsDFS 枚举：递归
// 另外，每个元素存在 选/不选，构成满二叉树，则可用 前/中/后序 遍历来实现
func SubsetsDFS(nums []int) [][]int {
	n := len(nums)
	N := 1 << n // 2^n
	ret := make([][]int, 0, N)
	arr := make([]int, 0, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ret = append(ret, append([]int(nil), arr...))
			return
		}
		dfs(i + 1)
		arr = append(arr, nums[i])
		dfs(i + 1)
		arr = arr[:len(arr)-1] // 回溯
	}
	dfs(0)
	return ret
}

// ====================子集 II====================

// SubsetsWithDup 枚举：迭代
// 整数数组 nums 中可能包含重复元素，返回该数组所有可能的子集（幂集）
func SubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums) // 有序
	n := len(nums)
	N := 1 << n
	ret := make([][]int, 1)
	temp := make([]int, 0, n)
out:
	for i := 1; i < N; i++ {
		for j := 0; j < n; j++ {
			if i>>j&1 != 0 { // 选择元素
				if j > 0 && nums[j] == nums[j-1] && i>>(j-1)&1 == 0 { // 重复元素：前一个没选，则此轮位重复子集
					temp = temp[:0]
					continue out
				}
				temp = append(temp, nums[j])
			}
		}
		ret = append(ret, append([]int(nil), temp...))
		temp = temp[:0]
	}
	return ret
}

// SubsetsWithDupDFS 枚举：递归
func SubsetsWithDupDFS(nums []int) [][]int {
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
