//给你两个整数 m 和 n ，表示一个下标从 0 开始的 m x n 的网格图。
//
// 给你一个下标从 0 开始的二维整数矩阵 coordinates ，其中 coordinates[i] = [x, y] 表示坐标为 [x, y] 的格子是
// 黑色的 ，所有没出现在 coordinates 中的格子都是 白色的。
//
// 一个块定义为网格图中 2 x 2 的一个子矩阵。更正式的，对于左上角格子为 [x, y] 的块，其中 0 <= x < m - 1 且 0 <= y <
//n - 1 ，包含坐标为 [x, y] ，[x + 1, y] ，[x, y + 1] 和 [x + 1, y + 1] 的格子。
//
// 请你返回一个下标从 0 开始长度为 5 的整数数组 arr ，arr[i] 表示恰好包含 i 个 黑色 格子的块的数目。
//
//
//
// 示例 1：
//
//
//输入：m = 3, n = 3, coordinates = [[0,0]]
//输出：[3,1,0,0,0]
//解释：网格图如下：
//
//只有 1 个块有一个黑色格子，这个块是左上角为 [0,0] 的块。
//其他 3 个左上角分别为 [0,1] ，[1,0] 和 [1,1] 的块都有 0 个黑格子。
//所以我们返回 [3,1,0,0,0] 。
//
//
// 示例 2：
//
//
//输入：m = 3, n = 3, coordinates = [[0,0],[1,1],[0,2]]
//输出：[0,2,2,0,0]
//解释：网格图如下：
//
//有 2 个块有 2 个黑色格子（左上角格子分别为 [0,0] 和 [0,1]）。
//左上角为 [1,0] 和 [1,1] 的两个块，都有 1 个黑格子。
//所以我们返回 [0,2,2,0,0] 。
//
//
//
//
// 提示：
//
//
// 2 <= m <= 10⁵
// 2 <= n <= 10⁵
// 0 <= coordinates.length <= 10⁴
// coordinates[i].length == 2
// 0 <= coordinates[i][0] < m
// 0 <= coordinates[i][1] < n
// coordinates 中的坐标对两两互不相同。
//
//
// Related Topics 数组 哈希表 枚举 👍 10 👎 0

package main

import "fmt"

func main() {
	m, n := 3, 3
	coordinates := [][]int{{0, 0}, {1, 1}, {0, 2}}
	blocks := countBlackBlocks(m, n, coordinates)
	fmt.Println(blocks)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countBlackBlocks(m int, n int, coordinates [][]int) []int64 {
	// lc
	type pair struct{ x, y int }
	set := make(map[pair]int, len(coordinates))
	for _, p := range coordinates {
		set[pair{p[0], p[1]}] = 1
	}

	arr := make([]int64, 5)
	vis := make(map[pair]bool, len(set)*4)
	for _, p := range coordinates {
		x, y := p[0], p[1]
		for i := max(x-1, 0); i <= x && i < m-1; i++ {
			for j := max(y-1, 0); j <= y && j < n-1; j++ {
				if !vis[pair{i, j}] {
					vis[pair{i, j}] = true
					cnt := set[pair{i, j}] + set[pair{i, j + 1}] +
						set[pair{i + 1, j}] + set[pair{i + 1, j + 1}]
					arr[cnt]++
				}
			}
		}
	}
	arr[0] = int64(m-1)*int64(n-1) - int64(len(vis))
	return arr

	// lc
	// n = len(coordinates)，memo 长度为 n，visited 长度为 4n
	// memo 把 coordinates 的点对 copy 一份，visited 记录某个位置是否被访问
	// 遍历 coordinates 获取点对，再对点对所在 2*2 格子遍历，数 memo 中有几个黑点在格子中
	// 个人：hash 表
	//ret, memo := make([]int64, 5), make(map[[2]int]int64) // hash表，防 out of memory
	//ret[0] = int64((m - 1) * (n - 1))
	//dx, dy := [4]int{0, -1, -1, 0}, [4]int{-1, -1, 0, 0}
	//for _, c := range coordinates {
	//	x, y := c[0], c[1]
	//	for i := 0; i < 4; i++ {
	//		if nx, ny := x+dx[i], y+dy[i]; nx >= 0 && nx < m-1 && ny >= 0 && ny < n-1 {
	//			key := [2]int{nx, ny}
	//			ret[memo[key]]-- // 只操作有效的变更，防超时
	//			memo[key]++
	//			ret[memo[key]]++
	//		}
	//	}
	//}
	//return ret

	// out of memory
	//ret := make([]int64, 5)
	//memo := make([][]int, m+1)
	//for i := range memo {
	//	memo[i] = make([]int, n+1)
	//}
	//ret[0] = int64((m - 1) * (n - 1))
	//dx, dy := [4]int{0, -1, -1, 0}, [4]int{-1, -1, 0, 0}
	//for _, c := range coordinates {
	//	x, y := c[0], c[1]
	//	for i := 0; i < 4; i++ {
	//		if nx, ny := x+dx[i], y+dy[i]; nx >= 0 && nx < m-1 && ny >= 0 && ny < n-1 {
	//			ret[memo[nx][ny]]--
	//			memo[nx][ny]++
	//			ret[memo[nx][ny]]++
	//		}
	//	}
	//}
	//return ret

	// 超时
	//ret := make([]int64, 5)
	//memo := make([][]int, m+1)
	//for i := range memo {
	//	memo[i] = make([]int, n+1)
	//}
	//for _, c := range coordinates {
	//	i, j := c[0], c[1]
	//	memo[i][j]++
	//	memo[i][j+1]++
	//	memo[i+1][j]++
	//	memo[i+1][j+1]++
	//}
	//for i := 1; i < m; i++ {
	//	for j := 1; j < n; j++ {
	//		ret[memo[i][j]]++
	//	}
	//}
	//return ret
}

//leetcode submit region end(Prohibit modification and deletion)
