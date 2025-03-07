package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 在 x，y 轴上的投影，有交集
	return max(rec1[0], rec2[0]) < min(rec1[2], rec2[2]) &&
		max(rec1[1], rec2[1]) < min(rec1[3], rec2[3])
}

//leetcode submit region end(Prohibit modification and deletion)
