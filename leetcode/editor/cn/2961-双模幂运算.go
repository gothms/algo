package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func getGoodIndices(variables [][]int, target int) []int {
	pow := func(x, n, mod int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	ans := make([]int, 0)
	for i, v := range variables {
		if pow(pow(v[0], v[1], 10), v[2], v[3]) == target {
			ans = append(ans, i)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
