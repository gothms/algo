package main

import (
	"fmt"
	"math/bits"
	"math/rand"
)

func main() {
	grid := [][]int{
		{9, 1, 7},
		{8, 9, 2},
		{3, 4, 6},
	}
	values := findMissingAndRepeatedValues(grid)
	fmt.Println(values)

	n := 10
	ret := make([]int, 0, n)
	memo := make([][2]int, 0, n)
	sum := make([]int, 0, n)
	for i := 2; i <= 10; i++ {
		a := rand.Intn(i*i) + 1
		b := a
		for a == b {
			b = rand.Intn(i*i) + 1
		}
		xor := a
		for j := i * i; j >= 1; j-- {
			if j != b {
				xor ^= j
			}
		}
		if i&1 == 0 {
			xor ^= i * i
		} else {
			xor ^= 1
		}
		ret = append(ret, xor)
		memo = append(memo, [2]int{a, b})
		sum = append(sum, a-b)
	}
	fmt.Println(ret)
	fmt.Println(sum)
	fmt.Println(memo)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findMissingAndRepeatedValues(grid [][]int) []int {
	// 位运算：参考 260
	xor, s, n := 0, 0, len(grid)*len(grid)
	for _, g := range grid {
		for _, v := range g {
			xor ^= v
			s += v
		}
	}
	if n&1 == 0 {
		xor ^= n
	} else {
		xor ^= 1
	}
	k := bits.TrailingZeros(uint(xor))
	ans := make([]int, 2)
	for _, g := range grid {
		for _, v := range g {
			ans[v>>k&1] ^= v
		}
	}
	for i := 1; i <= n; i++ {
		ans[i>>k&1] ^= i
	}
	if (s-n*(n+1)>>1)^(ans[0]-ans[1]) < 0 {
		ans[0], ans[1] = ans[1], ans[0]
	}
	return ans

	//n := len(grid) * len(grid)
	//memo := make([]bool, n+1)
	//var s, a int
	//for _, g := range grid {
	//	for _, v := range g {
	//		s += v
	//		if memo[v] {
	//			a = v
	//		} else {
	//			memo[v] = true
	//		}
	//	}
	//}
	//return []int{a, n*(n+1)>>1 - s + a}
}

//leetcode submit region end(Prohibit modification and deletion)
