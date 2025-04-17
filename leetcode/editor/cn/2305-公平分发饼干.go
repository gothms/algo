package main

import (
	"fmt"
)

func main() {
	cookies := []int{8, 15, 10, 20, 8}
	k := 2
	i := distributeCookies(cookies, k)
	fmt.Println(i)

	v := 26
	for i := v; i > 0; i = (i - 1) & v {
		fmt.Printf("i:%b\n", i)
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func distributeCookies(cookies []int, k int) int {
	// 状压 dp
	return 0

	// 二分：error
	//sort.Ints(cookies)
	//s, n := 0, len(cookies)
	//for _, v := range cookies {
	//	s += v
	//}
	//i := sort.Search(s, func(i int) bool {
	//	cnt, sum, l, r := 0, 0, 0, n-1
	//	if cookies[r] > i {
	//		return false
	//	}
	//	for l <= r {
	//		for l <= r && sum+cookies[r] <= i {
	//			sum += cookies[r]
	//			r--
	//		}
	//		for l <= r && sum+cookies[l] <= i {
	//			sum += cookies[l]
	//			l++
	//		}
	//		sum = 0
	//		cnt++
	//	}
	//	return cnt <= k
	//})
	//return i
}

//leetcode submit region end(Prohibit modification and deletion)
