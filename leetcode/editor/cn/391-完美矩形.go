package main

import (
	"math"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isRectangleCover(rectangles [][]int) bool {
	// 扫描线
	// lc-218 lc-391 lc-850

	memo := make(map[[2]int]int)
	area, minX, minY, maxX, maxY := 0, math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32
	for _, r := range rectangles {
		x1, y1, x2, y2 := r[0], r[1], r[2], r[3]
		area += (x2 - x1) * (y2 - y1) // 总面积
		memo[[2]int{x1, y1}]++        // 四个顶点
		memo[[2]int{x1, y2}]++
		memo[[2]int{x2, y1}]++
		memo[[2]int{x2, y2}]++

		minX = min(minX, x1) // 组合成的矩形
		minY = min(minY, y1)
		maxX = max(maxX, x2)
		maxY = max(maxY, y2)
	}
	if area != (maxX-minX)*(maxY-minY) || // 面积不重叠，不留空
		memo[[2]int{minX, minY}] != 1 || memo[[2]int{minX, maxY}] != 1 ||
		memo[[2]int{maxX, minY}] != 1 || memo[[2]int{maxX, maxY}] != 1 { // 四个顶点仅有一个点
		return false
	}
	delete(memo, [2]int{minX, minY}) // 清空大矩形的四个顶点
	delete(memo, [2]int{minX, maxY})
	delete(memo, [2]int{maxX, minY})
	delete(memo, [2]int{maxX, maxY})
	for _, v := range memo { // 小矩形邻接点，有 2/4 个顶点
		if v&1 > 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
