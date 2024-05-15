package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumWealth(accounts [][]int) int {
	ans := 0
	for _, a := range accounts {
		s := 0
		for _, v := range a {
			s += v
		}
		ans = max(ans, s)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
