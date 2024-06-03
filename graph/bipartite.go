package graph

/*
https://oi-wiki.org/graph/bi-graph/

lc
	785
	2493
*/

// BipartiteDFS 染色法判断二分图
func BipartiteDFS(n int, edges [][]int) bool {
	const (
		RED int = iota - 1 // 染色法
		UNCOLORED
		GREEN // 技巧：另种颜色互为负数
	)
	color, adj := make([]int, n), make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		adj[x], adj[y] = append(adj[x], y), append(adj[y], x)
	}
	valid := true // 标记是否为二分图
	var dfs func(int, int, [][]int)
	dfs = func(f, c int, graph [][]int) {
		color[f] = c
		for _, t := range graph[f] {
			switch color[t] {
			case UNCOLORED: // 未被标记
				dfs(t, -c, graph)
				if !valid {
					return
				}
			case c: // 已被标记，且颜色和 f 相同，说明环中有奇数个顶点
				valid = false
				return
			}
		}
	}
	for i := 0; i < n && valid; i++ { // 遍历每个顶点：在 dfs 中更改 valid
		if color[i] == UNCOLORED {
			dfs(i, RED, adj) // 默认从 RED 开始标记
		}
	}
	return valid
}

// BipartiteBFS BFS 判断二分图
func BipartiteBFS(n int, edges [][]int) bool {
	visited, adj := make([]int, n), make([][]int, n) // 0：未访问，1 & -1：标记两种颜色
	for _, e := range edges {
		x, y := e[0], e[1]
		adj[x], adj[y] = append(adj[x], y), append(adj[y], x)
	}
	queue := make([]int, 0)
	for i := 0; i < n; i++ { // 图中可能有多个连通分量，则需要访问每个顶点
		if visited[i] != 0 { // 已染色
			continue
		}
		queue = append(queue, i)
		visited[i] = 1
		for len(queue) > 0 {
			f := queue[0]
			queue = queue[1:]
			for _, t := range adj[f] {
				switch visited[t] {
				case visited[f]: // 已染色，且 f 和 t 颜色相同
					return false
				case 0: // 未染色
					visited[t] = -visited[f] // 染色（和 f 不同）
					queue = append(queue, t)
				}
			}
		}
	}
	return true
}

// BipartiteUnion 并查集判断二分图
func BipartiteUnion(n int, edges [][]int) bool {
	uni := make([]int, n)
	for i := range uni {
		uni[i] = i
	}
	find := func(p int) int {
		for p != uni[p] {
			p, uni[p] = uni[p], uni[uni[p]]
		}
		return uni[p]
	}
	union := func(p, q int) bool {
		rootP, rootQ := find(p), find(q)
		if rootP != rootQ {
			uni[rootQ] = rootP
			return false
		}
		return true
	}
	adj := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		adj[x], adj[y] = append(adj[x], y), append(adj[y], x)
	}
	for f, nodes := range adj { // 遍历邻接表
		for _, t := range nodes {
			if find(f) == find(t) { // f->t，却属于同一集合
				return false
			}
			union(nodes[0], t) // 与 f 有边的顶点 t，都分到同一集合
		}
	}
	return true
}
