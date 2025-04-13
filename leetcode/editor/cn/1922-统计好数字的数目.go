package main

import (
	"fmt"
	"math"
)

func main() {
	//formatInt := math.Log10(float64(106856))
	//fmt.Println(formatInt)

	var n int64 = 10000
	numbers := countGoodNumbers(n)
	fmt.Println(numbers)

	ans, mod := 1, int(1e9+7)
	for i := 0; i < 10000; i++ {
		ans = ans * (5 - i&1) % mod
	}
	fmt.Println(ans) // 564908303 325891746

	v := math.Log(math.MaxInt64) / math.Log(20) // 14.576827429064828
	fmt.Println(math.MaxInt64)
	fmt.Println(math.Pow(20, 14.576))
	fmt.Println(v)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countGoodNumbers(n int64) int {
	// 1-5：5 20 100 400 2000
	// 快速幂
	const mod = 1_000_000_007
	pow := func(x int, n int64) int {
		ret := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				ret = ret * x % mod
			}
			x = x * x % mod
		}
		return ret
	}
	return pow(5, (n+1)>>1) * pow(4, n>>1) % mod
}

//leetcode submit region end(Prohibit modification and deletion)
