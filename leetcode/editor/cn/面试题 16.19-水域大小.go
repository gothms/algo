//你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角连接的水域为池塘。池塘的大小是指
//相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。
// 示例：
// 输入：
//[
//  [0,2,1,0],
//  [0,1,0,1],
//  [1,1,0,1],
//  [0,1,0,1]
//]
//输出： [1,2,4]
//
// 提示：
//
// 0 < len(land) <= 1000
// 0 < len(land[i]) <= 1000
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 93 👎 0

package main

import (
	"sort"
)

func main() {

}

/*
思路：dfs
	1.深度优先搜索时，每遇到一个 0，就八连通搜索周围
	2.每个搜索过的点，都标记为 1，表示已访问过
思路：并查集
	1.uni 记录并查集，陆地集合都置为 n*m
	2.size 记录每个池塘当前的大小，合并鱼塘时，遵循大鱼塘合小鱼塘
	3.uni 中每个集合的根节点，记录到最终池塘集合中，且大小从size中查出
*/
// leetcode submit region begin(Prohibit modification and deletion)
func pondSizes(land [][]int) []int {
	// dfs
	//n, m := len(land), len(land[0])
	//pools, d := make([]int, 0), [][2]int{
	//	{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	//var dfs func(int, int) int
	//dfs = func(i, j int) (cnt int) {
	//	if i < 0 || i == n || j < 0 || j == m || land[i][j] > 0 {
	//		return 0 // 越界/陆地
	//	}
	//	land[i][j] = 1 // 已访问
	//	cnt++
	//	for k := 0; k < 8; k++ { // 8连通
	//		cnt += dfs(i+d[k][0], j+d[k][1])
	//	}
	//	return
	//}
	//for i := 0; i < n; i++ {
	//	for j := 0; j < m; j++ {
	//		if land[i][j] == 0 { // 可以养鱼
	//			pools = append(pools, dfs(i, j))
	//		}
	//	}
	//}
	//sort.Ints(pools) // 排序
	//return pools
	// 并查集
	n, m, um := len(land), len(land[0]), len(land)*len(land[0])
	d := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	uni, size := make([]int, um), make([]int, um)
	for i := range uni {
		uni[i], size[i] = i, 1 // 初始化并查集，每个池塘大小
	}
	find := func(uni []int, p int) int {
		if p != uni[p] {
			p, uni[p] = uni[p], uni[uni[p]]
		}
		return p
	}
	union := func(uni []int, p, q int) {
		rootP, rootQ := find(uni, p), find(uni, q)
		if rootP == rootQ {
			return
		}
		if size[rootP] >= size[rootQ] { // 大鱼塘合小鱼塘，被合并的鱼塘大小置为 0
			uni[rootQ], size[rootP], size[rootQ] = rootP, size[rootP]+size[rootQ], 0
		} else {
			uni[rootP], size[rootQ], size[rootP] = rootQ, size[rootP]+size[rootQ], 0
		}
	}
	for i, p := 0, 0; i < n; i++ {
		for j := 0; j < m; j++ {
			p = i*m + j
			if land[i][j] > 0 { // 不可以养鱼
				uni[p] = um // 陆地的集合
				continue
			}
			for k := 0; k < 8; k++ {
				ni, nj := i+d[k][0], j+d[k][1]
				if ni < 0 || ni == n || nj < 0 || nj == m || land[ni][nj] > 0 {
					continue
				}
				union(uni, p, ni*m+nj) // 合并鱼塘
			}
		}
	}
	pools := make([]int, 0)
	for i := range uni {
		if i == uni[i] { // 是鱼塘
			pools = append(pools, size[i]) // 查出池塘的大小
		}
	}
	sort.Ints(pools) // 排序
	return pools
}

//leetcode submit region end(Prohibit modification and deletion)
