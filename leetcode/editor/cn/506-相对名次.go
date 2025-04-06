package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	score := []int{10, 3, 8, 9, 4}
	ranks := findRelativeRanks(score)
	fmt.Println(ranks)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findRelativeRanks(score []int) []string {
	idx := make([]int, len(score))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return score[idx[i]] > score[idx[j]]
	})
	fmt.Println(idx)
	ans := make([]string, len(score))
	for mc, i := range idx {
		switch mc {
		case 0:
			ans[i] = "Gold Medal"
		case 1:
			ans[i] = "Silver Medal"
		case 2:
			ans[i] = "Bronze Medal"
		default:
			ans[i] = strconv.Itoa(mc + 1)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
