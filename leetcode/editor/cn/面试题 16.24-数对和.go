package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func pairSums(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	memo := make(map[int]int)
	for _, v := range nums {
		if memo[v] > 0 {
			memo[v]--
			ans = append(ans, []int{target - v, v})
		} else {
			memo[target-v]++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
