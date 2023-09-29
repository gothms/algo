//给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
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
// 1 <= s.length <= 10⁴
// s 由小写英文字母组成
//
//
//
//
// 注意：该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-
//distinct-characters 相同
//
// Related Topics 栈 贪心 字符串 单调栈 👍 986 👎 0

package main

import "fmt"

func main() {
	s := "cbacdcbc"
	letters := removeDuplicateLetters(s)
	fmt.Println(letters)
}

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicateLetters(s string) string {
	stack, inStack, cache := make([]int32, 0, 26), [26]bool{}, [26]int{}
	for _, c := range s {
		cache[c-'a']++
	}
	for _, c := range s {
		i := c - 'a'
		if !inStack[i] { // “当前首次”添加
			last := len(stack) - 1
			for last >= 0 && c < stack[last] { // 字典序更大，则尝试淘汰 stack[last]
				v := stack[last] - 'a'
				if cache[v] == 0 { // 最后一次出现了
					break
				}
				last--
				inStack[v] = false // 已出栈
			}
			stack = stack[:last+1]
			stack = append(stack, c)
			inStack[i] = true // 已入栈
		}
		cache[i]-- // 次数 -1
	}
	return string(stack)
}

//leetcode submit region end(Prohibit modification and deletion)
