package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	num := 2
	num = 1
	num = 2147483647
	num = 67 // [69,56]
	numbers := findClosedNumbers(num)
	fmt.Println(numbers)

	a, b := 2147483647, 3221225471
	fmt.Println(bits.OnesCount(uint(a)), bits.OnesCount(uint(b)))
	fmt.Printf("%b,%b\n", a, b)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findClosedNumbers(num int) []int {
	// 略小：100111 变 011110
	// 写法一
	//onesCnt := bits.TrailingZeros(uint(num + 1))
	//mn := num >> onesCnt << onesCnt // 100000
	//tzc := bits.TrailingZeros(uint(mn))
	//mn &= mn - 1                                      // 0：移除最后一个 1
	//mn |= (1<<(onesCnt+1) - 1) << (tzc - onesCnt - 1) // |= 1111 << 1
	// 写法二
	onesCnt := bits.TrailingZeros(uint(num + 1))
	mn := num >> onesCnt // 100
	zCnt := bits.TrailingZeros(uint(mn))
	mn = (1<<(onesCnt+2)-1)<<(zCnt-1) ^ mn<<onesCnt // 111110 ^ 100000
	if mn <= 0 {
		mn = -1
	}
	// 略大：101110 变 110011
	zeroCnt := bits.TrailingZeros(uint(num))
	mx := num + num&-num                                  // 110000
	mx |= 1<<(bits.TrailingZeros(uint(mx))-zeroCnt-1) - 1 // |= 11
	if mx >= math.MaxInt32 {
		mx = -1
	}
	return []int{mx, mn}
}

//leetcode submit region end(Prohibit modification and deletion)
