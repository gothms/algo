package main

import (
	"fmt"
	"math/bits"
)

func main() {
	//n, m := 2, 5
	//con := Constructor(n, m)
	//gather := con.Gather(4, 0)
	//fmt.Println(gather)
	//gather = con.Gather(2, 0)
	//fmt.Println(gather)
	//scatter := con.Scatter(5, 1)
	//fmt.Println(scatter)
	//scatter = con.Scatter(5, 1)
	//fmt.Println(scatter)

	n, m := 0, 37
	con := Constructor(n, m)
	gather := con.Gather(28, 8)
	fmt.Println(gather)
	scatter := con.Scatter(10, 15)
	fmt.Println(scatter)
	gather = con.Gather(47, 16)
	fmt.Println(gather)
	scatter = con.Scatter(45, 8)
	fmt.Println(scatter)
	scatter = con.Scatter(34, 8)
	fmt.Println(scatter)
	scatter = con.Scatter(28, 7) // true
	fmt.Println(scatter)
}

// ====================树状数组？无法向下更新====================

// leetcode submit region begin(Prohibit modification and deletion)

type seg2286 []struct {
	l, r int // 排区间 [l,r]，从 1 开始
	v    int // 叶子：所在排还剩 v 个座位，非叶子：包含的排剩下座位的最大值为 v
	s    int // 叶子：剩 v 个座位，非叶子：总座位
}

func (seg seg2286) build(l, r, i, v int) {
	seg[i].l, seg[i].r = l, r
	if l == r {
		seg[i].v, seg[i].s = v, v
		return
	}
	m, idx := (l+r)>>1, i<<1
	seg.build(l, m, idx, v)
	seg.build(m+1, r, idx|1, v)
	seg[i].v = max(seg[idx].v, seg[idx|1].v) // pushUp
	seg[i].s = seg[idx].s + seg[idx|1].s
}

func (seg seg2286) gather(f, t, i, k int) (int, int) {
	if seg[i].v < k {
		return 0, 0
	}
	if seg[i].l == seg[i].r { // 排座的行
		seg[i].v -= k
		seg[i].s -= k
		return seg[i].l, seg[i].v
	}
	row, col, m, idx := 0, 0, (seg[i].l+seg[i].r)>>1, i<<1
	if f <= m && seg[idx].v >= k { // 左行足够
		row, col = seg.gather(f, t, idx, k)
	} else if t > m { // 左行不够
		row, col = seg.gather(f, t, idx|1, k)
		//} else {
		//	return 0, 0 // 均不够
	}
	seg[i].v = max(seg[idx].v, seg[idx|1].v)
	seg[i].s = seg[idx].s + seg[idx|1].s
	return row, col
}

func (seg seg2286) scatter(f, t, i, k int) bool {
	if f <= seg[i].l && seg[i].r <= t { // 排座的分割点
		return seg[i].s >= k
	}
	m, idx := (seg[i].l+seg[i].r)>>1, i<<1
	if t > m && seg[idx].s < k { // 左不够
		return seg.scatter(f, t, idx|1, k-seg[idx].s)
	} else {
		return seg.scatter(f, t, idx, k) // 左够
	}
}

func (seg seg2286) scatterUpdate(f, t, i, k int) int {
	if seg[i].l == seg[i].r { // 排座的分割点
		seg[i].v -= k
		seg[i].s -= k
		if seg[i].v == 0 { // 更新 this.i
			return seg[i].l + 1
		}
		return seg[i].l
	}
	m, idx := (seg[i].l+seg[i].r)>>1, i<<1
	row := 0
	if t > m && seg[idx].s < k { // 左不够
		row = seg.scatterUpdate(f, t, idx|1, k-seg[idx].s)
		//seg.scatterUpdate(f, t, idx, seg[idx].s) // seg[idx].s 的值已为 0，则应该先计算 idx|1
		seg[idx].v, seg[idx].s = 0, 0
	} else {
		row = seg.scatterUpdate(f, t, idx, k) // 左够
	}
	seg[i].v = max(seg[idx].v, seg[idx|1].v) // pushUp
	seg[i].s = seg[idx].s + seg[idx|1].s
	return row
}

type BookMyShow struct {
	seg2286
	m, i int // i 记录当前开始查询的行，仅在 Scatter 操作后更新
}

func Constructor(n int, m int) BookMyShow {
	t := make(seg2286, func(n int) int {
		k := bits.Len(uint(n - 1))
		stLen := 1 << (k + 1)
		if n > 1 {
			stLen -= 1<<(k-bits.Len(uint(n-stLen>>2))+1) - 2
		}
		return stLen
	}(n))
	t.build(1, n, 1, m)
	return BookMyShow{t, m, 1}
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	if this.i > maxRow+1 || k > this.m { // fast fail
		return nil
	}
	t := this.seg2286
	i, j := t.gather(this.i, maxRow+1, 1, k) // 查询并修改：i 行号，j 排座后还剩下座位数
	if i == 0 {                              // fail
		return nil
	}
	return []int{i - 1, this.m - k - j}
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	if this.i > maxRow+1 { // fast fail
		return false
	}
	t := this.seg2286
	if !t.scatter(this.i, maxRow+1, 1, k) { // 仅查询
		return false
	}
	this.i = t.scatterUpdate(this.i, maxRow+1, 1, k) // 仅修改：可使用懒惰标记
	return true
}

/**
 * Your BookMyShow object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Gather(k,maxRow);
 * param_2 := obj.Scatter(k,maxRow);
 */
//leetcode submit region end(Prohibit modification and deletion)

// ====================线段树，个人====================

// ====================线段树，lc====================

//type seg []struct{ l, r, min, sum int }
//
//func (t seg) build(o, l, r int) {
//	t[o].l, t[o].r = l, r
//	if l == r {
//		return
//	}
//	m := (l + r) >> 1
//	t.build(o<<1, l, m)
//	t.build(o<<1|1, m+1, r)
//}
//
//// 将 idx 上的元素值增加 val
//func (t seg) add(o, idx, val int) {
//	if t[o].l == t[o].r {
//		t[o].min += val
//		t[o].sum += val
//		return
//	}
//	m := (t[o].l + t[o].r) >> 1
//	if idx <= m {
//		t.add(o<<1, idx, val)
//	} else {
//		t.add(o<<1|1, idx, val)
//	}
//	lo, ro := t[o<<1], t[o<<1|1]
//	t[o].min = min(lo.min, ro.min)
//	t[o].sum = lo.sum + ro.sum
//}
//
//// 返回区间 [l,r] 内的元素和
//func (t seg) querySum(o, l, r int) (sum int) {
//	if l <= t[o].l && t[o].r <= r {
//		return t[o].sum
//	}
//	m := (t[o].l + t[o].r) >> 1
//	if l <= m {
//		sum += t.querySum(o<<1, l, r)
//	}
//	if r > m {
//		sum += t.querySum(o<<1|1, l, r)
//	}
//	return
//}
//
//// 返回区间 [1,R] 中 <= val 的最靠左的位置，不存在时返回 0
//func (t seg) index(o, r, val int) int {
//	if t[o].min > val { // 说明整个区间的元素值都大于 val
//		return 0
//	}
//	if t[o].l == t[o].r {
//		return t[o].l
//	}
//	m := (t[o].l + t[o].r) >> 1
//	if t[o<<1].min <= val { // 看看左半部分
//		return t.index(o<<1, r, val)
//	}
//	if m < r { // 看看右半部分
//		return t.index(o<<1|1, r, val)
//	}
//	return 0
//}
//
//type BookMyShow_ struct {
//	seg
//	m int
//}
//
//func Constructor_(n, m int) BookMyShow_ {
//	t := make(seg, n*4)
//	t.build(1, 1, n)
//	return BookMyShow_{t, m}
//}
//
//func (t BookMyShow_) Gather_(k, maxRow int) []int {
//	i := t.index(1, maxRow+1, t.m-k)
//	if i == 0 { // 不存在
//		return nil
//	}
//	seats := t.querySum(1, i, i)
//	t.add(1, i, k) // 占据 k 个座位
//	return []int{i - 1, seats}
//}
//
//func (t BookMyShow_) Scatter_(k, maxRow int) bool {
//	if (maxRow+1)*t.m-t.querySum(1, 1, maxRow+1) < k { // 剩余座位不足 k 个
//		return false
//	}
//	// 从第一个没有坐满的排开始占座
//	for i := t.index(1, maxRow+1, t.m-1); ; i++ {
//		leftSeats := t.m - t.querySum(1, i, i)
//		if k <= leftSeats { // 剩余人数不够坐后面的排
//			t.add(1, i, k)
//			return true
//		}
//		k -= leftSeats
//		t.add(1, i, leftSeats)
//	}
//}
