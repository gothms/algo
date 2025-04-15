package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func wordSubsets(words1 []string, words2 []string) []string {
	memo := [26]int{}
	for _, w := range words2 {
		cur := [26]int{}
		for _, c := range w {
			cur[c-'a']++
		}
		for i, v := range memo {
			memo[i] = max(v, cur[i])
		}
	}
	ans := make([]string, 0)
out:
	for _, w := range words1 {
		cur := [26]int{}
		for _, c := range w {
			cur[c-'a']++
		}
		for i, v := range memo {
			if v > cur[i] {
				continue out
			}
		}
		ans = append(ans, w)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
