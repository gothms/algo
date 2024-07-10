package main

import (
	"fmt"
	"sort"
)

func main() {
	p1, p2, p3, p4 := []int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}
	//p1, p2, p3, p4 = []int{0, 0}, []int{0, 0}, []int{0, 0}, []int{0, 0}
	square := validSquare(p1, p2, p3, p4)
	fmt.Println(square)
}

// leetcode submit region begin(Prohibit modification and deletion)
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	distance := func(p1, p2 []int) int {
		dx, dy := p1[0]-p2[0], p1[1]-p2[1]
		return dx*dx + dy*dy
	}
	point := [][]int{p1, p2, p3, p4}
	sort.Slice(point, func(i, j int) bool { // 排序
		return point[i][0] < point[j][0] || point[i][0] == point[j][0] && point[i][1] < point[j][1]
	})
	edge := distance(point[0], point[1])
	return edge > 0 && // 边长
		edge == distance(point[2], point[3]) && // 对边
		edge == distance(point[0], point[2]) && // 邻边
		distance(point[0], point[3]) == distance(point[1], point[2]) // 对角线
}

//leetcode submit region end(Prohibit modification and deletion)
