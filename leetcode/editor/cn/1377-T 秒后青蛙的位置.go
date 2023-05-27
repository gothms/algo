//给你一棵由 n 个顶点组成的无向树，顶点编号从 1 到 n。青蛙从 顶点 1 开始起跳。规则如下：
//
//
// 在一秒内，青蛙从它所在的当前顶点跳到另一个 未访问 过的顶点（如果它们直接相连）。
// 青蛙无法跳回已经访问过的顶点。
// 如果青蛙可以跳到多个不同顶点，那么它跳到其中任意一个顶点上的机率都相同。
// 如果青蛙不能跳到任何未访问过的顶点上，那么它每次跳跃都会停留在原地。
//
//
// 无向树的边用数组 edges 描述，其中 edges[i] = [ai, bi] 意味着存在一条直接连通 ai 和 bi 两个顶点的边。
//
// 返回青蛙在 t 秒后位于目标顶点 target 上的概率。与实际答案相差不超过 10⁻⁵ 的结果将被视为正确答案。
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 7, edges = [[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]], t = 2, target = 4
//输出：0.16666666666666666
//解释：上图显示了青蛙的跳跃路径。青蛙从顶点 1 起跳，第 1 秒 有 1/3 的概率跳到顶点 2 ，然后第 2 秒 有 1/2 的概率跳到顶点 4，因此青蛙
//在 2 秒后位于顶点 4 的概率是 1/3 * 1/2 = 1/6 = 0.16666666666666666 。
//
//
// 示例 2：
//
//
//
//
//输入：n = 7, edges = [[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]], t = 1, target = 7
//输出：0.3333333333333333
//解释：上图显示了青蛙的跳跃路径。青蛙从顶点 1 起跳，有 1/3 = 0.3333333333333333 的概率能够 1 秒 后跳到顶点 7 。
//
//
//
//
//
//
// 提示：
//
//
// 1 <= n <= 100
// edges.length == n - 1
// edges[i].length == 2
// 1 <= ai, bi <= n
// 1 <= t <= 50
// 1 <= target <= n
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 图 👍 54 👎 0

package main

import "fmt"

func main() {
	n, t, target := 7, 2, 4
	n, t, target = 7, 20, 6
	n, t, target = 6, 3, 4
	edges := [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}
	edges = [][]int{{1, 2}, {1, 3}, {2, 5}, {2, 6}, {3, 5}, {3, 6}, {5, 4}, {6, 4}}
	position := frogPosition(n, edges, t, target)
	fmt.Println(position)

	//m := map[int]int{1: 1, 2: 2, 3: 3}
	//for k, v := range m {
	//	if k == 2 {
	//		continue
	//	}
	//	fmt.Println(k, v)
	//}
}

/*
思路：bfs
	1.构建顶点路径的 邻接表，和一个已访问过顶点的记录数组
	2.操作流程
		2.1.将第 i 次访问到的所有顶点，放入一个 cache 的 map 中
			key value 分别是 顶点编号 和 访问到该顶点的概率
		2.2.遍历 cache 取得 k v，分别从 邻接表 中取得 k 对应的第 i+1 次能访问到的顶点
			那么第 i+1 次的概率分别是 v * (k能访问的顶点数-1)，k 就是 -1 的顶点
	3.注意点
		3.1.用一个哨兵顶点 0，把节点 1 的处理和 其他节点 统一
		3.2.“原地跳”的情况时，可以剪枝（原地跳 是 目标顶点 时，说明找到了结果）
		3.3.如果提前访问到了 目标顶点，返回 0
思路：dfs
	1.先建立 邻接表，深度 遍历 邻接表
		整体思路和 bfs 类似，包括注意点
	2.本代码提供返回 int 的写法
坑：
	重要提示：顶点数为 n，edges.length == n - 1
	可知：顶点之间不会出现多对一
*/
//leetcode submit region begin(Prohibit modification and deletion)
func frogPosition(n int, edges [][]int, t int, target int) float64 {
	// dfs
	memo, used := make([][]int, n+1), make([]bool, n+1)
	memo[1], used[0] = append(memo[1], 0), true
	for _, e := range edges {
		memo[e[0]], memo[e[1]] = append(memo[e[0]], e[1]), append(memo[e[1]], e[0])
	}
	var dfs func(int, int) int
	dfs = func(i, ts int) int {
		l := len(memo[i]) - 1
		if ts == t || l == 0 {
			if i == target {
				return 1.0
			}
			return 0.0
		}
		used[i] = true
		for j := 0; j <= l; j++ {
			if used[memo[i][j]] {
				continue
			}
			if v := dfs(memo[i][j], ts+1); v > 0 {
				return v * l
			}
		}
		return 0.0
	}
	if v := dfs(1, 0); v > 0 {
		return 1.0 / float64(v)
	}
	return 0.0

	// 返回 float
	//memo, used := make([][]int, n+1), make([]bool, n+1)
	//memo[1], used[0] = append(memo[1], 0), true
	//for _, e := range edges {
	//	memo[e[0]], memo[e[1]] = append(memo[e[0]], e[1]), append(memo[e[1]], e[0])
	//}
	//var dfs func(int, int) float64
	//dfs = func(i, ts int) (v float64) {
	//	l := len(memo[i]) - 1
	//	if ts == t || l == 0 {
	//		if i == target {
	//			v = 1.0
	//		}
	//		return
	//	} else if i == target {
	//		return
	//	}
	//	used[i] = true
	//	for j := 0; j <= l; j++ {
	//		if !used[memo[i][j]] {
	//			v += dfs(memo[i][j], ts+1)
	//		}
	//	}
	//	return v / float64(l)
	//}
	//return dfs(1, 0)

	// bfs
	memo, used := make([][]int, n+1), make([]bool, n+1)
	memo[1], used[0], used[1] = append(memo[1], 0), true, true
	for _, e := range edges {
		memo[e[0]], memo[e[1]] = append(memo[e[0]], e[1]), append(memo[e[1]], e[0])
	}
	cache := map[int]float64{1: 1.0}
	for i := 1; i <= t; i++ {
		next := make(map[int]float64)
		for k, v := range cache {
			if len(memo[k]) == 1 {
				if k == target {
					return v
				}
				continue
			} else if k == target && i < t {
				return 0
			}
			for _, nk := range memo[k] {
				if !used[nk] {
					next[nk], used[nk] = v/float64(len(memo[k])-1), true
				}
			}
		}
		cache = next
	}
	return cache[target]

	// bfs：1
	//memo, used := make([][]int, n+1), make([]bool, n+1)
	//memo[1], used[1] = append(memo[1], 0), true
	//for _, e := range edges {
	//	memo[e[0]], memo[e[1]] = append(memo[e[0]], e[1]), append(memo[e[1]], e[0])
	//}
	//queue := [][2]int{{1, 1}}
	//for i := 0; i < t; i++ {
	//	l := len(queue)
	//	for j := 0; j < l; j++ {
	//		k := len(memo[queue[j][0]]) - 1
	//		if len(memo[queue[j][0]]) == 1 {
	//			if queue[j][0] == target {
	//				return 1.0 / float64(queue[j][1])
	//			}
	//			continue
	//		}
	//		for _, p := range memo[queue[j][0]] {
	//			if !used[p] {
	//				queue = append(queue, [2]int{p, queue[j][1] * k})
	//				used[p] = true
	//			}
	//		}
	//	}
	//	queue = queue[l:]
	//}
	//for i := 0; i < len(queue); i++ {
	//	if queue[i][0] == target {
	//		if queue[i][1] > 0 {
	//			return 1.0 / float64(queue[i][1])
	//		}
	//		break
	//	}
	//}
	//return 0

	// bfs：1
	//memo, l := make([]map[int]bool, n+1), len(edges)
	//for i := 0; i < l; i++ {
	//	m1 := memo[edges[i][0]]
	//	if m1 == nil {
	//		m1 = make(map[int]bool)
	//		memo[edges[i][0]] = m1
	//	}
	//	m1[edges[i][1]] = true
	//	m2 := memo[edges[i][1]]
	//	if m2 == nil {
	//		m2 = make(map[int]bool)
	//		memo[edges[i][1]] = m2
	//	}
	//	m2[edges[i][0]] = true
	//}
	////fmt.Println(memo)
	//used := make([]bool, n+1)
	//used[1] = true
	//curr := map[int]int{1: 1}
	//for i := 0; i < t; i++ {
	//	next := make(map[int]int)
	//	for k, v := range curr {
	//		p := 0
	//		for nk := range memo[k] {
	//			if !used[nk] {
	//				p++
	//				used[nk] = true
	//			} else {
	//				delete(memo[k], nk)
	//			}
	//		}
	//		//fmt.Println(k, p, len(memo[k]))
	//		if len(memo[k]) == 0 { // or p==0
	//			next[k] = v // 原地，但 k!=target，可以作废了
	//		} else {
	//			for nk := range memo[k] {
	//				next[nk] = v * p
	//			}
	//		}
	//	}
	//	curr = next
	//}
	////fmt.Println(curr)
	//fmt.Println(memo)
	//if curr[target] == 0 {
	//	return 0.0
	//}
	//return 1.0 / float64(curr[target])
}

//leetcode submit region end(Prohibit modification and deletion)
