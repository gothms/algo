package algo

import (
	"container/heap"
	"fmt"
	"math"
)

/*
lc
	1553
*/

// AStar 从顶点s到顶点t的路径
// n 为顶点的个数，edges 为有向有权图
func AStar(n, s, t int, edges [][]int) {
	var (
		adj         = make([][]edge, n)                            // 邻接表
		vertexes    = make([]*vertexAStar, n)                      // 顶点 s 与顶点的距离和 f 值，以及顶点坐标
		visited     = make([]bool, n)                              // 记录已访问过的顶点（已进入过队列）
		predecessor = make([]int, n)                               // 用于还原最短路径
		queue       = &priorityQueueAStar{make([]*vertexAStar, 0)} // 小顶堆：根据 vertexAStar.f
		prt         func(int, int, []int)                          // 用于打印最 短路径
	)
	for _, e := range edges { // 遍历有向有权图
		adj[e[0]] = append(adj[e[0]], edge{e[0], e[1], e[2]}) // 添加一条边
	}
	for i := 0; i < n; i++ {
		vertexes[i] = &vertexAStar{id: i, dist: math.MaxInt32} // x y 未模拟
	}
	vertexes[s].dist, vertexes[s].f = 0, 0 // 初始化
	heap.Push(queue, vertexes[s])
	//visited[s] = true
out:
	for queue.Len() > 0 {
		minVertex := heap.Pop(queue).(*vertexAStar)
		if visited[minVertex.id] { // 先出队列的顶点，dist 总是更小
			continue
		}
		visited[minVertex.id] = true
		for _, e := range adj[minVertex.id] {
			nextVertex := vertexes[e.tid] // 边 minVertex-->nextVertex
			if minVertex.dist+e.w >= nextVertex.dist {
				continue
			}
			nextVertex.dist = minVertex.dist + e.w                               // 更新 next 的 dist 和 f
			nextVertex.f = nextVertex.dist + hManhattan(nextVertex, vertexes[t]) // 2.A* 算法在更新顶点 dist 值的时候，会同步更新 f 值
			predecessor[nextVertex.id] = minVertex.id
			heap.Push(queue, nextVertex)
			if nextVertex.id == t { // 3.A* 算法是一旦遍历到终点就结束
				//queue.Clear()
				break out
			}
		}
	}
	prt = func(s, t int, predecessor []int) {
		if s == t {
			return
		}
		prt(s, predecessor[t], predecessor) // 上一个顶点
		fmt.Print("->", t)
	}
	fmt.Print(s)
	prt(s, t, predecessor)
}

// hManhattan 曼哈顿距离（Manhattan distance）
func hManhattan(v, t *vertexAStar) int {
	dx, dy := v.x-t.x, v.y-t.y
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

type vertexAStar struct {
	id   int // 顶点编号 ID
	dist int // 从起始顶点 s 到这个顶点的距离
	f    int // 新增：f(i)=g(i)+h(i)
	x, y int // 新增：顶点在地图中的坐标
}
type priorityQueueAStar struct {
	pq []*vertexAStar
}

func (p priorityQueueAStar) Len() int           { return len(p.pq) }
func (p priorityQueueAStar) Less(i, j int) bool { return p.pq[i].f < p.pq[j].f } // 1.优先级队列构建的方式不同，根据 f 值
func (p priorityQueueAStar) Swap(i, j int)      { p.pq[i], p.pq[j] = p.pq[j], p.pq[i] }
func (p *priorityQueueAStar) Push(x any)        { p.pq = append(p.pq, x.(*vertexAStar)) }
func (p *priorityQueueAStar) Pop() any {
	v := p.pq[len(p.pq)-1]
	p.pq = p.pq[:len(p.pq)-1]
	return v
}
