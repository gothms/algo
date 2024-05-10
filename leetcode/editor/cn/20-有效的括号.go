//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
//
//
//
// 示例 1：
//
//
//输入：s = "()"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "()[]{}"
//输出：true
//
//
// 示例 3：
//
//
//输入：s = "(]"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
//
// Related Topics 栈 字符串 👍 4443 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	st := make([]int32, 1)
	for _, c := range s {
		var t int32
		switch c {
		case ')':
			t = '('
		case ']':
			t = '['
		case '}':
			t = '{'
		default:
			st = append(st, c)
			continue
		}
		if st[len(st)-1] != t {
			return false
		}
		st = st[:len(st)-1]
	}
	return len(st) == 1
}

//leetcode submit region end(Prohibit modification and deletion)
