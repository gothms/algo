package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findPairs(nums []int, k int) int {
	memo := make(map[int]int)
	for _, v := range nums {
		memo[v]++
	}
	ans := 0
	if k == 0 {
		for _, c := range memo {
			if c > 1 {
				ans++
			}
		}
	} else {
		for v := range memo {
			if memo[v-k] > 0 {
				ans++
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
