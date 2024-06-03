package main

import (
	"fmt"
	"math"
)

func main() {
	graph := [][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}}
	graph = [][]int{{2},
		{2, 8},
		{0, 1, 3, 4},
		{2},
		{2},
		{8, 6},
		{5},
		{8},
		{1, 5, 7}} // 11
	graph = [][]int{{1, 2, 3, 4},
		{0, 2},
		{0, 1},
		{0, 5},
		{0, 6},
		{3},
		{4}} // 7
	//graph = [][]int{{1}, {0}}
	//graph = [][]int{{}}
	length := shortestPathLength(graph)
	fmt.Println(length)
}

// leetcode submit region begin(Prohibit modification and deletion)
func shortestPathLength(graph [][]int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func shortestPathLength_(graph [][]int) int {
	// Floyd + 状态dp
	const inf = math.MaxInt32 >> 1
	n := len(graph)
	m := 1 << n
	floyd := make([][]int, n) // floyd 算法
	for i := range floyd {
		floyd[i] = make([]int, n)
		for j := range floyd[i] {
			floyd[i][j] = inf // 初始化
		}
		floyd[i][i] = 0 // i->i
	}
	for i, g := range graph {
		for _, j := range g {
			floyd[i][j] = 1 // 权值设为 1
		}
	}
	for k := 0; k < n; k++ { // 最短距
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				floyd[i][j] = min(floyd[i][j], floyd[i][k]+floyd[k][j])
			}
		}
	}
	dp := make([][]int, n) // 状压 dp
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = inf
		}
		dp[i][1<<i] = 0
	}
	for mask := 1; mask < m; mask++ {
		if mask&(mask-1) == 0 {
			continue
		}
		// 重要：https://leetcode.cn/problems/shortest-path-visiting-all-nodes/solutions/918311/fang-wen-suo-you-jie-dian-de-zui-duan-lu-mqc2/
		for u := 0; u < n; u++ { // dp(u,mask)：从任一节点开始到节点 u 为止，并且经过的“关键节点”对应的二进制表示为 mask 时的最短路径长度
			if mask>>u&1 == 0 {
				continue
			}
			for v := 0; v < n; v++ { // 枚举“关键节点”进行状态转移
				if v != u && 1<<v&mask > 0 { // u 时最后一个关键节点，状态转移方程：dp[u][mask] = min(dp[u][mask], dp[v][mask&^(1<<u)]+floyd[v][u])
					dp[u][mask] = min(dp[u][mask], dp[v][mask&^(1<<u)]+floyd[v][u])
				}
			}
		}
	}
	ans := inf
	for _, d := range dp {
		ans = min(ans, d[m-1])
	}
	return ans

	// bfs
	//n := len(graph)
	//if n == 1 {
	//	return 0
	//}
	//m := 1 << n
	//memo := make([][]bool, n)
	//q := make([][3]int, 0, n)
	//for i := range memo {
	//	memo[i] = make([]bool, m)
	//	mask := (m - 1) &^ (1 << i)
	//	memo[i][mask] = true
	//	q = append(q, [3]int{i, mask, 0})
	//}
	//for {
	//	i, mask, d := q[0][0], q[0][1], q[0][2]+1
	//	q = q[1:]
	//	for _, j := range graph[i] {
	//		if next := mask &^ (1 << j); !memo[j][next] {
	//			if next == 0 {
	//				return d
	//			}
	//			q = append(q, [3]int{j, next, d})
	//			memo[j][next] = true
	//		}
	//	}
	//}

	// lc
	//n := len(graph)
	//m := 1 << n
	//q, memo := make([][3]int, n), make([][]bool, n)
	//for i := 0; i < n; i++ {
	//	memo[i] = make([]bool, m)
	//	memo[i][1<<i] = true
	//	q[i] = [3]int{i, 1 << i, 0}
	//}
	//for m--; ; q = q[1:] {
	//	i, mask, d := q[0][0], q[0][1], q[0][2]
	//	if mask == m {
	//		return d
	//	}
	//	d++
	//	for _, j := range graph[i] {
	//		if next := mask | 1<<j; !memo[j][next] {
	//			memo[j][next] = true
	//			q = append(q, [3]int{j, next, d})
	//		}
	//	}
	//}
}
