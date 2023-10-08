//给你一个方程，左边用 words 表示，右边用 result 表示。
//
// 你需要根据以下规则检查方程是否可解：
//
//
// 每个字符都会被解码成一位数字（0 - 9）。
// 每对不同的字符必须映射到不同的数字。
// 每个 words[i] 和 result 都会被解码成一个没有前导零的数字。
// 左侧数字之和（words）等于右侧数字（result）。
//
//
// 如果方程可解，返回 True，否则返回 False。
//
//
//
// 示例 1：
//
// 输入：words = ["SEND","MORE"], result = "MONEY"
//输出：true
//解释：映射 'S'-> 9, 'E'->5, 'N'->6, 'D'->7, 'M'->1, 'O'->0, 'R'->8, 'Y'->'2'
//所以 "SEND" + "MORE" = "MONEY" ,  9567 + 1085 = 10652
//
// 示例 2：
//
// 输入：words = ["SIX","SEVEN","SEVEN"], result = "TWENTY"
//输出：true
//解释：映射 'S'-> 6, 'I'->5, 'X'->0, 'E'->8, 'V'->7, 'N'->2, 'T'->1, 'W'->'3', 'Y'->
//4
//所以 "SIX" + "SEVEN" + "SEVEN" = "TWENTY" ,  650 + 68782 + 68782 = 138214
//
// 示例 3：
//
// 输入：words = ["THIS","IS","TOO"], result = "FUNNY"
//输出：true
//
//
// 示例 4：
//
// 输入：words = ["LEET","CODE"], result = "POINT"
//输出：false
//
//
//
//
// 提示：
//
//
// 2 <= words.length <= 5
// 1 <= words[i].length, results.length <= 7
// words[i], result 只含有大写英文字母
// 表达式中使用的不同字符数最大为 10
//
//
// Related Topics 数组 数学 字符串 回溯 👍 73 👎 0

package main

import (
	"fmt"
	"math/bits"
)

func main() {
	words := []string{"SIX", "SEVEN", "SEVEN"}
	result := "TWENTY"
	words = []string{"BUT", "ITS", "STILL"}
	result = "FUNNY"
	words = []string{"CBA", "CBA", "CBA", "CBA", "CBA"}
	result = "EDD"
	words = []string{"A", "B"}
	result = "A"
	solvable := isSolvable(words, result)
	fmt.Println(solvable)
	fmt.Printf("%b\n", 1023)
	fmt.Printf("%b\n", 2044)
	fmt.Printf("%b", 1023&^1<<1)
}

/*
剪枝
	0.leadingNum 和 result 前导数字的取值/范围
	1.低位优先：从低位向高位搜索，sum(words)-result，可计算出 result 的字符的值
	2.权值合并：先映射每个字符的权值，也可根据权值确定字符的范围
*/
// leetcode submit region begin(Prohibit modification and deletion)
func isSolvable(words []string, result string) bool {
	// 回溯
	var (
		match = [26]int{}      // 记录每个字符的取值
		start = [26]int{}      // 记录首字符的最小取值
		ch    = make([]int, 0) // 记录出现过的字符
	)
	t, h, n := 1<<10-1, 0, len(result)
	rangeStr := func(s string) { // 收集字符
		if len(result) > 1 {
			start[result[0]-'A'] = 1
		}
		for _, c := range s {
			if i := c - 'A'; match[i] == 0 {
				ch = append(ch, int(i))
				match[i] = -1 // 去重
			}
		}
	}
	rangeStr(result)
	for _, w := range words {
		if len(w) > 1 {
			start[w[0]-'A'] = 1
		}
		switch len(w) { // result 的第一个字符
		case n:
			h = 9
		case n - 1:
			h++
		}
		rangeStr(w)
	}
	if h == 0 { // 如 ABC 和 D
		return false
	} else if h == len(words) {
		h--
	} else if h > 9 {
		h = 9
	}
	getV := func(s string) (v int) { // 计算字符串的值
		for _, c := range s {
			v = v*10 + match[c-'A']
		}
		return
	}
	getRet := func() bool { // 验证和 = result
		sum := 0
		for _, w := range words {
			sum += getV(w)
		}
		return sum == getV(result)
	}
	var dfs func(int, int) bool
	dfs = func(i, t int) bool {
		if i == len(ch) {
			return getRet() // 验证
		}
		l := bits.Len(uint(t))
		for b := start[ch[i]]; b < l; b++ { // word 前导不能为 0
			if bit := 1 << b; bit&t != 0 { // 对每个字符赋值
				if match[ch[i]] = b; dfs(i+1, t&^bit) {
					return true
				}
			}
		}
		return false
	}
	for i := start[ch[0]]; i <= h; i++ {
		match[ch[0]] = i // result 的第一个字符
		if dfs(1, t&^(1<<i)) {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
