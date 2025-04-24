package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSum(nums []int) int {
	memo := [10]int{}
	ans := -1
	for _, v := range nums {
		mx := 0
		for i := v; i > 0; i /= 10 {
			mx = max(mx, i%10)
		}
		if memo[mx] > 0 {
			ans = max(ans, memo[mx]+v)
		}
		memo[mx] = max(memo[mx], v)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
