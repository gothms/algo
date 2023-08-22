//给你一个整数 n ，表示有 n 节课，课程编号从 1 到 n 。同时给你一个二维整数数组 relations ，其中 relations[j] = [
//prevCoursej, nextCoursej] ，表示课程 prevCoursej 必须在课程 nextCoursej 之前 完成（先修课的关系）。同时给你一个下
//标从 0 开始的整数数组 time ，其中 time[i] 表示完成第 (i+1) 门课程需要花费的 月份 数。
//
// 请你根据以下规则算出完成所有课程所需要的 最少 月份数：
//
//
// 如果一门课的所有先修课都已经完成，你可以在 任意 时间开始这门课程。
// 你可以 同时 上 任意门课程 。
//
//
// 请你返回完成所有课程所需要的 最少 月份数。
//
// 注意：测试数据保证一定可以完成所有课程（也就是先修课的关系构成一个有向无环图）。
//
//
//
// 示例 1:
//
//
//
// 输入：n = 3, relations = [[1,3],[2,3]], time = [3,2,5]
//输出：8
//解释：上图展示了输入数据所表示的先修关系图，以及完成每门课程需要花费的时间。
//你可以在月份 0 同时开始课程 1 和 2 。
//课程 1 花费 3 个月，课程 2 花费 2 个月。
//所以，最早开始课程 3 的时间是月份 3 ，完成所有课程所需时间为 3 + 5 = 8 个月。
//
//
// 示例 2：
//
//
//
// 输入：n = 5, relations = [[1,5],[2,5],[3,5],[3,4],[4,5]], time = [1,2,3,4,5]
//输出：12
//解释：上图展示了输入数据所表示的先修关系图，以及完成每门课程需要花费的时间。
//你可以在月份 0 同时开始课程 1 ，2 和 3 。
//在月份 1，2 和 3 分别完成这三门课程。
//课程 4 需在课程 3 之后开始，也就是 3 个月后。课程 4 在 3 + 4 = 7 月完成。
//课程 5 需在课程 1，2，3 和 4 之后开始，也就是在 max(1,2,3,7) = 7 月开始。
//所以完成所有课程所需的最少时间为 7 + 5 = 12 个月。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 5 * 10⁴
// 0 <= relations.length <= min(n * (n - 1) / 2, 5 * 10⁴)
// relations[j].length == 2
// 1 <= prevCoursej, nextCoursej <= n
// prevCoursej != nextCoursej
// 所有的先修课程对 [prevCoursej, nextCoursej] 都是 互不相同 的。
// time.length == n
// 1 <= time[i] <= 10⁴
// 先修课程图是一个有向无环图。
//
//
// Related Topics 图 拓扑排序 数组 动态规划 👍 32 👎 0

package main

import (
	"fmt"
)

func main() {
	n := 5
	relations := [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}
	time := []int{1, 2, 3, 4, 5}
	i := minimumTime(n, relations, time)
	fmt.Println(i)
}

/*
思路：kahn+优先队列
	1.课程之间的关系，可以用拓扑排序直观的展示
		我们可以贪心的选择当前所修课程中，最先能修完的课程，看它的拓扑排序关系中，哪些课程是可选的下一门课
		最先修完的课程，我们可以用 优先队列 来选择（根据 time 加上其开始修的时间）
		另外选择 kahn 拓扑排序算法，记录每门课程的入度，当入度为 0 时，说明这门课程当前可修
	2.具体实现
		adj：邻接表记录先修关系图的
		kahn：记录每门课程的入度
			当某个课程修完时，我们遍历这门课程的后修课程（邻接表）
			将每门后修课程 i 的入度 -1，kahn[i] 表示的就是 第i门课 的先修课程还有 kahn[i] 门
			当 kahn[i]==0 时，说明这门课程的先修课都修完了，此时将这门课加入 优先队列 中
		th：自定义的小顶堆，用于选择当前最先修完的课程
思路：kahn+dp
	1.对于任意课程 i
		因为必须修完它的所有先修课，所以我们可以从它的先修课程中，选出修完用时最多的课程
		作为 i 开始修的起点
		也就是对上述方案（kahn+优先队列）的一种优化，使用队列来代替优先队列
	2.状态转移方程
		dp[i] = maxVal(dp[i], dp[curr])
			dp[i] 即下一门可能要修的课程
			dp[curr] 即当前修完的课程，当 curr 修完后，我们尝试更新开始修下一门课程的时间
		i 的每一门先修课程结束时，都会尝试更新 i 的起修时间
*/
// leetcode submit region begin(Prohibit modification and deletion)
func minimumTime(n int, relations [][]int, time []int) int {
	// dp：kahn+优先队列 的优化
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	cnt := 0
	kahn, adj := make([]int, n), make([][]int, n)
	for _, edge := range relations {
		adj[edge[0]-1] = append(adj[edge[0]-1], edge[1]-1) // 初始化邻接表
		kahn[edge[1]-1]++                                  // 初始化 kahn
	}
	queue, dp := make([]int, 0), make([]int, n) // dp 记录每门课需修完的时间
	for i := 0; i < n; i++ {
		if kahn[i] == 0 {
			queue = append(queue, i) // 初始化队列
		}
	}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]           // 修完一门课
		dp[curr] += time[curr]      // 课程修完，更新时间
		cnt = maxVal(cnt, dp[curr]) // 记录当前最后一门课的时长
		for _, i := range adj[curr] {
			dp[i] = maxVal(dp[i], dp[curr]) // 更新 i 门的修完时间
			if kahn[i]--; kahn[i] == 0 {
				queue = append(queue, i) // 开始修一门课
			}
		}
	}
	return cnt

	// kahn+堆：个人
	//cnt := 0
	//h, kahn, adj := &th{}, make([]int, n+1), make([][]int, n+1)
	//for _, edge := range relations {
	//	adj[edge[0]] = append(adj[edge[0]], edge[1]) // 初始化邻接表
	//	kahn[edge[1]]++                              // 初始化 kahn
	//}
	//for i := 1; i <= n; i++ {
	//	if kahn[i] == 0 {
	//		heap.Push(h, [2]int{i, time[i-1]}) // 初始化优先队列
	//	}
	//}
	//for h.Len() > 0 {
	//	top := heap.Pop(h).([2]int)     // 最先修完的课程
	//	cnt = top[1]                    // 当前花费时间
	//	for _, i := range adj[top[0]] { // 下一门可修课程：拓扑排序
	//		kahn[i]--         // 入度 -1
	//		if kahn[i] == 0 { // 入度为 0
	//			heap.Push(h, [2]int{i, cnt + time[i-1]}) // 修下门课程
	//		}
	//	}
	//}
	//return cnt
}

type th [][2]int

func (t th) Len() int {
	return len(t)
}
func (t th) Less(i, j int) bool {
	return t[i][1] < t[j][1]
}
func (t th) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t *th) Push(x any) {
	*t = append(*t, x.([2]int))
}
func (t *th) Pop() any {
	last := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return last
}

//leetcode submit region end(Prohibit modification and deletion)
