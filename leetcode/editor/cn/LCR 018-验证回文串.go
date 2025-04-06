package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "race a car"
	palindrome := isPalindrome(s)
	fmt.Println(palindrome)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var valid func(int, int) bool
	valid = func(l, r int) bool {
		if l >= r {
			return true
		}
		if s[l] == s[r] {
			return valid(l+1, r-1)
		} else if !unicode.IsDigit(rune(s[l])) && !unicode.IsLetter(rune(s[l])) {
			return valid(l+1, r)
		} else if !unicode.IsDigit(rune(s[r])) && !unicode.IsLetter(rune(s[r])) {
			return valid(l, r-1)
		}
		return false
	}
	return valid(0, len(s)-1)
}

//leetcode submit region end(Prohibit modification and deletion)
