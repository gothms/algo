//给定一个整数，打印该整数的英文描述。
//
// 示例 1:
//
//
//输入: 123
//输出: "One Hundred Twenty Three"
//
//
// 示例 2:
//
//
//输入: 12345
//输出: "Twelve Thousand Three Hundred Forty Five"
//
// 示例 3:
//
//
//输入: 1234567
//输出: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
//
// 示例 4:
//
//
//输入: 1234567891
//输出: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven
//Thousand Eight Hundred Ninety One"
//
// 注意：本题与 273 题相同：https://leetcode-cn.com/problems/integer-to-english-words/
//
// Related Topics 递归 数学 字符串 👍 24 👎 0

package main

import (
	"fmt"
	"strings"
)

func main() {
	num := 12345
	words := numberToWords(num)
	fmt.Println(words)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numberToWords(num int) string {
	// 递归
	if num == 0 {
		return "Zero"
	}
	var (
		singles   = [10]string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
		teens     = [10]string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
		tens      = [10]string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
		thousands = [4]string{"", "Thousand", "Million", "Billion"}
	)
	var sb strings.Builder
	var dfs func(int)
	dfs = func(v int) {
		switch {
		case v == 0: // 终止
		case v < 10: // 小于 10
			sb.WriteString(singles[v])
			sb.WriteByte(' ')
		case v < 20: // [10,20)
			sb.WriteString(teens[v-10])
			sb.WriteByte(' ')
		case v < 100: // [20,100)
			sb.WriteString(tens[v/10])
			sb.WriteByte(' ')
			dfs(v % 10)
		default: // [100,1000)
			sb.WriteString(singles[v/100])
			sb.WriteString(" Hundred ")
			dfs(v % 100)
		}
	}
	for i, unit := 3, int(1e9); i >= 0; i-- { // 千的索引 & 除数
		if cur := num / unit; cur > 0 {
			num -= cur * unit // reduce
			dfs(cur)
			sb.WriteString(thousands[i]) // 千
			sb.WriteByte(' ')
		}
		unit /= 1000 // 三位为 dfs 数字
	}
	return strings.TrimSpace(sb.String())
}

//leetcode submit region end(Prohibit modification and deletion)
