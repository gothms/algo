package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countCompleteSubarrays(nums []int) int {
	memo := make(map[int]int)
	for _, v := range nums {
		memo[v]++
	}
	win := make(map[int]int)
	ans, k, j := 0, len(memo), 0
	for _, v := range nums {
		win[v]++
		for len(win) == k {
			win[nums[j]]--
			if win[nums[j]] == 0 {
				delete(win, nums[j])
			}
			j++
		}
		ans += j
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
