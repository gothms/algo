package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func specialPerm(nums []int) int {
	const mod = 1e9 + 7
	n := len(nums)
	m := 1 << n
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j, x := range nums {
			for k, y := range nums {
				if i>>k&1 > 0 && (x%y == 0 || y%x == 0) {
					dp[i][j] += dp[1<<k^i][k]
				}
			}
			dp[i][j] %= mod
		}
	}
	ans := 0
	for i := range nums {
		ans += dp[(m-1)^(1<<i)][i]
	}
	return ans % mod
}

//leetcode submit region end(Prohibit modification and deletion)

func specialPerm_(nums []int) int {
	// 状压DP
	// 记忆化
	const mod = 1e9 + 7
	n := len(nums)
	m := 1 << n
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1 // 初始值
		}
	}
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i == 0 { // 终止条件
			return 1
		}
		p := &memo[i][j]
		if *p >= 0 {
			return *p
		}
		ret := 0
		for k := 0; k < n; k++ {
			if 1<<k&i > 0 && (nums[j]%nums[k] == 0 || nums[k]%nums[j] == 0) {
				ret = (ret + dfs(1<<k^i, k)) % mod
			} // 遍历每个没用过，且符合条件的数
		}
		*p = ret // 记忆化
		return ret
	}
	nums = append(nums, 0) // 哨兵
	cnt := dfs(m-1, n)
	return cnt

	// 回溯：超时
	//nums = append(nums, 0)
	//const mod = 1e9 + 7
	//var dfs func([]int, int) int
	//dfs = func(arr []int, i int) int {
	//	if i == 0 {
	//		if arr[1]%arr[0] == 0 || arr[0]%arr[1] == 0 {
	//			return 1
	//		}
	//		return 0
	//	}
	//	cnt := 0
	//	for j := i; j >= 0; j-- {
	//		if arr[i+1]%arr[j] == 0 || arr[j]%arr[i+1] == 0 {
	//			arr[i], arr[j] = arr[j], arr[i]
	//			cnt = (cnt + dfs(arr, i-1)) % mod
	//			arr[i], arr[j] = arr[j], arr[i]
	//		}
	//	}
	//	return cnt
	//}
	//return dfs(nums, len(nums)-2)
}
