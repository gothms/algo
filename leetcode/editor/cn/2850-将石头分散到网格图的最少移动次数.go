package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		{1, 3, 0},
		{1, 0, 0},
		{1, 0, 3}} // 4
	grid = [][]int{
		{1, 2, 2},
		{1, 1, 0},
		{0, 1, 1}} // 4
	grid = [][]int{
		{3, 0, 0},
		{4, 1, 0},
		{1, 0, 0}} // 10
	moves := minimumMoves(grid)
	fmt.Println(moves)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumMoves(grid [][]int) int {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}
	// 最小费用最大流
	// https://leetcode.cn/problems/minimum-moves-to-spread-stones-over-grid/solutions/2435313/tong-yong-zuo-fa-zui-xiao-fei-yong-zui-d-iuw8/
	m, n := len(grid), len(grid[0])
	src := m * n   // 超级源点
	dst := src + 1 // 超级汇点
	type edge struct{ to, rid, cap, cost int }
	g := make([][]edge, m*n+2)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], edge{to, len(g[to]), cap, cost})
		g[to] = append(g[to], edge{from, len(g[from]) - 1, 0, -cost})
	}
	for x, row := range grid {
		for y, v := range row {
			if v > 1 {
				addEdge(src, x*n+y, v-1, 0)
				for i, r := range grid {
					for j, w := range r {
						if w == 0 {
							addEdge(x*n+y, i*n+j, 1, abs(x-i)+abs(y-j))
						}
					}
				}
			} else if v == 0 {
				addEdge(x*n+y, dst, 1, 0)
			}
		}
	}

	// 下面是最小费用最大流模板
	const inf int = 1e9
	dist := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	spfa := func() bool {
		for i := range dist {
			dist[i] = 1e9
		}
		dist[src] = 0
		inQ := make([]bool, len(g))
		inQ[src] = true
		q := []int{src}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				if newD := dist[v] + e.cost; newD < dist[w] {
					dist[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						q = append(q, w)
						inQ[w] = true
					}
				}
			}
		}
		return dist[dst] < inf
	}
	ek := func() (maxFlow, minCost int) {
		for spfa() {
			// 沿 st-end 的最短路尽量增广
			minF := inf
			for v := dst; v != src; {
				p := fa[v]
				if c := g[p.v][p.i].cap; c < minF {
					minF = c
				}
				v = p.v
			}
			for v := dst; v != src; {
				p := fa[v]
				e := &g[p.v][p.i]
				e.cap -= minF
				g[v][e.rid].cap += minF
				v = p.v
			}
			maxFlow += minF
			minCost += dist[dst] * minF
		}
		return
	}
	_, cost := ek()
	return cost

	// 状压 dp
	//f, t := make([][2]int, 0), make([][2]int, 0)
	//for i, g := range grid {
	//	for j, v := range g {
	//		if v > 1 {
	//			for k := 1; k < v; k++ {
	//				f = append(f, [2]int{i, j})
	//			}
	//		} else if v == 0 {
	//			t = append(t, [2]int{i, j})
	//		}
	//	}
	//}
	//mask := 1 << len(f)
	//dp := make([]int, mask)
	//for i := 1; i < mask; i++ {
	//	dp[i] = 18
	//}
	//for i := 1; i < mask; i++ {
	//	tv := t[bits.OnesCount(uint(i))-1]
	//	for j, fv := range f {
	//		if 1<<j&i != 0 {
	//			dp[i] = min(dp[i], dp[i^(1<<j)]+abs(fv[0]-tv[0])+abs(fv[1]-tv[1]))
	//		}
	//	}
	//}
	//return dp[mask-1]

	// 全排列：曼哈顿距离
	//f, t := make([][2]int, 0), make([][2]int, 0)
	//ans := 100
	//for i, g := range grid {
	//	for j, v := range g {
	//		if v > 1 {
	//			for k := 1; k < v; k++ {
	//				f = append(f, [2]int{i, j})
	//			}
	//		} else if v == 0 {
	//			t = append(t, [2]int{i, j})
	//		}
	//	}
	//}
	//var permute func(int, func())
	//permute = func(i int, do func()) {
	//	if i == len(f) {
	//		do()
	//		return
	//	}
	//	permute(i+1, do)
	//	for j := i + 1; j < len(f); j++ {
	//		f[i], f[j] = f[j], f[i]
	//		permute(i+1, do)
	//		f[i], f[j] = f[j], f[i]
	//	}
	//}
	//permute(0, func() {
	//	distSum := 0
	//	for i := range f {
	//		fv, tv := f[i], t[i]
	//		distSum += abs(fv[0]-tv[0]) + abs(fv[1]-tv[1])
	//	}
	//	ans = min(ans, distSum)
	//})
	//return ans

	// bfs：错误
}

//leetcode submit region end(Prohibit modification and deletion)
