package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumScore(a int, b int, c int) int {
	mx, tot := max(a, b, c), a+b+c
	if mx > tot-mx {
		return tot - mx
	}
	return tot >> 1
}

//leetcode submit region end(Prohibit modification and deletion)
