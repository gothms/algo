package main

import "fmt"

/*
未解决问题：

	LeetCode-1659
*/
func main() {
	vertexes := make([]vertex, 2)
	vertexes[0] = vertex{1, 2}
	next := vertexes[0]
	next.dist = 99
	fmt.Println(vertexes) // {1 2}

	pointVertexes := make([]*vertex, 2)
	pointVertexes[0] = &vertex{1, 2}
	pointVertexes[1] = &vertex{3, 4}
	pointNext := pointVertexes[0]
	pointNext.dist = 99
	for _, p := range pointVertexes {
		fmt.Println(*p) // {1 99}
	}
}

type vertex struct {
	id   int // 顶点编号 ID
	dist int // 从起始顶点 s 到这个顶点的距离
}
