package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func waysToSplitArray(nums []int) int {
	// 或者先记录总 sum
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	ans, last := 0, nums[n-1]
	for i, v := range nums[:n-1] {
		if v >= last-nums[i] {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
