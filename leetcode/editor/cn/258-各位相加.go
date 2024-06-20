package main

import (
	"fmt"
)

func main() {
	num := 2035950041
	digits := addDigits(num)
	fmt.Println(digits)

	fmt.Println((num-10)%9 + 10)
}

// leetcode submit region begin(Prohibit modification and deletion)
func addDigits(num int) int {
	return (num-1)%9 + 1
}

//leetcode submit region end(Prohibit modification and deletion)
