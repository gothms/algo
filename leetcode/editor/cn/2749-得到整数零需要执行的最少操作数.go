package main

import (
	"fmt"
	"math/bits"
)

func main() {
	num1, num2 := 16, 10
	num1, num2 = 85, 42
	num1, num2 = 16, 12
	zero := makeTheIntegerZero(num1, num2)
	fmt.Println(zero)
}

// leetcode submit region begin(Prohibit modification and deletion)
func makeTheIntegerZero(num1 int, num2 int) int {
	for i := 1; ; i++ { // 操作 i 次
		num1 -= num2
		// 1.bits.OnesCount(uint(num1)) <= i：如示例 1
		// 2.num1 >= i：选 0 则 2^i 值最小，为 1
		if bits.OnesCount(uint(num1)) <= i && num1 >= i {
			return i
		}
		if num2 >= -1 && num2 >= num1 { // num1 无限递减
			return -1
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
