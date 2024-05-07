//给你一个 n 个点组成的无向图边集 edgeList ，其中 edgeList[i] = [ui, vi, disi] 表示点 ui 和点 vi 之间有一条
//长度为 disi 的边。请注意，两个点之间可能有 超过一条边 。
//
// 给你一个查询数组queries ，其中 queries[j] = [pj, qj, limitj] ，你的任务是对于每个查询 queries[j] ，判断
//是否存在从 pj 到 qj 的路径，且这条路径上的每一条边都 严格小于 limitj 。
//
// 请你返回一个 布尔数组 answer ，其中 answer.length == queries.length ，当 queries[j] 的查询结果为
//true 时， answer 第 j 个值为 true ，否则为 false 。
//
//
//
// 示例 1：
//
//
//输入：n = 3, edgeList = [[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries = [[0,1,2],[0
//,2,5]]
//输出：[false,true]
//解释：上图为给定的输入数据。注意到 0 和 1 之间有两条重边，分别为 2 和 16 。
//对于第一个查询，0 和 1 之间没有小于 2 的边，所以我们返回 false 。
//对于第二个查询，有一条路径（0 -> 1 -> 2）两条边都小于 5 ，所以这个查询我们返回 true 。
//
//
// 示例 2：
//
//
//输入：n = 5, edgeList = [[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries = [[0,4,14],
//[1,4,13]]
//输出：[true,false]
//解释：上图为给定数据。
//
//
//
//
// 提示：
//
//
// 2 <= n <= 10⁵
// 1 <= edgeList.length, queries.length <= 10⁵
// edgeList[i].length == 3
// queries[j].length == 3
// 0 <= ui, vi, pj, qj <= n - 1
// ui != vi
// pj != qj
// 1 <= disi, limitj <= 10⁹
// 两个点之间可能有 多条 边。
//
//
// Related Topics 并查集 图 数组 双指针 排序 👍 182 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	edgeList := [][]int{{0, 1, 10},
		{1, 2, 5},
		{2, 3, 9},
		{3, 4, 13}}
	queries := [][]int{{0, 4, 14}, {1, 4, 13}}
	n := 5
	exist := distanceLimitedPathsExist(n, edgeList, queries)
	fmt.Println(exist)
}

// leetcode submit region begin(Prohibit modification and deletion)
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	// 优化：不重新申请空间
	// 执行耗时:247 ms,击败了100.00% 的Go用户
	// 内存消耗:18.3 MB,击败了25.00% 的Go用户
	idx := make([]int, len(queries)) // 优化
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool { // queries 排序
		return queries[idx[i]][2] < queries[idx[j]][2]
	})
	sort.Slice(edgeList, func(i, j int) bool { // edgeList 排序
		return edgeList[i][2] < edgeList[j][2]
	})
	uni := make([]int, n) // 并查集
	for i := range uni {
		uni[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if uni[x] != x {
			uni[x] = find(uni[x])
		}
		return uni[x]
	}
	merge := func(p, q int) {
		uni[find(q)] = find(p) // uni[find(id)] 将整个并查集合并到 p
	}
	ans := make([]bool, len(queries))
	k := 0
	for _, i := range idx { // 核心逻辑
		for ; k < len(edgeList) && edgeList[k][2] < queries[i][2]; k++ { // 易错点：k < len(edgeList)
			merge(edgeList[k][0], edgeList[k][1])
		}
		ans[i] = find(queries[i][0]) == find(queries[i][1]) // 在 id[2] 的限制下，id[0] 和 id[1] 的路径是否已通
	}
	return ans

	// 并查集 + 排序
	// 思路：请见代码
	// 执行耗时:279 ms,击败了25.00% 的Go用户
	// 内存消耗:16.6 MB,击败了100.00% 的Go用户
	//for i := range queries {
	//	queries[i] = append(queries[i], i) // i 用于操作 ans。但会重新申请空间，所以耗时
	//}
	//sort.Slice(queries, func(i, j int) bool { // queries 排序
	//	return queries[i][2] < queries[j][2]
	//})
	//sort.Slice(edgeList, func(i, j int) bool { // edgeList 排序
	//	return edgeList[i][2] < edgeList[j][2]
	//})
	//uni := make([]int, n) // 并查集
	//for i := range uni {
	//	uni[i] = i
	//}
	//var find func(int) int
	//find = func(x int) int {
	//	if uni[x] != x {
	//		uni[x] = find(uni[x])
	//	}
	//	return uni[x]
	//}
	//merge := func(p, id int) {
	//	uni[find(id)] = find(p) // uni[find(id)] 将整个并查集合并到 p
	//}
	//ans := make([]bool, len(queries))
	//k := 0
	//for _, id := range queries { // 核心逻辑
	//	for ; k < len(edgeList) && edgeList[k][2] < id[2]; k++ { // 易错点：k < len(edgeList)
	//		merge(edgeList[k][0], edgeList[k][1])
	//	}
	//	ans[id[3]] = find(id[0]) == find(id[1]) // 在 id[2] 的限制下，id[0] 和 id[1] 的路径是否已通
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
