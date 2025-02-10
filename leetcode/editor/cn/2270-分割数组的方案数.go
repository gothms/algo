package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 0} // 2
	array := waysToSplitArray(nums)
	fmt.Println(array)
}

// leetcode submit region begin(Prohibit modification and deletion)
func waysToSplitArray(nums []int) int {
	// 或者先记录总 sum
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	ans, last := 0, nums[n-1]
	for _, v := range nums[:n-1] {
		if v >= last-v {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
