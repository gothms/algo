package main

import "fmt"

func main() {
	nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 1}
	target := 1
	nums = []int{1, 2, 1}
	target = 0
	sumWays := findTargetSumWays(nums, target)
	fmt.Println(sumWays)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findTargetSumWays(nums []int, target int) int {
	//s := 0
	//for _, v := range nums {
	//	s += v
	//}
	//s -= func(x int) int {
	//	if x < 0 {
	//		return -x
	//	}
	//	return x
	//}(target)
	//if s < 0 || s&1 != 0 {
	//	return 0
	//}
	//t := s >> 1
	//f := make([]int, t+1)
	//f[0] = 1
	//for _, v := range nums {
	//	for i := t; i >= v; i-- {
	//		f[i] += f[i-v]
	//	}
	//}
	//return f[t]

	// dp
	//s, n := 0, len(nums)
	//for _, v := range nums {
	//	s += v
	//}
	//s -= func(x int) int {
	//	if x < 0 {
	//		return -x
	//	}
	//	return x
	//}(target)
	//if s < 0 || s&1 != 0 {
	//	return 0
	//}
	//t := s >> 1
	//dp := make([][]int, n+1)
	//for i := range dp {
	//	dp[i] = make([]int, t+1)
	//}
	//dp[0][0] = 1
	//for i, v := range nums {
	//	for j := 0; j <= t; j++ {
	//		dp[i+1][j] = dp[i][j]
	//		if j >= v {
	//			dp[i+1][j] += dp[i][j-v]
	//		}
	//	}
	//}
	//return dp[n][t]

	// 记忆化搜索
	//s,n := 0,len(nums)
	//for _, v := range nums {
	//	s += v
	//}
	//s -= func(x int) int {
	//	if x < 0 {
	//		return -x
	//	}
	//	return x
	//}(target)
	//if s < 0 || s&1 != 0 {
	//	return 0
	//}
	//t := s >> 1
	//memo := make([][]int, n)
	//for i := range memo {
	//	memo[i] = make([]int, t+1)
	//	for j := range memo[i] {
	//		memo[i][j] = -1
	//	}
	//}
	//var dfs func(int, int) int
	//dfs = func(i, w int) int {
	//	if i < 0 {
	//		if w == 0 {
	//			return 1
	//		}
	//		return 0
	//	}
	//	v := &memo[i][w]
	//	if *v != -1 {
	//		return *v
	//	}
	//	res := dfs(i-1, w)
	//	if nums[i] <= w {
	//		res += dfs(i-1, w-nums[i])
	//	}
	//	*v = res
	//	return res
	//}
	//return dfs(n-1, t)

	// 回溯
	s, n := 0, len(nums)
	for _, v := range nums {
		s += v
	}
	s -= func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}(target)
	if s < 0 || s&1 != 0 {
		return 0
	}
	t := s >> 1
	cnt := 0
	var dfs func(int, int)
	dfs = func(i, cur int) {
		if i == n {
			if cur == t {
				cnt++
			}
			return
		}
		dfs(i+1, cur)
		if cur+nums[i] <= t {
			dfs(i+1, cur+nums[i])
		}
	}
	dfs(0, 0)
	return cnt
}

// leetcode submit region end(Prohibit modification and deletion)
func findTargetSumWays_(nums []int, target int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	s := 0
	for _, x := range nums {
		s += x
	}
	s -= abs(target)
	if s < 0 || s%2 == 1 {
		return 0
	}
	m := s / 2

	n := len(nums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, c int) (res int) {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		p := &memo[i][c]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if c < nums[i] {
			return dfs(i-1, c) // 只能不选
		}
		return dfs(i-1, c) + dfs(i-1, c-nums[i]) // 不选 + 选
	}
	return dfs(n-1, m)
}
