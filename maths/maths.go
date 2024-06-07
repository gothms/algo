package maths

import (
	"fmt"
	"math"
)

/*
质数
	埃氏筛法

lc
	754：sn
*/

// sn 总和为 s，则 1+2+...+n >= s，1+2+...+m <= s
func sn(s int) (int, int) {
	n := int(math.Ceil((-1 + math.Sqrt(float64(8*s+1))) / 2))
	m := int(math.Sqrt(float64(2*n)+0.25) - 0.5)

	sum := n * (1 + n) >> 1
	fmt.Println(n, sum, s)
	sum = m * (1 + m) >> 1
	fmt.Println(n, sum, s)
	return n, m
}
