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

	y = x ^ (x - 1) // 小于等于“最低位的 1”的位，全置为 1

	y = x | (x + 1)  // 最低位的 0，变 1
	y = ^x & (x + 1) // 取出最低位的 0，用 1 标识
	y = (x + 1) &^ x // 同上

	y = x & (x - 1) // y==0 判断 x 为 2 的幂
	y = ^x + 1      // 即 y = -x，^x+1 = -x
	//(x ^ x >> size(x) - 1) - (x >> size(x) - 1)	// 取绝对值

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

// convertInteger 整数转换，确定需要改变几个位才能将整数A转成整数B
func convertInteger(A int, B int) int {
	//return bits.OnesCount(uint(A&0xffffffff ^ B&0xffffffff))
	//return bits.OnesCount(uint(A&math.MaxUint32 ^ B&math.MaxUint32))
	return bits.OnesCount(uint(uint32(A) ^ (uint32(B))))
}

// rangeBits 遍历值 bit 的二进制位 1 的所有子集
// 如 bit=13，k=13,12,9,8,5,4,1
func rangeBits(bit int) {
	for k := bit; k != 0; k = bit & (k - 1) {
		fmt.Println(k)
	}
}

// 一个 bug？
// https://github.com/golang/go/issues/24523
func uintBug() {
	var i = 7
	v := ^(^uint(0) << i) // yes
	v = ^(^uint(0) << 7)  // no
	fmt.Println(v)
}

// 返回 big 和 small，它们是与 x 二进制表达式中 1 的个数相同，且大小最接近的那两个数（一个略大，一个略小）
func bigAndSmall(x int) (int, int) {
	// 略大：101110 变 110011
	zeroCnt := bits.TrailingZeros(uint(x))
	big := x + x&-x                                         // 110000
	big |= 1<<(bits.TrailingZeros(uint(big))-zeroCnt-1) - 1 // |= 11

	// 略小：100111 变 011110
	onesCnt := bits.TrailingZeros(uint(x + 1))
	small := x >> onesCnt // 100
	zCnt := bits.TrailingZeros(uint(small))
	small = (1<<(onesCnt+2)-1)<<(zCnt-1) ^ small<<onesCnt // 111110 ^ 100000

	return big, small
}
func swap(x, y int) (int, int) {
	x ^= y
	y ^= x
	x ^= y
	return x, y
}
