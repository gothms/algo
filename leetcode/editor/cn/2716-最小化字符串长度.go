package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimizedStringLength(s string) int {
	cache := [260]bool{}
	ans := 0
	for _, c := range s {
		if !cache[c-'a'] {
			ans++
		}
		cache[c-'a'] = true
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
