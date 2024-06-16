package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxRotateFunction(nums []int) int {
	// 个人：观察案例一
	n := len(nums)
	cur, s := 0, 0
	for i, v := range nums {
		s += v
		cur += i * v
	}
	ans := cur
	for i := 1; i < n; i++ {
		cur += s - nums[n-i]*n
		ans = max(ans, cur)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
