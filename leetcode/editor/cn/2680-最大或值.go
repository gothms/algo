package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumOr(nums []int, k int) int64 {
	
}

//leetcode submit region end(Prohibit modification and deletion)

func maximumOr_(nums []int, k int) int64 {
	// lc
	or, fix := 0, 0
	for _, v := range nums {
		fix |= or & v
		or |= v
	}
	ans := 0
	for _, v := range nums {
		ans = max(ans, (or^v)|fix|v<<k)
	}
	return int64(ans)

	//n := len(nums)
	//memo := make([]int, n)
	//for i := n - 2; i >= 0; i-- {
	//	memo[i] = memo[i+1] | nums[i+1]
	//}
	//pre := 0
	//for i, v := range nums {
	//	memo[i] |= pre
	//	pre |= v
	//}
	//var ans int64
	//for i, v := range nums {
	//	ans = max(ans, int64(memo[i])|int64(v<<k))
	//}
	//return ans
}
