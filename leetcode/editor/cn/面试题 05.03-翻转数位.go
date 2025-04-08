package main

import (
	"fmt"
	"math/bits"
)

func main() {
	num := -1 // 32
	//num = -3  // 32
	//fmt.Printf("num %b\n", num)
	i := reverseBits(num)
	fmt.Println(i)

	//num := -3
	//cur := bits.TrailingZeros(uint(num)) // 正常
	//cur = bits.TrailingZeros(uint(-3))   // 报错
	//fmt.Println(cur)
}

// leetcode submit region begin(Prohibit modification and deletion)
func reverseBits(num int) int {
	// 迭代 32 位

	// 个人
	var v uint32
	if num < 0 {
		v = ^uint32(^num)
	} else {
		v = uint32(num)
	}
	for ans, last := 0, 0; ; {
		cur := bits.TrailingZeros(uint(v + 1)) // num = -1 时 ans = 65
		ans, last = max(ans, last+cur+1), cur
		if v >>= cur + 1; v == 0 {
			return min(ans, 32) // 所以最多是 32
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
