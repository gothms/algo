package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func satisfiesConditions(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	for i, g := range grid {
		for j, v := range g {
			if i < m-1 && v != grid[i+1][j] || j < n-1 && v == grid[i][j+1] {
				return false
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
