//给你一个二进制字符串 binary ，它仅有 0 或者 1 组成。你可以使用下面的操作任意次对它进行修改：
//
//
// 操作 1 ：如果二进制串包含子字符串 "00" ，你可以用 "10" 将其替换。
//
//
//
// 比方说， "00010" -> "10010"
//
//
// 操作 2 ：如果二进制串包含子字符串 "10" ，你可以用 "01" 将其替换。
//
// 比方说， "00010" -> "00001"
//
//
//
//
// 请你返回执行上述操作任意次以后能得到的 最大二进制字符串 。如果二进制字符串 x 对应的十进制数字大于二进制字符串 y 对应的十进制数字，那么我们称二进制
//字符串 x 大于二进制字符串 y 。
//
//
//
// 示例 1：
//
//
//输入：binary = "000110"
//输出："111011"
//解释：一个可行的转换为：
//"000110" -> "000101"
//"000101" -> "100101"
//"100101" -> "110101"
//"110101" -> "110011"
//"110011" -> "111011"
//
//
// 示例 2：
//
//
//输入：binary = "01"
//输出："01"
//解释："01" 没办法进行任何转换。
//
//
//
//
// 提示：
//
//
// 1 <= binary.length <= 10⁵
// binary 仅包含 '0' 和 '1' 。
//
//
// Related Topics 贪心 字符串 👍 41 👎 0

package main

import (
	"fmt"
)

func main() {
	s := "000110"
	s = "01"
	s = "1100"
	//s = "01111001100000110010"
	//s = "11"
	//s = "101010111011001101000110010001100001111"
	binaryString := maximumBinaryString(s)
	fmt.Println(binaryString)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumBinaryString(binary string) string {
	// 贪心
	cnt, i, n := 0, 0, len(binary)
	buf := make([]byte, n)
	for ; i < n && binary[i] == '1'; i++ {
		buf[i] = '1'
	}
	k := i
	for ; i < n; i++ {
		if binary[i] == '0' {
			cnt++
		}
		buf[i] = '1'
	}
	if cnt > 0 {
		buf[k+cnt-1] = '0' // 排除前导 1 后，所有 0 移到头部
	}
	return string(buf)

	// 错误：01111001100000110010
	//buf, cnt, n := []byte(binary), 0, len(binary)
	//for i := 0; i < n; i++ {
	//	if buf[i] == '0' {
	//		if i > 0 && buf[i-1] == '0' {
	//			buf[i-1] = '1'
	//		} else {
	//			cnt++
	//		}
	//	}
	//}
	//if cnt > 1 {
	//	i := n - 1
	//	for ; cnt > 0; i-- {
	//		if buf[i] == '0' {
	//			buf[i] = '1'
	//			cnt--
	//		}
	//	}
	//	buf[i+2] = '0'
	//}
	//return string(buf)
}

//leetcode submit region end(Prohibit modification and deletion)
