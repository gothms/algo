//给你一个二维整数数组 intervals ，其中 intervals[i] = [lefti, righti] 表示 闭 区间 [lefti,
//righti] 。
//
// 你需要将 intervals 划分为一个或者多个区间 组 ，每个区间 只 属于一个组，且同一个组中任意两个区间 不相交 。
//
// 请你返回 最少 需要划分成多少个组。
//
// 如果两个区间覆盖的范围有重叠（即至少有一个公共数字），那么我们称这两个区间是 相交 的。比方说区间 [1, 5] 和 [5, 8] 相交。
//
//
//
// 示例 1：
//
//
//输入：intervals = [[5,10],[6,8],[1,5],[2,3],[1,10]]
//输出：3
//解释：我们可以将区间划分为如下的区间组：
//- 第 1 组：[1, 5] ，[6, 8] 。
//- 第 2 组：[2, 3] ，[5, 10] 。
//- 第 3 组：[1, 10] 。
//可以证明无法将区间划分为少于 3 个组。
//
//
// 示例 2：
//
//
//输入：intervals = [[1,3],[5,6],[8,10],[11,13]]
//输出：1
//解释：所有区间互不相交，所以我们可以把它们全部放在一个组内。
//
//
//
// 提示：
//
//
// 1 <= intervals.length <= 10⁵
// intervals[i].length == 2
// 1 <= lefti <= righti <= 10⁶
//
//
// Related Topics 贪心 数组 双指针 前缀和 排序 堆（优先队列） 👍 52 👎 0

package main

import (
	"container/heap"
	"math"
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minGroups(intervals [][]int) int {
	// lc：堆
	h := mgHp{[]int{math.MaxInt32}}
	//h := mgHp{}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for _, p := range intervals {
		// 由于已经排序（左边界大于等于前面所有的左边界）所以 p[0] <= h.IntSlice[0] 说明有交集
		// 由于堆顶元素是最小的右边界，只要和它没交集，则可以成为独立的区间
		if p[0] <= h.IntSlice[0] { // 有交集
			heap.Push(&h, p[1])
		} else { // 成为独立区间
			h.IntSlice[0] = p[1]
			heap.Fix(&h, 0)
		}
	}
	return h.Len() - 1

	// 排序
	//maxVal := func(a, b int) int {
	//	if b > a {
	//		return b
	//	}
	//	return a
	//}
	//memo := make(map[int]int)
	//for _, iv := range intervals {
	//	memo[iv[0]]++
	//	memo[iv[1]+1]--
	//}
	//mg := make([][2]int, 0, len(memo))
	//for k, v := range memo {
	//	mg = append(mg, [2]int{k, v})
	//}
	//sort.Slice(mg, func(i, j int) bool {
	//	return mg[i][0] < mg[j][0]
	//})
	//ret, sum := 0, 0
	//for _, v := range mg {
	//	sum += v[1]
	//	ret = maxVal(ret, sum)
	//}
	//return ret
}

type mgHp struct{ sort.IntSlice }

func (h *mgHp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *mgHp) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
