package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func prefixCount(words []string, pref string) int {
	ans, n := 0, len(pref)
	for _, w := range words {
		if len(w) >= n && w[:n] == pref {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
