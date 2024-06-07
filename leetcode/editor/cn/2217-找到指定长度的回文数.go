package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func kthPalindrome(queries []int, intLength int) []int64 {
	n := len(queries)
	base := int(math.Pow10((intLength - 1) >> 1))
	maxV := 9 * base // 最多有 maxV 个回文数
	ans := make([]int64, n)
	for i, q := range queries {
		if q > maxV { // 超过范围
			ans[i] = -1
			continue
		}
		v := base + q - 1 // 主要算法：先算出前半部分
		val := v
		if intLength&1 != 0 { // 偶数为 1xVVx1，奇数为 1xVx1
			val /= 10
		}
		for ; val != 0; val /= 10 { // 计算后半部分
			v = v*10 + val%10
		}
		ans[i] = int64(v)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
