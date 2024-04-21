package math

import (
	"slices"
)

/*
组合 & 排列 & 子集
元素不重复				组合			排列			子集
	长度为 k 的集合		√			√			同组合
	任意长度的集合			同子集		√			√
	长度为 k 的总数		C(n,k)		A(n,k)		C(n,k)
	任意长度的总数			2^n			n!+n!/1+...	2^n
元素重复					组合			排列			子集
	长度为 k 的集合		√			√			同组合
	任意长度的集合			同子集		同上			√
	长度为 k 的总数		同子集		同下			类左
	任意长度的总数			同子集		√			√
		SubsetsAndPermuteCounter：所有子集的全排列总数
		SubsetsAndCombineCounterK：长度为 k 的组合总数
		SubsetsAndCombineCounter：所有子集的总数
		以上 3 个函数都在 subsets.go

排列
	46
	47：同 1079
	60
	784

其他
	31
	60
	2514：费马小定理
*/

func permuteCnt(n int) int {
	N := 1
	for i := 2; i <= n; i++ {
		N *= i
	}
	return N
}

// ====================排列====================

// Permute 给定一个不含重复数字的数组 nums ，返回其所有可能的全排列
func Permute(nums []int) [][]int {
	n := len(nums)
	ret := make([][]int, 0, permuteCnt(n))
	n--
	var dfs func(int)
	dfs = func(i int) {
		if i == n { // i == n：最后一轮分支数为 1
			ret = append(ret, slices.Clone(nums))
			return
		}
		dfs(i + 1) // 分支：n*(n-1)*(n-2)*...*1
		for j := i + 1; j <= n; j++ {
			nums[i], nums[j] = nums[j], nums[i]
			dfs(i + 1)
			nums[i], nums[j] = nums[j], nums[i] // 回溯
		}
	}
	dfs(0)
	return ret
}

// PermuteDFS 枚举（下一个要选的数） + 回溯
func PermuteDFS(nums []int) [][]int {
	n := len(nums)
	ret := make([][]int, 0, permuteCnt(n))
	path, visited := make([]int, n), make([]bool, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ret = append(ret, slices.Clone(path))
			return
		}
		for j, b := range visited { // 寻找下一个要选的数
			if !b { // 可选择
				path[i] = nums[j]
				visited[j] = true // 标记已选择
				dfs(i + 1)
				visited[j] = false // 回溯
			}
		}
	}
	dfs(0)
	return ret
}

// PermuteK 给定一个不含重复数字的数组 nums ，返回其所有可能的长度为 k 的全排列
func PermuteK(nums []int, k int) [][]int {
	n := len(nums)
	ret := make([][]int, 0, func() int {
		N := 1
		for i := n - k + 1; i <= n; i++ {
			N *= i
		}
		return N
	}())
	var dfs func(int)
	dfs = func(i int) {
		if i == k {
			ret = append(ret, slices.Clone(nums[:k]))
			return
		}
		dfs(i + 1)
		//for j := i + 1; j < n; j++ {
		for j := min(i+1, k); j < n; j++ { // 剪枝：min(i+1, k)
			nums[i], nums[j] = nums[j], nums[i]
			dfs(i + 1)
			nums[i], nums[j] = nums[j], nums[i] // 回溯
		}
	}
	dfs(0)
	return ret
}

// PermuteAll 给定一个不含重复数字的数组 nums ，返回其所有可能的任意长度的全排列
func PermuteAll(nums []int) [][]int {
	n := len(nums)
	ret := make([][]int, 0, func() int {
		N := 1
		for i := 2; i <= n; i++ {
			N = N*i + i
		}
		return N
	}())
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			return
		}
		ret = append(ret, slices.Clone(nums[:i+1]))
		dfs(i + 1)
		for j := i + 1; j < n; j++ {
			nums[i], nums[j] = nums[j], nums[i]
			ret = append(ret, slices.Clone(nums[:i+1]))
			dfs(i + 1)
			nums[i], nums[j] = nums[j], nums[i] // 回溯
		}
	}
	dfs(0)
	return ret
}

// ====================排列：重复元素====================
// 另参考 subsets.go 的函数 SubsetsAndPermute，但 PermuteUnique 更佳

// PermuteUnique 可包含重复数字的序列 nums ，按任意顺序返回所有不重复的全排列
// 总数计算：
// 集合长度为 n，全排列个数为 n!
// 重复元素 a 有 x 个，则需除以 a 的全排列个数 x!
func PermuteUnique(nums []int) [][]int {
	//now := time.Now()
	n := len(nums)
	ret := make([][]int, 0)
	check := func(i, j int) bool {
		for ; i < j; i++ {
			if nums[i] == nums[j] { // 以 j 为目标
				return false
			}
		}
		return true
	}
	var dfs func(int)
	dfs = func(i int) {
		if i == n-1 {
			ret = append(ret, slices.Clone(nums))
			return
		}
		//memo := make(map[int]struct{}) // 也可以每次使用 map 记录元素：第一次出现时才操作
		//memo[nums[i]] = struct{}{}
		dfs(i + 1) // 分支：n*(n-1)*(n-2)*...*1
		for j := i + 1; j < n; j++ {
			// 示例：[0 1 2 1 3]
			// 0 和第一个 1 已经交换过，如果 0 和第二个 1 再交换，就出现了重复的排列
			// 但在操作 [1 0 2 1 3] 时，0 和第二个 1 是可以交换的
			// 效率：PermuteUnique > PermuteUniqueMemo > subsets.go SubsetsAndPermute
			if check(i, j) { // [i,j] 存在和 j 重复的元素：即 nums[j] 在区间内第一次出现时才交换
				nums[i], nums[j] = nums[j], nums[i]
				dfs(i + 1)
				nums[i], nums[j] = nums[j], nums[i] // 回溯
			}
		}
	}
	dfs(0)
	//fmt.Println(time.Since(now))
	return ret
}

// PermuteUniqueK 可包含重复数字的序列 nums ，按任意顺序返回所有不重复的长度为 k 的全排列
func PermuteUniqueK(nums []int, k int) [][]int {
	n := len(nums)
	ret := make([][]int, 0)
	check := func(i, j int) bool {
		for ; i < j; i++ {
			if nums[i] == nums[j] { // 以 j 为目标
				return false
			}
		}
		return true
	}
	var dfs func(int)
	dfs = func(i int) {
		if i == k {
			ret = append(ret, slices.Clone(nums[:k]))
			return
		}
		dfs(i + 1) // 分支：n*(n-1)*(n-2)*...*1
		//for j := i + 1; j < n; j++ {
		for j := min(i+1, k); j < n; j++ { // 剪枝：min(i+1, k)
			if check(i, j) { // [i,j] 存在和 j 重复的元素：即 nums[j] 在区间内第一次出现时才交换
				nums[i], nums[j] = nums[j], nums[i]
				dfs(i + 1)
				nums[i], nums[j] = nums[j], nums[i] // 回溯
			}
		}
	}
	dfs(0)
	return ret
}

//func PermuteUniqueAll(nums []int, k int) [][]int
