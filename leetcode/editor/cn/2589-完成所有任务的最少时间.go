package main

import (
	"fmt"
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

}

//leetcode submit region end(Prohibit modification and deletion)

func findMinimumTime_(tasks [][]int) int {
	// 栈 + 二分
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})
	type space struct {
		l, r int // 栈中保存闭区间左右端点，栈底到栈顶的区间长度的和
		k    int // 统计已被选择的时间的总数
	}
	st := []space{{-1, -1, 0}} // 哨兵，保证不和任何区间相交
	for _, t := range tasks {  // st 中的区间严格单调递增
		s, e, d := t[0], t[1], t[2]
		i := sort.Search(len(st), func(i int) bool {
			return st[i].l >= s // 1.1.二分结果 i 的右边区间的所有运行中的时间点，都被选择
		}) - 1
		d -= st[len(st)-1].k - st[i].k // 前缀和，第一次去掉运行中的时间点
		if s <= st[i].r {              // 1.2.当前区间 t 和栈中的区间 st[i] 有交集，则选择交集中的所有时间点
			d -= st[i].r - s + 1 // 第二次去掉运行中的时间点
		}
		if d <= 0 { // 已经选够了
			continue
		}
		for e-st[len(st)-1].r <= d { // 2.1.st 最后一个区间和当前区间 t 的 end 时间的空隙将被全部选择（剩余的 d 填充区间后缀，<= 则可以合并区间）
			d += st[len(st)-1].r - st[len(st)-1].l + 1 // 间隙被全部选择
			st = st[:len(st)-1]                        // 区间被合并
		} // 因为 s <= st[i].r 且 st 中的区间有间隙，所有永远不会选到哨兵
		st = append(st, space{e - d + 1, e, d + st[len(st)-1].k}) // 2.2.合并后的区间
	}
	return st[len(st)-1].k

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
