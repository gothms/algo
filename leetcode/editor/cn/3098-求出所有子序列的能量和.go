package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4} // 4
	k := 3
	arr = []int{1, 2, 3, 4, 8, 2, 5, 4, 6} // 100
	k = 3
	arr = []int{-100000000, 100000000}
	k = 2 // 200000000
	//// [-999 -977 -959 -928 -921 -877 -855 -846 -653 -628 -599 -592 -423 -295 -294 -291 -119 -79 -65 -24 -10 44 67 85 119 274 349 372 554 759 826 833 857 962]
	//arr = []int{-24, -921, 119, -291, -65, -628, 372, 274, 962, -592, -10, 67, -977, 85, -294, 349, -119, -846, -959, -79, -877, 833, 857, 44, 826, -295, -855, 554, -999, 759, -653, -423, -599, -928}
	//k = 19 // 990202285
	//arr = []int{6, 738018, 1490730, 2258772, 3042408, 3842244, 4658364, 5491344, 6341316, 7208862, 8094180, 8997648, 9919602, 10860426, 11820252, 12799560, 13798482, 14817654, 15857328, 16917948, 17999964, 19103508, 20228772, 21376368, 22546968, 23740842, 24958596, 26200638, 27467316, 28759188, 30076824, 31420794, 32791794, 34190520, 35617296, 37072716, 38557308, 40071558, 41615832, 43190712, 44796402, 46433112, 48101148, 49800642, 51532176, 53295876, 55092324, 56921802, 58784718, 60681360}
	//k = 28 // 78733672
	powers := sumOfPowers(arr, k)
	fmt.Println(powers)
}

// leetcode submit region begin(Prohibit modification and deletion)
func sumOfPowers(nums []int, k int) int {
	// dp
	const mod = 1_000_000_007
	n := len(nums)
	sort.Ints(nums) // 排序
	dp := make([][]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]map[int]int, k+1) // 索引 i，子序列长度 l，最小值 d，数量 c
		for j := 1; j <= min(i+1, k); j++ {
			dp[i][j] = make(map[int]int)
		}
		dp[i][1][mod] = 1
		for l := 2; l <= min(i+1, k); l++ {
			for j := l - 2; j < i; j++ {
				cur := nums[i] - nums[j]
				for d, c := range dp[j][l-1] {
					dp[i][l][min(d, cur)] += c // 状态转移方程
				}
			}
		}
	}
	ans := 0
	for _, dk := range dp[k-1:] { // 统计结果
		for d, c := range dk[k] {
			ans = (ans + c%mod*d) % mod
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func sumOfPowers_(nums []int, k int) int {
	//	// dp
	//	const mod = 1_000_000_007
	//	sort.Ints(nums)
	//	ret, n := 0, len(nums)
	//	dp := make([][]map[int]int, n) // 索引 i，子序列长度 kk，最小值 v，数量 c
	//	for i := 1; i < n; i++ {       // 索引 i：不精简时，i 从 1 开始
	//		//for i := 0; i < n; i++ {       // 索引 i：不精简时，i 从 1 开始
	//		dp[i] = make([]map[int]int, k+1)
	//		for j := 1; j <= k; j++ {
	//			dp[i][j] = make(map[int]int)
	//		}
	//		for j := 0; j < i; j++ { // 长度为 2
	//			dp[i][2][nums[i]-nums[j]]++
	//		}
	//		for kk := 3; kk <= k; kk++ { // 目标子序列长度 kk+1，从 3 开始
	//			for j := kk - 2; j < i; j++ { // 索引 j
	//				for v, c := range dp[j][kk-1] {
	//					mi := min(v, nums[i]-nums[j])
	//					dp[i][kk][mi] += c
	//				}
	//			}
	//		}
	//
	//		// 精简上面的代码
	//		//dp[i][1][mod] = 1
	//		//for kk := 2; kk <= k; kk++ { // 目标子序列长度 kk+1，从 3 开始
	//		//	for j := kk - 2; j < i; j++ { // 索引 j
	//		//		for v, c := range dp[j][kk-1] {
	//		//			mi := min(v, nums[i]-nums[j])
	//		//			dp[i][kk][mi] += c
	//		//		}
	//		//	}
	//		//}
	//	}
	//	for i := k - 1; i < n; i++ {
	//		for v, c := range dp[i][k] {
	//			ret = (ret + c%mod*v) % mod // c 可能很大，所以先 mod
	//		}
	//	}
	//	return ret

	// 记忆化搜索
	sort.Ints(nums)
	const mod = 1_000_000_007
	n := len(nums)
	//memo := make([][]map[int]int, n)	// 3 维定位：i、c、mi 无法准确定位
	memo := make([][]map[int]map[int]int, n) // 4 维定位：当前索引 i，子序列元素个数 c，上一个元素值 last，最小差值 mi
	for i := 0; i < n; i++ {
		memo[i] = make([]map[int]map[int]int, k+1)
	}
	var dfs func(int, int, int, int) int
	dfs = func(i, c, last, mi int) int {
		if n-i < c { // 不够
			return 0
		}
		if c == 0 { // 找到一个子序列
			return mi
		}
		if memo[i][c] == nil {
			memo[i][c] = make(map[int]map[int]int)
			memo[i][c][last] = make(map[int]int)
		}
		if memo[i][c][last] == nil {
			memo[i][c][last] = make(map[int]int)
		}
		ret, ok := memo[i][c][last][mi]
		if !ok {
			ret = dfs(i+1, c-1, nums[i], min(mi, nums[i]-last)) // 选择：找子序列
			ret = (ret + dfs(i+1, c, last, mi)) % mod           // 不选择
			memo[i][c][last][mi] = ret                          // 记忆化
		}
		return ret
	}
	return dfs(0, k, -mod, mod) // 上一个元素 -mod
}

// 竞赛：
//func sumOfPowers(nums []int, k int) int {
//	const (
//		MOD int = 1e9 + 7
//		INF int = 1e18
//	)
//	n := len(nums)
//
//	sort.Ints(nums)
//	dp := make([][]map[int]map[int]int, n+10)
//	for i := range dp {
//		dp[i] = make([]map[int]map[int]int, k+10)
//	}
//	var dfs func(i, rem, lst, mn int) int
//	dfs = func(i, rem, lst, mn int) int {
//		if n-i < rem {
//			return 0
//		}
//		if rem == 0 {
//			return mn
//		}
//		if dp[i][rem] == nil {
//			dp[i][rem] = map[int]map[int]int{}
//			dp[i][rem][lst] = map[int]int{}
//		}
//		if dp[i][rem][lst] == nil {
//			dp[i][rem][lst] = map[int]int{}
//		}
//		if res, ok := dp[i][rem][lst][mn]; ok {
//			return res
//		}
//		res := 0
//		res = (res + dfs(i+1, rem-1, nums[i], min(mn, nums[i]-lst))) % MOD
//		res = (res + dfs(i+1, rem, lst, mn)) % MOD
//		dp[i][rem][lst][mn] = res
//		return res
//	}
//
//	ans := dfs(0, k, -INF, INF)
//	return ans
//}

// 读错题意：任意两个元素差值绝对值的最大值
//const mod = 1_000_000_009
//const n3098 = 50 - 2 + 1
//
//var m3098 [][]int
//
//func init() {
//	m3098 = make([][]int, n3098)
//	m3098[1] = make([]int, n3098)
//	m3098[1][0], m3098[1][1] = 1, 1
//	for i := 2; i < n3098; i++ {
//		m3098[i] = make([]int, n3098)
//		m3098[i][0], m3098[i][i] = 1, 1
//		for j := 1; j <= i>>1; j++ {
//			v := m3098[i][j-1] * (i - j + 1) / j % mod
//			m3098[i][j], m3098[i][i-j] = v, v
//		}
//	}
//	//for i, m := range m3098 {
//	//	fmt.Println(i, m)
//	//}
//}
//func sumOfPowers(nums []int, k int) int {
//	sort.Ints(nums)
//	ret, m, n := 0, k-2, len(nums)
//	for i := n - k; i >= 0; i-- {
//		for j := i + k - 1; j < n; j++ {
//			fmt.Println(i, j, (nums[j]-nums[i])*m3098[j-i-1][m])
//			ret = (ret + (nums[j]-nums[i])*m3098[j-i-1][m]%mod) % mod
//		}
//	}
//	return ret
//}
