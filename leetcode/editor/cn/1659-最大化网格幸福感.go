//给你四个整数 m、n、introvertsCount 和 extrovertsCount 。有一个 m x n 网格，和两种类型的人：内向的人和外向的人。总
//共有 introvertsCount 个内向的人和 extrovertsCount 个外向的人。
//
// 请你决定网格中应当居住多少人，并为每个人分配一个网格单元。 注意，不必 让所有人都生活在网格中。
//
// 每个人的 幸福感 计算如下：
//
//
// 内向的人 开始 时有 120 个幸福感，但每存在一个邻居（内向的或外向的）他都会 失去 30 个幸福感。
// 外向的人 开始 时有 40 个幸福感，每存在一个邻居（内向的或外向的）他都会 得到 20 个幸福感。
//
//
// 邻居是指居住在一个人所在单元的上、下、左、右四个直接相邻的单元中的其他人。
//
// 网格幸福感 是每个人幸福感的 总和 。 返回 最大可能的网格幸福感 。
//
//
//
// 示例 1：
//
//
//输入：m = 2, n = 3, introvertsCount = 1, extrovertsCount = 2
//输出：240
//解释：假设网格坐标 (row, column) 从 1 开始编号。
//将内向的人放置在单元 (1,1) ，将外向的人放置在单元 (1,3) 和 (2,3) 。
//- 位于 (1,1) 的内向的人的幸福感：120（初始幸福感）- (0 * 30)（0 位邻居）= 120
//- 位于 (1,3) 的外向的人的幸福感：40（初始幸福感）+ (1 * 20)（1 位邻居）= 60
//- 位于 (2,3) 的外向的人的幸福感：40（初始幸福感）+ (1 * 20)（1 位邻居）= 60
//网格幸福感为：120 + 60 + 60 = 240
//上图展示该示例对应网格中每个人的幸福感。内向的人在浅绿色单元中，而外向的人在浅紫色单元中。
//
//
// 示例 2：
//
//
//输入：m = 3, n = 1, introvertsCount = 2, extrovertsCount = 1
//输出：260
//解释：将内向的人放置在单元 (1,1) 和 (3,1) ，将外向的人放置在单元 (2,1) 。
//- 位于 (1,1) 的内向的人的幸福感：120（初始幸福感）- (1 * 30)（1 位邻居）= 90
//- 位于 (2,1) 的外向的人的幸福感：40（初始幸福感）+ (2 * 20)（2 位邻居）= 80
//- 位于 (3,1) 的内向的人的幸福感：120（初始幸福感）- (1 * 30)（1 位邻居）= 90
//网格幸福感为 90 + 80 + 90 = 260
//
//
// 示例 3：
//
//
//输入：m = 2, n = 2, introvertsCount = 4, extrovertsCount = 0
//输出：240
//
//
//
//
// 提示：
//
//
// 1 <= m, n <= 5
// 0 <= introvertsCount, extrovertsCount <= min(m * n, 6)
//
//
// Related Topics 位运算 记忆化搜索 动态规划 状态压缩 👍 63 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%b\n", 243)
	fmt.Printf("%b\n", 729)
	m := 2
	n := 3
	introvertsCount := 1
	extrovertsCount := 2
	//m = 3
	//n = 1
	//introvertsCount = 2
	//extrovertsCount = 1
	happiness := getMaxGridHappiness(m, n, introvertsCount, extrovertsCount)
	fmt.Println(happiness)
}

/*
思路：状压dp
*/
// leetcode submit region begin(Prohibit modification and deletion)
func getMaxGridHappiness(m int, n int, introvertsCount int, extrovertsCount int) int {
	// 状压dp
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 1.状态定义
	T := int(math.Pow(3, float64(n)))                            // 状态压缩，3进制
	neighbors := [][]int{{0, 0, 0}, {0, -60, -10}, {0, -10, 40}} // 邻居幸福感的组合情况
	iCount, eCount, innerScore :=
		make([]int, T), make([]int, T), make([]int, T) // 内向/外向人数，行幸福感
	maskBits, interScore :=
		make([][]int, T), make([][]int, T) // 行的每个位置可能的居住情况，行与行幸福感
	dp := make([][][][]int, T) // T：每一行居住的人的状态
	for i := 0; i < T; i++ {
		maskBits[i] = make([]int, n)
		interScore[i] = make([]int, T)
		dp[i] = make([][][]int, m) // 有m行
		for j := 0; j < m; j++ {
			dp[i][j] = make([][]int, introvertsCount+1) // 内向的人数
			for k := 0; k <= introvertsCount; k++ {
				dp[i][j][k] = make([]int, extrovertsCount+1) // 外向的人数
				for l := 0; l <= extrovertsCount; l++ {
					dp[i][j][k][l] = -1 // 标识未访问过
				}
			}
		}
	}
	// 2.初始化数据
	// 每一行的情况都是相同的，所以只用记录一行的状态
	for mask := 0; mask < T; mask++ { // 单行可能的居住情况
		for temp, i := mask, 0; i < n; i++ { // 每一个可居住的位置
			b := temp % 3
			maskBits[mask][i] = b // 第i个位置居住的人
			temp /= 3             // 右移一位
			switch b {
			case 1: // 居住了内向的人
				iCount[mask]++
				innerScore[mask] += 120
			case 2: // 居住了外向的人
				eCount[mask]++
				innerScore[mask] += 40
			}
			if i > 0 {
				innerScore[mask] += neighbors[b][maskBits[mask][i-1]] // 和左边的邻居
			}
		}
	}
	for i := 0; i < T; i++ {
		for j := 0; j < T; j++ {
			interScore[i][j] = 0     // 每一种可能的居住情况，初始化值为0
			for k := 0; k < n; k++ { // 每一个可居住的位置
				interScore[i][j] += neighbors[maskBits[i][k]][maskBits[j][k]] // 和上面的邻居
			}
		}
	}
	// 3.dp
	var dfs func(int, int, int, int) int
	dfs = func(preMask int, row int, ic int, ec int) int {
		if row == m || (ic == 0 && ec == 0) {
			return 0 // 边界条件：完成居住/没人可居住
		}
		p := &dp[preMask][row][ic][ec]
		if *p != -1 {
			return *p // 记忆化
		}
		v := 0 // 也可以只用 *p
		for mask := 0; mask < T; mask++ {
			if iCount[mask] > ic || eCount[mask] > ec { // 没有可居住的人了
				continue
			}
			v = maxVal(v, innerScore[mask]+interScore[preMask][mask]+ // 截止当前行的幸福感
				dfs(mask, row+1, ic-iCount[mask], ec-eCount[mask])) // 下一行的幸福感
		}
		*p = v
		return v
	}
	return dfs(0, 0, introvertsCount, extrovertsCount)

	// lc
	//T, N, M, L := 243, 5, 6, int(math.Pow(3, float64(n)))
	//neighbors := [][]int{{0, 0, 0}, {0, -60, -10}, {0, -10, 40}}
	//iCount, eCount, innerScore := make([]int, T), make([]int, T), make([]int, T)
	//maskBits, interScore := make([][]int, T), make([][]int, T)
	//dp := make([][][][]int, T)
	//for i := 0; i < T; i++ {
	//	maskBits[i] = make([]int, N)
	//	interScore[i] = make([]int, T)
	//	dp[i] = make([][][]int, N)
	//	for j := 0; j < N; j++ {
	//		dp[i][j] = make([][]int, M+1)
	//		for k := 0; k <= introvertsCount; k++ {
	//			dp[i][j][k] = make([]int, M+1)
	//			for l := 0; l <= extrovertsCount; l++ {
	//				dp[i][j][k][l] = -1
	//			}
	//		}
	//	}
	//}
	//for mask := 0; mask < L; mask++ {
	//	for temp, i := mask, 0; i < n; i++ {
	//		b := temp % 3
	//		maskBits[mask][i] = b
	//		temp /= 3
	//		switch b {
	//		case 1:
	//			iCount[mask]++
	//			innerScore[mask] += 120
	//		case 2:
	//			eCount[mask]++
	//			innerScore[mask] += 40
	//		}
	//		if i > 0 {
	//			innerScore[mask] += neighbors[b][maskBits[mask][i-1]]
	//		}
	//	}
	//	for i := 0; i < T; i++ {
	//		for j := 0; j < T; j++ {
	//			interScore[i][j] = 0
	//			for k := 0; k < n; k++ {
	//				interScore[i][j] += neighbors[maskBits[i][k]][maskBits[j][k]]
	//			}
	//		}
	//	}
	//}
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//var dfs func(int, int, int, int) int
	//dfs = func(preMask int, row int, ic int, ec int) int {
	//	if row == m || (ic == 0 && ec == 0) {
	//		return 0
	//	}
	//	p := &dp[preMask][row][ic][ec]
	//	if *p != -1 {
	//		return *p
	//	}
	//	v := 0
	//	for mask := 0; mask < L; mask++ {
	//		if iCount[mask] > ic || eCount[mask] > ec {
	//			continue
	//		}
	//		v = maxVal(v, innerScore[mask]+interScore[preMask][mask]+
	//			dfs(mask, row+1, ic-iCount[mask], ec-eCount[mask]))
	//	}
	//	*p = v
	//	return *p
	//}
	//return dfs(0, 0, introvertsCount, extrovertsCount)
}

//leetcode submit region end(Prohibit modification and deletion)
