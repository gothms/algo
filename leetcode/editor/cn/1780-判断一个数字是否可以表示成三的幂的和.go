//给你一个整数 n ，如果你可以将 n 表示成若干个不同的三的幂之和，请你返回 true ，否则请返回 false 。
//
// 对于一个整数 y ，如果存在整数 x 满足 y == 3ˣ ，我们称这个整数 y 是三的幂。
//
//
//
// 示例 1：
//
// 输入：n = 12
//输出：true
//解释：12 = 3¹ + 3²
//
//
// 示例 2：
//
// 输入：n = 91
//输出：true
//解释：91 = 3⁰ + 3² + 3⁴
//
//
// 示例 3：
//
// 输入：n = 21
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁷
//
//
// Related Topics 数学 👍 120 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func checkPowersOfThree(n int) bool {
	// 3 进制：模拟二进制
	// %3 == 0：该位有 0 个 3 的幂
	// %3 == 1：该位有 1 个 3 的幂
	for n > 0 && n%3 < 2 {
		n /= 3
	}
	return n == 0
}

//leetcode submit region end(Prohibit modification and deletion)
