package main

import (
	"fmt"
)

func main() {
	s := "010101" // 9
	//s = "000000001"
	//s = "0011"
	cost := minimumCost(s)
	fmt.Println(cost)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumCost(s string) int64 {
	// dp

	// 迭代
	n := len(s)
	ans := 0
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			ans += min(i, n-i)
		}
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

func minimumCost_(s string) int64 {
	// 终版
	// 1.如果 s[i−1]≠s[i]，那么必须反转，不然没法都相等
	// 要么翻转 s[0] 到 s[i−1]，成本为 i
	// 要么翻转 s[i] 到 s[n−1]，成本为 n−i
	// 2.从左到右遍历，反转后：
	// s[i] 及其左边的字符，都已经相等了
	// s[i] 右边的每对相邻字符，反转前不同的，反转后仍然不同。所以要继续反转
	//ans, n := 0, len(s)
	//for i := 1; i < n; i++ {
	//	if s[i] != s[i-1] {
	//		ans += min(i, n-i)
	//	}
	//}
	//return int64(ans)

	// 如何理解？
	// 结论：只是在终版的基础上画蛇添足，如果要 0 1 讨论，那不应该这么写。不如直接理解终版
	//zero, one, n := int(s[0]^'0'), int(s[0]^'1'), len(s)
	zero, one, n := 0, 0, len(s)
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			if s[i] == '0' {
				one++ // 无论往前或往后反转，都多反转一个字符？错
			} else {
				zero++
			}
		} else {
			zero, one = one+min(i, n-i), zero+min(i, n-i) // 重要
		}
	}
	return int64(min(zero, one))
}
