package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 3
	k := 2
	budget := 15
	composition := [][]int{{1, 1, 1}, {1, 1, 10}}
	stock := []int{0, 0, 100}
	stock = []int{0, 0, 0}
	cost := []int{1, 2, 3}
	alloys := maxNumberOfAlloys(n, k, budget, composition, stock, cost)
	fmt.Println(alloys)
}
func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	maxBudget, one := budget, 0
	for i, v := range cost {
		one += v // 生产一个合金的最少消耗
		maxBudget += stock[i] * v
	}
	mnoa, max := 0, maxBudget/one+1  // 所有和金的消耗和金属的 cost 都为 1，最大生产 maxBudget/one 个合金
	for _, cp := range composition { // composition[i] 条件下，最多生产多少合金
		cnt := sort.Search(max, func(i int) bool { // max+1 处理
			sum := 0
			for j, v := range cp { // 统计生产 i 个合金，需要多少消耗
				sum += maxVal(0, i*v-stock[j]) * cost[j]
			}
			return sum > budget
		}) - 1 // 修正二分
		mnoa = maxVal(mnoa, cnt)
	}
	return mnoa
}
