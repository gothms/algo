//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//
//
// 示例 2：
//
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//
//
// 示例 3：
//
//
//输入：s = ""
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'
//
//
// Related Topics 栈 字符串 动态规划 👍 2497 👎 0

package main

import "fmt"

func main() {
	s := "(()"
	//s = "()(()"
	//s = "()(())"
	parentheses := longestValidParentheses(s)
	fmt.Println(parentheses)
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {

}

//leetcode submit region end(Prohibit modification and deletion)

//func longestValidParentheses(s string) int {
//    // 类卡特兰思路
//    var ans, left, right int
//    for _, c := range s {
//        switch c { // 计数
//        case '(':
//            left++
//        case ')':
//            right++
//        }
//        if left == right { // 区间内完全匹配
//            ans = max(ans, right<<1)
//        } else if right > left { // 重置区间
//            left, right = 0, 0
//        }
//    }
//    left, right = 0, 0
//    for i := len(s) - 1; i >= 0; i-- { // (() 结果是 0，则倒序再遍历一遍
//        switch s[i] {
//        case '(':
//            left++
//        case ')':
//            right++
//        }
//        if left == right {
//            ans = max(ans, left<<1)
//        } else if left > right {
//            left, right = 0, 0
//        }
//    }
//    return ans
//
//    // stack
//    //st := []int{-1} // 哨兵
//    //var ans int
//    //for i, c := range s {
//    //	switch c {
//    //	case '(':
//    //		st = append(st, i)
//    //	case ')':
//    //		st = st[:len(st)-1]
//    //		if len(st) == 0 { // 新的哨兵
//    //			st = append(st, i)
//    //		} else {
//    //			ans = max(ans, i-st[len(st)-1]) // 此时的 st[len(st)-1] 肯定是哨兵
//    //		}
//    //	}
//    //}
//    //return ans
//
//    // dp
//    //dp := make([]int, len(s)+1)
//    //var ans int
//    //for i, c := range s {
//    //	if c == '(' {
//    //		continue
//    //	}
//    //	if i > 0 && s[i-1] == '(' { // 相邻匹配
//    //		dp[i+1] = dp[i-1] + 2
//    //	} else if j := i - dp[i] - 1; dp[i] > 0 && j >= 0 && s[j] == '(' { // “间隔”匹配
//    //		dp[i+1] = dp[i] + 2 + dp[j]
//    //	}
//    //	ans = max(ans, dp[i+1])
//    //}
//    //return ans
//}
