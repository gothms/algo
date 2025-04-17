package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumCardPickup(cards []int) int {
	memo := make(map[int]int)
	ans := 100_001
	for i, v := range cards {
		if l, ok := memo[v]; ok {
			ans = min(ans, i-l+1)
		}
		memo[v] = i
	}
	if ans == 100_001 {
		return -1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
