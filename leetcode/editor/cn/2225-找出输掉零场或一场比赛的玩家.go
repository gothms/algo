package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findWinners(matches [][]int) [][]int {
	ans, memo := make([][]int, 2), make(map[int]int)
	for _, e := range matches {
		memo[e[0]] += 0
		memo[e[1]]++
	}
	for k, v := range memo {
		if v <= 1 {
			ans[v] = append(ans[v], k)
		}
	}
    sort.Ints(ans[0])
    sort.Ints(ans[1])
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
