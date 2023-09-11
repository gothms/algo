//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返
//回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
//
//
//
// 示例 1：
//
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//
//
// 示例 2：
//
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
//
//
// 提示：
//
//
// 1 <= intervals.length <= 10⁴
// intervals[i].length == 2
// 0 <= starti <= endi <= 10⁴
//
//
// Related Topics 数组 排序 👍 2000 👎 0

package main

import "sort"

func main() {

}

/*
思路：排序
思路：归并
*/
// leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	// 排序
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ret, n := make([][]int, 0), len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	}) // 按区间左边界排序
	for i := 0; i < n; {
		s, e := intervals[i][0], intervals[i][1] // s 为左边界
		for i++; i < n && e >= intervals[i][0]; i++ {
			e = maxVal(e, intervals[i][1]) // 寻找右边界 e
		}
		ret = append(ret, []int{s, e}) // 合并后的区间 [s,e]
	}
	return ret

	// 归并
	//return mergeSort(intervals, 0, len(intervals)-1)
}

// 归并排序
func mergeSort(dst [][]int, l, r int) [][]int {
	if l >= r {
		return [][]int{dst[r]}
	}
	m := (l + r) >> 1
	return mergeInterval(mergeSort(dst, l, m), mergeSort(dst, m+1, r))
}

// 按区间左边界进行排序
func mergeInterval(l [][]int, r [][]int) [][]int {
	var ret [][]int
	i, j, n, m := 0, 0, len(l), len(r)
	if l[i][0] <= r[j][0] {
		ret, i = append(ret, l[i]), 1
	} else {
		ret, j = append(ret, r[j]), 1
	}
	for ; i < n; i++ {
		for ; j < m && l[i][0] > r[j][0]; j++ {
			appendInterval(&ret, r[j])
		}
		appendInterval(&ret, l[i])
	}
	for ; j < m; j++ {
		appendInterval(&ret, r[j])
	}
	return ret
}

// 尝试合并区间
func appendInterval(ret *[][]int, v []int) {
	last := len(*ret) - 1
	if (*ret)[last][1] < v[0] { // 拼接
		*ret = append(*ret, v) // 地址变化
	} else if (*ret)[last][1] < v[1] { // 合并
		(*ret)[last][1] = v[1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)

/*
{2, 3, 4, 5, 6}
{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 16, 17, 18, 19, 20}
{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 16, 17, 18, 19, 20, 30, 31, 32, 33, 34, 35, 36}
*/

// lc讨论区：位图表示坐标轴
//func java() {
//	class
//	Solution{
//		public, int[][] merge(int[][] intervals){
//		BitSet bitSet = new BitSet()
//		int max = 0
//		for (int[] interval: intervals){
//		int temp = interval[1] * 2 + 1
//		bitSet.set(interval[0] * 2, temp, true)
//		max = temp >= max ? temp: max
//		System.out.println(bitSet);
//	}
//
//		int index = 0, count = 0
//		while (index < max){
//		int start = bitSet.nextSetBit(index)	// 返回第一个设置为 true 的位的索引，这发生在指定的起始索引或之后的索引上
//		int end = bitSet.nextClearBit(start)	// 返回第一个设置为 false 的位的索引，这发生在指定的起始索引或之后的索引上
//
//		int[] item ={start / 2, (end - 1) / 2}
//		intervals[count++] = item
//
//		index = end
//	}
//		int[][] ret = new int[count][2]
//		for (int i = 0; i < count; i++){
//		ret[i] = intervals[i]
//	}
//
//		return ret
//	}
//	}
//}
