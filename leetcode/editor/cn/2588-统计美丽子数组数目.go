package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func beautifulSubarrays(nums []int) int64 {
	// 优化
	memo := make(map[int]int)
	memo[0] = 1
	var ans int64
	xor := 0
	for _, v := range nums {
		xor ^= v
		ans += int64(memo[xor])
		memo[xor]++
	}
	return ans

	//n := len(nums)
	//memo := make(map[int]int)
	//memo[0] = 1
	//dp := make([]int64, n+1)
	//xor := 0
	//for i, v := range nums {
	//	xor ^= v
	//	dp[i+1] = dp[i] + int64(memo[xor])
	//	memo[xor]++
	//}
	//return dp[n]

	// 非连续的子数组
	//n := len(nums)
	//var ans int64
	//path := make([]int, 0, n)
	//var dfs func(int, int)
	//dfs = func(i, v int) {
	//	if i == n {
	//		if v == 0 {
	//			ans++
	//		}
	//		return
	//	}
	//	dfs(i+1, v)
	//	v ^= nums[i]
	//	path = append(path, i)
	//	dfs(i+1, v)
	//	v ^= nums[i]
	//	path = path[:len(path)-1]
	//}
	//dfs(0, 0)
	//return ans - 1
}

//leetcode submit region end(Prohibit modification and deletion)
