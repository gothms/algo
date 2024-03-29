//给你一个整数 num ，返回 num 中能整除 num 的数位的数目。
//
// 如果满足 nums % val == 0 ，则认为整数 val 可以整除 nums 。
//
//
//
// 示例 1：
//
// 输入：num = 7
//输出：1
//解释：7 被自己整除，因此答案是 1 。
//
//
// 示例 2：
//
// 输入：num = 121
//输出：2
//解释：121 可以被 1 整除，但无法被 2 整除。由于 1 出现两次，所以返回 2 。
//
//
// 示例 3：
//
// 输入：num = 1248
//输出：4
//解释：1248 可以被它每一位上的数字整除，因此答案是 4 。
//
//
//
//
// 提示：
//
//
// 1 <= num <= 10⁹
// num 的数位中不含 0
//
//
// Related Topics 数学 👍 7 👎 0

package main

import "fmt"

func main() {
	num := 121
	digits := countDigits(num)
	fmt.Println(digits)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countDigits(num int) int {
	cnt := 0
	for v := num; v > 0; v /= 10 {
		if num%(v%10) == 0 {
			cnt++
		}
	}
	return cnt

	//cnt := 0
	//for n := num; n > 0; n /= 10 {
	//	if mod := n % 10; num%(mod) == 0 {
	//		cnt++
	//	}
	//}
	//return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
