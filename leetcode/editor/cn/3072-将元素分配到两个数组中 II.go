package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	nums := []int{5, 14, 3, 1, 2}                  // [5 3 1 2 14]
	nums = []int{2, 1, 3, 3}                       // [2 3 1 3]
	nums = []int{47, 104, 11, 37}                  // [47,11,104,37]
	nums = []int{11, 92, 25, 90}                   // [11,92,25,90]
	nums = []int{68, 824, 967, 240, 921, 751, 177} // [68,967,921,751,177,824,240]
	array := resultArray(nums)
	fmt.Println(array)
}

// leetcode submit region begin(Prohibit modification and deletion)
func resultArray(nums []int) []int {

}

//leetcode submit region end(Prohibit modification and deletion)

func resultArray_(nums []int) []int {
	// 有序集合：redblacktree
	// redblacktree 无法区间查询

	sorted := slices.Clone(nums)
	sort.Ints(sorted)
	//sorted = slices.Compact(sorted) // 去掉重复元素，也可以不去重
	v2i := make(map[int]int)
	for i, v := range sorted {
		v2i[v] = i + 1 // k-v：值-索引
	}
	n := len(sorted) + 1
	bit1, bit2 := make([]int, n), make([]int, n) // 两个树状数组，分别对应 arr1 和 arr2
	update3072(v2i[nums[0]], bit1)
	update3072(v2i[nums[1]], bit2)
	arr1, arr2 := append(make([]int, 0, len(nums)-1), nums[0]), append(make([]int, 0, len(nums)-1), nums[1])
	for _, v := range nums[2:] {
		i := v2i[v]
		// 查出小于等于 v 的元素个数，再求得大于 v 的元素个数
		v1, v2 := len(arr1)-rangeSum3072(i, bit1), len(arr2)-rangeSum3072(i, bit2)
		if v1 > v2 || v1 == v2 && len(arr1) <= len(arr2) {
			update3072(i, bit1)
			arr1 = append(arr1, v)
		} else {
			update3072(i, bit2)
			arr2 = append(arr2, v)
		}
	}
	return append(arr1, arr2...)

	// OOM
	//n := 0
	//for _, v := range nums {
	//	n = max(n, v)
	//}
	//n++
	//fmt.Println(n)
	//bit1, bit2 := make([]int, n), make([]int, n)
	//arr1, arr2 := append(make([]int, 0, len(nums)-1), nums[0]), append(make([]int, 0, len(nums)-1), nums[1])
	//update3072(n-nums[0]+1, bit1)
	//update3072(n-nums[1]+1, bit2)
	//for _, v := range nums[2:] {
	//	t := n - v
	//	v1, v2 := rangeSum3072(t, bit1), rangeSum3072(t, bit2)
	//	if v1 > v2 || v1 == v2 && len(arr1) <= len(arr2) {
	//		arr1 = append(arr1, v)
	//		update3072(t+1, bit1)
	//	} else {
	//		arr2 = append(arr2, v)
	//		update3072(t+1, bit2)
	//	}
	//}
	//return append(arr1, arr2...)
}

func update3072(i int, bit []int) {
	for ; i < len(bit); i += i & -i {
		bit[i]++
	}
}
func rangeSum3072(i int, bit []int) int {
	s := 0
	for ; i > 0; i &= i - 1 {
		s += bit[i]
	}
	return s
}
