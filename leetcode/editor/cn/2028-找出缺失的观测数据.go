package main

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	sum := mean * (n + m)
	for _, v := range rolls {
		sum -= v
	}
	if sum < n || sum > n*6 {
		return nil
	}
	val, k := sum/n, sum%n
	ans := make([]int, n)
	for i := 0; i < k; i++ {
		ans[i] = val + 1
	}
	for i := k; i < n; i++ {
		ans[i] = val
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
