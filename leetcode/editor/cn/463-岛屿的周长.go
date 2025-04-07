package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func islandPerimeter(grid [][]int) int {
	// lc：dfs
	// 2 标记已访问，0 则结果 +1

	// 个人
	//dx, dy := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}
	ans, r, c := 0, len(grid), len(grid[0])
	//var dfs func(x,y int)
	//dfs= func(x, y int)  {
	//}
	for i, g := range grid {
		for j, v := range g {
			if v == 1 {
				ans += 4
				if j+1 < c && grid[i][j+1] == 1 {
					ans -= 2
				}
				if i+1 < r && grid[i+1][j] == 1 {
					ans -= 2
				}
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
