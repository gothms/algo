//给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
//
// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。
//
//
//
// 示例 1：
//
//
//输入：num1 = "11", num2 = "123"
//输出："134"
//
//
// 示例 2：
//
//
//输入：num1 = "456", num2 = "77"
//输出："533"
//
//
// 示例 3：
//
//
//输入：num1 = "0", num2 = "0"
//输出："0"
//
//
//
//
//
//
// 提示：
//
//
// 1 <= num1.length, num2.length <= 10⁴
// num1 和num2 都只包含数字 0-9
// num1 和num2 都不包含任何前导零
//
//
// Related Topics 数学 字符串 模拟 👍 727 👎 0

package main

import "fmt"

func main() {
	fmt.Println('9') // 48 ~ 57
	num1 := "456"
	num2 := "77"
	strings := addStrings(num1, num2)
	fmt.Println(strings)
}

/*
思路：模拟
    1.从个位开始，倒序遍历处理每一位的和
	2.'0'=48 '1'=49，且最后需要修正最高位是否进位
*/
//leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n, m := len(num1)-1, len(num2)-1
	str, add := make([]uint8, maxVal(n, m)+2), uint8(0)
	for i := len(str) - 1; i > 0; i-- {
		if n >= 0 { // 处理nums1
			add += num1[n] - 48
			n--
		}
		if m >= 0 { // 处理nums2
			add += num2[m] - 48
			m--
		}
		str[i] = add%10 + 48 // 倒序第 i 位的值
		add /= 10            // 进位
	}
	if add > 0 { // 修正
		str[0] = 49
		return string(str)
	}
	return string(str[1:])
}

//leetcode submit region end(Prohibit modification and deletion)
