package main

import "fmt"

func main() {
	nums := []int{3, 5, 2, 6}
	k := 2
	//nums = []int{3, 5, 2, 1}
	competitive := mostCompetitive(nums, k)
	fmt.Println(competitive)
}

// leetcode submit region begin(Prohibit modification and deletion)
func mostCompetitive(nums []int, k int) []int {
	// 栈
	n := len(nums)
	st := make([]int, 0, k)
	for i, v := range nums {
		// max(0, k-n+i)：保证至少留 k 个元素
		for len(st) > max(0, k-n+i) && v < st[len(st)-1] {
			st = st[:len(st)-1]
		}
		if len(st) < k {
			st = append(st, v)
		}
	}
	return st
}

//leetcode submit region end(Prohibit modification and deletion)
