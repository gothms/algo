//给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
//
// 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
//
//
//
// 示例 1：
//
//
//输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
//输出：[[1,5],[6,9]]
//
//
// 示例 2：
//
//
//输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//输出：[[1,2],[3,10],[12,16]]
//解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
//
// 示例 3：
//
//
//输入：intervals = [], newInterval = [5,7]
//输出：[[5,7]]
//
//
// 示例 4：
//
//
//输入：intervals = [[1,5]], newInterval = [2,3]
//输出：[[1,5]]
//
//
// 示例 5：
//
//
//输入：intervals = [[1,5]], newInterval = [2,7]
//输出：[[1,7]]
//
//
//
//
// 提示：
//
//
// 0 <= intervals.length <= 10⁴
// intervals[i].length == 2
// 0 <= intervals[i][0] <= intervals[i][1] <= 10⁵
// intervals 根据 intervals[i][0] 按 升序 排列
// newInterval.length == 2
// 0 <= newInterval[0] <= newInterval[1] <= 10⁵
//
//
// Related Topics 数组 👍 739 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{3, 4}
	//intervals = [][]int{{1, 5}}
	//newInterval = []int{0, 0}
	//intervals = [][]int{{3, 5}, {12, 15}}
	//newInterval = []int{6, 6}
	//intervals = [][]int{}
	//newInterval = []int{6, 6}
	i := insert(intervals, newInterval)
	fmt.Println(i)
}

/*
思路：二分
	1.插入 newInterval 的情况可以分为两种：
		newInterval 与 intervals 某些区间有交集，则合并这些区间
		没有任何交集，则插入 newInterval。插入位置可能是：头/尾/中间
		所以关键问题是，找到插入位置或需要合并的一些区间的“首尾”
	2.两次二分查找
		一找：插入/合并的起始位置
			算法为，查询 intervals[i][1] 大于等于 newInterval[0] 的起始位置
		二找：插入/合并的终止位置，这里我们找终止位置的后一个位置
			算法为，查询 intervals[i][0] 大于 newInterval[1] 的起始位置
*/
// leetcode submit region begin(Prohibit modification and deletion)
func insert(intervals [][]int, newInterval []int) [][]int {
	// 二分
	n := len(intervals)
	l := sort.Search(n, func(i int) bool {
		return intervals[i][1] >= newInterval[0]
	}) // 查询大于等于 newInterval[0] 的起始位置
	r := sort.Search(n, func(i int) bool {
		return intervals[i][0] > newInterval[1]
	}) // 查询大于 newInterval[1] 的起始位置
	switch {
	case r == 0: // 添加到头部
		return append([][]int{newInterval}, intervals...)
	case l == n: // 添加到尾部
		return append(intervals, newInterval)
	case l == r: // 插入中间
		//ret := make([][]int, 0, n+1)	// 效率更高
		intervals = append(intervals, nil)
		copy(intervals[l+1:], intervals[l:])
		intervals[l] = newInterval
		return intervals
	default: // 合并区间
		if newInterval[0] < intervals[l][0] {
			intervals[l][0] = newInterval[0]
		}
		if newInterval[1] > intervals[r-1][1] {
			intervals[l][1] = newInterval[1]
		} else {
			intervals[l][1] = intervals[r-1][1]
		}
		//intervals[l][0], intervals[l][1] =
		//	minVal(intervals[l][0], newInterval[0]), maxVal(intervals[r-1][1], newInterval[1])
		return append(intervals[:l+1], intervals[r:]...)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
