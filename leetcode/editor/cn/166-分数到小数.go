//给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以 字符串形式返回小数 。
//
// 如果小数部分为循环小数，则将循环的部分括在括号内。
//
// 如果存在多个答案，只需返回 任意一个 。
//
// 对于所有给定的输入，保证 答案字符串的长度小于 10⁴ 。
//
//
//
// 示例 1：
//
//
//输入：numerator = 1, denominator = 2
//输出："0.5"
//
//
// 示例 2：
//
//
//输入：numerator = 2, denominator = 1
//输出："2"
//
//
// 示例 3：
//
//
//输入：numerator = 4, denominator = 333
//输出："0.(012)"
//
//
//
//
// 提示：
//
//
// -2³¹ <= numerator, denominator <= 2³¹ - 1
// denominator != 0
//
//
// Related Topics 哈希表 数学 字符串 👍 474 👎 0

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num, deno := 4, 333
	num, deno = -50, -8
	num, deno = 7, -12 // "-0.58(3)"
	//num, deno = -12, 4 // "-0.58(3)"
	decimal := fractionToDecimal(num, deno)
	fmt.Println(decimal)
}

// leetcode submit region begin(Prohibit modification and deletion)
func fractionToDecimal(numerator int, denominator int) string {
	// 细节：正负判断（包含余数 == 0），循环标记
	absVal := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	v, rem := numerator/denominator, numerator%denominator
	if rem == 0 { // fast path
		return strconv.Itoa(v)
	}
	mark, start := numerator*denominator < 0 && v == 0, 0 // mark：正负判断
	path, memo := make([]uint8, 0), make(map[int]int)     // memo 从 1 开始记录，方便 dfs 终止条件判断
	var dfs func(int)
	dfs = func(v int) {
		v *= 10                    // 下一位
		if v == 0 || memo[v] > 0 { // v == 0 时，start = 0
			start = memo[v] // 标记循环
			return
		}
		for ; v < denominator; v *= 10 {
			memo[v] = len(path) + 1
			path = append(path, '0')
		}
		memo[v] = len(path) + 1                       // 标记位置
		path = append(path, uint8(v/denominator)+'0') // 记录商
		dfs(v % denominator)                          // 余数
	}
	denominator = absVal(denominator) // 被除数转为正
	dfs(absVal(rem))                  // 余数为负，则转为正
	var sb strings.Builder
	if mark {
		sb.WriteByte('-')
	}
	sb.WriteString(strconv.Itoa(v))
	sb.WriteByte('.')
	if start == 0 {
		return sb.String() + string(path)
	}
	return sb.String() + string(path[:start-1]) + "(" + string(path[start-1:]) + ")"
}

//leetcode submit region end(Prohibit modification and deletion)
