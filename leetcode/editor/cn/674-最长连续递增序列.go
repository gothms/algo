package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findLengthOfLCIS(nums []int) int {
	ans, cur := 1, 1
	for i := range nums[:len(nums)-1] {
		if nums[i+1] > nums[i] {
			cur++
			ans = max(ans, cur)
		} else {
			cur = 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
