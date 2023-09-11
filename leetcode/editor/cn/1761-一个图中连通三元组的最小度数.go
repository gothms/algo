//给你一个无向图，整数 n 表示图中节点的数目，edges 数组表示图中的边，其中 edges[i] = [ui, vi] ，表示 ui 和 vi 之间有一条
//无向边。
//
// 一个 连通三元组 指的是 三个 节点组成的集合且这三个点之间 两两 有边。
//
// 连通三元组的度数 是所有满足此条件的边的数目：一个顶点在这个三元组内，而另一个顶点不在这个三元组内。
//
// 请你返回所有连通三元组中度数的 最小值 ，如果图中没有连通三元组，那么返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：n = 6, edges = [[1,2],[1,3],[3,2],[4,1],[5,2],[3,6]]
//输出：3
//解释：只有一个三元组 [1,2,3] 。构成度数的边在上图中已被加粗。
//
//
// 示例 2：
//
//
//输入：n = 7, edges = [[1,3],[4,1],[4,3],[2,5],[5,6],[6,7],[7,5],[2,6]]
//输出：0
//解释：有 3 个三元组：
//1) [1,4,3]，度数为 0 。
//2) [2,5,6]，度数为 2 。
//3) [5,6,7]，度数为 2 。
//
//
//
//
// 提示：
//
//
// 2 <= n <= 400
// edges[i].length == 2
// 1 <= edges.length <= n * (n-1) / 2
// 1 <= ui, vi <= n
// ui != vi
// 图中没有重复的边。
//
//
// Related Topics 图 👍 28 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 7
	edges := [][]int{{1, 3}, {4, 1}, {4, 3}, {2, 5}, {5, 6}, {6, 7}, {7, 5}, {2, 6}}
	degree := minTrioDegree(n, edges)
	fmt.Println(degree)
}

/*
思路：图
	1.假设顶点 i、j、k 是三元组，那么它们的联通三元组度数 degree 计算方法为：
		degree = degree[i]+degree[j]+degree[k]-6
		其中 degree[i] 是顶点 i 的度数，6代表三元组的必要条件：3条边（6个顶点）
	2.通过遍历任意顶点组合 i,j,k，计算它们的联通三元组度数，并从中求出最小值
		数据结构定义如下
		degree 数组：记录每个顶点的度
		adjMap map数组：代表“邻接矩阵”
		adjArr 二维数组：代表邻接表
*/
// leetcode submit region begin(Prohibit modification and deletion)
func minTrioDegree(n int, edges [][]int) int {
	minVal := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	minMax := func(a, b int) (int, int) {
		if a < b {
			return a, b
		}
		return b, a
	}
	degree, adjMap, adjArr :=
		make([]int, n+1), make([]map[int]struct{}, n+1), make([][]int, n+1)
	for i := 1; i <= n; i++ {
		adjMap[i] = make(map[int]struct{})
	}
	for _, e := range edges {
		degree[e[0]]++ // 顶点的度
		degree[e[1]]++
		min, max := minMax(e[0], e[1])
		adjMap[min][max] = struct{}{}          // “邻接矩阵”
		adjArr[min] = append(adjArr[min], max) // 邻接表
	}
	cnt := math.MaxInt32
	for i := 1; i <= n; i++ {
		for _, j := range adjArr[i] { // 边 [i,j]
			for _, k := range adjArr[j] { // 查询三元组 [i,j,k]
				if _, ok := adjMap[i][k]; ok { // 计算三元组的度
					cnt = minVal(cnt, degree[i]+degree[j]+degree[k]-6)
				}
			}
		}
	}
	if cnt < math.MaxInt32 {
		return cnt
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
