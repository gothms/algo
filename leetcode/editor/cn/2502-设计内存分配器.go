package main

import "math/bits"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type Allocator struct {
}

func Constructor(n int) Allocator {

}

func (this *Allocator) Allocate(size int, mID int) int {

}

func (this *Allocator) FreeMemory(mID int) int {

}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.FreeMemory(mID);
 */
//leetcode submit region end(Prohibit modification and deletion)

// ==========线段树 lazy + 二分==========
type segTree []struct {
	l, r     int
	max      int // 最长连续空间
	pre, suf int // 各有几个前后缀的空间
	todo     int // -1 已释放的空间，0 标记将释放的空间，1 已使用的空间
}

func (s segTree) build(i, l, r int) {
	s[i].l, s[i].r = l, r
	s.do(i, -1) // 无lazy标记
	if l == r {
		return
	}
	m := int(uint(l+r) >> 1)
	s.build(i<<1, l, m)
	s.build(i<<1|1, m+1, r)
}
func (s segTree) update(i, f, t, v int) {
	if f <= s[i].l && s[i].r <= t {
		s.do(i, v) // down
		return
	}
	s.lazy(i) // lazy
	m := int(uint(s[i].l+s[i].r) >> 1)
	if f <= m {
		s.update(i<<1, f, t, v)
	}
	if m < t {
		s.update(i<<1|1, f, t, v)
	}

	ls, rs := s[i<<1], s[i<<1|1] // 合并左右子树的信息
	s[i].pre = ls.pre            // 更新 pre
	if ls.pre == m-s[i].l+1 {
		s[i].pre += rs.pre // 和右子树的 pre 拼接
	}
	s[i].suf = rs.suf // 更新 suf
	if rs.suf == s[i].r-m {
		s[i].suf += ls.suf
	}
	s[i].max = max(ls.max, rs.max, ls.suf+rs.pre) // 更新 max
}
func (s segTree) find(i, size int) int {
	if s[i].max < size {
		return -1
	}
	if s[i].l == s[i].r { // 叶子节点
		return s[i].l
	}
	s.lazy(i)                 // lazy
	idx := s.find(i<<1, size) // 1.递归左子树
	if idx < 0 {
		if s[i<<1].suf+s[i<<1|1].pre >= size {
			idx = s[i<<1].r - s[i<<1].suf + 1 // 2.拼接
			//idx = int(uint(s[i].l+s[i].r)>>1) - s[i<<1].suf + 1
		} else {
			idx = s.find(i<<1|1, size) // 3.递归右子树
		}
	}
	return idx
}
func (s segTree) lazy(i int) {
	if v := s[i].todo; v != -1 { // lazy down
		s.do(i<<1, v)
		s.do(i<<1|1, v)
		s[i].todo = -1 // 已释放
	}
}
func (s segTree) do(i, v int) {
	t := &s[i]
	size := 0
	if v <= 0 { // -1 和 0 都表示未使用的空间
		size = t.r - t.l + 1
	}
	t.max, t.pre, t.suf = size, size, size // 更新变量
	t.todo = v                             // lazy
}

type Allocator struct {
	seg    segTree
	blocks map[int][][2]int
}

func Constructor(n int) Allocator {
	t := make(segTree, func(n int) int {
		k := bits.Len(uint(n - 1))
		stLen := 1 << (k + 1)
		if n > 1 {
			stLen -= 1<<(k-bits.Len(uint(n-stLen>>2))+1) - 2
		}
		return stLen
	}(n))
	t.build(1, 0, n-1) // 构建线段树
	return Allocator{seg: t, blocks: make(map[int][][2]int)}
}

func (this *Allocator) Allocate(size int, mID int) int {
	seg := this.seg
	i := seg.find(1, size) // 查找 size 个连续空间
	if i == -1 {
		return i
	}
	seg.update(1, i, i+size-1, 1)                                        // 分配连续的空间
	this.blocks[mID] = append(this.blocks[mID], [2]int{i, i + size - 1}) // 记录 mID
	return i
}

func (this *Allocator) FreeMemory(mID int) int {
	ans := 0
	for _, b := range this.blocks[mID] {
		ans += b[1] - b[0] + 1            // 统计释放的空间
		this.seg.update(1, b[0], b[1], 0) // 释放内存
	}
	delete(this.blocks, mID) // 删除 mID 记录
	return ans
}

// ==========模拟==========

//type Allocator struct {
//	arr []int
//}
//
//func Constructor(n int) Allocator {
//	return Allocator{make([]int, n)}
//}
//
//func (this *Allocator) Allocate(size int, mID int) int {
//	arr := this.arr
//	cnt := 0
//	for i, v := range arr {
//		if v != 0 {
//			cnt = 0
//			continue
//		}
//		if cnt++; cnt == size {
//			j := i
//			for ; j >= i-size+1; j-- {
//				arr[j] = mID
//			}
//			return j + 1
//		}
//	}
//	return -1
//}
//
//func (this *Allocator) FreeMemory(mID int) int {
//	cnt := 0
//	for i, v := range this.arr {
//		if v == mID {
//			this.arr[i] = 0
//			cnt++
//		}
//	}
//	return cnt
//}
