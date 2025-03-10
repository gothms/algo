package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxIncreaseKeepingSkyline(grid [][]int) int {
	ans := 0
	rowMax := make([]int, len(grid))
	colMax := make([]int, len(grid[0]))
	for i, row := range grid {
		for j, x := range row {
			rowMax[i] = max(rowMax[i], x)
			colMax[j] = max(colMax[j], x)
		}
	}

	for i, row := range grid {
		for j, x := range row {
			ans += min(rowMax[i], colMax[j]) - x
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
