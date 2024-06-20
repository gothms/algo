package main

import (
	"fmt"
	"sort"
)

func main() {
	costs := []int{6, 2, 8, 8, 5, 6, 6, 2, 2, 2}
	coins := 77 // 10
	//costs = []int{7, 3, 3, 6, 6, 6, 10, 5, 9, 2}
	//coins = 56 // 9
	//costs = []int{11}
	//coins = 12
	cream := maxIceCream(costs, coins)
	fmt.Println(cream)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxIceCream(costs []int, coins int) int {
	// 计数排序

	// lc
	sort.Ints(costs)
	for i, v := range costs {
		if coins >= v {
			coins -= v
			continue
		}
		return i
	}
	return len(costs)

	// 个人
	//n := len(costs)
	//var qSort func(int, int, int) (int, int)
	//qSort = func(l, r, coin int) (int, int) {
	//	if l >= r {
	//		if l == n-1 && costs[l] <= coin { // 只剩一个尾元素
	//			return n, 0
	//		}
	//		return l, 0
	//	}
	//	counter, s := l+1, 0
	//	for i := l + 1; i <= r; i++ {
	//		if costs[i] < costs[l] {
	//			s += costs[i]
	//			costs[i], costs[counter] = costs[counter], costs[i]
	//			counter++
	//		}
	//	}
	//	counter--
	//	costs[l], costs[counter] = costs[counter], costs[l]
	//	pivot := counter
	//
	//	if s <= coin {
	//		if d := coin - s - costs[pivot]; d == 0 {
	//			return pivot + 1, 0
	//		} else if d < 0 {
	//			return pivot, 0
	//		} else {
	//			return qSort(pivot+1, r, d)
	//		}
	//	} else {
	//		return qSort(l, pivot-1, coin)
	//	}
	//}
	//i, _ := qSort(0, n-1, coins)
	//return i
}

//leetcode submit region end(Prohibit modification and deletion)
