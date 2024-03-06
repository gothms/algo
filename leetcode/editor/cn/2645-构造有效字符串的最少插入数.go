//给你一个字符串 word ，你可以向其中任何位置插入 "a"、"b" 或 "c" 任意次，返回使 word 有效 需要插入的最少字母数。
//
// 如果字符串可以由 "abc" 串联多次得到，则认为该字符串 有效 。
//
//
//
// 示例 1：
//
// 输入：word = "b"
//输出：2
//解释：在 "b" 之前插入 "a" ，在 "b" 之后插入 "c" 可以得到有效字符串 "abc" 。
//
//
// 示例 2：
//
// 输入：word = "aaa"
//输出：6
//解释：在每个 "a" 之后依次插入 "b" 和 "c" 可以得到有效字符串 "abcabcabc" 。
//
//
// 示例 3：
//
// 输入：word = "abc"
//输出：0
//解释：word 已经是有效字符串，不需要进行修改。
//
//
//
//
// 提示：
//
//
// 1 <= word.length <= 50
// word 仅由字母 "a"、"b" 和 "c" 组成。
//
//
// Related Topics 栈 贪心 字符串 动态规划 👍 25 👎 0

package main

import "fmt"

func main() {
	word := "aaaaac"
	minimum := addMinimum(word)
	fmt.Println(minimum)
}

// leetcode submit region begin(Prohibit modification and deletion)
func addMinimum(word string) int {
	// dp
	ret := 0
	for i, c := range word {
		switch c {
		case 'a':
			if i > 0 && word[i-1] != 'c' {
				ret += int('c' - word[i-1])
			}
		case 'b':
			if i > 0 && word[i-1] != 'a' {
				ret += int('c'-word[i-1]) + 1
			} else if i == 0 {
				ret = 1
			}
		case 'c':
			if i == 0 {
				ret += 2
			} else {
				ret += 2 - (int(word[i-1]-'a')+1)%3
			}
		}
	}
	return ret + int('c'-word[len(word)-1])

	// lc 假设答案由 ttt 个 ‘abc’\text{`abc'}‘abc’ 组成，那么需要插入的字符个数为 3t−n3t-n3t−n
	//t := 1
	//for i := 1; i < len(word); i++ {
	//	if word[i-1] >= word[i] { // 必须生成一个新的 abc
	//		t++
	//	}
	//}
	//return t*3 - len(word)
}

//leetcode submit region end(Prohibit modification and deletion)
