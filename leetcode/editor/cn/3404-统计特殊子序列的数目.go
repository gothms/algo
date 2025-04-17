package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 3, 6, 1}
	nums = []int{3, 4, 3, 4, 3, 4, 3, 4}
	subsequences := numberOfSubsequences(nums)
	fmt.Println(subsequences)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfSubsequences(nums []int) int64 {
	// 枚举右，维护左

	// 前后缀
	// a*c=b*d -> a/b = d/c
	//gcd := func(a, b int) int {
	//	for b != 0 {
	//		a, b = b, a%b
	//	}
	//	return a
	//}
	//ans, n := 0, len(nums)
	//suf := make(map[[2]int]int)  // 后缀
	//for r, v := range nums[6:] { // 枚举 s
	//	for _, w := range nums[4 : r+5] { // 枚举r
	//		g := gcd(v, w)
	//		suf[[2]int{v / g, w / g}]++
	//	}
	//}
	//for r, v := range nums[2 : n-4] { // 枚举 q
	//	for _, w := range nums[:r+1] { // 枚举 p
	//		g := gcd(v, w)
	//		ans += suf[[2]int{w / g, v / g}]
	//	}
	//	l := r + 4
	//	y := nums[l]                   // r
	//	for _, x := range nums[l+2:] { // s
	//		g := gcd(x, y)
	//		suf[[2]int{x / g, y / g}]-- // 移除 r s
	//	}
	//}
	//return int64(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
