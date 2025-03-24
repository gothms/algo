package main

import "strings"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countPrefixes(words []string, s string) int {
	memo := make(map[string]struct{})
	var sb strings.Builder
	for _, c := range s {
		sb.WriteRune(c)
		memo[sb.String()] = struct{}{}
	}
	ans := 0
	for _, w := range words {
		if _, ok := memo[w]; ok {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
