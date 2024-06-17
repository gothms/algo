package main

/*
lc
	1851：一种并查集标记切片的遍历方式
	1697：“高级”用法
	839
	1627
	1559
	2612：集合数组
*/

// TODO 1697

var pa, size []int // 模拟数据

func find(x int) int {
	if pa[x] != x {
		pa[x] = find(pa[x]) // 路径压缩
	}
	return pa[x]
}
func unite(x, y int) {
	//pa[find(y)] = find(x)

	x, y = find(x), find(y) // 启发式合并
	if x == y {
		return
	}
	if size[x] < size[y] {
		x, y = y, x
	}
	pa[y] = x
	size[x] += size[y]
}

// erase 删除一个叶子节点，可以将其父亲设为自己
// 为了保证要删除的元素都是叶子，可以预先为每个节点制作副本，并将其副本作为父亲
func erase(x int) {
	size[find(x)]--
	pa[x] = x
}

// move 与删除类似，通过以副本作为父亲，保证要移动的元素都是叶子
func move(x, y int) {
	fx, fy := find(x), find(y)
	if fx == fy {
		return
	}
	pa[y] = fx
	size[fx]++
	size[fy]--
}
