//给你两组点，其中第一组中有 size1 个点，第二组中有 size2 个点，且 size1 >= size2 。
//
// 任意两点间的连接成本 cost 由大小为 size1 x size2 矩阵给出，其中 cost[i][j] 是第一组中的点 i 和第二组中的点 j 的连接
//成本。如果两个组中的每个点都与另一组中的一个或多个点连接，则称这两组点是连通的。换言之，第一组中的每个点必须至少与第二组中的一个点连接，且第二组中的每个点必须至
//少与第一组中的一个点连接。
//
// 返回连通两组点所需的最小成本。
//
//
//
// 示例 1：
//
//
//
// 输入：cost = [[15, 96], [36, 2]]
//输出：17
//解释：连通两组点的最佳方法是：
//1--A
//2--B
//总成本为 17 。
//
//
// 示例 2：
//
//
//
// 输入：cost = [[1, 3, 5], [4, 1, 1], [1, 5, 3]]
//输出：4
//解释：连通两组点的最佳方法是：
//1--A
//2--B
//2--C
//3--A
//最小成本为 4 。
//请注意，虽然有多个点连接到第一组中的点 2 和第二组中的点 A ，但由于题目并不限制连接点的数目，所以只需要关心最低总成本。
//
// 示例 3：
//
// 输入：cost = [[2, 5, 1], [3, 4, 7], [8, 1, 2], [6, 2, 4], [3, 8, 8]]
//输出：10
//
//
//
//
// 提示：
//
//
// size1 == cost.length
// size2 == cost[i].length
// 1 <= size1, size2 <= 12
// size1 >= size2
// 0 <= cost[i][j] <= 100
//
//
// Related Topics 位运算 数组 动态规划 状态压缩 矩阵 👍 68 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	cost := [][]int{{15, 96}, {36, 2}}
	cost = [][]int{{1, 3, 5}, {4, 1, 1}, {1, 5, 3}}
	//cost = [][]int{{93, 56, 92}, {53, 44, 18}, {86, 44, 69}, {54, 60, 30}} // 172
	groups := connectTwoGroups(cost)
	fmt.Println(groups)
}

/*
思路：状压DP
	1.有可能出现第二组点集没有被连接的情况，所以不能贪心的选择第一组点集的最优解
		而是要记录第一组点集的每个点，连接第二组点集的所有点的情况
	2.第一组点集的当前点 i，连接第二组点集中的每个可以被连接点 j
		2.1.状态压缩第二组点集的所有子集，二进制第k位为1，表示这个点可以被连接
		2.2.f(i,j) 的最优解，从三种情况中选出（这3种情况都要加上 i 连接 k 的值）：
			f(i,j) = f(i-1,j)：i-1 个点集，对应子集 j
			f(i,j) = f(i-1,j^k)：i-1 个点集，对应子集 j 除去 k
			f(i,j) = f(i,j^k)：i 个点集，对应子集 j 除去 k
			f(i,j) 都需要加上 cost[i][k]
	3.数据初始化
		f(0,0) = 0
		f(i,j) = MaxInt32

状压dp相似题目
	1879. 两个数组最小的异或值之和
	2172. 数组的最大与和
	1125. 最小的必要团队，题解
	1986. 完成任务的最少工作时间段
*/
// leetcode submit region begin(Prohibit modification and deletion)
func connectTwoGroups(cost [][]int) int {
	// 状压dp
	minVal := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n, m, bM := len(cost), len(cost[0]), 1<<len(cost[0])
	dp0, dpi := make([]int, bM), make([]int, bM)
	for i := 1; i < bM; i++ { // dp[0][0]=0
		dp0[i] = math.MaxInt32
	}
	for i := 0; i < n; i++ {
		dpi[0] = math.MaxInt32
		for j := 1; j < bM; j++ {
			dpi[j] = math.MaxInt32   // 初始化 dp[i][j]
			for k := 0; k < m; k++ { // 穷举第 i 个点连接每一个可以连接的点
				if b := 1 << k; j&b > 0 {
					dpi[j] = minVal(dpi[j],
						minVal(dpi[j^b], minVal(dp0[j], dp0[j^b]))+cost[i][k])
				} // 选出最优解
			}
		}
		dp0, dpi = dpi, dp0
	}
	return dp0[bM-1]
}

// 记忆化搜索

// 记忆化：可能多出一条连接，而导致值增大
//minVal := func(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//n, m := len(cost), len(cost[0])
//nl, ml := 1<<n, 1<<m
//memo := make([][]int, nl)
//for i := range memo {
//	memo[i] = make([]int, ml)
//	for j := range memo[i] {
//		memo[i][j] = -1
//	}
//}
//getMin := func(i int) (idx, v int) {
//	arr := cost[i]
//	for j := 1; j < m; j++ {
//		if arr[j] < arr[idx] {
//			idx = j
//		}
//	}
//	return idx, arr[idx]
//}
//var dfs func(int, int) int
//dfs = func(i, j int) int {
//	if i == 0 {
//		v := 0
//		for k := 0; k < m; k++ {
//			if b := 1 << k; b&j > 0 {
//				vk := math.MaxInt32
//				for x := 0; x < n; x++ {
//					vk = minVal(vk, cost[x][k])
//				}
//				v += vk
//			}
//		}
//		return v
//	}
//	if p := memo[i][j]; p != -1 {
//		return p
//	}
//	sum := math.MaxInt32
//	for k := 0; k < n; k++ {
//		if b := 1 << k; b&i > 0 {
//			idx, v := getMin(k)
//			sum = minVal(sum, v+dfs(i&^b, j&^(1<<idx)))
//		}
//	}
//	memo[i][j] = sum
//	return sum
//}
//return dfs(nl-1, ml-1)
// 回溯
}

//leetcode submit region end(Prohibit modification and deletion)
