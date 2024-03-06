//给你两个长度为 n 、下标从 0 开始的整数数组 nums1 和 nums2 ，另给你一个下标从 1 开始的二维数组 queries ，其中
//queries[i] = [xi, yi] 。
//
// 对于第 i 个查询，在所有满足 nums1[j] >= xi 且 nums2[j] >= yi 的下标 j (0 <= j < n) 中，找出 nums1
//[j] + nums2[j] 的 最大值 ，如果不存在满足条件的 j 则返回 -1 。
//
// 返回数组 answer ，其中 answer[i] 是第 i 个查询的答案。
//
//
//
// 示例 1：
//
// 输入：nums1 = [4,3,1,2], nums2 = [2,4,9,5], queries = [[4,1],[1,3],[2,5]]
//输出：[6,10,7]
//解释：
//对于第 1 个查询：xi = 4 且 yi = 1 ，可以选择下标 j = 0 ，此时 nums1[j] >= 4 且 nums2[j] >= 1 。
//nums1[j] + nums2[j] 等于 6 ，可以证明 6 是可以获得的最大值。
//对于第 2 个查询：xi = 1 且 yi = 3 ，可以选择下标 j = 2 ，此时 nums1[j] >= 1 且 nums2[j] >= 3 。
//nums1[j] + nums2[j] 等于 10 ，可以证明 10 是可以获得的最大值。
//对于第 3 个查询：xi = 2 且 yi = 5 ，可以选择下标 j = 3 ，此时 nums1[j] >= 2 且 nums2[j] >= 5 。
//nums1[j] + nums2[j] 等于 7 ，可以证明 7 是可以获得的最大值。
//因此，我们返回 [6,10,7] 。
//
//
// 示例 2：
//
// 输入：nums1 = [3,2,5], nums2 = [2,3,4], queries = [[4,4],[3,2],[1,1]]
//输出：[9,9,9]
//解释：对于这个示例，我们可以选择下标 j = 2 ，该下标可以满足每个查询的限制。
//
//
// 示例 3：
//
// 输入：nums1 = [2,1], nums2 = [2,3], queries = [[3,3]]
//输出：[-1]
//解释：示例中的查询 xi = 3 且 yi = 3 。对于每个下标 j ，都只满足 nums1[j] < xi 或者 nums2[j] < yi 。因此，不
//存在答案。
//
//
//
//
// 提示：
//
//
// nums1.length == nums2.length
// n == nums1.length
// 1 <= n <= 10⁵
// 1 <= nums1[i], nums2[i] <= 10⁹
// 1 <= queries.length <= 10⁵
// queries[i].length == 2
// xi == queries[i][1]
// yi == queries[i][2]
// 1 <= xi, yi <= 10⁹
//
//
// Related Topics 栈 树状数组 线段树 数组 二分查找 排序 单调栈 👍 22 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{4, 3, 1, 2}
	nums2 := []int{2, 4, 9, 5}
	queries := [][]int{{4, 1}, {1, 3}, {2, 5}}
	nums1 = []int{72, 44}
	nums2 = []int{60, 86}
	queries = [][]int{{33, 14}}
	sumQueries := maximumSumQueries(nums1, nums2, queries)
	fmt.Println(sumQueries)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	n, m := len(nums1), len(queries)
	p, q := make([][2]int, n), make([]int, m)
	for i := 0; i < n; i++ {
		p[i] = [2]int{nums1[i], nums2[i]}
	}
	for i := 1; i < m; i++ {
		q[i] = i
	}
	sort.Slice(p, func(i, j int) bool { // 按 nums1 排序
		return p[i][0] > p[j][0]
	})
	sort.Slice(q, func(i, j int) bool { // 按 queries[i][0] 排序
		return queries[q[i]][0] > queries[q[j]][0]
	})
	ret, stack := make([]int, m), make([]int, 0, m)
	i := 0
	for _, v := range q { // 查询
		x, y := queries[v][0], queries[v][1]
		for ; i < n && p[i][0] >= x; i++ { // nums1 符合
			j := len(stack)
			for s := p[i][0] + p[i][1]; j > 0 && s >= p[stack[j-1]][0]+p[stack[j-1]][1]; j-- { // 维护单调栈：递减
				stack = stack[:j-1]
			}
			if j == 0 || p[i][1] >= p[stack[j-1]][1] { // nums2 符合
				stack = append(stack, i)
			}
		}
		idx := sort.Search(len(stack), func(i int) bool { // 二分查 nums2
			return p[stack[i]][1] >= y
		})
		if idx < len(stack) {
			ret[v] = p[stack[idx]][0] + p[stack[idx]][1]
		} else {
			ret[v] = -1
		}
	}
	return ret

	// lc
	//type pair struct{ x, y int }
	//a := make([]pair, len(nums1))
	//for i, x := range nums1 {
	//	a[i] = pair{x, nums2[i]}
	//}
	//qid := make([]int, len(queries))
	//for i := range qid {
	//	qid[i] = i
	//}
	//sort.Slice(a, func(i, j int) bool {
	//	return a[i].x > a[j].x
	//})
	//sort.Slice(qid, func(i, j int) bool {
	//	return queries[qid[i]][0] > queries[qid[j]][0]
	//})
	//fmt.Println(a, qid)
	//ans := make([]int, len(queries))
	//type data struct{ y, s int }
	//st := []data{}
	//j := 0
	//for _, i := range qid {
	//	x, y := queries[i][0], queries[i][1]
	//	for ; j < len(a) && a[j].x >= x; j++ { // 下面只需关心 a[j].y
	//		for len(st) > 0 && st[len(st)-1].s <= a[j].x+a[j].y { // a[j].y >= st[len(st)-1].y
	//			st = st[:len(st)-1]
	//		}
	//		if len(st) == 0 || st[len(st)-1].y < a[j].y {
	//			st = append(st, data{a[j].y, a[j].x + a[j].y})
	//		}
	//	}
	//	p := sort.Search(len(st), func(i int) bool { return st[i].y >= y })
	//	if p < len(st) {
	//		ans[i] = st[p].s
	//	} else {
	//		ans[i] = -1
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
