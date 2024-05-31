package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func sumOfSquares(nums []int) int {
	// i+1 满足 则 n/(i+1) 也满足
	ans, n := 0, len(nums)
	for i, v := range nums {
		if n%(i+1) == 0 {
			ans += v * v
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
