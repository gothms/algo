package main

import "fmt"

func main() {
	a, b := 2, 3
	i := getSum(a, b)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, a&b<<1
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
