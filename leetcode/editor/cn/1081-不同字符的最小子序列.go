//返回 s 字典序最小的子序列，该子序列包含 s 的所有不同字符，且只包含一次。
//
//
//
// 示例 1：
//
//
//输入：s = "bcabc"
//输出："abc"
//
//
// 示例 2：
//
//
//输入：s = "cbacdcbc"
//输出："acdb"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 由小写英文字母组成
//
//
//
//
// 注意：该题与 316 https://leetcode.cn/problems/remove-duplicate-letters/ 相同
//
// Related Topics 栈 贪心 字符串 单调栈 👍 199 👎 0

package main

import "fmt"

func main() {
	s := "cbacdcbc" // acdb
	subsequence := smallestSubsequence(s)
	fmt.Println(subsequence)
}

// leetcode submit region begin(Prohibit modification and deletion)
func smallestSubsequence(s string) string {
	stack, in, cnt := make([]int32, 0, len(s)), [26]bool{}, [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	for _, c := range s {
		idx := c - 'a'
		if !in[idx] { // 没在栈中
			last := len(stack) - 1
			for last >= 0 && stack[last] > idx { // 可能有更小的字典序
				if cnt[stack[last]] == 0 { // 没有该字符了
					break
				}
				in[stack[last]] = false // 出栈：并在后面添加
				last--
			}
			stack = stack[:last+1]
			stack = append(stack, idx)
			in[idx] = true // 入栈
		}
		cnt[idx]-- // 次数 -1
	}
	for i := range stack { // 修正字符
		stack[i] += 'a'
	}
	return string(stack)
}

//leetcode submit region end(Prohibit modification and deletion)
