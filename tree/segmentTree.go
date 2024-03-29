package tree

/*
Segment Tree 长度计算：1 << (bits.Len(uint(n-1)) + 1)
	n=8：16
	n=[9,16]：32
	stLen := 1 << (bits.Len(uint(n-1)) + 1)

lc
	2493
	1157
*/

// ====================构建线段树，从 0/1 开始====================
// m, idx := int(uint(l+r)>>1), i<<1+1	// 0
// m, idx := int(uint(l+r)>>1), i<<1	// 1

// Build 构建线段树
// l r 为 arr 的当前区间，i 为 st 当前索引
// 不需要扩容，所以直接传入 st
// 从索引 1 开始，效率更高
func Build(l, r int, arr []int, i int, st []int) {
	if l == r {
		st[i] = arr[l]
		return
	}
	//m, idx := int(uint(l+r)>>1), i<<1+1 // 从 0 开始
	m, idx := int(uint(l+r)>>1), i<<1 // 从 1 开始
	Build(l, m, arr, idx, st)
	Build(m+1, r, arr, idx+1, st)
	st[i] = st[idx] + st[idx+1]
}

// STRangeSum 查询区间和
// f t 为要查询的区间和，l r 为 arr 的当前区间，i 为线段树当前索引（起始为 0）
func STRangeSum(f, t int, l, r int, i int, st []int) int {
	if f <= l && r <= t {
		return st[i] // 查询区间落在 [l,r] 范围内
	}
	//sum, m, idx := 0, int(uint(l+r)>>1), i<<1+1 // 从 0 开始
	sum, m, idx := 0, int(uint(l+r)>>1), i<<1 // 从 1 开始
	if f <= m {                               // 和左节点代表的区间 [l,m] 有交集
		sum = STRangeSum(f, t, l, m, idx, st)
	}
	if t > m { // 和右节点代表的区间 [m+1,r] 有交集
		sum += STRangeSum(f, t, m+1, r, idx+1, st)
	}
	return sum
}

// STUpdate 更新原数组某个元素
// t 为要更新的索引（v 为被修改的元素的变化量），l r 为 arr 的当前区间，i 为线段树当前索引（起始为 0）
// 如果需求是“更新某个元素”，而不是“更新某个区间”，应该使用树状数组：binaryIndexedTree.go
func STUpdate(t, v int, l, r int, i int, st []int) {
	if l == r {
		st[i] += v
		return
	}
	//m, idx := int(uint(l+r)>>1), i<<1+1 // 从 0 开始
	m, idx := int(uint(l+r)>>1), i<<1 // 从 1 开始
	if i <= m {
		STUpdate(t, v, l, m, idx, st)
	} else {
		STUpdate(t, v, m+1, r, idx+1, st)
	}
	st[i] = st[idx] + st[idx+1]
}

// ====================线段树，懒惰标记====================

// STRangeSumLazy 查询区间和 & 懒惰标记
// f t 为要查询的区间和，l r 为 arr 的当前区间，i 为线段树当前索引（起始为 0）
// lazy 为懒惰标记数组
func STRangeSumLazy(f, t int, l, r int, i int, st, lazy []int) int {
	if f <= l && r <= t { // 查询区间落在 [l,r] 范围内
		return st[i]
	}
	//sum, m, idx := 0, int(uint(l+r)>>1), i<<1+1 // 从 0 开始
	sum, m, idx := 0, int(uint(l+r)>>1), i<<1 // 从 1 开始
	if lazy[i] != 0 {                         // 更新懒惰标记
		lazyDown(l, m, r, i, idx, st, lazy)
	}
	if f <= m { // 和左节点代表的区间 [l,m] 有交集
		sum = STRangeSumLazy(f, t, l, m, idx, st, lazy)
	}
	if t > m { // 和右节点代表的区间 [m+1,r] 有交集
		sum += STRangeSumLazy(f, t, m+1, r, idx+1, st, lazy)
	}
	return sum
}

// STUpdateLazy 更新原数组某个区间 & 懒惰标记
// f t 为要更新的区间（v 为被修改的元素的变化量），l r 为 arr 的当前区间，i 为线段树当前索引（起始为 0）
// lazy 为懒惰标记数组
func STUpdateLazy(f, t, v int, l, r int, i int, st, lazy []int) {
	if f <= l && r <= t {
		st[i] += (r - l + 1) * v
		lazy[i] += v
		return
	}
	//m, idx := int(uint(l+r)>>1), i<<1+1 // 从 0 开始
	m, idx := int(uint(l+r)>>1), i<<1 // 从 1 开始
	if lazy[i] != 0 && l != r {       // 更新懒惰标记，lazy[i] 有可能 <0
		lazyDown(l, m, r, i, idx, st, lazy)
	}
	if f <= m {
		STUpdateLazy(f, t, v, l, m, idx, st, lazy)
	}
	if t > m {
		STUpdateLazy(f, t, v, m+1, r, idx+1, st, lazy)
	}
	st[i] = st[idx] + st[idx+1]
}

// lazyDown 懒惰标记
func lazyDown(l, m, r int, i, idx int, st, lazy []int) { // 封装方法，下推一层
	st[idx] += lazy[i] * (m - l + 1)
	st[idx+1] += lazy[i] * (r - m)
	lazy[idx] += lazy[i]
	lazy[idx+1] += lazy[i]
	lazy[i] = 0
}
