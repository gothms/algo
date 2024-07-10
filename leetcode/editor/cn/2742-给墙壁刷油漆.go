package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func paintWalls(cost []int, time []int) int {
	n := len(cost)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1 // 没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int { // j 表示剩余需要的体积
		if j <= 0 { // 没有约束，后面什么也不用选了
			return 0
		}
		if i < 0 { // 此时 j>0，但没有物品可选，不合法
			return math.MaxInt / 2 // 防止加法溢出
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		*p = min(dfs(i-1, j-time[i]-1)+cost[i], dfs(i-1, j))
		return *p
	}
	return dfs(n-1, n)
}

//leetcode submit region end(Prohibit modification and deletion)
