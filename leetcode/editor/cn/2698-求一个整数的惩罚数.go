//给你一个正整数 n ，请你返回 n 的 惩罚数 。
//
// n 的 惩罚数 定义为所有满足以下条件 i 的数的平方和：
//
//
// 1 <= i <= n
// i * i 的十进制表示的字符串可以分割成若干连续子字符串，且这些子字符串对应的整数值之和等于 i 。
//
//
//
//
// 示例 1：
//
//
//输入：n = 10
//输出：182
//解释：总共有 3 个整数 i 满足要求：
//- 1 ，因为 1 * 1 = 1
//- 9 ，因为 9 * 9 = 81 ，且 81 可以分割成 8 + 1 。
//- 10 ，因为 10 * 10 = 100 ，且 100 可以分割成 10 + 0 。
//因此，10 的惩罚数为 1 + 81 + 100 = 182
//
//
// 示例 2：
//
//
//输入：n = 37
//输出：1478
//解释：总共有 4 个整数 i 满足要求：
//- 1 ，因为 1 * 1 = 1
//- 9 ，因为 9 * 9 = 81 ，且 81 可以分割成 8 + 1 。
//- 10 ，因为 10 * 10 = 100 ，且 100 可以分割成 10 + 0 。
//- 36 ，因为 36 * 36 = 1296 ，且 1296 可以分割成 1 + 29 + 6 。
//因此，37 的惩罚数为 1 + 81 + 100 + 1296 = 1478
//
//
//
//
// 提示：
//
//
// 1 <= n <= 1000
//
//
// Related Topics 数学 回溯 👍 28 👎 0

package main

import "fmt"

func main() {
	fmt.Printf("%b, %b %b %b\n", 1296, 1, 29, 6)
	fmt.Println(memo)
}

// leetcode submit region begin(Prohibit modification and deletion)
func punishmentNumber(n int) int {
	return memo[n]
}

var memo [1001]int

func init() {
	var dfs func(int, int, int) bool
	dfs = func(i, t, sum int) bool {
		if sum+t == i { // sum + t
			return true
		}
		for d, s := 10, t; s > 0; d *= 10 {
			if v := sum + t%d; v > i {
				break
			} else if dfs(i, t/d, v) { // t/10,t/100,t/1000 ...
				return true
			}
			s /= 10 // 标记是否继续循环
		}
		return false
	}
	for i, last := 1, 0; i <= 1000; i++ {
		if t := i * i; dfs(i, t, 0) {
			memo[i], last = last+t, last+t
		} else {
			memo[i] = last // i 不满足
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
