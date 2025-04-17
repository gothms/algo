package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findMaxLength(nums []int) int {
	// 前缀和
	memo := make(map[int]int)
	memo[0] = -1
	ans, s := 0, 0
	for i, v := range nums {
		s += v<<1 - 1 // 0 改为 -1
		if j, ok := memo[s]; ok {
			ans = max(ans, i-j)
		} else {
			memo[s] = i
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
