//给你两个正整数 n 和 limit 。
//
// 请你将 n 颗糖果分给 3 位小朋友，确保没有任何小朋友得到超过 limit 颗糖果，请你返回满足此条件下的 总方案数 。
//
//
//
// 示例 1：
//
//
//输入：n = 5, limit = 2
//输出：3
//解释：总共有 3 种方法分配 5 颗糖果，且每位小朋友的糖果数不超过 2 ：(1, 2, 2) ，(2, 1, 2) 和 (2, 2, 1) 。
//
//
// 示例 2：
//
//
//输入：n = 3, limit = 3
//输出：10
//解释：总共有 10 种方法分配 3 颗糖果，且每位小朋友的糖果数不超过 3 ：(0, 0, 3) ，(0, 1, 2) ，(0, 2, 1) ，(0, 3,
// 0) ，(1, 0, 2) ，(1, 1, 1) ，(1, 2, 0) ，(2, 0, 1) ，(2, 1, 0) 和 (3, 0, 0) 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁶
// 1 <= limit <= 10⁶
//
//
// Related Topics 数学 组合数学 枚举 👍 13 👎 0

package main

import "fmt"

func main() {
	n := 5
	limit := 2
	n = 3
	limit = 3
	n = 2
	limit = 1
	n = 4
	limit = 1
	candies := distributeCandies(n, limit)
	fmt.Println(candies)
}

// leetcode submit region begin(Prohibit modification and deletion)
const n2929, m2929 = 1e6 + 3, 3

var comb2929 [n2929][m2929]int64

func init() {
	for i := 0; i < n2929; i++ {
		comb2929[i][0] = 1
		if i < m2929 {
			comb2929[i][i] = 1
		}
		for j := min(i, m2929-1); j >= 1; j-- {
			comb2929[i][j] = comb2929[i-1][j-1] + comb2929[i-1][j]
		}
	}
	//for i, c := range comb2929 {
	//	fmt.Println(i, c)
	//}
}

func distributeCandies(n int, limit int) int64 {
	// 容斥原理：公式 P(AUBUC）= P(A)+P(B)+P(C) - P(AB)-P(BC)-P(AC)+P(ABC)
	// 不合法：
	// 至少一个人分到 limit+1 个
	// 至少两个人分到 limit+1 个
	// 至少三个人分到 limit+1 个
	// 合法数 = 总数 - 不合法数
	// 合法数 = 总数 - (至少一 - 至少二 + 至少三)
	ret := comb2929[n+2][2]
	if n >= limit {
		ret -= 3 * comb2929[n-limit+1][2]
		if n > limit<<1 {
			ret += 3 * comb2929[n-limit<<1][2]
			if n > 3*limit+1 {
				ret -= comb2929[n-3*limit-1][2]
			}
		}
	}
	return ret

	// 隔板法 & 容斥原理
	//limit++ // 容许分 0 个，所以 limit+1 才超标
	//dp := make([]int64, n+1)
	//dp[0] = 1                 // 全分 0 个，所以默认为 dp[0] = 1
	//for i := 1; i <= n; i++ { // i+3 个糖，分 3 份
	//	dp[i] = comb2929[i+2][2]          // 容许分 0 个，所以 comb2929[i+2][2]
	//	for j := 1; i-j*limit >= 0; j++ { // j 个数容许分 0，分 3 份
	//		dp[i] -= comb2929[j+2][2] * dp[i-j*limit]
	//	}
	//}
	//return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
