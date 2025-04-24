package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans, j, s := n+1, 0, 0
	for i, v := range nums {
		if s += v; s < target {
			continue
		}
		for ; j <= i && s >= target; j++ {
			s -= nums[j]
		}
		ans = min(ans, i-j+2)
	}
	if ans > n {
		return 0
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
