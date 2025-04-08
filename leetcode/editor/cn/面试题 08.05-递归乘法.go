package main

import "fmt"

func main() {
	A, B := 3, 4
	i := multiply(A, B)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func multiply(A int, B int) int {
	ans := 0
	for B > 0 {
		if B&1 == 1 {
			ans += A
		}
		A <<= 1
		B >>= 1
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
