package main

import (
	"unicode"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func shortestCompletingWord(licensePlate string, words []string) string {
	memo := [26]int{}
	for _, c := range licensePlate {
		if unicode.IsLetter(c) {
			memo[unicode.ToLower(c)-'a']++
		}
	}
	var ans string
out:
	for _, w := range words {
		cache := [26]int{}
		for _, c := range w {
			cache[c-'a']++
		}
		for i := range memo {
			if memo[i] > cache[i] {
				continue out
			}
		}
		if len(ans) == 0 || len(ans) > len(w) {
			ans = w
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
