package main

import (
	"fmt"
)

func main() {
	num := 4 // 2
	num = 30 // 14
	num = 63 // 31
	num = 13 // 6
	num = 38 // 18
	even := countEven(num)
	fmt.Println(even)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countEven(num int) int {
	ans := num >> 1
	if func(i int) bool {
		s := 0
		for ; i > 0; i /= 10 {
			s += i % 10
		}
		return s&1 == 1
	}(num) && num&1 == 0 {
		ans--
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
func countEven_(num int) int {
	ans := num >> 1
	is := func(i int) bool {
		sum := 0
		for ; i > 0; i /= 10 {
			sum += i % 10
		}
		return sum&1 == 0
	}
	//if !is(num) && (num&1 == 0 || !is(num-1)) { // x 不满足且是偶数 / x 和 x-1 都不满足
	// 修正：x 和 x-1 都不满足，则 x 的个位一定是 0
	if !is(num) && num&1 == 0 { // x 不满足且是偶数
		return ans - 1
	}
	return ans
}
