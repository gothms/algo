package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 1, 2}
	nums = []int{2, 1, 4}
	shifts := minimumRightShifts(nums)
	fmt.Println(shifts)
}
func minimumRightShifts(nums []int) int {
	i, t, n := 1, nums[0], len(nums)
	for i < n && nums[i] > nums[i-1] { // 计算旋转开始位置
		i++
	}
	for j := i + 1; j < n; j++ {
		if nums[j] > t || nums[j] < nums[j-1] {
			return -1 // 后半截错误
		}
	}
	return n - i
}
