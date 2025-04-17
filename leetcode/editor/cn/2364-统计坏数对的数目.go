package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countBadPairs(nums []int) int64 {
	// j - i != nums[j] - nums[i]：则 nums[i] - i != nums[j] - j
	memo := make(map[int]int)
	ans := 0
	for i, v := range nums {
		ans += i - memo[v-i]
		memo[v-i]++
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
