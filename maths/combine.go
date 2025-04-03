package maths

import (
	"fmt"
	"math/bits"
	"sort"
)

/*
组合数恒等式
	C(n,k) = C(n-1,k-1) + C(n-1,k)

组合
	77：四种方案，主要体现思想的不同
	1079
组合总和
	39
	40
	216
	377：dp

隔板法：三种类型
	1.标准型：C(n-1, m-1)
		被分配的 n 个元素无差别
		n 个元素分给 m 个不同对象
		每个对象至少分一个元素
	2.多分型：C(n-(x-1)*m, m-1)
		被分配的 n 个元素无差别
		n 个元素分给 m 个不同对象
		每个对象至少分 x 个元素
	3.少分型：C(n+m-1, m-1)
		被分配的 n 个元素无差别
		n 个元素分给 m 个不同对象
		每个对象可以有任意个元素（包括 0）
	https://zhuanlan.zhihu.com/p/144997851

其他类似题型
	https://leetcode.cn/problems/permutations/solutions/9914/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/
*/

//const CombineN = 8
//
//var comb [CombineN][CombineN]int
//
//func init() { // 预处理组合数
//	for i := 0; i < CombineN; i++ {
//		comb[i][0], comb[i][i] = 1, 1
//		for j := 1; j < i; j++ {
//			comb[i][j] = comb[i-1][j-1] + comb[i-1][j] // 组合数恒等式
//		}
//	}
//	for _, v := range comb {
//		fmt.Println(v)
//	}
//}

func combineCnt(n int, k int) int {
	N := n
	for i := 2; i <= k; i++ {
		N = (n - i + 1) * N / i
	}
	return N
}

// Combine 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合
// 示例：n,k = 4,2，arr = [1,2,5]
// [1 3 5]
// [2 3 5]
// [1 4 5]
// [2 4 5]
// [3 4 5]
// [1 2 6]
func Combine(n int, k int) [][]int {
	ret := make([][]int, 0, combineCnt(n, k))
	arr := make([]int, k+1)
	arr[k] = n + 1 // 哨兵
	for i := 0; i < k; i++ {
		arr[i] = i + 1
	}
	for i := 0; i < k; {
		ret = append(ret, append([]int(nil), arr[:k]...))
		for i = 0; i < k && arr[i]+1 == arr[i+1]; i++ { // i 范围内已查找完
			arr[i] = i + 1 // 回到起点
		}
		arr[i]++ // 扩大范围
	}
	return ret
}

// CombineBit 位运算：2^n 个数的二进制位，标记了对应位置的元素是否被选择
// 换种思路：n 个二进制位，每位存在 选/不选
func CombineBit(n int, k int) [][]int {
	ret := make([][]int, 0, combineCnt(n, k))
	arr := make([]int, k)
	for i := uint(0); i < 1<<n; i++ {
		if bits.OnesCount(i) == k { // k 个元素
			v := i
			for j := 0; j < k; j++ {
				b := v & -v          // 取出最后一个 1
				v &^= b              // 去掉最后一个 1
				arr[j] = bits.Len(b) // 左移的位数 +1
			}
			ret = append(ret, append([]int(nil), arr...))
		}
	}
	return ret
}

// CombineDFS 深度优先搜索：枚举（下一个要选的数） + 回溯
func CombineDFS(n int, k int) [][]int {
	ret := make([][]int, 0, combineCnt(n, k))
	arr := make([]int, 0, k)
	var dfs func(int)
	dfs = func(i int) {
		if len(arr) == k {
			ret = append(ret, append([]int(nil), arr...))
			return
		}
		for last := len(arr); i <= n-(k-last)+1; i++ { // i 取值位 [1,n]，所以 +1
			arr = append(arr, i)
			dfs(i + 1)
			arr = arr[:last] // 回溯
		}
	}
	dfs(1)
	return ret
}

// CombineDFSChoose 深度优先搜索：选 / 不选
func CombineDFSChoose(n int, k int) [][]int {
	ret := make([][]int, 0, combineCnt(n, k))
	arr := make([]int, k)
	var dfs func(int, int)
	dfs = func(i, j int) {
		if j == k {
			ret = append(ret, append([]int(nil), arr...))
			return
		}
		if i > n-(k-j)+1 { // 剪枝
			return
		}
		dfs(i+1, j)   // 不选 i
		arr[j] = i    // 不用回溯：用 j 标识了当前的索引
		dfs(i+1, j+1) // 选 i
	}
	dfs(1, 0)
	return ret
}

// CombineDFSChooseBit 深度优先搜索：选 / 不选 + 位运算，耗时！
func CombineDFSChooseBit(n int, k int) [][]int {
	m := combineCnt(n, k)
	ret := make([][]int, 0, m)
	arr := make([]int, 0, m)
	var dfs func(int, int, int)
	dfs = func(i, v, cnt int) {
		if cnt == k {
			arr = append(arr, v)
			return
		}
		if i >= n { // i 标识左移位数
			return
		}
		dfs(i+1, v|1<<i, cnt+1) // 选择
		dfs(i+1, v, cnt)        // 不选择
	}
	dfs(0, 0, 0)
	temp := make([]int, k)
	for _, v := range arr {
		for i, idx := 1, 0; v > 0; i++ {
			if v&1 == 1 {
				temp[idx] = i
				idx++
			}
			v >>= 1
		}
		ret = append(ret, append([]int(nil), temp...))
	}
	return ret
}

// ====================组合：重复元素====================

// CombineWithDup 整数数组 nums 中可能包含重复元素，返回该数组所有可能的长度为 k 的组合
// 总数计算方式：也符合 C(m,n) = C(m,m-n)
// 示例：nums = [1 1 1 2 3 4 5 6 6], k=3
// C(6,3) + C(5,1) + C(3,3) + C(5,1) = 31
// [1,6] 中选 3 个，组合 [1,1,x]，组合 [1,1,1]，组合 [6,6,x]
func CombineWithDup(nums []int, k int) [][]int {
	sort.Ints(nums) // 有序
	n := len(nums)
	ret := make([][]int, 0)
	temp := make([]int, 0, k) // 复用 path
	var dfs func(int)
	cnt := 0
	dfs = func(i int) {
		if i == n {
			if len(temp) == k {
				ret = append(ret, append([]int(nil), temp...))
			}
			return
		}
		if len(temp) > k { // i==n 才添加，所以 ==k 放行
			return
		}
		cnt++
		temp = append(temp, nums[i])
		dfs(i + 1)
		temp = temp[:len(temp)-1]             // 回溯
		for i+1 < n && nums[i] == nums[i+1] { // 重复元素：上一个 dfs 已选择，此处全跳过
			i++
		}
		dfs(i + 1)
	}
	dfs(0)
	fmt.Println("CombineWithDup:", cnt)
	return ret
}

// CombineWithDupDFS 深度优先搜索：选 / 不选
// 优于 CombineWithDup
func CombineWithDupDFS(nums []int, k int) [][]int {
	sort.Ints(nums) // 有序
	n := len(nums)
	ret := make([][]int, 0)
	temp := make([]int, k) // 复用 path
	var dfs func(int, int)
	cnt := 0
	dfs = func(i, j int) { // j 记录了 temp 的索引
		if j == k {
			ret = append(ret, append([]int(nil), temp...))
			return
		}
		if i > n-(k-j) { // 剪枝
			return
		}
		cnt++
		temp[j] = nums[i] // 选
		dfs(i+1, j+1)
		for i+1 <= n-(k-j) && nums[i] == nums[i+1] { // 不选：跳过相同元素
			i++
		}
		dfs(i+1, j)
	}
	dfs(0, 0)
	fmt.Println("CombineWithDupDFS cnt:", cnt)
	return ret
}

// BFS：参考 subsets.go SubsetsAndCombineCounter
