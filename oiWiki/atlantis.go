package wiki

import (
	"sort"
)

/*
扫描线
	https://oi-wiki.org/geometry/scanning/
lc
	218
	391
	850
	1851
	2589
*/

// Atlantis lc 850
func Atlantis(rectangles [][]int) int {
	const mod = 1_000_000_007
	var ans int
	n := len(rectangles) << 1
	var (
		xs = make([]rect, 0, n)
		ys = make([]int, 0, n)
	)
	for _, r := range rectangles {
		x1, y1, x2, y2 := r[0], r[1], r[2], r[3]
		xs = append(xs, rect{x1, y1, y2, 1}, rect{x2, y1, y2, -1})
		ys = append(ys, y1, y2)
	}
	sort.Slice(xs, func(i, j int) bool { // 把矩形的边的横坐标从小到大排序
		return xs[i].x < xs[j].x
	})
	sort.Ints(ys) // 离散化
	m := 0
	for _, y := range ys[1:] { // 去重
		if ys[m] != y {
			m++
			ys[m] = y
		}
	}
	ys = ys[:m+1]
	var (
		stLen = m << 3 // stLen 需要足够大？ TODO
		st    = make([]atSeg, stLen)
		lazy  = make([]int, stLen) // 标记线段树每个节点
	)
	build(0, m, ys, 1, st)                              // 建树
	update(xs[0].y1, xs[0].y2, xs[0].flag, 1, st, lazy) // 更新最左边的 x 坐标
	for j := 1; j < n; j++ {
		//if xs[j].x != xs[j-1].x {	// x 坐标相等，则不用计算
		ans += (xs[j].x - xs[j-1].x) * st[1].sum % mod // st[1] 为当前 xs[j] 到 xs[j-1] 区间，矩形的总高度
		//}
		update(xs[j].y1, xs[j].y2, xs[j].flag, 1, st, lazy) // 更新第 j 个 x 坐标的纵坐标 y 区间 [y1,y2]
	}
	return ans % mod
}

type rect struct {
	x, y1, y2 int
	flag      int
}
type atSeg struct {
	l, r int
	sum  int
}

func build(l, r int, ys []int, i int, st []atSeg) {
	st[i].l, st[i].r = ys[l], ys[r]
	if r-l <= 1 { // l==r 说明只是一条线，而不是区间
		st[i].sum = 0
		return
	}
	m := (l + r) >> 1
	build(l, m, ys, i<<1, st)
	build(m, r, ys, i<<1+1, st) // 是 m 而非 m+1
}
func update(y1, y2 int, flag int, i int, st []atSeg, lazy []int) {
	if st[i].l == y1 && st[i].r == y2 { // 精确区间
		lazy[i] += flag
		pushUp(i, st, lazy)
		return
	}
	idx := i << 1
	if y1 < st[idx].r { // 等同于 update(l, m)，m = (l+r)>>1
		update(y1, min(st[idx].r, y2), flag, idx, st, lazy)
	}
	if y2 > st[i<<1+1].l {
		update(max(st[idx+1].l, y1), y2, flag, idx+1, st, lazy) // 易错：max
	}
	pushUp(i, st, lazy)
}
func pushUp(i int, st []atSeg, lazy []int) {
	if lazy[i] > 0 { // 加进来了矩形的长
		st[i].sum = st[i].r - st[i].l
	} else {
		st[i].sum = st[i<<1].sum + st[i<<1+1].sum
	}
}
