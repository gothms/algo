package main

import (
	"fmt"
)

func main() {
	difficulty := []int{13, 37, 58}
	profit := []int{4, 90, 96}
	worker := []int{34, 73, 45} // 190
	assignment := maxProfitAssignment(difficulty, profit, worker)
	fmt.Println(assignment)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	// 个人：排序 difficulty + profit，则不使用二分而使用迭代记录当前最大 profit[i]，则更优化
	//n := len(difficulty)
	//idx, maxProfit := make([]int, n), make([]int, n)
	//for i := range idx {
	//	idx[i] = i
	//}
	//sort.Slice(idx, func(i, j int) bool {
	//	return difficulty[idx[i]] < difficulty[idx[j]]
	//})
	//last := 0
	//for i, j := range idx {
	//	maxProfit[i] = max(last, profit[j])
	//	last = maxProfit[i]
	//}
	//ans := 0
	//for _, v := range worker {
	//	i := sort.Search(n, func(i int) bool {
	//		return difficulty[idx[i]] > v
	//	}) - 1
	//	if i >= 0 {
	//		ans += maxProfit[i]
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
