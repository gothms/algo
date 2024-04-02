//你的笔记本键盘存在故障，每当你在上面输入字符 'i' 时，它会反转你所写的字符串。而输入其他字符则可以正常工作。
//
// 给你一个下标从 0 开始的字符串 s ，请你用故障键盘依次输入每个字符。
//
// 返回最终笔记本屏幕上输出的字符串。
//
//
//
// 示例 1：
//
// 输入：s = "string"
//输出："rtsng"
//解释：
//输入第 1 个字符后，屏幕上的文本是："s" 。
//输入第 2 个字符后，屏幕上的文本是："st" 。
//输入第 3 个字符后，屏幕上的文本是："str" 。
//因为第 4 个字符是 'i' ，屏幕上的文本被反转，变成 "rts" 。
//输入第 5 个字符后，屏幕上的文本是："rtsn" 。
//输入第 6 个字符后，屏幕上的文本是： "rtsng" 。
//因此，返回 "rtsng" 。
//
//
// 示例 2：
//
// 输入：s = "poiinter"
//输出："ponter"
//解释：
//输入第 1 个字符后，屏幕上的文本是："p" 。
//输入第 2 个字符后，屏幕上的文本是："po" 。
//因为第 3 个字符是 'i' ，屏幕上的文本被反转，变成 "op" 。
//因为第 4 个字符是 'i' ，屏幕上的文本被反转，变成 "po" 。
//输入第 5 个字符后，屏幕上的文本是："pon" 。
//输入第 6 个字符后，屏幕上的文本是："pont" 。
//输入第 7 个字符后，屏幕上的文本是："ponte" 。
//输入第 8 个字符后，屏幕上的文本是："ponter" 。
//因此，返回 "ponter" 。
//
//
//
// 提示：
//
//
// 1 <= s.length <= 100
// s 由小写英文字母组成
// s[0] != 'i'
//
//
// Related Topics 字符串 模拟 👍 10 👎 0

package main

import (
	"fmt"
	"slices"
)

func main() {
	s := "string"
	s = "poiinter"
	s = "poiabcinter"
	s2 := finalString(s)
	fmt.Println(s2)
}

// leetcode submit region begin(Prohibit modification and deletion)
func finalString(s string) string {
	// 可优化：odd，even 定义为 [2][]uint8{}，通过 ^1 来判断奇偶
	last, cnt, n := 0, 0, len(s)
	odd, even := make([]uint8, 0, n), make([]uint8, 0, n) // 奇、偶
	for cur := 0; cur < n; cur++ {
		if s[cur] == 'i' {
			if cnt&1 == 0 {
				even = append(even, s[last:cur]...)
			} else {
				odd = append(odd, s[last:cur]...)
			}
			cnt++
			last = cur + 1
		}
	}
	if cnt&1 == 0 {
		even = append(even, s[last:n]...)
		slices.Reverse(odd)
		return string(odd) + string(even)
	} else {
		odd = append(odd, s[last:n]...)
		slices.Reverse(even)
		return string(even) + string(odd)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
