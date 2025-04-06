package main

import "math/bits"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)

var prime762 []bool

func init() {
	n := 32
	prime762 = make([]bool, n)
	prime762[1] = true
	for i := 2; i < n; i++ {
		if prime762[i] {
			continue
		}
		for j := i << 1; j < n; j += i {
			prime762[j] = true
		}
	}
}
func countPrimeSetBits(left int, right int) int {
	// lc
	// 1 <= left <= right <= 106，则二进制中 1 的个数不会超过 19 个
	// 用 mask=665772=10100010100010101100 来存储这些质数 2,3,5,7,11,13,17,19
	// 1<<bits.OnesCount(uint(x))&665772 != 0 判断是否为质数

	ans := 0
	for i := left; i <= right; i++ {
		if !prime762[bits.OnesCount(uint(i))] {
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
