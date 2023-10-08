//一个公司准备组织一场会议，邀请名单上有 n 位员工。公司准备了一张 圆形 的桌子，可以坐下 任意数目 的员工。
//
// 员工编号为 0 到 n - 1 。每位员工都有一位 喜欢 的员工，每位员工 当且仅当 他被安排在喜欢员工的旁边，他才会参加会议。每位员工喜欢的员工 不会
//是他自己。
//
// 给你一个下标从 0 开始的整数数组 favorite ，其中 favorite[i] 表示第 i 位员工喜欢的员工。请你返回参加会议的 最多员工数目 。
//
//
//
//
// 示例 1：
//
//
//
// 输入：favorite = [2,2,1,2]
//输出：3
//解释：
//上图展示了公司邀请员工 0，1 和 2 参加会议以及他们在圆桌上的座位。
//没办法邀请所有员工参与会议，因为员工 2 没办法同时坐在 0，1 和 3 员工的旁边。
//注意，公司也可以邀请员工 1，2 和 3 参加会议。
//所以最多参加会议的员工数目为 3 。
//
//
// 示例 2：
//
// 输入：favorite = [1,2,0]
//输出：3
//解释：
//每个员工都至少是另一个员工喜欢的员工。所以公司邀请他们所有人参加会议的前提是所有人都参加了会议。
//座位安排同图 1 所示：
//- 员工 0 坐在员工 2 和 1 之间。
//- 员工 1 坐在员工 0 和 2 之间。
//- 员工 2 坐在员工 1 和 0 之间。
//参与会议的最多员工数目为 3 。
//
//
// 示例 3：
//
//
//
// 输入：favorite = [3,0,1,4,1]
//输出：4
//解释：
//上图展示了公司可以邀请员工 0，1，3 和 4 参加会议以及他们在圆桌上的座位。
//员工 2 无法参加，因为他喜欢的员工 0 旁边的座位已经被占领了。
//所以公司只能不邀请员工 2 。
//参加会议的最多员工数目为 4 。
//
//
//
//
// 提示：
//
//
// n == favorite.length
// 2 <= n <= 10⁵
// 0 <= favorite[i] <= n - 1
// favorite[i] != i
//
//
// Related Topics 深度优先搜索 图 拓扑排序 👍 88 👎 0

package main

import "fmt"

func main() {
	favorite := []int{3, 0, 1, 4, 1} // 4
	favorite = []int{1, 0, 3, 2, 5, 6, 7, 4, 9, 8, 11, 10, 11, 12, 10}
	invitations := maximumInvitations(favorite)
	fmt.Println()
	fmt.Println(invitations)
}

/*
内向基环树
	从 i 向 favorite[i] 连边，我们可以得到一张有向图
	由于每个大小为 k 的连通块都有 k 个点和 k 条边，所以每个连通块必定有且仅有一个环
	且由于每个点的出度均为 1，这样的有向图又叫做内向基环树 (pseudotree)，由基环树组成的森林叫基环树森林 (pseudoforest)

*/

// leetcode submit region begin(Prohibit modification and deletion)
func maximumInvitations(favorite []int) int {
	// 内向基环树
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	chainSum, maxRing, n := 0, 0, len(favorite)
	kahn, queue, depth := make([]int, n), make([]int, 0), make([]int, n)
	for _, t := range favorite {
		kahn[t]++ // 统计入度
	}
	for i, v := range kahn {
		if v == 0 {
			queue = append(queue, i) // 入度为 0
		}
	}
	for len(queue) > 0 { // 拓扑排序：剪去树枝，只剩下环
		f := queue[0]
		queue = queue[1:]
		depth[f]++
		t := favorite[f]                      // 唯一的出
		depth[t] = maxVal(depth[t], depth[f]) // 最深的树枝
		if kahn[t]--; kahn[t] == 0 {
			queue = append(queue, t)
		}
	}
	for i, v := range kahn {
		if v == 0 {
			continue
		}
		ringSize := 1 // 环的大小
		for t := favorite[i]; t != i; t = favorite[t] {
			kahn[t] = 0 // 标记已遍历的环
			ringSize++
		}
		if ringSize == 2 { // 基环大小为 2，可拼接
			chainSum += 2 + depth[i] + depth[favorite[i]]
		} else { // 基环大小大于 2，不可拼接
			maxRing = maxVal(maxRing, ringSize)
		}
	}
	return maxVal(chainSum, maxRing) // 拼接长度 / 最大环

	// 拓扑排序：判断有误
	//	maxVal := func(a, b int) int {
	//		if b > a {
	//			return b
	//		}
	//		return a
	//	}
	//	n := len(favorite)
	//	cache := make([]int, n)
	//	adj := make([][]int, n)
	//	for f, t := range favorite {
	//		adj[t] = append(adj[t], f)
	//	}
	//	var dfs func(int) int
	//	dfs = func(i int) int {
	//		v := &cache[i]
	//		switch *v {
	//		case -1:
	//			return 0
	//		case 0:
	//			*v = -1
	//		default:
	//			return *v
	//		}
	//		max := 0
	//		for _, j := range adj[i] {
	//			max = maxVal(max, dfs(j))
	//		}
	//		max++
	//		*v = max
	//		return max
	//	}
	//	cnt := 0
	//	for i := 0; i < n; i++ {
	//		if cache[i] == 0 {
	//			cnt = maxVal(cnt, dfs(i))
	//			fmt.Println(i, cnt)
	//		}
	//	}
	//	fmt.Println(cache)
	//	return cnt
	//}

	//func checkCircle(favorite []int) int {
	//	// 拓扑排序：检测环
	//	n := len(favorite)
	//	visited := make([]int, n)
	//	var dfs func(int) bool
	//	dfs = func(i int) bool {
	//		if visited[i] != 0 {
	//			return visited[i] == 1
	//		}
	//		visited[i] = -1
	//		fmt.Print("->", i)
	//		if !dfs(favorite[i]) { // 环
	//			return false
	//		}
	//		visited[i] = 1
	//		return true
	//	}
	//	for i := 0; i < n; i++ {
	//		if visited[i] == 0 && !dfs(i) {
	//			return i // 有环
	//		}
	//		fmt.Println()
	//	}
	//	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
