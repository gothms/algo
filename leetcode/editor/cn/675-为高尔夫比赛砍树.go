package main

import "sort"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func cutOffTree(forest [][]int) int {
	// A* 启发式搜索

	// Dijkstra

	// bfs
	//di, dj := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}
	//m, n := len(forest), len(forest[0])
	//trees := make([][3]int, 0)
	//for i, row := range forest {
	//	for j, h := range row {
	//		if h > 1 {
	//			trees = append(trees, [3]int{h, i, j})
	//		}
	//	}
	//}
	//sort.Slice(trees, func(i, j int) bool {
	//	return trees[i][0] < trees[j][0]
	//})
	//bfs := func(x, y, tx, ty int) int {
	//	vis := make([][]bool, m)
	//	for i := range vis {
	//		vis[i] = make([]bool, n)
	//	}
	//	vis[x][y] = true
	//	q := [][3]int{{0, x, y}}
	//	for len(q) > 0 {
	//		cur := q[0]
	//		d, i, j := cur[0], cur[1], cur[2]
	//		if i == tx && j == ty {
	//			return d
	//		}
	//		d++
	//		for k := 0; k < 4; k++ {
	//			ni, nj := i+di[k], j+dj[k]
	//			if ni >= 0 && ni < m && nj >= 0 && nj < n &&
	//				forest[ni][nj] != 0 && !vis[ni][nj] {
	//				vis[ni][nj] = true
	//				q = append(q, [3]int{d, ni, nj})
	//			}
	//		}
	//		q = q[1:]
	//	}
	//	return -1
	//}
	//ans, px, py := 0, 0, 0
	//for _, t := range trees {
	//	d := bfs(px, py, t[1], t[2])
	//	if d < 0 {
	//		return -1
	//	}
	//	ans += d
	//	px, py = t[1], t[2]
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
