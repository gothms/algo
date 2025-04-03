package main

import (
	"fmt"
	"math"
)

func main() {
	n := 100
	drop := twoEggDrop(n)
	fmt.Println(drop)
}

// leetcode submit region begin(Prohibit modification and deletion)
func twoEggDrop(n int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func twoEggDrop_(n int) int {
	// dp
	return tedDP[n]

	// 记忆化搜索
	//k := 2 // k 个鸡蛋
	//memo := make([][]int, n+1)
	//for i := 0; i <= n; i++ {
	//	memo[i] = make([]int, k+1)
	//}
	//var dfs func(int, int) int
	//dfs = func(n, k int) int {
	//	if n == 0 || k == 1 { // 到顶层 / 最后一个蛋
	//		return n
	//	}
	//	v := &memo[n][k]
	//	if *v > 0 {
	//		return *v
	//	}
	//	ret := math.MaxInt32
	//	for i := 1; i <= n; i++ {
	//		ret = min(ret, 1+max(dfs(i-1, k-1), dfs(n-i, k))) // 碎 / 没碎
	//	}
	//	*v = ret
	//	return ret
	//}
	//return dfs(n, k)

	// 数学
	// 让公式 1 + 2 + 3 + 4 + .... + x >= n 成立的最小的x，就是 f 的最小次数
	//return int(math.Ceil((math.Sqrt(float64(n*8+1)) - 1.0) / 2.0))
}

var tedDP []int

func init() {
	const N = 1000
	tedDP = make([]int, N+1)
	//tedDP[0], tedDP[1], tedDP[2] = 0, 1, 2
	//for i := 3; i <= N; i++ {
	tedDP[1] = 1
	for i := 2; i <= N; i++ {
		tedDP[i] = math.MaxInt32
		for j := 1; j < i; j++ {
			tedDP[i] = min(tedDP[i], max(tedDP[i-j]+1, j)) // 没碎 / 碎
		}
	}
}
