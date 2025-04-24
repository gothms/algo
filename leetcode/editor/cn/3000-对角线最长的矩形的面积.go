package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func areaOfMaxDiagonal(dimensions [][]int) int {
	area, dia := 0, 0
	for _, d := range dimensions {
		x, y := d[0], d[1]
		if cur := x*x + y*y; cur > dia || cur == dia && x*y > area {
			area, dia = x*y, cur
		}
	}
	return area
}

//leetcode submit region end(Prohibit modification and deletion)
