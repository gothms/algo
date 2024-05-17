package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 1
	a, b := 2, 3
	number := nthMagicalNumber(n, a, b)
	fmt.Println(number)
}

// leetcode submit region begin(Prohibit modification and deletion)
func nthMagicalNumber(n int, a int, b int) int {
	c := a * b / func(a, b int) int {
		for a != 0 {
			b, a = a, b%a
		}
		return b
	}(a, b)
	return sort.Search(min(a, b)*n, func(i int) bool {
		return i/a+i/b-i/c >= n
	}) % 1_000_000_007

	// 个人
	//const mod = 1_000_000_007
	//gcd := func(a, b int) int {
	//	for b != 0 {
	//		a, b = b, a%b
	//	}
	//	return a
	//}
	//g := gcd(a, b)
	//c := a / g * b
	//var l, r int64 = int64(min(a, b)), math.MaxInt64
	//var i, j, k, m = int64(a), int64(b), int64(c), int64(n)
	//for l < r {
	//	mid := (r-l)/2 + l
	//	if mid/i+mid/j-mid/k >= m {
	//		r = mid
	//	} else {
	//		l = mid + 1
	//	}
	//}
	//return int(l % mod)
}

//leetcode submit region end(Prohibit modification and deletion)
