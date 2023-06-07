//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1：
//
//
//输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//输出：1
//
//
// 示例 2：
//
//
//输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//输出：3
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 2190 👎 0

package main

func main() {

}

/*
每日一题：lc-2455
	今天日刷较简单，直接贴代码
		另：代码中，可以直接选择 %6
	补一经典题型：lc-200 岛屿数量
思路：并查集
	1.并查集模板
		1.1.parent[]：一般初始化 parent[i]=i，每个元素都是一个集合
		1.2.find(x)：找到元素 x 所在的集合的代表，也可用于判断两个元素是否位于同一集合
		1.3.union(x,y)：合并 x y 所在的集合，如果不相交则不合并
	2.遍历二维数组，找出岛屿数量
		2.1.假设每个 grid[i][j]=='1' 都是一个岛屿，也就是一个集合
			每合并一个 '1'，岛屿数量就 -1
		2.2.常用的方法：将二维数组的索引 i j 映射为一维数组的索引
			i*n+j：n=len(grid[0])
	3.其他：该问题使用 dfs，会很简洁
*/
//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == '1'
	}
	cnt, parent, d := 0, make([]int, m*n), [5]int{0, -1, 0, 1, 0} // 1
	for i := 1; i < len(parent); i++ {                            // 1.1
		parent[i] = i
	}
	for i := 0; i < m; i++ { // 2
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			cnt++
			for k := 0; k < 4; k++ {
				ni, nj := i+d[k], j+d[k+1]
				if ok(ni, nj) && union(parent, i*n+j, ni*n+nj) { // 2.2
					cnt-- // 2.1
				}
			}
		}
	}
	return cnt
}
func find(parent []int, p int) int { // 1.2
	for parent[p] != p {
		p, parent[p] = parent[p], parent[parent[p]]
	}
	return p
}
func union(parent []int, p, q int) bool { // 1.3
	rootP, rootQ := find(parent, p), find(parent, q)
	if rootP != rootQ {
		parent[rootQ] = rootP
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
