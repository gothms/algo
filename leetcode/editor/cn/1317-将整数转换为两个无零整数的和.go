package main

import "fmt"

func main() {
	n := 11
	n = 10000
	n = 1010
	n = 19
	integers := getNoZeroIntegers(n)
	fmt.Println(integers)
}

// leetcode submit region begin(Prohibit modification and deletion)
func getNoZeroIntegers(n int) []int {
	// lc：枚举 a，得到 b，同时验证 a b 是否包含 0

	// 个人
	a, b, r := 0, 0, 0
	for pow := 1; n != 0; pow *= 10 {
		v := n%10 - r
		n /= 10
		if n == 0 && v == 0 { // 已结束
			break
		}
		r = 1 // 默认高位借位
		switch v {
		case -1: // 该位本是 0，同时低位已向该位借位
			a += pow
			b += 8 * pow
		case 0:
			a += pow
			b += 9 * pow
		case 1:
			if n == 0 { // n 的最高位是 1
				b += pow
			} else {
				a += 2 * pow
				b += 9 * pow
			}
		default:
			r = 0 // 不向高位借位
			a += pow
			b += (v - 1) * pow
		}
	}
	return []int{a, b}
}

//leetcode submit region end(Prohibit modification and deletion)
