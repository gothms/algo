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
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	courses := [][]int{{100, 200}, {200, 450}, {300, 400}, {2000, 3000}}
	courses = [][]int{{5, 5}, {4, 6}, {2, 6}}
	course := scheduleCourse(courses)
	fmt.Println(course)
}

// leetcode submit region begin(Prohibit modification and deletion)
func scheduleCourse(courses [][]int) int {
	// 排序 + 优先队列
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1] // 按期限排序
	})
	costTime, h := 0, sc{0}
	for _, c := range courses {
		if t := c[0]; costTime+t <= c[1] { // 还能学就继续加课
			costTime += t    // 总耗时
			heap.Push(&h, t) // “已学”课程
		} else if t <= h[0] { // 不能学就尝试优化
			costTime -= h[0] - t // 修正耗时
			h[0] = t             // 替换：用时更短，期限也更长
			heap.Fix(&h, 0)
		}
	}
	return h.Len() - 1

	//sort.Slice(courses, func(i, j int) bool {
	//	return courses[i][1] < courses[j][1] // 按期限排序
	//})
	//time, h := 0, &sc{sort.IntSlice{0}}
	//for _, c := range courses {
	//	// 写法二
	//	if t := c[0]; time+t <= c[1] { // 能修一门课的条件
	//		time += t       // 总耗时
	//		heap.Push(h, t) // 修这门课
	//	} else if c[0] < h.IntSlice[0] { // c 相较堆顶课程，耗时短，期限更宽裕，则替换它
	//		time -= h.IntSlice[0] - t
	//		h.IntSlice[0] = t
	//		heap.Fix(h, 0)
	//	}
	//
	//	// 写法一
	//	//time += c[0]
	//	//heap.Push(h, c[0]) // 修这门课
	//	//if time > c[1] {
	//	//	time -= heap.Pop(h).(int)
	//	//}
	//}
	//return h.Len() - 1
}

type sc sort.IntSlice

func (s sc) Len() int           { return len(s) }
func (s sc) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sc) Less(i, j int) bool { return s[i] > s[j] }
func (s *sc) Push(x any)        { *s = append(*s, x.(int)) }
func (s *sc) Pop() any {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

//type sc struct {
//	sort.IntSlice
//}
//
//func (s sc) Less(i, j int) bool { return s.IntSlice[i] > s.IntSlice[j] }
//func (s *sc) Push(x any)        { s.IntSlice = append(s.IntSlice, x.(int)) }
//func (s *sc) Pop() any {
//	v := s.IntSlice[len(s.IntSlice)-1]
//	s.IntSlice = s.IntSlice[:len(s.IntSlice)-1]
//	return v
//}

//leetcode submit region end(Prohibit modification and deletion)
