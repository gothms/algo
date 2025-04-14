package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func escapeGhosts(ghosts [][]int, target []int) bool {
	// 个人：曼哈顿距离
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	x, y := target[0], target[1]
	dis := abs(x) + abs(y)
	for _, g := range ghosts {
		if abs(g[0]-x)+abs(g[1]-y) <= dis {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
