package main

import (
	"fmt"
	"sort"
)

func main() {
	points := [][]int{{2, 3}, {1, 2}}
	w := 0
	points = [][]int{{2, 1}, {1, 0}, {1, 4}, {1, 8}, {3, 5}, {4, 6}}
	w = 1
	coverPoints := minRectanglesToCoverPoints(points, w)
	fmt.Println(coverPoints)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	ans, n := 0, len(points)
	for i := 0; i < n; i++ {
		x := points[i][0]
		ans++
		for i+1 < n && x+w >= points[i+1][0] {
			i++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
