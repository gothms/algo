package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfSets(n int, maxDistance int, roads [][]int) int {
	ans := 0
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = math.MaxInt / 2 // 防止加法溢出
		}
	}
	for _, e := range roads {
		x, y, wt := e[0], e[1], e[2]
		g[x][y] = min(g[x][y], wt)
		g[y][x] = min(g[y][x], wt)
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
next:
	for s := 0; s < 1<<n; s++ { // 枚举子集
		for i, row := range g {
			if s>>i&1 == 0 {
				continue
			}
			copy(f[i], row)
		}

		// Floyd 算法（只考虑在 s 中的节点）
		for k := range f {
			if s>>k&1 == 0 {
				continue
			}
			for i := range f {
				if s>>i&1 == 0 {
					continue
				}
				for j := range f {
					f[i][j] = min(f[i][j], f[i][k]+f[k][j])
				}
			}
		}

		// 判断保留的节点之间的最短路是否均不超过 maxDistance
		for i, di := range f {
			if s>>i&1 == 0 {
				continue
			}
			for j, dij := range di[:i] {
				if s>>j&1 > 0 && dij > maxDistance {
					continue next
				}
			}
		}
		ans++
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
