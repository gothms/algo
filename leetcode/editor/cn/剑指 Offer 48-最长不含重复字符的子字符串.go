//请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
//
//
//
// 示例 1:
//
// 输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
// 输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
// 输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//
//
// 提示：
//
//
// s.length <= 40000
//
//
// 注意：本题与主站 3 题相同：https://leetcode-cn.com/problems/longest-substring-without-
//repeating-characters/
//
// Related Topics 哈希表 字符串 滑动窗口 👍 601 👎 0

package main

import "fmt"

func main() {
	fmt.Printf("%c\n", 191+'a')
	s := "abcabcbb"
	s = "dvdf"
	substring := lengthOfLongestSubstring(s)
	fmt.Println(substring)
}

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	// 滑动窗体
	// dp
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	max, dp, n := 0, 0, len(s)
	cache := make(map[uint8]int)
	for i := 0; i < n; i++ {
		cache[s[i]]++
		if cache[s[i]] > 1 { // 字符重复
			j := i - dp
			for ; j < i && cache[s[i]] > 1; j++ { // 直到滑出重复的 s[i]
				cache[s[j]]--
			}
			dp = i + 1 - j // 新的长度
		} else { // 长度 +1
			dp++
		}
		//fmt.Printf("%c,%d\n", s[i], dp)
		max = maxVal(max, dp)
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
