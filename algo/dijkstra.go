package algo

import (
	"container/heap"
	"fmt"
	"math"
)

/*
lc-2699
*/

// Dijkstra 从顶点s到顶点t的最短路径
// n 为顶点的个数，edges 为有向有权图
func Dijkstra(n, s, t int, edges [][]int) {
	var (
		adj         = make([][]edge, n)                 // 邻接表
		vertexes    = make([]*vertex, n)                // 顶点和 s 与顶点的距离
		visited     = make([]bool, n)                   // 记录已访问过的顶点（已进入过队列）
		predecessor = make([]int, 0)                    // 用于还原最短路径
		queue       = &priorityQueue{make([]vertex, 0)} // 小顶堆：根据 vertex.dist
		print       func(int, int, []int)               // 用于打印最 短路径
	)
	for _, e := range edges { // 遍历有向有权图
		adj[e[0]] = append(adj[e[0]], edge{e[0], e[1], e[2]}) // 添加一条边
		//adj[e[1]] = append(adj[e[1]], edge{e[0], e[1], e[2]}) // 无向有权边
	}
	for i := 0; i < n; i++ {
		vertexes[i] = &vertex{i, math.MaxInt32}
	}
	heap.Push(queue, vertexes[s])
	for queue.Len() > 0 { // 从顶点
		minVertex := heap.Pop(queue).(vertex) // 取出堆顶元素并删除
		// 两种方式：1.入队列时，记录已访问。2.出队列后，记录已访问
		// 第一种方式（参考专栏）需要 heap.Fix(queue, i)，但 golang 中往往无法获取到 i，所以采用第二种方式
		if visited[minVertex.id] { // 先出队列的顶点，dist 总是更小
			continue
		}
		visited[minVertex.id] = true
		if minVertex.id == t { // 找到最短路径
			break
		}
		for _, e := range adj[minVertex.id] { // 遍历顶点 minVertex.id 的所有边
			nextVertex := vertexes[e.tid]              // 下一个顶点
			if minVertex.dist+e.w >= nextVertex.dist { // 边 minVertex -> nextVertex
				continue // 否则，更新 next 的 dist
			}
			nextVertex.dist = minVertex.dist + e.w // vertexes 是指针数组，所以修改的是内存值
			predecessor[nextVertex.id] = minVertex.id
			heap.Push(queue, nextVertex)
		}
	}
	print = func(s, t int, predecessor []int) {
		if s == t {
			return
		}
		print(s, predecessor[t], predecessor)
		fmt.Print("->", t)
	}
	fmt.Print(s)
	print(s, t, predecessor) // 输出最短路径
}

type edge struct {
	sid, tid int // 边 sid->tid
	w        int // 权重
}
type vertex struct {
	id   int // 顶点编号 ID
	dist int // 从起始顶点 s 到这个顶点的距离
}
type priorityQueue struct {
	qp []vertex
}

func (p priorityQueue) Len() int           { return len(p.qp) }
func (p priorityQueue) Less(i, j int) bool { return p.qp[i].dist < p.qp[j].dist }
func (p priorityQueue) Swap(i, j int)      { p.qp[i], p.qp[j] = p.qp[j], p.qp[i] }
func (p *priorityQueue) Push(x any)        { p.qp = append(p.qp, x.(vertex)) }
func (p *priorityQueue) Pop() any {
	v := p.qp[len(p.qp)-1]
	p.qp = p.qp[:len(p.qp)-1]
	return v
}
