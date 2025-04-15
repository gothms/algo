package main

import (
	"strings"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findAndReplacePattern(words []string, pattern string) []string {
	// lc：双射，words[i] 映射 pattern，且 pattern 映射 words[i]

	// 个人
	var sb strings.Builder
	match := func(s string) {
		sb.Reset()
		var i byte = 'a' // 自增变量
		memo := [26]byte{}
		for _, c := range s { // 待优化：实时判断 words[i] 的每个字符是否匹配
			idx := int(c - 'a')
			if memo[idx] == 0 {
				memo[idx] = i
				i++
			}
			sb.WriteByte(memo[idx])
		}
	}
	match(pattern)
	pattern = sb.String()

	ans := make([]string, 0)
	for _, w := range words {
		match(w)
		if sb.String() == pattern {
			ans = append(ans, w)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
