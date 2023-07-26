//给你一个正整数 n ，找出满足下述条件的 中枢整数 x ：
//
//
// 1 和 x 之间的所有元素之和等于 x 和 n 之间所有元素之和。
//
//
// 返回中枢整数 x 。如果不存在中枢整数，则返回 -1 。题目保证对于给定的输入，至多存在一个中枢整数。
//
//
//
// 示例 1：
//
//
//输入：n = 8
//输出：6
//解释：6 是中枢整数，因为 1 + 2 + 3 + 4 + 5 + 6 = 6 + 7 + 8 = 21 。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//解释：1 是中枢整数，因为 1 = 1 。
//
//
// 示例 3：
//
//
//输入：n = 4
//输出：-1
//解释：可以证明不存在满足题目要求的整数。
//
//
//
// 提示：
//
//
// 1 <= n <= 1000
//
//
// Related Topics 数学 前缀和 👍 12 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 8
	integer := pivotInteger(n)
	fmt.Println(integer)
}

/*
思路：数学
	1.[1,n] 的和 sum = n*(n+1)/2
	2.假若存在中枢整数 i，则 i + sum = 2*i*(i+1)
		化简得：n*(n+1)/2 = i*i
*/
// leetcode submit region begin(Prohibit modification and deletion)
func pivotInteger(n int) int {
	sum := n * (n + 1) >> 1
	if i := int(math.Sqrt(float64(sum))); i*i == sum {
		return i
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
