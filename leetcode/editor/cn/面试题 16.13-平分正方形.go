package main

import (
	"fmt"
	"math"
)

func main() {
	square1 := []int{68, 130, 64}
	square2 := []int{-230, 194, 7} // [-230.0,197.88055,132.0,158.52067]
	squares := cutSquares(square1, square2)
	fmt.Println(squares)
}

// leetcode submit region begin(Prohibit modification and deletion)
func cutSquares(square1 []int, square2 []int) []float64 {
	// 几何
	d1, d2 := float64(square1[2])/2, float64(square2[2])/2
	xm1, ym1 := float64(square1[0])+d1, float64(square1[1])+d1 // s1 的中心
	xm2, ym2 := float64(square2[0])+d2, float64(square2[1])+d2 // s2 的中心
	if xm1 == xm2 {                                            // 分割线与 y 平行
		return []float64{xm1, float64(min(square1[1], square2[1])), xm1, float64(max(square1[1]+square1[2], square2[1]+square2[2]))}
	} else if ym1 == ym2 { // 分割线与 x 平行
		return []float64{float64(min(square1[0], square2[0])), ym1, float64(max(square1[0]+square1[2], square2[0]+square2[2])), ym1}
	}
	dx, dy := xm2-xm1, ym2-ym1                      // 用于求斜率 / 反斜率 / 判断与 x 轴的夹角与 45° 的关系
	reverse := dx > 0 && dy < 0 || dx < 0 && dy > 0 // 斜率是否为负值
	dx, dy = math.Abs(dx), math.Abs(dy)
	if dx > dy { // 与 x 轴夹角 < 45
		slope := dy / dx // 斜率
		// 两条与 y 平行的线，将与分割线相交
		xMin, xMax := float64(min(square1[0], square2[0])), float64(max(square1[0]+square1[2], square2[0]+square2[2]))
		if reverse { // 斜率为负
			ya := ym1 + slope*xm1 // 与 y 轴的交点
			return []float64{xMin, ya - slope*xMin, xMax, ya - slope*xMax}
		} else {
			ya := ym1 - slope*xm1
			return []float64{xMin, ya + slope*xMin, xMax, ya + slope*xMax}
		}
	} else { // 与 x 轴夹角 > 45
		reSlope := dx / dy
		yMin, yMax := float64(min(square1[1], square2[1])), float64(max(square1[1]+square1[2], square2[1]+square2[2]))
		if reverse {
			xa := xm1 + reSlope*ym1
			return []float64{xa - reSlope*yMax, yMax, xa - reSlope*yMin, yMin} // 更小的 x 在前：xa - reSlope*yMax
		} else {
			xa := xm1 - reSlope*ym1
			return []float64{xa + reSlope*yMin, yMin, xa + reSlope*yMax, yMax}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
