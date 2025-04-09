package main

import (
	"fmt"
	"math/bits"
)

func main() {
	con := Operations{}
	ret := con.Divide(2147483647, -1) // -2147483647
	ret = con.Divide(-2147483648, 1)  // -2147483648
	fmt.Println(ret)
}

// leetcode submit region begin(Prohibit modification and deletion)
type Operations struct {
}

func Constructor() Operations {
	return Operations{}
}

func (this *Operations) Minus(a int, b int) int {
	return a + ^b + 1
}

func (this *Operations) Multiply(a int, b int) int {
	var (
		product int
		fu      = a^b < 0
	)
	if a < 0 {
		a = ^a + 1
	}
	if b < 0 {
		b = ^b + 1
	}
	for b > 0 {
		if b&1 == 1 {
			product += a
		}
		a <<= 1
		b >>= 1
	}
	if fu {
		return ^product + 1
	}
	return product
}

func (this *Operations) Divide(a int, b int) int {
	var (
		quotient int
		fu       = a^b < 0
	)
	if a < 0 {
		a = ^a + 1
	}
	if b < 0 {
		b = ^b + 1
	}
	for i := bits.Len(uint(a)) - 1; i >= 0; i = this.Minus(i, 1) {
		if a>>i >= b {
			quotient += 1 << i
			a = this.Minus(a, b<<i)
		}
	}
	if fu {
		return ^quotient + 1
	}
	return quotient
}

/**
 * Your Operations object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Minus(a,b);
 * param_2 := obj.Multiply(a,b);
 * param_3 := obj.Divide(a,b);
 */
//leetcode submit region end(Prohibit modification and deletion)
