package main

import "fmt"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	// 参见 836：矩形重叠
	area := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	if dx, dy := min(ax2, bx2)-max(ax1, bx1), min(ay2, by2)-max(ay1, by1); dx > 0 && dy > 0 {
		area -= dx * dy
	}
	return area
}

//leetcode submit region end(Prohibit modification and deletion)
