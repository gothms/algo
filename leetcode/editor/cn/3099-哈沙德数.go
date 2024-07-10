package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func sumOfTheDigitsOfHarshadNumber(x int) int {
	s := 0
	for v := x; v > 0; v /= 10 {
		s += v % 10
	}
	if x%s == 0 {
		return s
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
