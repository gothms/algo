package math

import "math/bits"

/*
组合
	77：四种方案，主要体现思想的不同
组合总和
	39
	40
	216
	377
*/

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

// CombineDFS 深度优先搜索：回溯
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
