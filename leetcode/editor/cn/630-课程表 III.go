//这里有 n 门不同的在线课程，按从 1 到 n 编号。给你一个数组 courses ，其中 courses[i] = [durationi,
//lastDayi] 表示第 i 门课将会 持续 上 durationi 天课，并且必须在不晚于 lastDayi 的时候完成。
//
// 你的学期从第 1 天开始。且不能同时修读两门及两门以上的课程。
//
// 返回你最多可以修读的课程数目。
//
//
//
// 示例 1：
//
//
//输入：courses = [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
//输出：3
//解释：
//这里一共有 4 门课程，但是你最多可以修 3 门：
//首先，修第 1 门课，耗费 100 天，在第 100 天完成，在第 101 天开始下门课。
//第二，修第 3 门课，耗费 1000 天，在第 1100 天完成，在第 1101 天开始下门课程。
//第三，修第 2 门课，耗时 200 天，在第 1300 天完成。
//第 4 门课现在不能修，因为将会在第 3300 天完成它，这已经超出了关闭日期。
//
// 示例 2：
//
//
//输入：courses = [[1,2]]
//输出：1
//
//
// 示例 3：
//
//
//输入：courses = [[3,2],[4,3]]
//输出：0
//
//
//
//
// 提示:
//
//
// 1 <= courses.length <= 10⁴
// 1 <= durationi, lastDayi <= 10⁴
//
//
// Related Topics 贪心 数组 排序 堆（优先队列） 👍 411 👎 0

package main

import (
	"fmt"
)

func main() {
	courses := [][]int{{100, 200}, {200, 450}, {300, 400}, {2000, 3000}}
	courses = [][]int{{3, 2}, {4, 3}}
	//courses = [][]int{{5, 5}, {4, 6}, {2, 6}}
	course := scheduleCourse(courses)
	fmt.Println(course)
}

// leetcode submit region begin(Prohibit modification and deletion)
func scheduleCourse(courses [][]int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

//func scheduleCourse(courses [][]int) int {
//	sort.Slice(courses, func(i, j int) bool {
//		return courses[i][1] < courses[j][1] // 按期限排序
//	})
//	h := &hp630{sort.IntSlice{0}} // 大顶堆
//	var curTimes int              // 总耗时
//	for _, c := range courses {   // 还能学就继续加课
//		if c[0]+curTimes <= c[1] {
//			curTimes += c[0]
//			heap.Push(h, c[0])
//		} else if c[0] < h.IntSlice[0] { // 优化耗时最多的课程
//			curTimes -= h.IntSlice[0] - c[0]
//			h.IntSlice[0] = c[0] // 替换耗时最多的课程
//			heap.Fix(h, 0)
//		}
//	}
//	return h.Len() - 1
//}
//
//type hp630 struct {
//	sort.IntSlice
//}
//
//func (h hp630) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
//func (h *hp630) Push(x any)        { h.IntSlice = append(h.IntSlice, x.(int)) }
//func (h *hp630) Pop() any {
//	v := h.IntSlice[len(h.IntSlice)-1]
//	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
//	return v
//}
