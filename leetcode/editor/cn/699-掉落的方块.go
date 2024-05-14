//在二维平面上的 x 轴上，放置着一些方块。
//
// 给你一个二维整数数组 positions ，其中 positions[i] = [lefti, sideLengthi] 表示：第 i 个方块边长为
//sideLengthi ，其左侧边与 x 轴上坐标点 lefti 对齐。
//
// 每个方块都从一个比目前所有的落地方块更高的高度掉落而下。方块沿 y 轴负方向下落，直到着陆到 另一个正方形的顶边 或者是 x 轴上 。一个方块仅仅是擦过另
//一个方块的左侧边或右侧边不算着陆。一旦着陆，它就会固定在原地，无法移动。
//
// 在每个方块掉落后，你必须记录目前所有已经落稳的 方块堆叠的最高高度 。
//
// 返回一个整数数组 ans ，其中 ans[i] 表示在第 i 块方块掉落后堆叠的最高高度。
//
//
//

// 示例 1：
//
//
//输入：positions = [[1,2],[2,3],[6,1]]
//输出：[2,5,5]
//解释：
//第 1 个方块掉落后，最高的堆叠由方块 1 组成，堆叠的最高高度为 2 。
//第 2 个方块掉落后，最高的堆叠由方块 1 和 2 组成，堆叠的最高高度为 5 。
//第 3 个方块掉落后，最高的堆叠仍然由方块 1 和 2 组成，堆叠的最高高度为 5 。
//因此，返回 [2, 5, 5] 作为答案。
//
//
// 示例 2：
//
//
//输入：positions = [[100,100],[200,100]]
//输出：[100,100]
//解释：
//第 1 个方块掉落后，最高的堆叠由方块 1 组成，堆叠的最高高度为 100 。
//第 2 个方块掉落后，最高的堆叠可以由方块 1 组成也可以由方块 2 组成，堆叠的最高高度为 100 。
//因此，返回 [100, 100] 作为答案。
//注意，方块 2 擦过方块 1 的右侧边，但不会算作在方块 1 上着陆。
//
//
//
//
// 提示：
//
//
// 1 <= positions.length <= 1000
// 1 <= lefti <= 10⁸
// 1 <= sideLengthi <= 10⁶
//
//
// Related Topics 线段树 数组 有序集合 👍 188 👎 0

package main

import (
	"fmt"
)

func main() {
	positions := [][]int{{1, 2}, {2, 3}, {6, 1}}
	positions = [][]int{{9, 7}, {1, 9}, {3, 1}}                 // [7,16,17]
	positions = [][]int{{100, 100}, {200, 100}}                 // [100,100]
	positions = [][]int{{7, 2}, {1, 7}, {9, 5}, {1, 8}, {3, 1}} // [2,9,9,17,18]
	squares := fallingSquares(positions)
	fmt.Println(squares)
}

// leetcode submit region begin(Prohibit modification and deletion)
func fallingSquares(positions [][]int) []int {
	// 线段树：动态开点
	type stNode struct {
		left, right *stNode
		l, r        int
		h           int
		d           int
	}
	n := len(positions)
	ans := make([]int, n+1)
	root := &stNode{l: 1, r: 1e8 + 1e6 - 1} // 可先遍历找出 r 最大值
	//var maxX int
	//for _, p := range positions {
	//	maxX = max(maxX, p[0]+p[1])
	//}
	//root := &stNode{l: 1, r: maxX - 1} // 可先遍历找出 r 最大值
	down := func(m int, cur *stNode) {
		if cur.left == nil {
			cur.left = &stNode{l: cur.l, r: m, h: cur.h}
		}
		if cur.right == nil {
			cur.right = &stNode{l: m + 1, r: cur.r, h: cur.h}
		}
		if cur.d > 0 { // 重要：如果 left / right 已存在，也要更新懒惰标记
			cur.left.d, cur.right.d = cur.d, cur.d
			cur.h, cur.d = cur.d, 0
		}
	}
	var update func(int, int, int, *stNode)
	update = func(f, t, h int, cur *stNode) {
		if f <= cur.l && cur.r <= t {
			cur.h = h
			cur.d = h
			return
		}
		m := (cur.l + cur.r) >> 1
		down(m, cur)
		if f <= m {
			update(f, t, h, cur.left)
		}
		if t > m {
			update(f, t, h, cur.right)
		}
		cur.h = max(cur.left.h, cur.right.h)
	}
	var query func(int, int, *stNode) int
	query = func(f, t int, cur *stNode) int {
		if f <= cur.l && cur.r <= t {
			//down((cur.l+cur.r)>>1, cur) // 可能 cur.d 没被更新到 cur.h
			if cur.d > 0 {
				cur.h, cur.d = cur.d, 0
			}
			return cur.h
		}
		h, m := 0, (cur.l+cur.r)>>1
		down(m, cur)
		if f <= m {
			h = query(f, t, cur.left)
		}
		if t > m {
			h = max(h, query(f, t, cur.right))
		}
		return h
	}
	for i, p := range positions {
		h := query(p[0], p[0]+p[1]-1, root) + p[1] // 最高高度 = 查出的高度 + p[1]
		//ans[i+1] = max(ans[i], h)
		update(p[0], p[0]+p[1]-1, h, root) // 将最高高度覆盖的区间全部更新
		ans[i+1] = root.h                  // root 即为最高高度
	}
	return ans[1:]

	// 有序集合：treemap

	// 暴力枚举
	//n := len(positions)
	//ret := make([]int, n)
	//for i, p := range positions {
	//	l, r := p[0], p[0]+p[1]-1
	//	ret[i] = p[1]
	//	for j, q := range positions[:i] {
	//		lp, rp := q[0], q[0]+q[1]-1
	//		if r >= lp && rp >= l { // 重点：有交集，4个点构成 “X” 型
	//			ret[i] = max(ret[i], ret[j]+p[1]) // 在前者的基础上堆叠
	//		}
	//	}
	//}
	//for i := 1; i < n; i++ {
	//	ret[i] = max(ret[i], ret[i-1])
	//}
	//return ret

	// 线段树：out of memory，stLen = 268435450
	//last, n := 0, len(positions)
	//for _, p := range positions {
	//	last = max(last, p[0]+p[1])
	//}
	//k := bits.Len(uint(last))
	//stLen := 1 << (k + 1)
	//if last > 0 {
	//	stLen -= 1<<(k-bits.Len(uint(last+1-stLen>>2))+1) - 2
	//}
	//
	//ret, st := make([]int, n), make([][2]int, stLen)
	//down := func(i, v int) {
	//	st[i][0], st[i][1] = v, v
	//}
	//var query func(f, t, l, r, i int) int
	//query = func(f, t, l, r, i int) int {
	//	if f <= l && r <= t {
	//		return st[i][0]
	//	}
	//	h, m, idx := 0, (l+r)>>1, i<<1
	//	if ad := st[i][1]; ad > 0 { // pushDown：查询时需要更新“懒惰标记”
	//		down(idx, ad)
	//		down(idx+1, ad)
	//		st[i][1] = 0
	//	}
	//	if f <= m {
	//		h = query(f, t, l, m, idx)
	//	}
	//	if t > m {
	//		h = max(h, query(f, t, m+1, r, idx+1))
	//	}
	//	//st[i][0] = max(st[idx][0], st[idx+1][0])
	//	return h
	//}
	//var update func(f, t, l, r, i, v int)
	//update = func(f, t, l, r, i, v int) {
	//	if f <= l && r <= t {
	//		down(i, v) // 区间 [f, t] 内，高度不一，所以将老数据放弃，统一更新为 v
	//		return
	//	}
	//	m, idx := (l+r)>>1, i<<1
	//	//if ad := st[i][1]; ad > 0 {
	//	//	down(idx, ad)
	//	//	down(idx+1, ad)
	//	//	st[i][1] = 0
	//	//}
	//	if f <= m {
	//		update(f, t, l, m, idx, v)
	//	}
	//	if t > m {
	//		update(f, t, m+1, r, idx+1, v)
	//	}
	//	st[i][0] = max(st[idx][0], st[idx+1][0]) // pushUp
	//}
	//hMax := 0
	//for i, p := range positions {
	//	h := query(p[0], p[0]+p[1]-1, 0, last, 1) + p[1]
	//	hMax = max(hMax, h)
	//	update(p[0], p[0]+p[1]-1, 0, last, 1, h) // “全”更新
	//	ret[i] = hMax
	//}
	//return ret
}

//leetcode submit region end(Prohibit modification and deletion)
