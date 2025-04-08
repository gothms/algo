package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func canPermutePalindrome(s string) bool {
	memo := make(map[rune]int)
	for _, c := range s {
		memo[c]++
	}
	odd := 0
	for _, v := range memo {
		odd += v & 1
	}
	return odd <= 1
}

//leetcode submit region end(Prohibit modification and deletion)
