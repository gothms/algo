//给你两个下标从 0 开始的数组 nums1 和 nums2 ，和一个二维数组 queries 表示一些操作。总共有 3 种类型的操作：
//
//
// 操作类型 1 为 queries[i] = [1, l, r] 。你需要将 nums1 从下标 l 到下标 r 的所有 0 反转成 1 或将 1 反转成
//0 。l 和 r 下标都从 0 开始。
// 操作类型 2 为 queries[i] = [2, p, 0] 。对于 0 <= i < n 中的所有下标，令 nums2[i] = nums2[i] +
// nums1[i] * p 。
// 操作类型 3 为 queries[i] = [3, 0, 0] 。求 nums2 中所有元素的和。
//
//
// 请你返回一个数组，包含所有第三种操作类型的答案。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,0,1], nums2 = [0,0,0], queries = [[1,1,1],[2,1,0],[3,0,0]]
//输出：[3]
//解释：第一个操作后 nums1 变为 [1,1,1] 。第二个操作后，nums2 变成 [1,1,1] ，所以第三个操作的答案为 3 。所以返回 [3] 。
//
//
//
// 示例 2：
//
//
//输入：nums1 = [1], nums2 = [5], queries = [[2,0,0],[3,0,0]]
//输出：[5]
//解释：第一个操作后，nums2 保持不变为 [5] ，所以第二个操作的答案是 5 。所以返回 [5] 。
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length,nums2.length <= 10⁵
// nums1.length = nums2.length
// 1 <= queries.length <= 10⁵
// queries[i].length = 3
// 0 <= l <= r <= nums1.length - 1
// 0 <= p <= 10⁶
// 0 <= nums1[i] <= 1
// 0 <= nums2[i] <= 10⁹
//
//
// Related Topics 线段树 数组 👍 16 👎 0

package main

import "fmt"

func main() {
	nums1 := []int{1, 0, 1}
	nums2 := []int{0, 0, 0}
	queries := [][]int{{1, 1, 1}, {2, 1, 0}, {3, 0, 0}}
	nums1 = []int{1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 0}
	nums2 = []int{33, 13, 37, 28, 42, 5, 39, 31, 12, 1, 7, 23, 21}
	// [292,292,292,292,292,292]
	queries = [][]int{{1, 1, 3}, {3, 0, 0}, {1, 0, 9}, {3, 0, 0}, {3, 0, 0}, {3, 0, 0}, {3, 0, 0},
		{1, 0, 1}, {3, 0, 0}, {1, 9, 9}, {2, 3, 0}, {2, 27, 0}, {1, 0, 9}, {1, 3, 5}}
	nums1 = []int{0, 1, 0, 0, 0, 0}
	//nums1 = []int{1, 2, 3, 4, 5, 6}
	nums2 = []int{14, 4, 13, 13, 47, 18}
	// [109,109,197,197]
	queries = [][]int{{3, 0, 0}, {1, 4, 4}, {1, 1, 4}, {1, 3, 4}, {3, 0, 0}, {2, 5, 0}, {1, 1, 3}, {2, 16, 0},
		{2, 10, 0}, {3, 0, 0}, {3, 0, 0}, {2, 6, 0}}
	query := handleQuery(nums1, nums2, queries)
	fmt.Println(query)
}

/*
思路：线段树 & 懒惰标记
	1.分析 queries[i][0] 的3种行为
		1.1.queries[i][0]==1：[f,t] 区间更新 nums1 中的元素
			f，t 分别为更新的起始点和终点，更新方式为 nums1[i] ^= 1
			那怎么高效更新这个区间中的元素呢？线段树就是区间更新的数据结构（单元素更新请使用树状数组）
		1.2.queries[i][0]==2：先计算 nums2 的初始总和 sum，当 nums1 更新后，总和 sum 有什么变化呢？
			由题意 nums2[i] = nums2[i] + nums1[i] * p，可知 nums2 的总和 sum 的更新
			与 sum 的初始值无关，只与每次 nums1 更新后增加值有关（nums1[i] * p）
			即增加值 = 当前 nums1 中 1 的总数 * p
		1.3.queries[i][0]==3：
			将 sum 追加到结果集中
	2.线段树 + 懒惰标记
		2.1.线段树的每个节点记录它这棵子树中 1 的总数
		2.2.如何高效的更新线段树的区间值？使用懒惰标记
			懒惰标记的数组 lazy，记录了当前 update 行为下，某个节点对应的子树，是否需要更新
			当下次访问到这个节点时，就更新这棵子树（仅更新访问到这个节点的左右子树，而不是整棵树更新）
*/
// leetcode submit region begin(Prohibit modification and deletion)
func handleQuery(nums1 []int, nums2 []int, queries [][]int) []int64 {
	// 线段树 & 懒惰标记
	sum, n := int64(0), len(nums1)
	segmentTree, lazy := make([]int, n<<2), make([]bool, n<<2)
	for i := 0; i < n; i++ {
		sum += int64(nums2[i]) // 初始和
	}
	build(0, n-1, nums1, 0, segmentTree, lazy) // 构建线段树
	sums := make([]int64, 0)
	for _, q := range queries {
		switch q[0] {
		case 1: // 区间更新线段树
			updateLazy(q[1], q[2], 0, n-1, 0, segmentTree, lazy)
		case 2: // 计算当前和
			sum += int64(segmentTree[0]) * int64(q[1])
		case 3: // 结果集
			sums = append(sums, sum)
		}
	}
	return sums
}
func build(l, r int, arr []int, i int, st []int, lazy []bool) { // 构建线段树
	if l == r {
		st[i] = arr[l]
		lazy[i] = true // 懒惰标记
		return
	}
	m, idx := (l+r)>>1, i<<1+1
	build(l, m, arr, idx, st, lazy)
	build(m+1, r, arr, idx+1, st, lazy)
	st[i] = st[idx] + st[idx+1]
}
func updateLazy(f, t int, l, r int, i int, st []int, lazy []bool) {
	if r <= t && f <= l {
		st[i] = r - l + 1 - st[i] // 查询区间落在 [l,r] 范围内
		lazy[i] = !lazy[i]        // 更新当前节点的懒惰标记
		return
	}
	m, idx := (l+r)>>1, i<<1+1
	if lazy[i] {
		lazyDown(l, m, r, i, idx, st, lazy) // 更新子节点的懒惰标记
	}
	if f <= m { // 和左节点代表的区间 [l,m] 有交集
		updateLazy(f, t, l, m, idx, st, lazy)
	}
	if t > m { // 和右节点代表的区间 [m+1,r] 有交集
		updateLazy(f, t, m+1, r, idx+1, st, lazy)
	}
	st[i] = st[idx] + st[idx+1] // 更新 1 总数
}
func lazyDown(l, m, r int, i, idx int, st []int, lazy []bool) {
	lazy[i] = false                                   // 更新当前节点的懒惰标记
	lazy[idx], lazy[idx+1] = !lazy[idx], !lazy[idx+1] // 更新左右节点的懒惰标记
	st[idx], st[idx+1] = m-l+1-st[idx], r-m-st[idx+1] // 更新左右子树的 1 总数
}

// leetcode submit region end(Prohibit modification and deletion)

// 个人实现：没优化
func handleQuery_old(nums1 []int, nums2 []int, queries [][]int) []int64 {
	n := len(nums1)
	segmentTree, lazy := make([]int, n<<2), make([]int, n<<2)
	for i := range lazy {
		lazy[i] = -1
	}
	build_old(0, n-1, nums1, 0, segmentTree, lazy)
	//fmt.Println(segmentTree[:13])
	//fmt.Println(lazy[:13])
	var sum int64
	sums := make([]int64, 0)
	for i := 0; i < n; i++ {
		sum += int64(nums2[i])
	}
	for _, q := range queries {
		switch q[0] {
		case 1:
			updateLazy_old(q[1], q[2], 0, n-1, 0, segmentTree, lazy)
			//st := make([]int, n<<2)
			//for i := q[1]; i <= q[2]; i++ {
			//	nums1[i] ^= 1
			//}
			//build(0, n-1, nums1, 0, st, lazy)
			//fmt.Println("st:", st[:13])
			//fmt.Println("cr:", segmentTree[:13])
			//fmt.Println(lazy[:13])
		case 2:
			sum += int64(segmentTree[0]) * int64(q[1])
		case 3:
			sums = append(sums, sum)
		}
	}
	return sums
}
func build_old(l, r int, arr []int, i int, st, lazy []int) {
	if l == r {
		st[i] = arr[l]
		lazy[i] = 0
		return
	}
	m, idx := (l+r)>>1, i<<1+1
	build_old(l, m, arr, idx, st, lazy)
	build_old(m+1, r, arr, idx+1, st, lazy)
	st[i] = st[idx] + st[idx+1]
	lazy[i] = 0
}
func rangeSumLazy_old(st []int) int {
	return st[0]
}
func updateLazy_old(f, t int, l, r int, i int, st, lazy []int) {
	if r <= t && f <= l {
		st[i] = r - l + 1 - st[i]
		lazy[i] ^= 1
		return
	}
	m, idx := (l+r)>>1, i<<1+1
	if lazy[i] > 0 {
		lazyDown_old(l, m, r, i, idx, st, lazy)
	}
	if f <= m {
		updateLazy_old(f, t, l, m, idx, st, lazy)
	}
	if t > m {
		updateLazy_old(f, t, m+1, r, idx+1, st, lazy)
	}
	st[i] = st[idx] + st[idx+1]
}
func lazyDown_old(l, m, r int, i, idx int, st, lazy []int) {
	//st[idx], st[idx+1] = m-l+1-st[idx], r-m-st[idx+1]
	//lazy[idx], lazy[idx+1] = 1, 1
	//lazy[i] = 0

	//lazy[i] = 0
	//if lazy[idx] < 0 {
	//	return
	//}
	//lazy[idx] ^= 1
	//if lazy[idx] > 0 {
	//	st[idx] = m - l + 1 - st[idx]
	//}
	//if lazy[idx+1] < 0 {
	//	return
	//}
	//lazy[idx+1] ^= 1
	//if lazy[idx+1] > 0 {
	//	st[idx+1] = r - m - st[idx+1]
	//}

	lazy[i] = 0
	lazy[idx] ^= 1
	lazy[idx+1] ^= 1
	st[idx] = m - l + 1 - st[idx]
	st[idx+1] = r - m - st[idx+1]
	fmt.Println(lazy)
}
