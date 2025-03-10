package main

import (
	"fmt"
	"math"
)

func main() {
	num := 240
	k := 2
	substrings := divisorSubstrings(num, k)
	fmt.Println(substrings)
}

// leetcode submit region begin(Prohibit modification and deletion)
func divisorSubstrings(num int, k int) int {
	ans, p := 0, int(math.Pow10(k-1))
	n, v := num, num%p
	for n /= p; n != 0; n /= 10 {
		v += n % 10 * p
		if v != 0 && num%v == 0 {
			ans++
		}
		v /= 10
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
