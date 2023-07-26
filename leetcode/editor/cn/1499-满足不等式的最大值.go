//给你一个数组 points 和一个整数 k 。数组中每个元素都表示二维平面上的点的坐标，并按照横坐标 x 的值从小到大排序。也就是说 points[i] =
// [xi, yi] ，并且在 1 <= i < j <= points.length 的前提下， xi < xj 总成立。
//
// 请你找出 yi + yj + |xi - xj| 的 最大值，其中 |xi - xj| <= k 且 1 <= i < j <= points.
//length。
//
// 题目测试数据保证至少存在一对能够满足 |xi - xj| <= k 的点。
//
//
//
// 示例 1：
//
// 输入：points = [[1,3],[2,0],[5,10],[6,-10]], k = 1
//输出：4
//解释：前两个点满足 |xi - xj| <= 1 ，代入方程计算，则得到值 3 + 0 + |1 - 2| = 4 。第三个和第四个点也满足条件，得到值 1
//0 + -10 + |5 - 6| = 1 。
//没有其他满足条件的点，所以返回 4 和 1 中最大的那个。
//
// 示例 2：
//
// 输入：points = [[0,0],[3,0],[9,2]], k = 3
//输出：3
//解释：只有前两个点满足 |xi - xj| <= 3 ，代入方程后得到值 0 + 0 + |0 - 3| = 3 。
//
//
//
//
// 提示：
//
//
// 2 <= points.length <= 10^5
// points[i].length == 2
// -10^8 <= points[i][0], points[i][1] <= 10^8
// 0 <= k <= 2 * 10^8
// 对于所有的1 <= i < j <= points.length ，points[i][0] < points[j][0] 都成立。也就是说，xi 是严格
//递增的。
//
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 66 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	// -6
	points := [][]int{{-19, 9}, {-15, -19}, {-5, -8}}
	k := 10
	equation := findMaxValueOfEquation(points, k)
	fmt.Println(equation)
}

/*
思路：优先队列
	1.定义大顶堆作为优先队列 type eq [][2]int
		[2]int：第一个元素定义为 xi+yi，第二个元素为 i
	2.大顶堆的更新
		我们将满足 |xi - xj| <= k（也就是满足 xj-xi<=k）的元素都放入一个优先队列中
		此处我们以 xi 为当前要处理的元素，把后面满足条件的 j 都招出来，放入优先队列
		当然，也可以以 xj 为当前要处理的元素
	3.尝试更新最大值
		当处理 xi 时，优先队列中小于等于 索引i 的元素，都需要淘汰掉
		然后取堆顶元素作为 j，尝试更新 maxV
		maxV = maxVal(maxV, yj+xj + yi-xi)
思路：双端队列
	1.上面方法是采用优先队列“挑选”当前最优 yj+xj + yi-xi
		我们可以用单调队列来挑选（假设此处我们以 xj 为当前要处理的元素）
		因为对于 yj+xj + yi-xi，当前元素 j 是固定的，那么 yj+xj 就是固定的
		则只需要考虑从满足 yi-xi<=k 条件的元素中，挑选出最大的 yi-xi
	2.我们将满足 xj-xi<=k 的元素都放入一个双端队列 window 中
		2.1.当前要处理的元素为 j 时，先淘汰掉满足 xj - xi > k 的元素 i
		2.2.以 j 为当前元素，尝试更新最大值
			i 元素为 window 中的第一个元素
		2.3.此时还需要把 j 放入 window 中，为后续机选中作为元素 i 使用
			而根据 1. 分析，window 队尾满足 yi-xi <= yj-yi 的都需要淘汰掉
			也就是说 window 其实是一个单调递减的双端队列
*/
// leetcode submit region begin(Prohibit modification and deletion)
func findMaxValueOfEquation(points [][]int, k int) int {
	// 双端队列
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxV, window := math.MinInt32, make([]int, 0)
	for idx, p := range points {
		i, j := 0, len(window)
		for sub := p[0] - k; i < j && sub > points[window[i]][0]; {
			i++
		} // 左淘汰：xj - xi > k（注释中的 i j 指的是题意中的 i j）
		if i < j { // 尝试更新最大值
			maxV = maxVal(maxV, p[0]+p[1]-points[window[i]][0]+points[window[i]][1])
		}
		for sub := p[1] - p[0]; j > i && sub >= points[window[j-1]][1]-points[window[j-1]][0]; {
			j--
		} // 右淘汰：yj - xj >= yj - yi
		window = append(window[i:j], idx) // 右添加：当前元素的索引 j
	}
	return maxV

	// 优先队列
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//maxV, j, n := math.MinInt32, 1, len(points)
	//h := eq{}
	//for i, p := range points {
	//	for h.Len() > 0 && h[0][1] <= i { // 已“过期”的元素需要淘汰
	//		heap.Pop(&h)
	//	}
	//	j = maxVal(j, i+1) // 更新 j
	//	for v := p[0] + k; j < n && v >= points[j][0]; j++ {
	//		heap.Push(&h, [2]int{points[j][0] + points[j][1], j})
	//	} // 满足 xj-xi<=k 的元素 j，都放入优先队列
	//	if h.Len() > 0 { // 尝试更新最大值
	//		maxV = maxVal(maxV, h[0][0]+p[1]-p[0])
	//	}
	//}
	//return maxV
}

type eq [][2]int

func (e eq) Len() int {
	return len(e)
}

func (e eq) Less(i, j int) bool {
	return e[i][0] > e[j][0]
}

func (e eq) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e *eq) Push(x any) {
	*e = append(*e, x.([2]int))
}

func (e *eq) Pop() any {
	v := (*e)[len(*e)-1]
	*e = (*e)[:len(*e)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
