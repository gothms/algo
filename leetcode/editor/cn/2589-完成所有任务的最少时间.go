package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	tasks := [][]int{{1, 3, 2},
		{2, 5, 3},
		{5, 6, 2}}
	time := findMinimumTime(tasks)
	fmt.Println(time)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findMinimumTime(tasks [][]int) int {
	// 线段树
	n := len(tasks)
	type segment []struct {
		l, r int
		cnt  int
		todo bool
	}
	st := make(segment, func() {

	})
}

//leetcode submit region end(Prohibit modification and deletion)

func findMinimumTime_(tasks [][]int) int {
	// lc：栈 + 二分
	slices.SortFunc(tasks, func(a, b []int) int { return a[1] - b[1] })
	// 栈中保存闭区间左右端点，栈底到栈顶的区间长度的和
	type tuple struct{ l, r, s int }
	st := []tuple{{-2, -2, 0}} // 哨兵，保证不和任何区间相交
	for _, p := range tasks {
		start, end, d := p[0], p[1], p[2]
		i := sort.Search(len(st), func(i int) bool { return st[i].l >= start }) - 1 // -1：st[i].l < start，找全所有已运行的时间点
		d -= st[len(st)-1].s - st[i].s                                              // 去掉运行中的时间点
		if start <= st[i].r {                                                       // start 在区间 st[i] 内
			d -= st[i].r - start + 1 // 去掉运行中的时间点
		}
		if d <= 0 {
			continue
		}
		for end-st[len(st)-1].r <= d { // 剩余的 d 填充区间后缀，<= 则可以合并区间
			top := st[len(st)-1]
			st = st[:len(st)-1]
			d += top.r - top.l + 1 // 合并区间
		}
		st = append(st, tuple{end - d + 1, end, st[len(st)-1].s + d})
	}
	return st[len(st)-1].s

	// 线段树

	// 贪心
	//sort.Slice(tasks, func(i, j int) bool { // 贪心的关键
	//	return tasks[i][1] < tasks[j][1]
	//})
	//run := make([]bool, tasks[len(tasks)-1][1]+1)
	//ans := 0
	//for _, t := range tasks {
	//	cnt := t[2]
	//	for _, b := range run[t[0] : t[1]+1] { // 统计区间内的已运行的电脑运行时间点
	//		if b {
	//			cnt--
	//		}
	//	}
	//	for j := t[1]; cnt > 0; j-- { // 贪心：尽量把新增的时间点安排在区间的后缀上
	//		if !run[j] {
	//			run[j] = true
	//			cnt--
	//			ans++
	//		}
	//	}
	//}
	//return ans
}
