package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countBalls(lowLimit int, highLimit int) int {
	ans := 0
	memo := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		v := 0
		for j := i; j > 0; j /= 10 {
			v += j % 10
		}
		memo[v]++
		ans = max(ans, memo[v])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
