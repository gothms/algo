package main

import (
	"fmt"
)

func main() {
	s1 := "1100011000"
	s2 := "0101001010"
	x := 2
	s1 = "100010010100111100001110101111100001001101011010100111101011100100011111110001011001001"
	s2 = "000001100010010011111101100101111011101110010001001010100101011100011110000111010011010"
	x = 6 // 38
	operations := minOperations(s1, s2, x)
	fmt.Println(operations)
}
func minOperations(s1 string, s2 string, x int) int {
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	var (
		cnt  = 0    // 不同的总数（若为奇数，返回 -1）
		id   = -500 // 上一个不同的位置
		pre  = 0    // 上上个不同的最小代价
		last = 0    // 上个不同的最小代价
	)
	for i := range s1 {
		if s1[i] == s2[i] {
			continue
		}
		cnt++
		//pre, last = last, minVal(last+x, pre+minVal(i-id, x)) // 重点：两个成对，计算会错误
		pre, last = last, minVal(last+x, pre+(i-id)<<1) // 和“最远”的一个不同互换 / 和上一个不同互换。重点：记录单个的不同的 cost
		id = i
	}
	if cnt&1 == 1 {
		return -1
	}
	return last >> 1
}
