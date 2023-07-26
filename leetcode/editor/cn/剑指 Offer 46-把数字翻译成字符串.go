//给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可
//能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
//
//
//
// 示例 1:
//
// 输入: 12258
//输出: 5
//解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
//
//
//
// 提示：
//
//
// 0 <= num < 2³¹
//
//
// Related Topics 字符串 动态规划 👍 581 👎 0

package main

import "fmt"

func main() {
	num := 12258
	num = 25
	i := translateNum(num)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func translateNum(num int) int {
	// dp
	h, l := 1, 1 // 预处理 个位
	for curr, last := 0, num%10; num > 0; last = curr {
		num /= 10
		if curr = num % 10; curr == 1 || curr == 2 && last <= 5 {
			h, l = h+l, h // curr 和 last 能合并
		} else {
			l = h // 不能合并
		}
	}
	return h
}

//leetcode submit region end(Prohibit modification and deletion)
