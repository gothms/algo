//n 位格雷码序列 是一个由 2ⁿ 个整数组成的序列，其中：
//
//
// 每个整数都在范围 [0, 2ⁿ - 1] 内（含 0 和 2ⁿ - 1）
// 第一个整数是 0
// 一个整数在序列中出现 不超过一次
// 每对 相邻 整数的二进制表示 恰好一位不同 ，且
// 第一个 和 最后一个 整数的二进制表示 恰好一位不同
//
//
// 给你一个整数 n ，返回任一有效的 n 位格雷码序列 。
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出：[0,1,3,2]
//解释：
//[0,1,3,2] 的二进制表示是 [00,01,11,10] 。
//- 00 和 01 有一位不同
//- 01 和 11 有一位不同
//- 11 和 10 有一位不同
//- 10 和 00 有一位不同
//[0,2,3,1] 也是一个有效的格雷码序列，其二进制表示是 [00,10,11,01] 。
//- 00 和 10 有一位不同
//- 10 和 11 有一位不同
//- 11 和 01 有一位不同
//- 01 和 00 有一位不同
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[0,1]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 16
//
//
// Related Topics 位运算 数学 回溯 👍 648 👎 0

package main

import "fmt"

func main() {
	n := 3
	code := grayCode(n)
	fmt.Println(code)
}

/*
思路：
	n = 0	1	2	3	...
		0	0	00	000
			1	01	001
				11	011
				10	010
					110
					111
					101
					100
*/
// leetcode submit region begin(Prohibit modification and deletion)
func grayCode(n int) []int {
	ret := make([]int, 1, 1<<(n+1))    // n = 0
	for i, head := 1, 1; i <= n; i++ { // 1-n
		for j := head - 1; j >= 0; j-- { // 逆序：长度为 head - 1
			ret = append(ret, head+ret[j]) // 头部添 1
		}
		head <<= 1
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
