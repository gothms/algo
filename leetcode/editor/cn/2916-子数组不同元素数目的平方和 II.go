package main

import (
	"fmt"
	"math/bits"
)

func main() {
	nums := []int{1, 2, 1}
	//nums = []int{2, 2}
	//nums = []int{1}
	counts := sumCounts(nums)
	fmt.Println(counts)

	//for i, v := range nums {
	//	i++
	//	fmt.Println(i, v)
	//}
}

// leetcode submit region begin(Prohibit modification and deletion)
func sumCounts(nums []int) int {
	// 参考：lc-2262
	// 重点：增量的计算
	// 比如位置 x+1 对 x 的增量：(x+1)^2 - x^2 = 2*x + 1
	// 所以 last = j 时，增量为：如 [1 2 2 3 1]，j=0，i=4
	// (v1 + v2 + v3) * 2 + i-j = (2+2+1)*2 + 4-0
	// 则采用 线段树+懒加载 实现增量的计算

	const mod = 1_000_000_007
	ret, s, n := 0, 0, len(nums)
	k := bits.Len(uint(n - 1))
	stLen := 1 << (k + 1)
	if n > 1 {
		stLen -= 1<<(k-bits.Len(uint(n-stLen>>2))+1) - 2
	}
	st := make([][2]int, stLen) // 懒惰标记：和、增量，索引从 1 开始
	last := make(map[int]int)
	down := func(l, r, i, d int) { // lazy 更新
		st[i][0] += d * (r - l + 1) // 更新区间和
		st[i][1] += d               // 更新增量
	}
	var rangeSumSegmentTree func(f, t, l, r, i, d int) int
	rangeSumSegmentTree = func(f, t, l, r, i, d int) int { // d 为增量
		var cur int
		if f <= l && r <= t {
			cur = st[i][0]   // 错误点：增量 d 是更新 线段树 后，用于“下一次”的计算
			down(l, r, i, d) // 先保存 st[i][0]，再 down
			return cur
		}
		m, idx := (l+r)>>1, i<<1
		if ad := st[i][1]; ad != 0 { // 更新懒惰标记
			down(l, m, idx, ad)
			down(m+1, r, idx+1, ad)
			st[i][1] = 0 // 清除懒惰标记
		}
		if f <= m {
			cur = rangeSumSegmentTree(f, t, l, m, idx, 1)
		}
		if t > m {
			cur += rangeSumSegmentTree(f, t, m+1, r, idx+1, 1)
		}
		st[i][0] = st[idx][0] + st[idx+1][0] // 区间和
		return cur
	}
	for i, v := range nums {
		l, t := last[v], i+1                                    // 更新区间 [l+1,t]
		s += rangeSumSegmentTree(l+1, t, 1, n, 1, 1)<<1 + t - l // 重点：懒惰标记区间 [1,n]，4个索引都从 0 / 1 开始均可
		//s += rangeSumSegmentTree(l, t-1, 0, n-1, 1, 1)<<1 + t - l
		ret = (ret + s) % mod
		last[v] = t // 默认索引由 0 改为从 1 开始，且 线段树 也是从 1 开始
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
