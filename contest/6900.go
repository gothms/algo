package main

import "fmt"

func main() {
	nums := []int{1, 3, 1, 2, 2}
	subarrays := countCompleteSubarrays(nums)
	fmt.Println(subarrays)
}
func countCompleteSubarrays(nums []int) int {
	// 滑动窗体
	cache := make(map[int]bool)
	for _, v := range nums {
		cache[v] = true
	}
	cnt, i, n := 0, 0, len(cache)
	window := make(map[int]int)
	for _, v := range nums {
		window[v]++
		for len(window) == n {
			window[nums[i]]-- // 刚好 n-1 个元素的子数组
			if window[nums[i]] == 0 {
				delete(window, nums[i])
			}
			i++ // 但是索引从 0 开始的，则刚好 i 个完全子数组
		}
		cnt += i
	}
	return cnt
}
