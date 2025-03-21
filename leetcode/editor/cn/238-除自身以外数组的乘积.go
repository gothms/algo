package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		ans[i] = ans[i+1] * nums[i+1]
	}
	pre := 1
	for i, v := range nums {
		ans[i] *= pre
		pre *= v
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
