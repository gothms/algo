//给你数字 k ，请你返回和为 k 的斐波那契数字的最少数目，其中，每个斐波那契数字都可以被使用多次。
//
// 斐波那契数字定义为：
//
//
// F1 = 1
// F2 = 1
// Fn = Fn-1 + Fn-2 ， 其中 n > 2 。
//
//
// 数据保证对于给定的 k ，一定能找到可行解。
//
//
//
// 示例 1：
//
// 输入：k = 7
//输出：2
//解释：斐波那契数字为：1，1，2，3，5，8，13，……
//对于 k = 7 ，我们可以得到 2 + 5 = 7 。
//
// 示例 2：
//
// 输入：k = 10
//输出：2
//解释：对于 k = 10 ，我们可以得到 2 + 8 = 10 。
//
//
// 示例 3：
//
// 输入：k = 19
//输出：3
//解释：对于 k = 19 ，我们可以得到 1 + 5 + 13 = 19 。
//
//
//
//
// 提示：
//
//
// 1 <= k <= 10^9
//
//
// Related Topics 贪心 数学 👍 151 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	k := 19
	numbers := findMinFibonacciNumbers(k)
	fmt.Println(numbers)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findMinFibonacciNumbers(k int) int {
	var dfs func(int, int) int
	dfs = func(n, v int) int {
		if v == 0 {
			return 0
		}
		idx := sort.Search(n, func(i int) bool {
			return fibMFN[i] > v // 查第一个大于 v 的元素
		})
		return dfs(idx, v-fibMFN[idx-1]) + 1 // 贪心：数目 +1
	}
	return dfs(len(fibMFN), k)
}

var fibMFN []int

func init() {
	const M = 1_000_000_000
	f, s := 1, 1
	for s <= M {
		fibMFN = append(fibMFN, s) // fib 元素集合
		f, s = s, f+s
	}
}

//leetcode submit region end(Prohibit modification and deletion)
