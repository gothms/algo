//你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表
//示如果要学习课程 ai 则 必须 先学习课程 bi 。
//
//
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
//
//
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：numCourses = 2, prerequisites = [[1,0]]
//输出：true
//解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
//
// 示例 2：
//
//
//输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
//输出：false
//解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
//
//
//
// 提示：
//
//
// 1 <= numCourses <= 2000
// 0 <= prerequisites.length <= 5000
// prerequisites[i].length == 2
// 0 <= ai, bi < numCourses
// prerequisites[i] 中的所有课程对 互不相同
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 👍 1688 👎 0

package main

import "fmt"

func main() {
	n := 4
	prerequisites := [][]int{{1, 0}, {0, 3}, {2, 3}}
	//n = 20
	//prerequisites = [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}
	finish := canFinish(n, prerequisites)
	fmt.Println(finish)
}

/*
思路：拓扑排序
	1.拓扑排序的两种实现：kahn 和 dfs，都有固定的模板代码
	2.该题是在拓扑排序的基础上检测是否有环
		kahn 的检测方法是：遍历后，入度为 0 的顶点数，是否等于总的顶点数
			相等则不存在环，反之
		dfs 的检测方法是：将顶点标记为三种状态，未遍历、遍历中、已遍历
			如果 dfs 时找到 “遍历中” 的顶点，则存在环，反之
*/

// leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 拓扑排序
	kahn, adj := make([]int, numCourses), make([][]int, numCourses)
	for _, e := range prerequisites {
		adj[e[1]] = append(adj[e[1]], e[0]) // 邻接表
		kahn[e[0]]++                        // 统计入度
	}
	cnt, queue := 0, make([]int, 0)
	for i, degree := range kahn {
		if degree == 0 {
			queue = append(queue, i) // 入度为 0 的顶点
		}
	}
	for len(queue) > 0 {
		f := queue[0] // 此轮要遍历的顶点
		queue = queue[1:]
		cnt++ // 顶点统计
		for _, t := range adj[f] {
			kahn[t]--         // 入度 -1
			if kahn[t] == 0 { // 入度为 0
				queue = append(queue, t)
			}
		}
	}
	return cnt == numCourses // 遍历的顶点数 < numCourses，则有环

	// dfs
	//adj, visited := make([][]int, numCourses), make([]int8, numCourses)
	//for _, p := range prerequisites {
	//	adj[p[0]] = append(adj[p[0]], p[1]) // 逆邻接表
	//}
	//var dfs func(int, []int8) bool
	//dfs = func(i int, visited []int8) bool {
	//	if visited[i] != 0 {
	//		return visited[i] == 1 // -1：检测到环，1：已检测且没换
	//	}
	//	visited[i] = -1 // 检测中
	//	for _, v := range adj[i] {
	//		if !dfs(v, visited) {
	//			return false // 存在环
	//		}
	//	}
	//	visited[i] = 1 // 已检测
	//	return true
	//}
	//for i := 0; i < numCourses; i++ {
	//	if !dfs(i, visited) { // 访问每个顶点
	//		return false
	//	}
	//}
	//return true

	// 拓扑排序：dfs
	//inverseAdj := make([][]int, numCourses)
	//for _, edge := range prerequisites {
	//	s, t := edge[0], edge[1]                 // s->t：s先于t，t依赖于s
	//	inverseAdj[t] = append(inverseAdj[t], s) // dfs 方案适合逆邻接表：inverseAdj
	//}
	//visited := make([]int8, numCourses)
	//for i := 0; i < numCourses; i++ { // 遍历每个顶点
	//	if visited[i] == 0 && !topoDfsCircle(inverseAdj, visited, i) { // 未访问的顶点
	//		return false
	//	}
	//}
	//return true
}
func topoDfsCircle(inverseAdj [][]int, visited []int8, t int) bool {
	if visited[t] != 0 { // 已访问的顶点
		return visited[t] == 1
	}
	visited[t] = -1 // 检测中
	for _, s := range inverseAdj[t] {
		if visited[s] <= 0 && !topoDfsCircle(inverseAdj, visited, s) {
			return false // 存在环
		}
	}
	visited[t] = 1 // 已检测
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
