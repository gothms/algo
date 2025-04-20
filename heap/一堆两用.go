package heap

import (
	"container/heap"
	"sort"
)

/*
一堆双用
	方法一：元素取反
	方法二：直接修改逻辑
		自定义结构体，调整堆化函数，如自定义 up / down 的 Less 函数逻辑
	方法三：动态比较函数
		扩展结构体支持动态比较

lc
	1801
*/

// 可以随时更改堆的类型
func heapTest(isMax bool) *hp {
	cmp := func(x, y int) bool {
		return x < y
	}
	h := &hp{cmp: cmp} // 默认为小顶堆
	if isMax {         // 修改为大顶堆
		cmp = func(x, y int) bool {
			return x > y
		}
	}
	return h
}

type hp struct {
	sort.IntSlice
	cmp func(int, int) bool // 抽象
	//less func(int, int) bool
}

func (h hp) Less(i, j int) bool { return h.cmp(h.IntSlice[i], h.IntSlice[j]) }

// func (h hp) Less(i, j int) bool { return h.less(i, j) }

func (h *hp) Push(x any) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hp) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

var _ heap.Interface = (*hp)(nil)
