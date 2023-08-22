package main

import "fmt"

func main() {
	nums := []int{2, 1, 3, 2, 1}
	operations := minimumOperations(nums)
	fmt.Println(operations)
}
func minimumOperations(nums []int) int {
	// 状态机 DP
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	cnt := [3]int{} // 已排序元素末尾是 1/2/3 时，最少步数
	for _, v := range nums {
		cnt[1] = minVal(cnt[0], cnt[1]) + 1 // 末尾是 2
		cnt[2] = minVal(cnt[1], cnt[2]+1)   // 末尾是 3
		cnt[0]++                            // 末尾是 1
		cnt[v-1]--                          // 修正末尾是 v
	}
	return minVal(cnt[0], minVal(cnt[1], cnt[2]))

	// LIS
	//one, two, three := 0, 0, 0
	//for _, v := range nums {
	//	switch v {
	//	case 1: // 插入 1
	//		one++
	//		if two > 0 { // 如果不是末尾，则 2/3 数量 -1
	//			two--
	//		} else if three > 0 {
	//			three--
	//		}
	//	case 2: // 插入 2
	//		two++
	//		if three > 0 { // 如果不是末尾，则 3 数量 -1
	//			three--
	//		}
	//	case 3: // 末尾插入 3
	//		three++
	//	}
	//}
	//return len(nums) - one - two - three
}
