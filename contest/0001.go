package main

import "fmt"

func main() {
	nums := []int{1, 2, 7, 8}
	//nums = []int{6, 5, 7, 8}
	count := incremovableSubarrayCount(nums)
	fmt.Println(count)
}
func incremovableSubarrayCount(nums []int) int {
	ret, n := 1, len(nums)-1 // n 为最后一个元素
	if n == 0 {
		return ret
	}
	left, i := 0, 1
	for ; i < n; i++ {
		if nums[i] <= nums[i-1] {
			break
		}
	}
	ret, left = i+1, i-1
	for i = n; i >= 0; i-- {
		if i < n && nums[i] >= nums[i+1] {
			break
		}
		ret++
		if left >= i-1 {
			left = i - 2
		}
		for left >= 0 && nums[i] <= nums[left] {
			left--
		}
		ret += left + 1
	}
	return ret
}
func incremovableSubarrayCount(nums []int) int64 {
	ret, n := int64(1), len(nums)-1 // n 为最后一个元素
	if n == 0 {
		return ret
	}
	left, i := 0, 1
	for ; i < n; i++ {
		if nums[i] <= nums[i-1] {
			break
		}
	}
	ret, left = int64(i+1), i-1
	for i = n; i >= 0; i-- {
		if i < n && nums[i] >= nums[i+1] {
			break
		}
		ret++
		if left >= i-1 {
			left = i - 2
		}
		for left >= 0 && nums[i] <= nums[left] {
			left--
		}
		ret += int64(left + 1)
	}
	return ret
}
