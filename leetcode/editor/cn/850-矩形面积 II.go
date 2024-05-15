package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{0, 0, 2, 2},
		{1, 0, 2, 3},
		{1, 0, 3, 1}} // 6
	arr = [][]int{{0, 0, 1000000000, 1000000000}} // 49
	arr = [][]int{{0, 0, 2, 2},
		{1, 1, 3, 3}} // 7
	arr = [][]int{{0, 0, 1, 1},
		{2, 2, 3, 3}} // 2
	arr = [][]int{{93516, 44895, 94753, 69358},
		{13141, 52454, 59740, 71232},
		{22877, 11159, 85255, 61703},
		{11917, 8218, 84490, 36637},
		{75914, 29447, 83941, 64384},
		{22490, 71433, 64258, 74059},
		{18433, 51177, 87595, 98688},
		{70854, 80720, 91838, 92304},
		{46522, 49839, 48550, 94096},
		{95435, 37993, 99139, 49382},
		{10618, 696, 33239, 45957},
		{18854, 2818, 57522, 78807},
		{61229, 36593, 76550, 41271},
		{99381, 90692, 99820, 95125}} // 971243962
	area := rectangleArea(arr)
	fmt.Println(area)
}

// leetcode submit region begin(Prohibit modification and deletion)
func rectangleArea(rectangles [][]int) int {
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
		//stLen = func(n int) int {
		//	k := bits.Len(uint(n - 1))
		//	stLen := 1 << (k + 1)
		//	if n > 1 {
		//		stLen -= 1<<(k-bits.Len(uint(n-stLen>>2))+1) - 2
		//	}
		//	return stLen
		//}(n)
		stLen = m << 3 // stLen 需要足够大？
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

// 离散化 + 扫描线 + 线段树
//ans := 0
//n := len(rectangles) * 2
//hBound := make([]int, 0, n)
//for _, r := range rectangles {
//	hBound = append(hBound, r[1], r[3])
//}
//// 排序，方便下面去重
//sort.Ints(hBound)
//m := 0
//for _, b := range hBound[1:] {
//	if hBound[m] != b {
//		m++
//		hBound[m] = b
//	}
//}
//hBound = hBound[:m+1]
//t := make(seg, m*4)
//fmt.Println(hBound, m)
//t.init(hBound, 1, 1, m)
//
//type tuple struct{ x, i, d int }
//sweep := make([]tuple, 0, n)
//for i, r := range rectangles {
//	sweep = append(sweep, tuple{r[0], i, 1}, tuple{r[2], i, -1})
//}
//sort.Slice(sweep, func(i, j int) bool { return sweep[i].x < sweep[j].x })
//
//for i := 0; i < n; i++ {
//	j := i
//	for j+1 < n && sweep[j+1].x == sweep[i].x {
//		j++
//	}
//	if j+1 == n {
//		break
//	}
//	// 一次性地处理掉一批横坐标相同的左右边界
//	for k := i; k <= j; k++ {
//		idx, diff := sweep[k].i, sweep[k].d
//		// 使用二分查找得到完整覆盖的线段的编号范围
//		left := sort.SearchInts(hBound, rectangles[idx][1]) + 1
//		right := sort.SearchInts(hBound, rectangles[idx][3])
//		t.update(1, 1, m, left, right, diff)
//	}
//	ans += t[1].len * (sweep[j+1].x - sweep[j].x)
//	i = j
//}
//return ans % (1e9 + 7)

// 离散化 + 扫描线 + 数组

// 排序

// 暴力：求出总面积，再减去每两个矩形相交的面积

//type seg []struct{ cover, len, maxLen int }
//
//func (t seg) init(hBound []int, idx, l, r int) {
//	if l == r {
//		t[idx].maxLen = hBound[l] - hBound[l-1]
//		return
//	}
//	mid := (l + r) / 2
//	t.init(hBound, idx*2, l, mid)
//	t.init(hBound, idx*2+1, mid+1, r)
//	t[idx].maxLen = t[idx*2].maxLen + t[idx*2+1].maxLen
//}
//
//func (t seg) update(idx, l, r, ul, ur, diff int) {
//	if l > ur || r < ul {
//		return
//	}
//	if ul <= l && r <= ur {
//		t[idx].cover += diff
//		t.pushUp(idx, l, r)
//		return
//	}
//	mid := (l + r) / 2
//	t.update(idx*2, l, mid, ul, ur, diff)
//	t.update(idx*2+1, mid+1, r, ul, ur, diff)
//	t.pushUp(idx, l, r)
//}
//
//func (t seg) pushUp(idx, l, r int) {
//	if t[idx].cover > 0 {
//		t[idx].len = t[idx].maxLen
//	} else if l == r {
//		t[idx].len = 0
//	} else {
//		t[idx].len = t[idx*2].len + t[idx*2+1].len
//	}
//}

//leetcode submit region end(Prohibit modification and deletion)
