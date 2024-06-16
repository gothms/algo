package main

import (
	"fmt"
	"math"
)

func main() {
	n := 12 // 32
	primes := []int{2, 7, 13, 19}
	number := nthSuperUglyNumber(n, primes)
	fmt.Println(number)
}

// leetcode submit region begin(Prohibit modification and deletion)
//var prime313 []int
//
//func init() {
//	const n, m = 1001, 168
//	prime313 = make([]int, 0, m)
//	memo := make([]bool, n)
//	for i := 2; i < n; i++ {
//		if memo[i] {
//			continue
//		}
//		prime313 = append(prime313, i)
//		for j := i << 1; j < n; j += i {
//			memo[j] = true
//		}
//	}
//}

func nthSuperUglyNumber(n int, primes []int) int {
	// 个人
	m := len(primes)
	ids := make([]int, m)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
		temp := make([]int, m)
		for j, k := range ids {
			v := dp[k] * primes[j]
			temp[j] = v
			dp[i] = min(dp[i], v)
		}
		for j, v := range temp {
			if v == dp[i] {
				ids[j]++
			}
		}
	}
	return dp[n-1]

	// 个人：那么 memo 的初始长度是多少呢？
	//memo := make([]bool, 100)
	//v := 1
	//for i, j, c := 0, 0, 1; c < n; {
	//	v++
	//	if memo[v] {
	//		continue
	//	}
	//	if i < len(primes) && v == primes[i] {
	//		i++
	//		j++
	//		c++
	//	} else if v == prime313[j] {
	//		j++
	//		for k := v << 1; k < len(memo); k += v {
	//			memo[k] = true
	//		}
	//	} else {
	//		c++
	//	}
	//}
	//return v
}

//leetcode submit region end(Prohibit modification and deletion)
