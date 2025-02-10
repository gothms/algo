package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumDistance(points [][]int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	//dis, x, y := 0, -1, -1
	//for i, p1 := range points {
	//	for j, p2 := range points[i+1:] {
	//		d := abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
	//	}
	//}
}

//leetcode submit region end(Prohibit modification and deletion)
