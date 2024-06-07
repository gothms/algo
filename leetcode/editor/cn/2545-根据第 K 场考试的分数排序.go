package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})
	return score
}

//leetcode submit region end(Prohibit modification and deletion)
