package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isSubsequence(s string, t string) bool {

}

//leetcode submit region end(Prohibit modification and deletion)

func isSubsequence_(s string, t string) bool {
	// 优雅
	i, n := 0, len(s)
	if n == 0 {
		return true
	}
	for _, c := range t {
		if s[i] == byte(c) {
			i++
			if i == n {
				return true
			}
		}
	}
	return false
}
