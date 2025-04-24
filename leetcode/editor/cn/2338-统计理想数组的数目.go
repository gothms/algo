package main

import "fmt"

func main() {
	n := 5
	maxValue := 3 // 11
	n = 5
	maxValue = 4 // 26
	n = 5
	maxValue = 6 // 56
	maxValue = 7 // 61
	maxValue = 8 // 96
	maxValue = 9 // 111
	//maxValue = 10 // 136
	//maxValue = 1e4 // 6313916
	n = 184
	maxValue = 389 // 510488787
	arrays := idealArrays(n, maxValue)
	fmt.Println(arrays)
}

// leetcode submit region begin(Prohibit modification and deletion)
//const mod = 1_000_000_007
//const maxN = 10_000
//const maxE = 13
//
//var exp [maxN + 1][]int
//var c [maxN + maxE][maxE + 1]int
//
//func init() {
//	// exp[x] 为 x 分解质因数后，每个质因数的指数
//	for x := 2; x <= maxN; x++ {
//		t := x
//		for i := 2; i*i <= t; i++ {
//			e := 0
//			for ; t%i == 0; t /= i {
//				e++
//			}
//			if e > 0 {
//				exp[x] = append(exp[x], e)
//			}
//		}
//		if t > 1 {
//			exp[x] = append(exp[x], 1)
//		}
//	}
//
//	// 预处理组合数
//	for i := range c {
//		c[i][0] = 1
//		for j := 1; j <= min(i, maxE); j++ {
//			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % mod
//		}
//	}
//}

func idealArrays(n int, maxValue int) int {
	// 商分
	// 组合
	// 分解质因子
	// 模运算：https://leetcode.cn/discuss/post/3584387/fen-xiang-gun-mo-yun-suan-de-shi-jie-dan-7xgu/
	//ans := 0
	//for i := 1; i <= maxValue; i++ {
	//	ret := 1
	//	for _, k := range exp[i] {
	//		ret = ret * c[n+k-1][k] % mod
	//	}
	//	ans = (ans + ret) % mod
	//}
	//return ans

	// 超时
	//dp := make([]int, maxValue+1)
	//for i := 1; i <= maxValue; i++ {
	//	dp[i] = 1
	//}
	//for k := 2; k <= n; k++ {
	//	for i := maxValue >> 1; i >= 1; i-- {
	//		for j := i << 1; j <= maxValue; j += i {
	//			dp[j] = (dp[j] + dp[i]) % mod
	//		}
	//	}
	//}
	//ans := 0
	//for _, v := range dp {
	//	ans = (ans + v) % mod
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
