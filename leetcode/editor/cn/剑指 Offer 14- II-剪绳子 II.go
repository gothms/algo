//给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m - 1]
// 。请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘
//积是18。
//
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
//
//
//
// 示例 1：
//
// 输入: 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1
//
// 示例 2:
//
// 输入: 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
//
//
//
// 提示：
//
//
// 2 <= n <= 1000
//
//
// 注意：本题与主站 343 题相同：https://leetcode-cn.com/problems/integer-break/
//
// Related Topics 数学 动态规划 👍 260 👎 0

package main

import "fmt"

func main() {
	//n := 120 // 953271190
	//n = 10   // 36
	//n = 5    // 6
	//n = 6    // 9
	//n = 7    // 12
	//n = 8    // 18
	//n = 9    // 27
	//rope := cuttingRope(n)
	//fmt.Println(rope)

	for i := 3; i < 90; i++ {
		rope := cuttingRope(i)
		t := test(i)
		if rope != t {
			fmt.Println(i, rope, t)
		}
	}
}

/*
思路：dp
	同 LeetCode-343
	说明
		dp 的思路和递推公式都是ok的
		只是每次都要 % 1e9+7，所以和贪心算法的结果会不一样
		而此题的结果是以贪心算法为标准，所以 dp 方案不能 AC
思路：贪心
	1.对于任意 n，讨论 n%3 的情况
		=0：全部剪为 3，乘积最大
		=1：全部剪为 3，剩 4 剪为 2 和 2，乘积最大
		=2：全部剪为 3，剩 2，乘积最大
	2.上述讨论建立在 n >= 4 的前提
		n=2：乘积为 1
		n=3：乘积为 2
	3.base 的 rem 次幂的技巧（在Rabin-Karp等算法中都有应用）
		for rem > 0 {
			if rem&1 > 0 {
				ret = ret * base % mod
			}
			base = base * base % mod
			rem >>= 1
		}
*/
// leetcode submit region begin(Prohibit modification and deletion)
func cuttingRope(n int) int {
	// 贪心
	if n < 4 {
		return n - 1 // 2和3的补丁
	}
	const mod = 1000000007
	ret, rem, base := 1, n/3-1, 3 // 初始值 rem = n/3-1
	for rem > 0 {                 // 幂运算技巧
		if rem&1 > 0 {
			ret = ret * base % mod
		}
		base = base * base % mod
		rem >>= 1
	}
	switch n % 3 {
	case 0:
		ret = ret * 3 % mod // *3
	case 1:
		ret = ret << 2 % mod // *2*2
	case 2:
		ret = ret * 6 % mod // *3*2
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
func test(n int) int {
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//const mod = 1000000007
	//dp3, dp2, dp1 := 0, 0, 1
	//for i := 3; i <= n; i++ {
	//	if i < 7 {
	//		dp1, dp2, dp3 = maxVal((i-2)<<1, (i-3)*3), dp1, dp2
	//	} else {
	//		dp1, dp2, dp3 = maxVal(dp2<<1%mod, dp3*3%mod), dp1, dp2
	//	}
	//	// 和上面写法是一样的
	//	//dp1, dp2, dp3 = maxVal(maxVal((i-2)<<1%mod, (i-3)*3%mod), maxVal(dp2<<1%mod, dp3*3%mod)), dp1, dp2
	//}
	//return dp1

	if n < 4 {
		return n - 1
	}
	const mod = 1000000007
	ret, rem, base := 1, n/3-1, 3
	for rem > 0 {
		if rem&1 > 0 {
			ret = ret * base % mod
		}
		base = base * base % mod
		rem >>= 1
	}
	switch n % 3 {
	case 0:
		ret = ret * 3 % mod
	case 1:
		ret = ret << 2 % mod
	case 2:
		ret = ret * 6 % mod
	}
	return ret
}
