package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	nums := []int{1, 5, 3, 9, 8}
	limit := 2
	//nums = []int{1, 60, 34, 84, 62, 56, 39, 76, 49, 38}
	//limit = 4 // [1,56,34,84,60,62,38,76,49,39]
	array := lexicographicallySmallestArray(nums, limit)
	fmt.Println(array)
}

// leetcode submit region begin(Prohibit modification and deletion)
func lexicographicallySmallestArray(nums []int, limit int) []int {
	// 并查集

	// lc：即对排序方法的优化
	n := len(nums)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return nums[idx[i]] < nums[idx[j]]
	})

	ans := slices.Clone(nums)
	for i := 0; i < n; {
		j := i
		for i++; i < n && nums[idx[i]]-nums[idx[i-1]] <= limit; {
			i++
		}
		if j+1 == i {
			continue
		}
		sub := slices.Clone(idx[j:i])
		sort.Ints(sub)
		for k, index := range sub {
			ans[index] = nums[idx[j+k]]
		}
	}
	return ans

	// 个人：排序
	//n := len(nums)
	//idx := make([]int, n)
	//for i := range idx {
	//	idx[i] = i
	//}
	//sort.Slice(idx, func(i, j int) bool {
	//	return nums[idx[i]] < nums[idx[j]]
	//})
	//match := slices.Clone(idx)
	//j := 0
	//for i := 1; i < n; i++ {
	//	if nums[match[i]]-nums[match[i-1]] <= limit { // 同一子集
	//		continue
	//	}
	//	if j+1 == i {
	//		match[j] = -1
	//	} else {
	//		sort.Ints(match[j:i])
	//	}
	//	j = i
	//}
	//if j == n-1 {
	//	match[j] = -1
	//} else {
	//	sort.Ints(match[j:])
	//}
	//ans := slices.Clone(nums)
	//for i := 0; i < n; i++ {
	//	for ; i < n && match[i] >= 0; i++ {
	//		ans[match[i]] = nums[idx[i]]
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
