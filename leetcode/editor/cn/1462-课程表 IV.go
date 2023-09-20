//你总共需要上
// numCourses 门课，课程编号依次为 0 到 numCourses-1 。你会得到一个数组 prerequisite ，其中
// prerequisites[i] = [ai, bi] 表示如果你想选
// bi 课程，你 必须 先选
// ai 课程。
//
//
// 有的课会有直接的先修课程，比如如果想上课程 1 ，你必须先上课程 0 ，那么会以 [0,1] 数对的形式给出先修课程数对。
//
//
// 先决条件也可以是 间接 的。如果课程 a 是课程 b 的先决条件，课程 b 是课程 c 的先决条件，那么课程 a 就是课程 c 的先决条件。
//
// 你也得到一个数组
// queries ，其中
// queries[j] = [uj, vj]。对于第 j 个查询，您应该回答课程
// uj 是否是课程
// vj 的先决条件。
//
// 返回一个布尔数组 answer ，其中 answer[j] 是第 j 个查询的答案。
//
//
//
// 示例 1：
//
//
//
//
//输入：numCourses = 2, prerequisites = [[1,0]], queries = [[0,1],[1,0]]
//输出：[false,true]
//解释：课程 0 不是课程 1 的先修课程，但课程 1 是课程 0 的先修课程。
//
//
// 示例 2：
//
//
//输入：numCourses = 2, prerequisites = [], queries = [[1,0],[0,1]]
//输出：[false,false]
//解释：没有先修课程对，所以每门课程之间是独立的。
//
//
// 示例 3：
//
//
//
//
//输入：numCourses = 3, prerequisites = [[1,2],[1,0],[2,0]], queries = [[1,0],[1,2]
//]
//输出：[true,true]
//
//
//
//
// 提示：
//
//
//
//
//
// 2 <= numCourses <= 100
// 0 <= prerequisites.length <= (numCourses * (numCourses - 1) / 2)
// prerequisites[i].length == 2
// 0 <= ai, bi <= n - 1
// ai != bi
// 每一对
// [ai, bi] 都 不同
// 先修课程图中没有环。
// 1 <= queries.length <= 10⁴
// 0 <= ui, vi <= n - 1
// ui != vi
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 109 👎 0

package main

import "fmt"

func main() {
	prerequisites := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
	queries := [][]int{{0, 4}, {4, 0}, {1, 3}, {3, 0}}
	numCourses := 5
	prerequisite := checkIfPrerequisite(numCourses, prerequisites, queries)
	fmt.Println(prerequisite)
}

// leetcode submit region begin(Prohibit modification and deletion)
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// 拓扑排序
	var (
		adj   = make([][]int, numCourses)  // 1.邻接表
		kahn  = make([]int, numCourses)    // 2.统计入度
		queue = make([]int, 0, numCourses) // 3.待遍历顶点的队列
		memo  = make([][]bool, numCourses) // 4.课程先修关系的映射表
		cip   = make([]bool, len(queries)) // 5.查询结果集
	)
	for _, p := range prerequisites { // 初始化 1&2
		adj[p[0]] = append(adj[p[0]], p[1])
		kahn[p[1]]++
	}
	for i := 0; i < numCourses; i++ { // 初始化 4
		memo[i] = make([]bool, numCourses)
	}
	for i, v := range kahn { // 初始化 3
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 { // 映射表 4
		f := queue[0]
		queue = queue[1:]
		for _, t := range adj[f] {
			memo[f][t] = true                 // 已存在 f->t
			for i := 0; i < numCourses; i++ { // 若存在 i->f，则存在 i->t
				memo[i][t] = memo[i][f] || memo[i][t]
			}
			kahn[t]--
			if kahn[t] == 0 {
				queue = append(queue, t)
			}
		}
	}
	for i, q := range queries { // 结果集 5
		cip[i] = memo[q[0]][q[1]] // check: from q[0] to q[1]
	}
	return cip

	// dfs
	//var (
	//	adj     = make([][]int, numCourses)  // 1.邻接表
	//	visited = make([]bool, numCourses)   // 2.标记已遍历的顶点
	//	memo    = make([][]bool, numCourses) // 4.课程先修关系的映射表
	//	cip     = make([]bool, len(queries)) // 5.查询结果集
	//	dfs     func(int)                    // 3.dfs 逻辑
	//)
	//for _, p := range prerequisites { // 初始化 1
	//	adj[p[1]] = append(adj[p[1]], p[0])
	//}
	//for i := 0; i < numCourses; i++ { // 初始化 4
	//	memo[i] = make([]bool, numCourses)
	//}
	//dfs = func(t int) { // 实现 3
	//	if visited[t] {
	//		return
	//	}
	//	visited[t] = true // 标记已遍历
	//	for _, f := range adj[t] {
	//		memo[f][t] = true                 // 已存在 f->t
	//		dfs(f)                            // f 作为 t
	//		for i := 0; i < numCourses; i++ { // 若存在 i->f，则存在 i->t
	//			memo[i][t] = memo[i][f] || memo[i][t]
	//		}
	//	}
	//}
	//for i := 0; i < numCourses; i++ { // 映射表 4
	//	dfs(i)
	//}
	//for i, q := range queries { // 结果集 5
	//	cip[i] = memo[q[0]][q[1]] // check: from q[0] to q[1]
	//}
	//return cip

	// 拓扑排序：dfs 超时
	//cip, visited, adj :=
	//	make([]bool, len(queries)), make([]bool, numCourses), make([][]int, numCourses)
	//for _, p := range prerequisites {
	//	adj[p[1]] = append(adj[p[1]], p[0])
	//}
	//var dfs func(int, int, []bool) bool
	//dfs = func(t, f int, visit []bool) bool {
	//	for _, next := range adj[f] {
	//		if next == t {
	//			return true
	//		}
	//		visit[next] = true
	//		if dfs(t, next, visit) {
	//			return true
	//		}
	//		visit[next] = false
	//	}
	//	return false
	//}
	//for i, q := range queries {
	//	visited[q[0]] = true
	//	cip[i] = dfs(q[0], q[1], visited)
	//	visited[q[0]] = false
	//}
	//return cip
}

//leetcode submit region end(Prohibit modification and deletion)
