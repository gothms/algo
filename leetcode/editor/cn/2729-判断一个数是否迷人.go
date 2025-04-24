package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isFascinating(n int) bool {
	memo := [10]bool{}
	memo[0] = true
	for i := 1; i <= 3; i++ {
		for v := i * n; v != 0; v /= 10 {
			idx := v % 10
			if memo[idx] {
				return false
			}
			memo[idx] = true
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
