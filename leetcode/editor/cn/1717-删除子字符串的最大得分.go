package main

import "fmt"

func main() {
	s := "cdbcbbaaabab"
	x, y := 4, 5
	gain := maximumGain(s, x, y)
	fmt.Println(gain)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumGain(s string, x int, y int) int {
	// 贪心
	del := func(src []uint8, st *[]uint8, v int, f func(c1, c2 uint8) bool) int {
		sum := 0
		for _, c := range src {
			if f((*st)[len(*st)-1], c) {
				sum += v
				*st = (*st)[:len(*st)-1]
			} else {
				*st = append(*st, c)
			}
		}
		return sum
	}
	fx := func(c1, c2 uint8) bool { // ab
		return c1 == 'a' && c2 == 'b'
	}
	fy := func(c1, c2 uint8) bool { // ba
		return c1 == 'b' && c2 == 'a'
	}
	stXY := make([]uint8, 1, len(s)+1)
	if x > y {
		sX := del([]uint8(s), &stXY, x, fx)       // 优先 ab
		return sX + del(stXY, &[]uint8{0}, y, fy) // 再 ba
	} else {
		sY := del([]uint8(s), &stXY, y, fy)       // 优先 ba
		return sY + del(stXY, &[]uint8{0}, x, fx) // 再 ab
	}

	// 个人：优化
	//del := func(src []uint8, st *[]uint8, v int, f func(c1, c2 uint8) bool) int {
	//	sum := 0
	//	for _, c := range src {
	//		if f((*st)[len(*st)-1], c) {
	//			sum += v
	//			*st = (*st)[:len(*st)-1]
	//		} else {
	//			*st = append(*st, c)
	//		}
	//	}
	//	return sum
	//}
	//fx := func(c1, c2 uint8) bool {
	//	return c1 == 'a' && c2 == 'b'
	//}
	//fy := func(c1, c2 uint8) bool {
	//	return c1 == 'b' && c2 == 'a'
	//}
	//n := len(s)
	//stX, stY := make([]uint8, 1, n+1), make([]uint8, 1, n+1)
	//sX, sY := del([]uint8(s), &stX, x, fx), del([]uint8(s), &stY, y, fy)
	//
	//st := make([]uint8, 1, max(len(stX), len(stY))+1)
	//sX += del(stX[1:], &st, y, fy)
	//st = st[:1]
	//sY += del(stY[1:], &st, x, fx)
	//return max(sX, sY)
}

//leetcode submit region end(Prohibit modification and deletion)
