//你在一个城市里，城市由 n 个路口组成，路口编号为 0 到 n - 1 ，某些路口之间有 双向 道路。输入保证你可以从任意路口出发到达其他任意路口，且任意两
//个路口之间最多有一条路。
//
// 给你一个整数 n 和二维整数数组 roads ，其中 roads[i] = [ui, vi, timei] 表示在路口 ui 和 vi 之间有一条需要花费
// timei 时间才能通过的道路。你想知道花费 最少时间 从路口 0 出发到达路口 n - 1 的方案数。
//
// 请返回花费 最少时间 到达目的地的 路径数目 。由于答案可能很大，将结果对 10⁹ + 7 取余 后返回。
//
//
//
// 示例 1：
// 输入：n = 7, roads = [[0,6,7],[0,1,2],[1,2,3],[1,3,3],[6,3,3],[3,5,1],[6,5,1],[2
//,5,1],[0,4,5],[4,6,2]]
//输出：4
//解释：从路口 0 出发到路口 6 花费的最少时间是 7 分钟。
//四条花费 7 分钟的路径分别为：
//- 0 ➝ 6
//- 0 ➝ 4 ➝ 6
//- 0 ➝ 1 ➝ 2 ➝ 5 ➝ 6
//- 0 ➝ 1 ➝ 3 ➝ 5 ➝ 6
//
//
// 示例 2：
//
// 输入：n = 2, roads = [[1,0,10]]
//输出：1
//解释：只有一条从路口 0 到路口 1 的路，花费 10 分钟。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 200
// n - 1 <= roads.length <= n * (n - 1) / 2
// roads[i].length == 3
// 0 <= ui, vi <= n - 1
// 1 <= timei <= 10⁹
// ui != vi
// 任意两个路口之间至多有一条路。
// 从任意路口出发，你能够到达其他任意路口。
//
//
// Related Topics 图 拓扑排序 动态规划 最短路 👍 132 👎 0

package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	roads := [][]int{{0, 1, 1000000000},
		{1, 2, 1000000000}}
	n := 2
	paths := countPaths(n, roads)
	fmt.Println(paths)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countPaths(n int, roads [][]int) int {
	const mod = 1_000_000_007
	var (
		adj = make([][][2]int, n) // 邻接表
		dis = make([]int, n)      // 0 到 其他顶点 的距离
		cnt = make([]int, n)      // 顶点的最短路径的路径数
		h   = &cpHp{{}}           // 优先队列
	)
	for _, r := range roads {
		adj[r[0]] = append(adj[r[0]], [2]int{r[1], r[2]})
		adj[r[1]] = append(adj[r[1]], [2]int{r[0], r[2]})
	}
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt // 不能是 math.MaxInt32
	}
	for cnt[0] = 1; ; { // 初始化
		cur := heap.Pop(h).([2]int)
		if cur[0] == n-1 {
			return cnt[n-1]
		}
		if cur[1] > dis[cur[0]] {
			continue
		}
		for _, e := range adj[cur[0]] {
			if d := cur[1] + e[1]; d < dis[e[0]] { // 更新距离
				dis[e[0]] = d
				cnt[e[0]] = cnt[cur[0]]
				heap.Push(h, [2]int{e[0], d})
			} else if d == dis[e[0]] { // 到顶点 e[0] 的距离相等
				cnt[e[0]] = (cnt[e[0]] + cnt[cur[0]]) % mod
			}
		}
	}
}

type cpHp [][2]int

func (c cpHp) Len() int           { return len(c) }
func (c cpHp) Less(i, j int) bool { return c[i][1] < c[j][1] }
func (c cpHp) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c *cpHp) Push(x any)        { *c = append(*c, x.([2]int)) }

func (c *cpHp) Pop() any {
	v := (*c)[c.Len()-1]
	*c = (*c)[:c.Len()-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)

//func countPaths(n int, roads [][]int) int {
//	type edge struct{ to, d int }
//	g := make([][]edge, n) // 邻接表
//	for _, r := range roads {
//		x, y, d := r[0], r[1], r[2]
//		g[x] = append(g[x], edge{y, d})
//		g[y] = append(g[y], edge{x, d})
//	}
//
//	dis := make([]int, n)
//	for i := 1; i < n; i++ {
//		dis[i] = math.MaxInt
//	}
//	f := make([]int, n)
//	f[0] = 1
//	h := &hp{{}}
//	for {
//		fmt.Println(h)
//		p := heap.Pop(h).(pair)
//		x := p.x
//		if x == n-1 {
//			// 不可能找到比 dis[n-1] 更短，或者一样短的最短路了（注意本题边权都是正数）
//			return f[n-1]
//		}
//		if p.dis > dis[x] {
//			continue
//		}
//		for _, e := range g[x] { // 尝试更新 x 的邻居的最短路
//			y := e.to
//			newDis := p.dis + e.d
//			if newDis < dis[y] {
//				// 就目前来说，最短路必须经过 x
//				dis[y] = newDis
//				f[y] = f[x]
//				heap.Push(h, pair{newDis, y})
//			} else if newDis == dis[y] {
//				// 和之前求的最短路一样长
//				f[y] = (f[y] + f[x]) % 1_000_000_007
//			}
//		}
//	}
//}
//
//type pair struct{ dis, x int }
//type hp []pair
//
//func (h hp) Len() int           { return len(h) }
//func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
//func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
//func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
