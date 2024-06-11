package main

import (
	"fmt"
	"sort"
)

func main() {
	mapping := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // [9,8,7,6,5,4,3,2,1,0]
	jumbled := sortJumbled(mapping, nums)
	fmt.Println(jumbled)
}

// leetcode submit region begin(Prohibit modification and deletion)
func sortJumbled(mapping []int, nums []int) []int {
	n := len(nums)
	mapNums := make([][2]int, n)
	for i, v := range nums {
		val := 0
		for base, cur := 1, v; ; {
			val += mapping[cur%10] * base
			base *= 10
			cur /= 10
			if cur == 0 { // 防止 nums[i]==0
				break
			}
		}
		mapNums[i] = [2]int{val, v}
	}
	sort.SliceStable(mapNums, func(i, j int) bool { // 稳定排序
		return mapNums[i][0] < mapNums[j][0]
	})
	for i, v := range mapNums {
		nums[i] = v[1]
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
