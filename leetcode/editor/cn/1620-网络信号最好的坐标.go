package main

import (
	"fmt"
	"math"
)

func main() {
	towers := [][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}}
	radius := 2 // [2,1]
	towers = [][]int{{42, 0, 0}}
	radius = 7 // [0,0]
	towers = [][]int{{0, 1, 2}, {2, 1, 2}, {1, 0, 2}, {1, 2, 2}}
	radius = 1 // [1,1]
	coordinate := bestCoordinate(towers, radius)
	fmt.Println(coordinate)
}

// leetcode submit region begin(Prohibit modification and deletion)
func bestCoordinate(towers [][]int, radius int) []int {
	x1, x2, y1, y2 := 50, 0, 50, 0
	for _, t := range towers {
		x, y := t[0], t[1]
		x1, x2 = min(x1, x), max(x2, x)
		y1, y2 = min(y1, y), max(y2, y)
	}
	xv, yv, maxV, r := 0, 0, 0, radius*radius
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			cur := 0
			for _, t := range towers {
				dx, dy := i-t[0], j-t[1]
				if d := dx*dx + dy*dy; d <= r { // 等于也算在信号范围内
					cur += int(float64(t[2]) / (1 + math.Sqrt(float64(d))))
				}
			}
			if cur > maxV {
				xv, yv, maxV = i, j, cur
			}
		}
	}
	return []int{xv, yv}
}

//leetcode submit region end(Prohibit modification and deletion)
