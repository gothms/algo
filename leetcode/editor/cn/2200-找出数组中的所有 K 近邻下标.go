package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findKDistantIndices(nums []int, key int, k int) []int {
	ans := make([]int, 0)
	cur, n := 0, len(nums)
	for i, v := range nums {
		if v == key {
			for j := max(i-k, cur); j <= min(i+k, n-1); j++ {
				ans = append(ans, j)
			}
			cur = i + k + 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
