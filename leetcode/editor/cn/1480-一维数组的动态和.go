package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func runningSum(nums []int) []int {
	ans := make([]int, len(nums)+1)
	for i, v := range nums {
		ans[i+1] = ans[i] + v
	}
	return ans[1:]
}

//leetcode submit region end(Prohibit modification and deletion)
