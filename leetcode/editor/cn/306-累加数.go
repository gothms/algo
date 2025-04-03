package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(len(strconv.Itoa(math.MaxInt64)))

	num := "199100199"
	num = "0"
	num = "123"
	num = "211738" // true
	number := isAdditiveNumber(num)
	fmt.Println(number)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isAdditiveNumber(num string) bool {
	// lc
	// 1.一个累加序列，当它的第一个数字和第二个数字以及总长度确定后，这整个累加序列也就确定了
	// 2.first + second = third，则 third 的计算可以仅计算 l := max(len.first, len.second) 以及 l+1
	// 3.如何处理由过大的整数输入导致的溢出？可以使用字符串的加法

	// 个人
	n := len(num)
	if n < 3 { // fast fail
		return false
	}
	// dfs：检验 a + b = c
	var dfs func(int, uint64, uint64) bool
	dfs = func(i int, a, b uint64) bool {
		if i == n {
			return true
		}
		c := a + b
		m := n
		if num[i] == '0' { // 0 开头的数，只能是 0
			m = i + 1
		}
		var v uint64
		for j := i; j < m; j++ { // 可优化：仅计算 l := max(len.first, len.second) 以及 l+1
			v = v*10 + uint64(num[j]-'0')
			if v > c { // 已经不可能，失败
				return false
			} else if v == c { // 可能成功
				return dfs(j+1, b, v)
			}
		}
		return false
	}

	a, an := uint64(0), (n-1)>>1 // 剪枝：an 的取值区间，如：num = "9990999"
	if num[0] == '0' {
		an = 1
	}
	for i := 0; i < an; i++ {
		a = a*10 + uint64(num[i]-'0')              // 计算 a
		b, bn := uint64(0), min(n-i-1, (n+i+1)>>1) // 剪枝：bn 的取值区间，如：num = "211738"
		if num[i+1] == '0' {
			bn = i + 2
		}
		for j := i + 1; j < bn; j++ {
			b = b*10 + uint64(num[j]-'0') // 计算 b
			if dfs(j+1, a, b) {
				return true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
