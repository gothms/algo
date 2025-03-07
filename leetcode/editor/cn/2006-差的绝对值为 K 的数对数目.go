package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5, 4}
	k := 2
	difference := countKDifference(nums, k)
	fmt.Println(difference)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countKDifference(nums []int, k int) int {
	memo := make(map[int]int, len(nums))
	cnt := 0
	for _, v := range nums {
		cnt += memo[v+k] + memo[v-k]
		memo[v]++
	}
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
