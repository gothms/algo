package main

import (
	"fmt"
)

func main() {
	arr := [][]int{{0, 0, 2, 2},
		{1, 0, 2, 3},
		{1, 0, 3, 1}} // 6
	arr = [][]int{{0, 0, 1000000000, 1000000000}} // 49
	arr = [][]int{{0, 0, 2, 2},
		{1, 1, 3, 3}} // 7
	arr = [][]int{{0, 0, 1, 1},
		{2, 2, 3, 3}} // 2
	arr = [][]int{{93516, 44895, 94753, 69358},
		{13141, 52454, 59740, 71232},
		{22877, 11159, 85255, 61703},
		{11917, 8218, 84490, 36637},
		{75914, 29447, 83941, 64384},
		{22490, 71433, 64258, 74059},
		{18433, 51177, 87595, 98688},
		{70854, 80720, 91838, 92304},
		{46522, 49839, 48550, 94096},
		{95435, 37993, 99139, 49382},
		{10618, 696, 33239, 45957},
		{18854, 2818, 57522, 78807},
		{61229, 36593, 76550, 41271},
		{99381, 90692, 99820, 95125}} // 971243962
	area := rectangleArea(arr)
	fmt.Println(area)
}

// leetcode submit region begin(Prohibit modification and deletion)
func rectangleArea(rectangles [][]int) int {
	// 参考：E:\gothmslee\algo\oiWiki\atlantis.go

	// 离散化 + 扫描线 + 线段树

	// 离散化 + 扫描线 + 数组

	// 排序

	// 暴力：求出总面积，再减去每两个矩形相交的面积
}

//leetcode submit region end(Prohibit modification and deletion)
