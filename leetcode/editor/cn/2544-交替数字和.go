//给你一个正整数 n 。n 中的每一位数字都会按下述规则分配一个符号：
//
//
// 最高有效位 上的数字分配到 正 号。
// 剩余每位上数字的符号都与其相邻数字相反。
//
//
// 返回所有数字及其对应符号的和。
//
//
//
// 示例 1：
//
//
//输入：n = 521
//输出：4
//解释：(+5) + (-2) + (+1) = 4
//
// 示例 2：
//
//
//输入：n = 111
//输出：1
//解释：(+1) + (-1) + (+1) = 1
//
//
// 示例 3：
//
//
//输入：n = 886996
//输出：0
//解释：(+8) + (-8) + (+6) + (-9) + (+9) + (-6) = 0
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁹
//
//
//
//
// Related Topics 数学 👍 9 👎 0

package main

import "fmt"

func main() {
	n := 521
	//n = 1112
	sum := alternateDigitSum(n)
	fmt.Println(sum)
}

/*
思路：
	1.从个位开始往高位遍历，假设个位的的符号是 +，十位的符号是 -
	2.记录 n 有多少位
		奇数位：return v
		偶数位：return -v
*/
// leetcode submit region begin(Prohibit modification and deletion)
func alternateDigitSum(n int) int {
	v, i := 0, 0
	for ; n > 0; n /= 10 {
		if b := n % 10; i&1 > 0 {
			v -= b // 奇数位符号 -
		} else {
			v += b // 偶数位符号 +
		}
		i++ // 从个位开始计数
	}
	if i&1 > 0 {
		return v
	}
	return -v // 修正符号
}

//leetcode submit region end(Prohibit modification and deletion)
