package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	duplicate := containsNearbyDuplicate(nums, k)
	fmt.Println(duplicate)
}

// leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyDuplicate(nums []int, k int) bool {
	memo := make(map[int]int, k+1)
	for i, v := range nums {
		if i > k {
			memo[nums[i-k-1]]--
		}
		if memo[v] == 1 {
			return true
		}
		memo[v]++
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func containsNearbyDuplicate_(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	memo := make(map[int]int)
	for i, v := range nums {
		if i > k {
			memo[nums[i-k-1]]--
		}
		if memo[v] == 1 {
			return true
		}
		memo[v]++
	}
	return false
}
