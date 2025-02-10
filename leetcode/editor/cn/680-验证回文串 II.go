package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	valid := func(x, y int) bool {
		for ; x < y; x, y = x+1, y-1 {
			if s[x] != s[y] {
				return false
			}
		}
		return true
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return valid(i+1, j) || valid(i, j-1)
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
