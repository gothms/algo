package main

import (
	"fmt"
	"strconv"
)

func main() {
	v, _ := strconv.Atoi("-1")
	fmt.Println(v)
}

// leetcode submit region begin(Prohibit modification and deletion)
func complexNumberMultiply(num1 string, num2 string) string {
	var i1, i2 int
	for i, c := range num1 {
		if c == '+' {
			i1 = i
			break
		}
	}
	for i, c := range num2 {
		if c == '+' {
			i2 = i
			break
		}
	}
	a, _ := strconv.Atoi(num1[:i1])
	ai, _ := strconv.Atoi(num1[i1+1 : len(num1)-1])
	b, _ := strconv.Atoi(num2[:i2])
	bi, _ := strconv.Atoi(num2[i2+1 : len(num2)-1])
	return fmt.Sprintf("%d+%di", a*b-ai*bi, a*bi+b*ai)
}

//leetcode submit region end(Prohibit modification and deletion)
