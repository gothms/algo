//二维矩阵 grid 由 0 （土地）和 1 （水）组成。岛是由最大的4个方向连通的 0 组成的群，封闭岛是一个 完全 由1包围（左、上、右、下）的岛。
//
// 请返回 封闭岛屿 的数目。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,
//0,1],[1,1,1,1,1,1,1,0]]
//输出：2
//解释：
//灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。
//
// 示例 2：
//
//
//
//
//输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
//输出：1
//
//
// 示例 3：
//
//
//输入：grid = [[1,1,1,1,1,1,1],
//             [1,0,0,0,0,0,1],
//             [1,0,1,1,1,0,1],
//             [1,0,1,0,1,0,1],
//             [1,0,1,1,1,0,1],
//             [1,0,0,0,0,0,1],
//             [1,1,1,1,1,1,1]]
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= grid.length, grid[0].length <= 100
// 0 <= grid[i][j] <=1
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 196 👎 0

package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 0, 1, 1, 0, 1, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 0, 1, 1, 1, 0},
		{1, 0, 1, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 1, 0, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 1, 0, 1, 0, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 0},
	}
	//grid = [][]int{
	//{1, 1, 0, 1, 1, 1, 1, 1, 1, 1},
	//{0, 0, 1, 0, 0, 1, 0, 1, 1, 1},
	//{1, 0, 1, 0, 0, 0, 1, 0, 1, 0},
	//{1, 1, 1, 1, 1, 0, 0, 1, 0, 0},
	//{1, 0, 1, 0, 1, 1, 1, 1, 1, 0},
	//{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	//{1, 0, 1, 0, 0, 0, 0, 1, 1, 0},
	//{1, 1, 0, 0, 1, 1, 0, 0, 0, 0},
	//{0, 0, 0, 1, 1, 0, 1, 1, 1, 0},
	//{1, 1, 0, 1, 0, 1, 0, 0, 1, 0},
	//}
	island := closedIsland(grid)
	fmt.Println(island)
}

/*
思路：dfs
	1.同类型 lc-200，不同的是边界上的岛屿有效与否
		按照 200 的解法，只需要去除边界上的岛屿即可
	2.当dfs查找到的 i j 索引越界，说明发起此次 dfs 的 i j 与边界联通
		返回 false，表示无效的岛屿
	3.所以发起 dfs 的 i j 必须在 4 个方向查找的结果都是 true，这个岛屿才有效
思路：并查集
	关键在于并集的记录
*/
// leetcode submit region begin(Prohibit modification and deletion)
func closedIsland(grid [][]int) int {
	// dfs
	cnt, n, m := 0, len(grid), len(grid[0])
	d := [5]int{0, -1, 0, 1, 0}
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i < 0 || i == n || j < 0 || j == m {
			return false // 边界上的 0
		}
		if grid[i][j] > 0 { // 已搜索 / 碰到 1
			return true
		}
		grid[i][j] = 2 // 已搜索
		island := true
		for k := 0; k < 4; k++ {
			island = dfs(i+d[k], j+d[k+1]) && island // island 要在后面
		}
		return island
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 && dfs(i, j) {
				cnt++
			}
		}
	}
	return cnt

// 并查集，没写完
//n, m := len(grid), len(grid[0])
////ok := func(i, j int) bool {
////	return i >= 0 && i < n && j >= 0 && j < m && grid[i][j] == 0
////}
//isEdge := func(i, j int) bool {
//	return i == 0 || i == n-1 || j == 0 || j == m-1
//}
//uni := make([]int, n*m)
//for i := range uni {
//	uni[i] = i
//}
////d := [5]int{0, -1, 0, 1, 0}
//for i := 0; i < n; i++ {
//	for j := 0; j < m; j++ {
//		if grid[i][j] > 0 {
//			continue
//		}
//		if isEdge(i, j) {
//			uni[i*m+j] = 0
//			//continue
//		}
//		//x := i*m + j
//		//for k := 0; k < 4; k++ {
//		//	if ni, nj := i+d[k], j+d[k+1]; ok(ni, nj) {
//		//y := ni*m + nj
//		//if isEdge(ni, nj) {
//		//	uni[y] = 0
//		//}
//		//if uni[y] == 0 {
//		//	x, y = y, x
//		//}
//		//islandUnion(uni, x, y)
//		//islandUnion(uni, i*m+j, ni*m+nj)
//		//if i == 1 {
//		//	uni[(i-1)*n+j] = 0
//		//}
//		//if j == 1 {
//		//	uni[i*n+j-1] = 0
//		//}
//		if i > 0 && grid[i-1][j] == 0 {
//			islandUnion(uni, (i-1)*n+j, i*n+j)
//		}
//		if j > 0 && grid[i][j-1] == 0 {
//			islandUnion(uni, i*n+j-1, i*n+j)
//		}
//		//	}
//		//}
//	}
//}
//counter := make(map[int]bool)
//for i := 0; i < n; i++ {
//	for j := 0; j < m; j++ {
//		if grid[i][j] == 0 {
//			counter[islandFind(uni, i*m+j)] = true
//		}
//	}
//}
//for i := 0; i < n; i++ {
//	fmt.Println(uni[i*m : (i+1)*m])
//}
//fmt.Println(counter)
//cnt := len(counter)
//if counter[0] {
//	cnt--
//}
//return cnt
}
func islandFind(uni []int, p int) int {
	//if p != uni[p] {
	//	p, uni[p] = uni[p], uni[uni[p]]
	//}
	for p != uni[p] {
		p, uni[p] = uni[p], uni[uni[p]]
	}
	return p
}
func islandUnion(uni []int, p, q int) bool {
	rootP, rootQ := islandFind(uni, p), islandFind(uni, q)
	if rootP != rootQ {
		uni[rootQ] = rootP
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
