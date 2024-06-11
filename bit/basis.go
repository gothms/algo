package bit

import (
	"fmt"
	"math/bits"
)

/*
异或
	x ^ 0 = x
	x ^ 1s = ~x	// 注意 1s = ~0
	x ^ (~x) = 1s
	x ^ x = 0
	c = a ^ b => a ^ c = b, b ^ c = a	// 交换两个数
	a ^ b ^ c = a ^ (b ^ c) = (a ^ b) ^ c	// associative

按位置零
	&^
*/

func Basis(x, n int) int {
	y := 0
	y = x & (^0 << n)               // 最右边n位清0
	y = (x >> n) & 1                // 获取 x 的第 n 位值（0 或者 1）
	y = x & (1 << n)                // 获取 x 的第 n 位的幂值
	y = x | (1 << n)                // 仅将第 n 位置为 1
	y = x & (^(1 << n))             // 仅将第 n 位置为 0
	y = x & ((1 << n) - 1)          // 将 x 最高位至第 n 位（含）清零
	y = x & (^((1 << (n + 1)) - 1)) // 将第 n 位至第 0 位（含）清零

	if ^x&x == 0 {
		// forever true
	}
	if -x == ^x+1 {
		// forever true
	}
	y = x & (x - 1) // 清零最低位的 1
	y = x & -x      // 得到最低位的 1

	y = x ^ (x - 1) // 低于等于“最低位的 1”的位，全置为 1

	y = x | (x + 1)  // 最低位的 0，变 1
	y = ^x & (x + 1) // 取出最低位的 0，用 1 标识
	y = (x + 1) &^ x // 同上

	return y
}

// 更多做法请参考：big.Int func (x *Int) Sign() int
// (x-1)>>31
// x 为正数：0
// 0：-1
// 负数：-1
// x-1 是为了将 0 划归为 -1
func format(x int) int {
	return (x-1)>>31 + 1
}

func api() int {
	var x uint = 4
	y := bits.LeadingZeros(x) // 61
	y = bits.TrailingZeros(x) // 2
	y = bits.Len(x)           // 3
	return y
}

// 一个 bug？
// https://github.com/golang/go/issues/24523
func uintBug() {
	var i int
	v := ^(^uint(0) << i) // yes
	v = ^(^uint(0) << 7)  // no
	fmt.Println(v)
}
