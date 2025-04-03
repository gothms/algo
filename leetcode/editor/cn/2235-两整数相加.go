package main

import "fmt"

func main() {
	num1 := 10
	num2 := -4
	i := sum(num1, num2)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func sum(num1 int, num2 int) int {
	for num2 != 0 {
		num1, num2 = num1^num2, num1&num2<<1
	}
	return num1
}

//leetcode submit region end(Prohibit modification and deletion)
