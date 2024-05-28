package main

import "fmt"

func main() {
	n := 5
	limit := 2
	//n = 3
	//limit = 3
	//n = 2
	//limit = 1
	//n = 4
	//limit = 1
	candies := distributeCandies(n, limit)
	fmt.Println(candies)
}

// leetcode submit region begin(Prohibit modification and deletion)
func distributeCandies(n int, limit int) int64 {

}

//leetcode submit region end(Prohibit modification and deletion)

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

func distributeCandies_(n int, limit int) int64 {
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
	//for i := 0; i <= n; i++ { // i+3 个糖，分 3 份
	//	dp[i] = comb2929[i+2][2]          // 容许分 0 个，所以 comb2929[i+2][2]
	//	for j := 1; i-j*limit >= 0; j++ { // j 个数容许分 0，分 3 份
	//		dp[i] -= comb2929[j+2][2] * dp[i-j*limit]
	//	}
	//}
	//return dp[n]
}
