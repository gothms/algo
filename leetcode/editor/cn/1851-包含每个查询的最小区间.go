package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}
	queries := []int{2, 3, 4, 5} // [3,3,1,4]
	interval := minInterval(intervals, queries)
	fmt.Println(interval)
}

/*
思路：排序+二分+并查集
	1.设 l,r = intervals[i][0],intervals[i][1]，区间长度 cl = r-l+1
		满足 l<=queries[j]<=r 的长度最小区间i的长度，肯定是按cl升序排序后，最先满足的区间
		所以我们可以对 intervals 按 r-l 进行升序排序
	2.对于任意区间 l,r，落在区间内的 queries[j] 肯定大于等于 l
		所以我们可以对 queries 升序排序
		又由于最后记录 queries[j] 的最小区间长度 cl 时，需要元素的索引 j
		所以，我们需要对 queries 排序时，映射它的索引 j，记录为 vIdx(queries[j], j)
	3.通过二分查找 vIdx 第一个大于等于 l 的索引 k，验证 k 及后面的索引是否落在区间内
		如果某个索引 j 已经找到过 最小区间长度，则跳过
		时间复杂度为 O(n*m)，n m 为 intervals 和 queries 的长度，此方法会超时
		ps：超时代码贴在最后
	4.上面方法中，验证 k 及后面的索引是否落在区间内，时间复杂度为 O(n)
		有优化空间吗？如果我们找到一种方法，能连续跳过 x 个 已经找到过 最小区间长度的索引
		就可以避免重复遍历，均摊时间复杂度简化为 O(1)
	5.使用并查集跳过连续空间
		当某个索引已经 找到过，那么我们就把这个索引合并到“下一个索引”集中
		详见code
*/
// leetcode submit region begin(Prohibit modification and deletion)
func minInterval(intervals [][]int, queries []int) []int {

}

//leetcode submit region end(Prohibit modification and deletion)

func minInterval_(intervals [][]int, queries []int) []int {
	// 排序+二分：超时
	//m := len(queries)
	//sort.Slice(intervals, func(i, j int) bool { // 区间左边界升序排序
	//	return intervals[i][1]-intervals[i][0] < intervals[j][1]-intervals[j][0]
	//})
	//ans, vi := make([]int, m), make([]vIdx, m)
	//for i := 0; i < m; i++ { // 初始化并查集 & 长度最小区间为 -1
	//	ans[i] = -1
	//}
	//for idx, v := range queries { // queries[i] 和 i 映射
	//	vi[idx] = vIdx{v, idx}
	//}
	//sort.Slice(vi, func(i, j int) bool { // queries[i]升序排序
	//	return vi[i].v < vi[j].v
	//})
	//for _, in := range intervals {
	//	l, r, cl := in[0], in[1], in[1]-in[0]+1
	//	i := sort.Search(m, func(i int) bool {
	//		return vi[i].v >= l
	//	}) // 二分查找：queries[i]大于等于区间左边界 l 的起始位置
	//	j := sort.Search(m-i, func(k int) bool {
	//		return vi[i+k].v > r
	//	}) + i // 二分：同上，右边界
	//	for ; i < j && vi[i].v <= r; i++ {
	//		if ans[vi[i].idx] > 0 {
	//			continue
	//		}
	//		ans[vi[i].idx] = cl
	//	} // 遍历可能落在 [l,r] 之间的queries[i]
	//}
	//return ans

	// 排序+二分+并查集
	type vIdx struct {
		v   int
		idx int
	}
	m := len(queries)
	sort.Slice(intervals, func(i, j int) bool { // 区间左边界升序排序
		return intervals[i][1]-intervals[i][0] < intervals[j][1]-intervals[j][0]
	})
	ans, uni, vi := make([]int, m), make([]int, m+1), make([]vIdx, m)
	uni[m] = m               // 哨兵
	for i := 0; i < m; i++ { // 初始化并查集 & 长度最小区间为 -1
		uni[i], ans[i] = i, -1
	}
	for idx, v := range queries { // queries[i] 和 i 映射
		vi[idx] = vIdx{v, idx}
	}
	sort.Slice(vi, func(i, j int) bool { // queries[i]升序排序
		return vi[i].v < vi[j].v
	})
	//var fd func(int) int
	//fd = func(p int) int { // 并查集的 find 方法
	//	if p != uni[p] {
	//		uni[p] = fd(uni[p])
	//	}
	//	return uni[p]
	//}
	fd := func(p int) int { // 更喜欢这种写法：但是这就没有回溯功能了
		if p == uni[p] {
			return p
		}
		for uni[p] != uni[uni[p]] {
			uni[p] = uni[uni[p]]
		}
		return uni[p]
	}
	for _, in := range intervals {
		l, r, cl := in[0], in[1], in[1]-in[0]+1
		i := sort.Search(m, func(i int) bool {
			return vi[i].v >= l
		}) // 二分查找：queries[i]大于等于区间左边界 l 的起始位置
		for i = fd(i); i < m && vi[i].v <= r; i = uni[i] {
			ans[vi[i].idx], uni[i] = cl, fd(i+1) // 合并已查询到的索引
		} // 遍历可能落在 [l,r] 之间的queries[i]
	}
	return ans

	// 按照区间长度由小到大排序，这样每次回答的时候用的就是长度最小的区间
	//sort.Slice(intervals, func(i, j int) bool { a, b := intervals[i], intervals[j]; return a[1]-a[0] < b[1]-b[0] })
	//
	//m := len(queries)
	//type pair struct{ pos, i int }
	//qs := make([]pair, m)
	//for i, q := range queries {
	//	qs[i] = pair{q, i}
	//}
	//// 离线：按查询位置排序
	//sort.Slice(qs, func(i, j int) bool { return qs[i].pos < qs[j].pos })
	//
	//// 初始化并查集
	//fa := make([]int, m+1)
	//for i := range fa {
	//	fa[i] = i
	//}
	//var find func(int) int
	//find = func(x int) int {
	//	if fa[x] != x {
	//		fa[x] = find(fa[x])
	//	}
	//	return fa[x]
	//}
	//
	//ans := make([]int, m)
	//for i := range ans {
	//	ans[i] = -1
	//}
	//fmt.Println(qs)
	//fmt.Println(intervals)
	//// 对每个区间，回答所有在 [l,r] 范围内的询问
	//// 由于每次回答询问之后，都将其指向了下一个询问
	//// 所以若 i = find(i) 符合 i < m && qs[i].pos <= r 的条件，则必然是一个在 [l,r] 范围内的还没有回答过的询问
	//for _, p := range intervals {
	//	l, r := p[0], p[1]
	//	length := r - l + 1
	//	// 二分找大于等于区间左端点的最小询问
	//	i := sort.Search(m, func(i int) bool { return qs[i].pos >= l })
	//	// 回答所有询问位置在 [l,r] 范围内的还没有被回答过的询问
	//	//for i = find(i); i < m && qs[i].pos <= r; i = find(i + 1) {
	//	for i = find(i); i < m && qs[i].pos <= r; {
	//		ans[qs[i].i] = length
	//		//fa[i] = i + 1
	//		fa[i] = find(i+1)
	//		i = fa[i]
	//		//fmt.Println(l, r, i, fa)
	//	}
	//}
	//return ans
}
