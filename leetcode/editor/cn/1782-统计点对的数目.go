//给你一个无向图，无向图由整数 n ，表示图中节点的数目，和 edges 组成，其中 edges[i] = [ui, vi] 表示 ui 和 vi 之间有一条
//无向边。同时给你一个代表查询的整数数组 queries 。
//
// 第 j 个查询的答案是满足如下条件的点对 (a, b) 的数目：
//
//
// a < b
// cnt 是与 a 或者 b 相连的边的数目，且 cnt 严格大于 queries[j] 。
//
//
// 请你返回一个数组 answers ，其中 answers.length == queries.length 且 answers[j] 是第 j 个查询的答
//案。
//
// 请注意，图中可能会有 重复边 。
//
//
//
// 示例 1：
//
//
//输入：n = 4, edges = [[1,2],[2,4],[1,3],[2,3],[2,1]], queries = [2,3]
//输出：[6,5]
//解释：每个点对中，与至少一个点相连的边的数目如上图所示。
//
//
// 示例 2：
//
//
//输入：n = 5, edges = [[1,5],[1,5],[3,4],[2,5],[1,3],[5,1],[2,3],[2,5]], queries =
// [1,2,3,4,5]
//输出：[10,10,9,8,6]
//
//
//
//
// 提示：
//
//
// 2 <= n <= 2 * 10⁴
// 1 <= edges.length <= 10⁵
// 1 <= ui, vi <= n
// ui != vi
// 1 <= queries.length <= 20
// 0 <= queries[j] < edges.length
//
//
// Related Topics 图 数组 双指针 二分查找 排序 👍 43 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	//arr := []int{3, 4, 4, 4, 5, 5}
	//tar := 3
	//i := sort.Search(len(arr), func(i int) bool {
	//	return arr[i] > tar
	//})
	//fmt.Println(len(arr) - i)
	n := 5
	edges := [][]int{{1, 5}, {1, 5}, {3, 4}, {2, 5}, {1, 3}, {5, 1}, {2, 3}, {2, 5}}
	queries := []int{1, 2, 3, 4, 5} // [10,10,9,8,6]
	ints := countPairs(n, edges, queries)
	fmt.Println(ints)
}

/*
思路：双指针
	1.读懂题意后，本题思路并不难，难在时空的优化上
		例如使用 hash 代替邻接表
	2.代码中的容器释义
		cntEdge：记录与顶点相连的边的总数
		es：记录两个顶点之间边的总数，key 是两个顶点 {编号小的顶点, 编号大的顶点}
		cntSort：cntEdge 排序后的数组
	3.核心算法：计算 queries[i] 的结果值
		3.1.统计大于 queries[i] 的点对的数量，即双指针搜索 cntSort
			示例：cntSort = [1,2,3,4]，queries[i] = 3
			最终搜索结果为 [1,4] [2,4] [3,4] [1,3] [2,3]，一共5个点对
		3.2.对点对进行修正：3.1.的计算中，对点对之间的连线计算了两次，所以要修正
			修正思路为：减去 cntEdge 中标记的两个之间的边数，仍然大于 queries[i]
*/
// leetcode submit region begin(Prohibit modification and deletion)
func countPairs(n int, edges [][]int, queries []int) []int {
	// 优化：终版
	// https://leetcode.cn/problems/count-pairs-of-nodes/solutions/2400682/ji-bai-100cong-shuang-zhi-zhen-dao-zhong-yhze/

	// 双指针
	minMax := func(a, b int) (int, int) {
		if a <= b {
			return a, b
		}
		return b, a
	}
	type e struct {
		x, y int
	}
	cntEdge, es := make([]int, n+1), make(map[e]int)
	for _, edge := range edges {
		i, j := minMax(edge[0], edge[1])
		es[e{i, j}]++ // 点对的边数统计
		cntEdge[i]++  // 点的边数统计
		cntEdge[j]++
	}
	ret, cntSort := make([]int, len(queries)), make([]int, n+1)
	copy(cntSort, cntEdge)
	sort.Ints(cntSort) // 点的边数排序
	for i, t := range queries {
		for l, r := 1, n; l < r; {
			if cntSort[l]+cntSort[r] > t {
				ret[i] += r - l // 统计：总边数大于 queries[i] 点对数量
				r--
			} else {
				l++
			}
		}
		for k, v := range es {
			// 修正：总边数大于 queries[i] 点对数量
			if sum := cntEdge[k.x] + cntEdge[k.y]; sum > t && sum-v <= t {
				ret[i]--
			}
		}
	}
	return ret

	// Time Limit Exceeded
	//maxMin := func(a, b int) (int, int) {
	//	if a > b {
	//		return a, b
	//	}
	//	return b, a
	//}
	//cnt, cp := make([]int, n+1), make(map[int]int)
	//for _, e := range edges {
	//	j, i := maxMin(e[0], e[1])
	//	cp[i*n+j]++
	//	cnt[i]++
	//	cnt[j]++
	//}
	//m := (n - 1) * n >> 1
	//ps := make([]int, m)
	//for i, k := 1, 0; i < n; i++ {
	//	for j := i + 1; j <= n; j++ {
	//		ps[k] = cnt[i] + cnt[j] - cp[i*n+j]
	//		k++
	//	}
	//}
	//sort.Ints(ps)
	//ret := make([]int, len(queries))
	//for i, query := range queries {
	//	ret[i] = m - sort.Search(m, func(i int) bool {
	//		return ps[i] > query
	//	})
	//}
	//return ret

	// ugly 统计方式
	//cnt, cp := make([]int, n+1), make([][]int, n+1)
	//for i := 1; i <= n; i++ {
	//	cp[i] = make([]int, n+1)
	//}
	//for _, e := range edges {
	//
	//	cp[e[0]][e[1]]++
	//	cnt[e[0]]++
	//	cnt[e[1]]++
	//}
	//m := (n - 1) * n >> 1
	//ps := make([]int, m)
	//for i, k := 1, 0; i < n; i++ {
	//	for j := i + 1; j <= n; j++ {
	//		//ps[(n-2)*(i-1)+j-2] = cnt[i] + cnt[j] - cp[i][j] - cp[j][i]
	//		ps[k] = cnt[i] + cnt[j] - cp[i][j] - cp[j][i]
	//		k++
	//	}
	//}
	//sort.Ints(ps)
	//ret := make([]int, len(queries))
	//for i, query := range queries {
	//	ret[i] = m - sort.Search(m, func(i int) bool {
	//		return ps[i] > query
	//	})
	//}
	////fmt.Println(ps)
	////fmt.Println(ps)
	////fmt.Println(cnt)
	////for _, c := range cp {
	////	fmt.Println(c)
	////}
	//return ret
}

//leetcode submit region end(Prohibit modification and deletion)
