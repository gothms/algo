package main

import "fmt"

func main() {
	answers := []int{0, 0, 1, 1, 1}
	rabbits := numRabbits(answers)
	fmt.Println(rabbits)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numRabbits(answers []int) int {
	memo := make(map[int]int)
	for _, a := range answers {
		memo[a]++
	}
	ans := 0
	for k, v := range memo {
		ans += (v + k) / (k + 1) * (k + 1) // 核心
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
