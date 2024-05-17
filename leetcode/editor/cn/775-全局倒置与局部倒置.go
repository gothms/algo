package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isIdealPermutation(nums []int) bool {
	// 归纳证明：相邻两个数差的绝对值 == 1

	// 维护前缀最大值/后缀最小值：先出现的数中的最大值，小于后出现的数

	// 个人
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i-1] > nums[i] {
			nums[i-1], nums[i] = nums[i], nums[i-1]
			i--
		}
	}
	for i := n - 1; i > 0; i-- {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
