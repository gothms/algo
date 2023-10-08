package main

import (
	"fmt"
)

func main() {
	//fmt.Println(1<<30 > 1_000_000_000)
	//fmt.Println(bits.Len(1_000_000_000))
	//fmt.Printf("%b\n", 1_000_000_000)
	//fmt.Printf("%b\n", 1<<30)

	nums := []int{2, 6, 5, 8}
	sum := maxSum(nums, 2)
	fmt.Println(sum)
}

/*
分析
	& 运算：总 1 不变 / 变少
	| 运算：总 1 变少 / 不变，把 1 挪到一个数字上
*/

func maxSum(nums []int, k int) int {
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	const N, mod = 30, 1_000_000_007
	topK, bit := make([]int, k), [N]int{} // 1_000_000_000 有 30 位
	for _, v := range nums {
		for i := 0; v != 0; i++ { // 统计位的数目
			bit[i] += v & 1
			v >>= 1
		}
	}
	ret, i := 0, N-1
	for i >= 0 && bit[i] == 0 { // 最大二进制位
		i--
	}
	for j := i; j >= 0; j-- {
		v := 1 << j                              // 位值
		for c := 0; c < minVal(bit[j], k); c++ { // 上限：minVal(bit[j], k)
			topK[c] |= v // 重组 k 个最大值
		}
	}
	for _, v := range topK { // 求平方和
		ret = (ret + v*v) % mod
	}
	return ret
}
