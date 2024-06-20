package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 25
	n = 99  // 20
	n = 100 // 20
	inRange := numberOf2sInRange(n)
	fmt.Println(inRange)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOf2sInRange(n int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func numberOf2sInRange_(n int) int {
	// 数位 dp
	// https://leetcode.cn/problems/number-of-2s-in-range-lcci/solutions/1750395/by-endlesscheng-x4mf/
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var dfs func(int, int, bool) int
	dfs = func(i, cnt int, isLimit bool) int {
		if i == m {
			return cnt
		}
		ret := 0
		if !isLimit { // 重要
			v := &dp[i][cnt]
			if *v >= 0 {
				return *v
			}
			defer func() { *v = ret }()
		}
		up := 9
		if isLimit { // 重要
			up = int(s[i] - '0')
		}
		for d := 0; d <= up; d++ {
			c := cnt
			if d == 2 { // 重要
				c++
			}
			ret += dfs(i+1, c, isLimit && d == up)
		}
		return ret
	}
	return dfs(0, 0, true)
}
