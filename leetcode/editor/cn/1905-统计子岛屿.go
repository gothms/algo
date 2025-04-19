package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	// bfs

	// dfs
	//dx, dy := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}
	//m, n := len(grid2), len(grid2[0])
	//ans := 0
	//var dfs func(int, int) bool
	//dfs = func(x, y int) bool {
	//	grid2[x][y] = 0
	//	var ret bool
	//	if grid1[x][y] == 1 { // 判断子岛屿
	//		ret = true
	//	}
	//	for k := 0; k < 4; k++ {
	//		nx, ny := x+dx[k], y+dy[k]
	//		if nx >= 0 && nx < m && ny >= 0 && ny < n &&
	//			grid2[nx][ny] == 1 {
	//			//ret = ret && dfs(nx, ny)	// error：dfs 很可能不运行
	//			ret = dfs(nx, ny) && ret
	//		}
	//	}
	//	return ret
	//}
	//for i, row := range grid2 {
	//	for j, v := range row {
	//		if v == 1 && dfs(i, j) {
	//			ans++
	//		}
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
