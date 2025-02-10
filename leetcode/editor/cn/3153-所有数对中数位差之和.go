package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func sumDigitDifferences(nums []int) int64 {
	memo, counter := [10][10]int{}, [10]int{}
	ans := 0
	for _, v := range nums {
		for i := 0; v != 0; i++ {
			j := v % 10
			v /= 10
			memo[i][j]++
			counter[i]++
			ans += counter[i] - memo[i][j]
		}
	}
	return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
