//给你一个字符串数组 words ，找出并返回 length(words[i]) * length(words[j]) 的最大值，并且这两个单词不含有公共字母
//。如果不存在这样的两个单词，返回 0 。
//
//
//
// 示例 1：
//
//
//输入：words = ["abcw","baz","foo","bar","xtfn","abcdef"]
//输出：16
//解释：这两个单词为 "abcw", "xtfn"。
//
// 示例 2：
//
//
//输入：words = ["a","ab","abc","d","cd","bcd","abcd"]
//输出：4
//解释：这两个单词为 "ab", "cd"。
//
// 示例 3：
//
//
//输入：words = ["a","aa","aaa","aaaa"]
//输出：0
//解释：不存在这样的两个单词。
//
//
//
//
// 提示：
//
//
// 2 <= words.length <= 1000
// 1 <= words[i].length <= 1000
// words[i] 仅包含小写字母
//
//
// Related Topics 位运算 数组 字符串 👍 432 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(words []string) int {
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	mp, n := 0, len(words)
	bitWords := make([]int, n)
	for i, w := range words {
		for _, c := range w {
			bitWords[i] |= 1 << int(c-'a') // 转为二进制
		}
	}
	for i, a := range bitWords {
		for j, b := range bitWords[:i] { // 遍历 length(words[i]) * length(words[j])
			if a&b == 0 {
				mp = maxVal(mp, len(words[i])*len(words[j]))
			}
		}
	}
	return mp
}

//leetcode submit region end(Prohibit modification and deletion)
