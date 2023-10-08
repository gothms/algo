//你的朋友正在使用键盘输入他的名字 name。偶尔，在键入字符 c 时，按键可能会被长按，而字符可能被输入 1 次或多次。
//
// 你将会检查键盘输入的字符 typed。如果它对应的可能是你的朋友的名字（其中一些字符可能被长按），那么就返回 True。
//
//
//
// 示例 1：
//
//
//输入：name = "alex", typed = "aaleex"
//输出：true
//解释：'alex' 中的 'a' 和 'e' 被长按。
//
//
// 示例 2：
//
//
//输入：name = "saeed", typed = "ssaaedd"
//输出：false
//解释：'e' 一定需要被键入两次，但在 typed 的输出中不是这样。
//
//
//
//
// 提示：
//
//
// 1 <= name.length, typed.length <= 1000
// name 和 typed 的字符都是小写字母
//
//
// Related Topics 双指针 字符串 👍 289 👎 0

package main

import "fmt"

func main() {
	name := "leelee"
	typed := "lleeelee"
	pressedName := isLongPressedName(name, typed)
	fmt.Println(pressedName)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isLongPressedName(name string, typed string) bool {
	j, n := 0, len(typed)
	for i := 0; i < n; i++ {
		if j < len(name) && typed[i] == name[j] { // i j 匹配
			j++
		} else if j == 0 || typed[i] != name[j-1] { // i j-1 不匹配
			return false
		}
	}
	return j == len(name)
}

//leetcode submit region end(Prohibit modification and deletion)
