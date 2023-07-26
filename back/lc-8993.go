package main

func specialPerm(nums []int) int {
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
		*p++ // -1 变 0
		for k := 0; k < n; k++ {
			if 1<<k&i > 0 && (nums[j]%nums[k] == 0 || nums[k]%nums[j] == 0) {
				*p = (*p + dfs(1<<k^i, k)) % mod
			} // 遍历每个没用过，且符合条件的数
		}
		memo[i][j] = *p // 记忆化
		return *p
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
