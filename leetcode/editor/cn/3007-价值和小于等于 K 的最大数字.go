package main

import (
	"fmt"
	"sort"
)

func main() {
	k, x := 9, 1              // 6
	k, x = 7, 2               // 9
	k, x = 1, 7               // 64
	k, x = 1, 8               // 128
	k, x = 282288420250713, 8 // 112971204339755
	number := findMaximumNumber(int64(k), x)
	fmt.Println(number)

	//v := int64(math.MaxInt64)
	//val := int(v)
	//fmt.Println(v)
	//fmt.Println(val)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findMaximumNumber(k int64, x int) int64 {
	// 试填法（逐位构造）

	// 二分 + 数学
	return int64(sort.Search(int(k+1)<<x, func(v int) bool {
		v++
		s, i := 0, x-1
		for n := v >> i; n > 0; n >>= x {
			s += n >> 1 << i
			if n&1 > 0 {
				mask := 1<<i - 1
				s += v&mask + 1
			}
			i += x
		}
		return int64(s) > k
	}))

	// 个人：如果不用 uint64，则 s += v / j >> 1 * j 时，j 为 0
	//return int64(sort.Search(int(k+1)<<x, func(i int) bool {
	//	s, v := uint64(0), uint64(i+1)
	//	for j := uint64(1 << (x - 1)); j <= v; j <<= x {
	//		s += v / j >> 1 * j
	//		if m := v%(j<<1) + 1; m > j {
	//			s += m - j
	//		}
	//	}
	//	return int64(s) > k
	//}))

	// 二分 + 数位 dp
}

//leetcode submit region end(Prohibit modification and deletion)
