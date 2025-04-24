package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	ans := 0
	for i, p := range points[:len(points)-1] {
		ans = max(ans, points[i+1][0]-p[0])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
