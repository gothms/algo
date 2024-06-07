package graph

/*
Kruskal 算法是一种用于寻找加权无向图的最小生成树（MST）的算法。该算法的优化主要集中在以下几个方面：

1. **使用高效的数据结构来管理边和集合**：
   - **边的排序**：Kruskal 算法首先将图中的所有边按权重从小到大排序。这一步通常使用快速排序或其他高效排序算法，时间复杂度为 \(O(E \log E)\)，其中 \(E\) 是边的数量。
   - **并查集（Union-Find）**：并查集是一种高效的数据结构，用于管理元素之间的合并（Union）和查找（Find）操作。通过路径压缩（Path Compression）和按秩合并（Union by Rank）等优化技术，可以将每次查找和合并操作的时间复杂度优化到接近常数时间，即 \(O(\alpha(V))\)，其中 \(\alpha\) 是阿克曼函数的逆， \(V\) 是顶点的数量。

2. **并查集的路径压缩和按秩合并**：
   - **路径压缩（Path Compression）**：在查找（Find）操作时，将树的路径上的所有节点直接连接到根节点，减少树的高度。
   - **按秩合并（Union by Rank）**：在合并（Union）操作时，总是将较小的树合并到较大的树上，进一步减少树的高度。

这些优化技术使得 Kruskal 算法在处理大规模图时更加高效。

### Kruskal 算法的步骤

1. **将所有边按权重从小到大排序**。
2. **初始化一个并查集**，每个顶点作为一个独立的集合。
3. **遍历排序后的边**，对于每一条边，如果它连接的两个顶点属于不同的集合，则将这条边加入最小生成树，并合并这两个顶点所在的集合。
4. **重复步骤3**，直到最小生成树包含 \(V-1\) 条边（其中 \(V\) 是顶点的数量）。

### 代码示例

下面是一个用 Go 实现的 Kruskal 算法示例：

```go
package main

import (
	"fmt"
	"sort"
)

// Edge represents an edge with a source, destination and weight
type Edge struct {
	u, v, weight int
}

// Graph represents a graph with vertices and edges
type Graph struct {
	vertices int
	edges    []Edge
}

// Subset represents a subset for union-find
type Subset struct {
	parent, rank int
}

// NewGraph creates a new graph with given vertices and edges
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(u, v, weight int) {
	g.edges = append(g.edges, Edge{u, v, weight})
}

// Find finds the subset of an element i using path compression
func find(subsets []Subset, i int) int {
	if subsets[i].parent != i {
		subsets[i].parent = find(subsets, subsets[i].parent)
	}
	return subsets[i].parent
}

// Union unites two subsets x and y using union by rank
func union(subsets []Subset, x, y int) {
	xroot := find(subsets, x)
	yroot := find(subsets, y)

	if subsets[xroot].rank < subsets[yroot].rank {
		subsets[xroot].parent = yroot
	} else if subsets[xroot].rank > subsets[yroot].rank {
		subsets[yroot].parent = xroot
	} else {
		subsets[yroot].parent = xroot
		subsets[xroot].rank++
	}
}

// KruskalMST finds and returns the minimum spanning tree of the graph
func (g *Graph) KruskalMST() []Edge {
	var result []Edge

	// Sort all the edges in non-decreasing order of their weight
	sort.Slice(g.edges, func(i, j int) bool {
		return g.edges[i].weight < g.edges[j].weight
	})

	subsets := make([]Subset, g.vertices)
	for v := 0; v < g.vertices; v++ {
		subsets[v] = Subset{v, 0}
	}

	for _, edge := range g.edges {
		u := find(subsets, edge.u)
		v := find(subsets, edge.v)

		if u != v {
			result = append(result, edge)
			union(subsets, u, v)
		}
	}

	return result
}

func main() {
	g := NewGraph(4)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 6)
	g.AddEdge(0, 3, 5)
	g.AddEdge(1, 3, 15)
	g.AddEdge(2, 3, 4)

	mst := g.KruskalMST()
	fmt.Println("Edges in the MST:")
	for _, edge := range mst {
		fmt.Printf("%d -- %d == %d\n", edge.u, edge.v, edge.weight)
	}
}
```

### 解释：

1. **Edge 结构体**：表示一条边，包含起点、终点和权重。
2. **Graph 结构体**：表示图，包含顶点数量和边列表。
3. **Subset 结构体**：表示并查集的一个子集，包含父节点和秩（用于按秩合并）。
4. **NewGraph 函数**：创建一个新的图。
5. **AddEdge 函数**：向图中添加一条边。
6. **find 函数**：查找并查集的根节点，使用路径压缩优化。
7. **union 函数**：合并两个子集，使用按秩合并优化。
8. **KruskalMST 函数**：实现 Kruskal 算法，返回最小生成树中的边。
9. **main 函数**：创建一个示例图，调用 Kruskal 算法并输出结果。

通过这些优化，Kruskal 算法能够高效地找到图的最小生成树，适用于处理边数量较多的稀疏图。

*/
