package main

import (
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numRescueBoats(people []int, limit int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func numRescueBoats_(people []int, limit int) int {
	// 贪心
	// 重点：最多可同时载两人
	sort.Ints(people)
	ans := 0
	i, j := 0, len(people)-1
	for i < j {
		if people[i]+people[j] <= limit {
			i++
		}
		j--
		ans++
	}
	if i == j {
		ans++
	}
	return ans
}
