//给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。
//
// 例如：
//
//
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28
//...
//
//
//
//
// 示例 1：
//
//
//输入：columnNumber = 1
//输出："A"
//
//
// 示例 2：
//
//
//输入：columnNumber = 28
//输出："AB"
//
//
// 示例 3：
//
//
//输入：columnNumber = 701
//输出："ZY"
//
//
// 示例 4：
//
//
//输入：columnNumber = 2147483647
//输出："FXSHRXW"
//
//
//
//
// 提示：
//
//
// 1 <= columnNumber <= 2³¹ - 1
//
//
// Related Topics 数学 字符串 👍 652 👎 0

package main

import (
	"fmt"
)

func main() {
	fmt.Println('A' - 1)
	columnNumber := 2147483647
	title := convertToTitle(columnNumber)
	fmt.Println(title)
}

// leetcode submit region begin(Prohibit modification and deletion)
func convertToTitle(columnNumber int) string {
	var buf []byte
	for ; columnNumber > 0; columnNumber = (columnNumber - 1) / 26 {
		buf = append(buf, 'A'+byte((columnNumber-1)%26))
	}
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

//leetcode submit region end(Prohibit modification and deletion)
