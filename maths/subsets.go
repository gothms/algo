package maths

import (
	"fmt"
	"slices"
	"sort"
)

/*
子集
	78
	90
	1079：子集 & 排列
		func SubsetsAndPermuteCounter(arr []int) int
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
		ret[i] = slices.Clone(arr) // O(n)：每次都要拷贝数组
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
			ret = append(ret, slices.Clone(arr))
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

// ====================子集 II：重复元素====================

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
		ret = append(ret, slices.Clone(temp))
		temp = temp[:0]
	}
	return ret
}

// SubsetsWithDupDFS 整数数组 arr 中可能包含重复元素，返回该数组所有可能的组合（即子集）
func SubsetsWithDupDFS(nums []int) [][]int {
	sort.Ints(nums) // 有序
	n := len(nums)
	ret := make([][]int, 0)
	temp := make([]int, 0, n) // 复用 path
	var dfs func(int)
	cnt := 0
	dfs = func(i int) {
		if i == n {
			ret = append(ret, slices.Clone(temp))
			return
		}
		cnt++
		temp = append(temp, nums[i]) // 选
		//ret = append(ret, append([]int(nil), temp...)) // 两处之一收集组合
		dfs(i + 1)
		temp = temp[:len(temp)-1]             // 回溯
		for i+1 < n && nums[i] == nums[i+1] { // 重复元素：第一个元素已选择，全跳过
			i++
		}
		dfs(i + 1)
	}
	dfs(0)
	fmt.Println("SubsetsAndCombine cnt:", cnt)
	return ret
}

// SubsetsAndCombineCounter 整数数组 arr 中可能包含重复元素，返回该数组所有可能的组合的总数
func SubsetsAndCombineCounter(arr []int) int {
	sort.Ints(arr) // 必须排序
	//dp, cur, n := 0, 0, len(arr) // i := 0 开始遍历，但需要判断 i > 0
	dp, cur, n := 1, 1, len(arr) // i := 1 开始遍历
	for i := 1; i < n; i++ {
		if arr[i] == arr[i-1] { // 相同：在 arr[i-1] 结尾的结果集后面追加 arr[i]
			dp += cur
		} else { // 不相同：在所有结果集后面追加 arr[i]，并添加单元素的组合 {arr[i]}
			cur = dp + 1 // 包含 arr[i] 组合的总数
			dp += dp + 1 // 所有组合的累加
		}
	}
	return dp

	// Hash 代替排序：适用于重复元素很多时
	//counts := make(map[int]int)
	//for _, v := range arr {
	//	counts[v]++
	//}
	//dp, cur := 0, 0
	//for _, cnt := range counts {
	//	cur = dp + 1
	//	dp += cur * cnt
	//}
	//return dp
}

// SubsetsAndCombineCounterK 整数数组 arr 中可能包含重复元素，返回该数组所有可能的长度为 k 的组合的总数
func SubsetsAndCombineCounterK(arr []int, k int) int {
	counts := make(map[int]int)
	for _, v := range arr {
		counts[v]++ // 统计每个字母的出现次数
	}
	dp := make([]int, k+1)
	dp[0] = 1                    // 构造空序列的方案数
	for _, cnt := range counts { // 枚举第 x 种字母
		for i := k; i > 0; i-- { // 枚举序列长度 i，但长度限制为 k（注意：j=0 表示选择 0 个 x 字母，已经在选择第 x-1 种字母时计算）
			for j := 1; j <= min(i, cnt); j++ { // 枚举第 x 种字母选了 j 个，最多选择 i / cnt 个
				dp[i] += dp[i-j] // 长度为 i，选出 j 个位置放置字母 x
			}
		}
	}
	fmt.Println(counts)
	return dp[k]
}

// ====================子集 & 排列：重复元素====================
// SubsetsAndPermute & SubsetsAndPermuteCounter 两者题意相同，所求不同
// 另参考 permute.go 的函数 PermuteUnique

// SubsetsAndPermute 整数数组 arr 中可能包含重复元素，返回该数组所有可能的子集（幂集）的排列
func SubsetsAndPermute(arr []int) [][]int {
	//now := time.Now()
	sort.Ints(arr)
	n := len(arr)
	ret := make([][]int, 0)
	temp := make([]int, 0, n)
	visited := make([]bool, n) // 标识某位是否已选择
	var dfs func()
	//cnt := 0
	dfs = func() {
		if len(temp) == n { // 所有不重复的全排列
			ret = append(ret, slices.Clone(temp))
			return
		}
		for i := 0; i < n; i++ {
			if visited[i] || i > 0 && !visited[i-1] && arr[i] == arr[i-1] {
				continue // 相同数字：第一次出现被标识选择了，后面才继续选择
			}
			//cnt++
			visited[i] = true
			temp = append(temp, arr[i])
			//ret = append(ret, slices.Clone(temp))
			dfs()
			visited[i] = false // 回溯
			temp = temp[:len(temp)-1]
		}
	}
	dfs()
	//fmt.Println(time.Since(now))
	//fmt.Println("SubsetsAndPermute cnt:", cnt)
	return ret
}

const CombineN = 11

var comb [CombineN][CombineN]int

func init() { // 预处理组合数
	for i := 0; i < CombineN; i++ {
		comb[i][0], comb[i][i] = 1, 1
		for j := 1; j < i; j++ {
			comb[i][j] = comb[i-1][j-1] + comb[i-1][j] // 组合数恒等式
		}
	}
	//for _, v := range comb {
	//	fmt.Println(v)
	//}
}

// SubsetsAndPermuteCounter 整数数组 arr 中可能包含重复元素，返回该数组所有可能的子集（幂集）的排列的总数
func SubsetsAndPermuteCounter(arr []int) int {
	counts := make(map[int]int)
	for _, v := range arr {
		counts[v]++ // 统计每个字母的出现次数
	}
	dp := make([]int, len(arr)+1)
	dp[0] = 1 // 构造空序列的方案数
	ret, sum := 0, 0
	for _, cnt := range counts { // 枚举第 x 种字母
		sum += cnt                 // 当前总长度
		for i := sum; i > 0; i-- { // 枚举序列长度 i（注意：j=0 表示选择 0 个 x 字母，已经在选择第 x-1 种字母时计算）
			for j := 1; j <= min(i, cnt); j++ { // 枚举第 x 种字母选了 j 个，最多选择 i / cnt 个
				dp[i] += dp[i-j] * comb[i][j] // 长度为 i，选出 j 个位置放置字母 x
			}
		}
	}
	for _, v := range dp[1:] {
		ret += v
	}
	return ret
}
