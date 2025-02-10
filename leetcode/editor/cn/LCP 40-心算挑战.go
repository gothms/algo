package main

import "slices"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxmiumScore(cards []int, cnt int) int {
	slices.SortFunc(cards, func(a, b int) int { return b - a })
	s := 0
	for _, v := range cards[:cnt] {
		s += v
	}
	if s%2 == 0 { // s 是偶数
		return s
	}

	replacedSum := func(x int) int {
		for _, v := range cards[cnt:] {
			if v%2 != x%2 { // 找到一个最大的奇偶性和 x 不同的数
				return s - x + v // 用 v 替换 s
			}
		}
		return 0
	}

	x := cards[cnt-1]
	ans := replacedSum(x)           // 替换 x
	for i := cnt - 2; i >= 0; i-- { // 前 cnt-1 个数
		if cards[i]%2 != x%2 { // 找到一个最小的奇偶性和 x 不同的数
			ans = max(ans, replacedSum(cards[i])) // 替换
			break
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
