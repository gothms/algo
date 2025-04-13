package main

import (
	"fmt"
	"strconv"
)

func main() {
	low := fmt.Sprintf("%0*d", 6, 12345)
	fmt.Println(low)

	var start, finish int64 = 1829505, 1255574165
	limit := 7
	s := "11223" // 5470
	powerfulInt := numberOfPowerfulInt(start, finish, limit, s)
	fmt.Println(powerfulInt)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	// 数位 dp
	high := strconv.FormatInt(finish, 10) // strconv.FormatInt
	n := len(high)
	low := fmt.Sprintf("%0*d", n, start) // fmt.Sprintf("%0*d", n, start)
	dif := n - len(s)

	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int, bool, bool) int
	dfs = func(i int, lowLimit, highLimit bool) int {
		if i == n {
			return 1
		}
		if !lowLimit && !highLimit && memo[i] >= 0 { // 记忆化的条件
			return memo[i]
		}
		ret := 0
		f, t := 0, 9
		if lowLimit { // 是 start 的前缀
			f = int(low[i] - '0')
		}
		if highLimit {
			t = int(high[i] - '0') // 是 finish 的前缀
		}
		if i < dif {
			for j := f; j <= min(t, limit); j++ {
				ret += dfs(i+1, lowLimit && j == f, highLimit && j == t)
			}
		} else {
			v := int(s[i-dif] - '0') // 指定选择 s 的某一位
			if f <= v && v <= t {
				ret += dfs(i+1, lowLimit && v == f, highLimit && v == t)
			}
		}
		if !lowLimit && !highLimit { // 记忆化的条件
			memo[i] = ret
		}
		return ret
	}
	return int64(dfs(0, true, true))
}

//leetcode submit region end(Prohibit modification and deletion)
