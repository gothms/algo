package bit

func Add(a, b int) int {
	for b != 0 {
		a, b = a^b, a&b<<1
	}
	return a
}
func Substract(a, b int) int {
	return Add(a, ^b+1)
}
func Multiply(a, b int) int {
	product, x, y := 0, abs(a), abs(b) // 被乘数、乘数，取绝对值
	for y > 0 {
		if y&1 > 0 { // 被乘数二进制位为 0，不需要加；为 1，则加
			product = Add(product, x)
		}
		x <<= 1
		y >>= 1
	}
	if a^b < 0 { // 确定乘积的符号
		product = Add(^product, 1)
	}
	return product
}

/*
取模 vs 取余
	5%3 = -2
*/

func Divide(a, b int) (int, int) {
	quotient, remainder, x, y := 0, 0, abs(a), abs(b) // 商，余数，被除数和除数的绝对值
	for i := GetBitCount(x) - 1; i >= 0; i = Substract(i, 1) {
		if x>>i >= y {
			quotient = Add(quotient, 1<<i)
			x = Substract(x, y<<i)
		}
	}
	if a^b < 0 {
		quotient = Add(^quotient, 1)
	}
	if a < 0 { // 余数的符号，和被除数相同
		remainder = Add(^x, 1)
	} else {
		remainder = x
	}
	return quotient, remainder
}
func abs(n int) int {
	if n < 0 {
		return Add(^n, 1)
	}
	return n
}
