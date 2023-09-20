package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{6, 0, 3, 3, 6, 7, 2, 7}
	nums = []int{1, 1} // 2
	//nums = []int{1, 1, 0, 1} // 1
	ways := countWays(nums)
	fmt.Println(ways)
}
func countWays(nums []int) int {
	cw, last := 0, len(nums)-1
	sort.Ints(nums)
	for i, v := range nums {
		if v < i+1 && (i == last || i+1 < nums[i+1]) {
			cw++
		} // v < i+1 < nums[i+1]，则 i+1 符合
	}
	if nums[0] > 0 {
		return cw + 1
	}
	return cw

	// 个人写法：没 AC
	//n := len(nums)
	//idx := make([][2]int, n)
	//for i := 0; i < n; i++ {
	//	idx[i] = [2]int{nums[i], i}
	//}
	//sort.Slice(idx, func(i, j int) bool {
	//	return idx[i][0] < idx[j][0]
	//})
	////fmt.Println(idx)
	//memo := make(map[int]bool)
	//for _, v := range idx {
	//	if v[0] >= v[1] {
	//		continue
	//	}
	//	k := sort.Search(v[1], func(i int) bool {
	//		return idx[i][0] >= v[1]
	//	})
	//	fmt.Println(v, k, v[1])
	//	//if k > v[0] {
	//	//	continue
	//	//}
	//	memo[k] = true
	//}
	//if idx[0][0] > 0 {
	//	memo[0] = true
	//}
	//fmt.Println(memo)
	//return len(memo) + 1
}
