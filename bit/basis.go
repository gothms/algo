package bit

/*
异或
	x ^ 0 = x
	x ^ 1s = ~x	// 注意 1s = ~0
	x ^ (~x) = 1s
	x ^ x = 0
	c = a ^ b => a ^ c = b, b ^ c = a	// 交换两个数
	a ^ b ^ c = a ^ (b ^ c) = (a ^ b) ^ c	// associative
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
	if x == ^x+1 {
		// forever true
	}
	y = x & (x - 1) // 清零最低位的 1
	y = x & -x      // 得到最低位的 1

	y = x | (x + 1)  // 最低位的0，变1
	y = ^x & (x + 1) // 取出最低位的0，用1标识
	y = (x + 1) &^ x // 同上
	return y
}
