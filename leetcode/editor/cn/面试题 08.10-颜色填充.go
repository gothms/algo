package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	v := image[sr][sc]
	if v == newColor {
		return image
	}
	dx, dy := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}
	m, n := len(image), len(image[0])
	var dfs func(int, int)
	dfs = func(i, j int) {
		image[i][j] = newColor
		for k := 0; k < 4; k++ {
			if ni, nj := i+dx[k], j+dy[k]; ni >= 0 && ni < m && nj >= 0 && nj < n &&
				image[ni][nj] == v {
				dfs(ni, nj)
			}
		}
	}
	dfs(sr, sc)
	return image
}

//leetcode submit region end(Prohibit modification and deletion)
