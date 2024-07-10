package main

import "fmt"

func main() {
	nums := []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1}
	subarrays := countAlternatingSubarrays(nums)
	fmt.Println(subarrays)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countAlternatingSubarrays(nums []int) int64 {
	ans, last := 1, -1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			ans++
			last = i - 1
		} else {
			ans += i - last
		}
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
