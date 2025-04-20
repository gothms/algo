package main

import "fmt"

func main() {
	answers := []int{0, 0, 1, 1, 1}
	rabbits := numRabbits(answers)
	fmt.Println(rabbits)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numRabbits(answers []int) int {
	ans := 0
	memo := make(map[int]int)
	for _, v := range answers {
		v++
		memo[v] = (memo[v] + 1) % v
		if v == 1 || memo[v] == 1 {
			ans += v
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func numRabbits_(answers []int) int {
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
