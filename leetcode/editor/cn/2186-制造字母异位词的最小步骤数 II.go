package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minSteps(s string, t string) int {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	memo := make(map[uint8]int)
	for i := range s {
		memo[s[i]]++
	}
	for i := range t {
		memo[t[i]]--
	}
	ans := 0
	for _, v := range memo {
		ans += abs(v)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
