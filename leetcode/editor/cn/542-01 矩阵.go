package main

import "fmt"

func main() {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	matrix := updateMatrix(mat)
	fmt.Println(matrix)
}

// leetcode submit region begin(Prohibit modification and deletion)
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	q := make([][2]int, 0)
	for i, r := range mat {
		for j, v := range r {
			if v == 0 {
				q = append(q, [2]int{i, j})
				vis[i][j] = true
			}
		}
	}
	dx, dy := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}
	for v, l := 1, len(q); l > 0; v, l = v+1, len(q) {
		for _, p := range q[:l] {
			x, y := p[0], p[1]
			for i := 0; i < 4; i++ {
				if nx, ny := x+dx[i], y+dy[i]; nx >= 0 && nx < m && ny >= 0 && ny < n &&
					!vis[nx][ny] {
					vis[nx][ny] = true
					q = append(q, [2]int{nx, ny})
					mat[nx][ny] = v
				}
			}
		}
		q = q[l:]
	}
	return mat
}

//leetcode submit region end(Prohibit modification and deletion)
