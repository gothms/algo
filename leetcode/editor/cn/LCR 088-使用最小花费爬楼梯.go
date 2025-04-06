package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	p, c := 0, 0
	for i := 2; i <= n; i++ {
		p, c = c, min(p+cost[i-2], c+cost[i-1])
	}
	return c
}

//leetcode submit region end(Prohibit modification and deletion)
