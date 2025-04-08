package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findClosest(words []string, word1 string, word2 string) int {
	n := len(words)
	l1, l2 := -n, -n
	ans := n
	for i, w := range words {
		if w == word1 {
			l1 = i
			ans = min(ans, l1-l2)
		} else if w == word2 {
			l2 = i
			ans = min(ans, l2-l1)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
