package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//s := "1001000110000100111001110010101"
	//fmt.Println(len(s))
	//v := 1
	//for i := 0; v < math.MaxInt32; i++ {
	//	fmt.Printf("%0*b,%[2]d\n", 31, v)
	//	v *= 5
	//}

	s := "1011"
	s = "111"
	s = "100111000110111" // 4
	substrings := minimumBeautifulSubstrings(s)
	fmt.Println(substrings)

	//fmt.Printf("%b,%b\n", 5005, 5005%5)
	//fmt.Println(math.Log(5005) / math.Log(5))
	//fmt.Println(math.Pow(5, 5.292650698694229))
}

// leetcode submit region begin(Prohibit modification and deletion)

var pow2767 []string

func init() {
	for v := 1; v < math.MaxInt32; v *= 5 {
		pow2767 = append(pow2767, strconv.FormatUint(uint64(v), 2))
	}
}

// var pow2767 map[int]struct{}
//
//	func init() {
//		pow2767 = map[int]struct{}{}
//		for v := 1; v <= math.MaxInt32; v *= 5 {
//			pow2767[v] = struct{}{}
//		}
//	}
func minimumBeautifulSubstrings(s string) int {
	// lc
	n := len(s)
	dp := make([]int, n+1)
	for i, c := range s {
		if c == '0' {
			dp[i+1] = n + 1
			continue
		}
		dp[i+1] = dp[i] + 1
		for _, t := range pow2767 {
			m := len(t)
			if i+1-m < 0 {
				break
			}
			if s[i+1-m:i+1] == t {
				dp[i+1] = min(dp[i+1], dp[i+1-m]+1)
			}
		}
	}
	if dp[n] > n {
		return -1
	}
	return dp[n]

	// dp
	// 前导和结尾都不能是 0
	//if s[0] == '0' {
	//	return -1
	//}
	//n := len(s)
	//dp := make([]int, n+1)
	//for i, c := range s {
	//	if c == '0' {
	//		dp[i+1] = n + 1
	//		continue
	//	}
	//	dp[i+1] = dp[i] + 1 // c == '1'
	//	v := 1
	//	for j := i - 1; j >= 0; j-- {
	//		if s[j] == '0' {
	//			continue
	//		}
	//		v |= 1 << (i - j)
	//		if _, ok := pow2767[v]; ok {
	//			dp[i+1] = min(dp[i+1], dp[j]+1)
	//		}
	//	}
	//}
	//if dp[n] > n {
	//	return -1
	//}
	//return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
