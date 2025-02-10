package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	nums = []int{6, 5, 7, 8}
	subarrayCount := incremovableSubarrayCount(nums)
	fmt.Println(subarrayCount)
}

// leetcode submit region begin(Prohibit modification and deletion)
func incremovableSubarrayCount(nums []int) int {
	// lc
	i, n := 0, len(nums)
	for i < n-1 && nums[i] < nums[i+1] {
		i++
	}
	if i == n-1 {
		return n * (n + 1) >> 1
	}
	ans := i + 2 // 无后缀
	for j := n - 1; j == n-1 || nums[j] < nums[j+1]; j-- {
		for i >= 0 && nums[i] >= nums[j] {
			i--
		}
		ans += i + 2 // 有后缀
	}
	return ans

	// 个人
	//nums = append([]int{0}, nums...)
	//i, n := 1, len(nums)
	//for i < n && nums[i] > nums[i-1] {
	//	i++
	//}
	//if i == n { // 已排序
	//	return n * (n - 1) >> 1
	//}
	//j := n - 1 // nums[j,n) 有序
	//for j > 1 && nums[j] > nums[j-1] {
	//	j--
	//}
	//ans := 0
	//for k := 0; k < i; k++ {
	//	idx := sort.Search(n-j, func(i int) bool {
	//		return nums[j+i] > nums[k]
	//	})
	//	ans += n - j - idx + 1
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
