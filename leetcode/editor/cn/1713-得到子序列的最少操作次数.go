package main

import (
	"fmt"
	"sort"
)

func main() {
	target := []int{6, 4, 8, 1, 3, 2}
	arr := []int{4, 7, 6, 2, 3, 8, 6, 1} // 3
	//target = []int{16, 7, 20, 11, 15, 13, 10, 14, 6, 8}
	//arr = []int{11, 14, 15, 7, 5, 5, 6, 10, 11, 6} // 6
	operations := minOperations(target, arr)
	fmt.Println(operations)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(target []int, arr []int) int {
	n := len(target)
	memo := make(map[int]int, n)
	for i, v := range target {
		memo[v] = i // 值映射索引
	}
	st := make([]int, 0)
	for _, v := range arr {
		if i, ok := memo[v]; ok {
			idx := sort.SearchInts(st, i) // 二分
			if idx == len(st) {
				st = append(st, i) // 追加子序列
			} else {
				st[idx] = i // 更新子序列
			}
		}
	}
	return n - len(st)
}

//leetcode submit region end(Prohibit modification and deletion)
