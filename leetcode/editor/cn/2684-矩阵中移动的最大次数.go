//给你一个下标从 0 开始、大小为 m x n 的矩阵 grid ，矩阵由若干 正 整数组成。
//
// 你可以从矩阵第一列中的 任一 单元格出发，按以下方式遍历 grid ：
//
//
// 从单元格 (row, col) 可以移动到 (row - 1, col + 1)、(row, col + 1) 和 (row + 1, col + 1)
//三个单元格中任一满足值 严格 大于当前单元格的单元格。
//
//
// 返回你在矩阵中能够 移动 的 最大 次数。
//
//
//
// 示例 1：
// 输入：grid = [[2,4,3,5],[5,4,9,3],[3,4,2,11],[10,9,13,15]]
//输出：3
//解释：可以从单元格 (0, 0) 开始并且按下面的路径移动：
//- (0, 0) -> (0, 1).
//- (0, 1) -> (1, 2).
//- (1, 2) -> (2, 3).
//可以证明这是能够移动的最大次数。
//
// 示例 2：
//
//
//输入：grid = [[3,2,4],[2,1,9],[1,1,7]]
//输出：0
//解释：从第一列的任一单元格开始都无法移动。
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 2 <= m, n <= 1000
// 4 <= m * n <= 10⁵
// 1 <= grid[i][j] <= 10⁶
//
//
// Related Topics 数组 动态规划 矩阵 👍 10 👎 0

package main

import "fmt"

func main() {
	grid := [][]int{
		{187, 167, 209, 251, 152, 236, 263, 128, 135},
		{267, 249, 251, 285, 73, 204, 70, 207, 74},
		{189, 159, 235, 66, 84, 89, 153, 111, 189},
		{120, 81, 210, 7, 2, 231, 92, 128, 218},
		{193, 131, 244, 293, 284, 175, 226, 205, 245}}
	//grid = [][]int{
	//	{1000000, 92910, 1021, 1022, 1023, 1024, 1025, 1026, 1027, 1028, 1029, 1030, 1031, 1032, 1033, 1034, 1035, 1036, 1037, 1038, 1039, 1040, 1041, 1042, 1043, 1044, 1045, 1046, 1047, 1048, 1049, 1050, 1051, 1052, 1053, 1054, 1055, 1056, 1057, 1058, 1059, 1060, 1061, 1062, 1063, 1064, 1065, 1066, 1067, 1068},
	//	{1069, 1070, 1071, 1072, 1073, 1074, 1075, 1076, 1077, 1078, 1079, 1080, 1081, 1082, 1083, 1084, 1085, 1086, 1087, 1088, 1089, 1090, 1091, 1092, 1093, 1094, 1095, 1096, 1097, 1098, 1099, 1100, 1101, 1102, 1103, 1104, 1105, 1106, 1107, 1108, 1109, 1110, 1111, 1112, 1113, 1114, 1115, 1116, 1117, 1118}}
	//grid = [][]int{
	//	{131, 1, 95, 208, 38, 257, 123, 204, 101},
	//	{185, 165, 292, 109, 266, 259, 97, 234, 60},
	//	{55, 281, 38, 61, 204, 243, 32, 54, 164},
	//	{106, 230, 202, 4, 65, 243, 89, 139, 211},
	//	{192, 246, 11, 294, 119, 43, 250, 161, 110},
	//	{71, 279, 288, 173, 64, 48, 216, 26, 276},
	//	{23, 206, 152, 106, 288, 286, 270, 131, 12},
	//	{115, 64, 251, 108, 194, 295, 131, 249, 121}}
	grid = [][]int{
		{102, 168, 168, 86, 228},
		{209, 210, 182, 153, 55},
		{99, 76, 168, 40, 262},
		{260, 257, 227, 97, 153},
		{189, 280, 257, 239, 93},
		{300, 108, 68, 220, 76}}
	moves := maxMoves(grid)
	fmt.Println(moves)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxMoves(grid [][]int) int {
	// dp
	m, n := len(grid), len(grid[0])
	cache := make([]bool, m)
	for i := range cache {
		cache[i] = true
	}
	for c := 1; c < n; c++ {
		next, temp := false, make([]bool, m)
		for r, b := range cache {
			if b {
				for k := max(0, r-1); k < min(m, r+2); k++ {
					if grid[k][c] > grid[r][c-1] { // 右上、同行、右下 > grid[r][c-1]
						next, temp[k] = true, true
					}
				}
			}
		}
		if !next {
			return c - 1
		}
		cache = temp
	}
	return n - 1

	// dp
	//maxVal := func(a, b int) int {
	//	if b > a {
	//		return b
	//	}
	//	return a
	//}
	//minVal := func(a, b int) int {
	//	if b < a {
	//		return b
	//	}
	//	return a
	//}
	//m, n := len(grid), len(grid[0])
	//col := make([]bool, m)
	//for i := 0; i < m; i++ { // 第一列
	//	col[i] = true
	//}
	//for c := 1; c < n; c++ {
	//	cnt, temp := 0, make([]bool, m)
	//	for r := 0; r < m; r++ {
	//		if !col[r] { // 上一个位置不可以移动
	//			continue
	//		}
	//		// 下限和上限
	//		for i, t := maxVal(r-1, 0), grid[r][c-1]; i < minVal(r+2, m); i++ {
	//			if grid[i][c] > t {
	//				temp[i] = true // 标记可以移动的位置
	//				cnt++
	//			}
	//		}
	//	}
	//	if cnt == 0 {
	//		return c - 1
	//	}
	//	col = temp
	//}
	//return n - 1

	// lc
	//m, n := len(grid), len(grid[0])
	//memo := make([][]int, m)
	//for i := range memo {
	//	memo[i] = make([]int, n)
	//	for j := range memo[i] {
	//		memo[i][j] = -1 // -1 表示没被计算过
	//	}
	//}
	//var dfs func(int, int) int
	//dfs = func(i, j int) (res int) {
	//	if j == n-1 {
	//		return
	//	}
	//	p := &memo[i][j]
	//	if *p != -1 {
	//		return *p
	//	}
	//	for k := maxVal(i-1, 0); k < minVal(i+2, m); k++ {
	//		if grid[k][j+1] > grid[i][j] {
	//			res = maxVal(res, dfs(k, j+1)+1)
	//		}
	//	}
	//	*p = res // 记忆化
	//	return
	//}
	//ans := 0
	//for i := 0; i < m; i++ {
	//	ans = maxVal(ans, dfs(i, 0))
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
