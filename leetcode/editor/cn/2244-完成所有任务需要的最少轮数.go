package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumRounds(tasks []int) int {
	memo := make(map[int]int)
	for _, v := range tasks {
		memo[v]++
	}
	ans := 0
	for _, v := range memo {
		if v == 1 {
			return -1
		}
		ans += (v + 2) / 3
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
