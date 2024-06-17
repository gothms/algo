package tree

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"math/bits"
)

/*
https://pkg.go.dev/github.com/emirpasic/gods/trees/redblacktree
	go get github.com/emirpasic/gods

lc
	2034
	3161
	2612
*/

// getResults lc 3161
func getResults(queries [][]int) []bool {
	// 有序集合 + 线段树
	var (
		rbt    *redblacktree.Tree
		st     []int
		update func(int, int, int, int, int)
		query  func(int, int, int, int) int
	)
	update = func(x, l, r, i, val int) {
		if l == r {
			st[i] = val
			return
		}
		m, idx := (l+r)>>1, i<<1
		if x <= m {
			update(x, l, m, idx, val)
		} else {
			update(x, m+1, r, idx+1, val)
		}
		st[i] = max(st[idx], st[idx+1])
	}
	query = func(x, l, r, i int) int {
		if x == r { // 递归终止条件，也可以写 x>=r
			return st[i]
		}
		m, idx := (l+r)>>1, i<<1
		if x <= m {
			return query(x, l, m, idx) // 找左边
		} else {
			return max(st[idx], query(x, m+1, r, idx+1)) // 左边最大长度 / 找右边
		}
	}

	ans := make([]bool, 0)
	maxX := 0
	for _, q := range queries {
		maxX = max(maxX, q[1]) // 右边界
	}
	rbt = redblacktree.NewWithIntComparator()
	st = make([]int, func(n int) int {
		k := bits.Len(uint(n - 1))
		stLen := 1 << (k + 1)
		if n > 1 {
			stLen -= 1<<(k-bits.Len(uint(n-stLen>>2))+1) - 2
		}
		return stLen
	}(maxX+1)) // [0,maxX]
	rbt.Put(0, struct{}{}) // 初始化数据
	rbt.Put(maxX, struct{}{})
	update(maxX, 0, maxX, 1, maxX)

	for _, q := range queries {
		x := q[1]
		left, _ := rbt.Floor(x) // 前驱节点
		lv := left.Key.(int)
		switch q[0] {
		case 1:
			if lv != x { // 不重复分隔
				right, _ := rbt.Ceiling(x) // 后驱节点
				rv := right.Key.(int)
				update(x, 0, maxX, 1, x-lv) // 更新线段树
				update(rv, 0, maxX, 1, rv-x)
				rbt.Put(x, struct{}{}) // 插入分割点
			}
		case 2:
			ans = append(ans, x-lv >= q[2] ||
				query(lv, 0, maxX, 1) >= q[2]) // 找左边
		}
	}
	return ans
}
