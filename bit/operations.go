package bit

import "math/bits"

/*
lc：面试题 16.09-运算
*/

func Add(a, b int) int {
	for b != 0 {
		a, b = a^b, a&b<<1
	}
	return a
}
func Substract(a, b int) int {
	return Add(a, Add(^b, 1))
}

func Multiply(a, b int) int {
	var (
		product int      // 乘积
		x       = abs(a) // 被乘数的绝对值
		y       = abs(b) // 乘数的绝对值
	)
	for y > 0 {
		if y&1 > 0 { // 被乘数二进制位为 0，不需要加；为 1，则加
			product = Add(product, x) // 乘积 += 1 * x
		}
		x <<= 1 // 由于 y >>= 1，所以 x <<= 1 方便计算
		y >>= 1 // 重点：遍历 y 的每一个二进制 1
	}
	if a^b < 0 { // 确定乘积的符号
		product = Add(^product, 1)
	}
	return product
}

func Divide(a, b int) (int, int) {
	var (
		quotient  int      // 商
		remainder int      // 余数
		x         = abs(a) // 被除数的绝对值
		y         = abs(b) // 除数的绝对值
	)
	//for i := GetBitCount(x) - 1; i >= 0; i = Substract(i, 1) {
	for i := bits.Len(uint(x)) - 1; i >= 0; i = Substract(i, 1) { // 从高位开始
		if x>>i >= y { // 够除
			quotient = Add(quotient, 1<<i) // 商 += 1 * 1<<i
			x = Substract(x, y<<i)         // 被除数 -= 1 * y<<i
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
