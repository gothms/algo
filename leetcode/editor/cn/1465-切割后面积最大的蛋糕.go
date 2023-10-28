//矩形蛋糕的高度为 h 且宽度为 w，给你两个整数数组 horizontalCuts 和 verticalCuts，其中：
//
//
// horizontalCuts[i] 是从矩形蛋糕顶部到第 i 个水平切口的距离
// verticalCuts[j] 是从矩形蛋糕的左侧到第 j 个竖直切口的距离
//
//
// 请你按数组 horizontalCuts 和 verticalCuts 中提供的水平和竖直位置切割后，请你找出 面积最大 的那份蛋糕，并返回其 面积 。由
//于答案可能是一个很大的数字，因此需要将结果 对 10⁹ + 7 取余 后返回。
//
//
//
// 示例 1：
//
//
//
//
//输入：h = 5, w = 4, horizontalCuts = [1,2,4], verticalCuts = [1,3]
//输出：4
//解释：上图所示的矩阵蛋糕中，红色线表示水平和竖直方向上的切口。切割蛋糕后，绿色的那份蛋糕面积最大。
//
//
// 示例 2：
//
//
//
//
//输入：h = 5, w = 4, horizontalCuts = [3,1], verticalCuts = [1]
//输出：6
//解释：上图所示的矩阵蛋糕中，红色线表示水平和竖直方向上的切口。切割蛋糕后，绿色和黄色的两份蛋糕面积最大。
//
// 示例 3：
//
//
//输入：h = 5, w = 4, horizontalCuts = [3], verticalCuts = [3]
//输出：9
//
//
//
//
// 提示：
//
//
// 2 <= h, w <= 10⁹
// 1 <= horizontalCuts.length <= min(h - 1, 10⁵)
// 1 <= verticalCuts.length <= min(w - 1, 10⁵)
// 1 <= horizontalCuts[i] < h
// 1 <= verticalCuts[i] < w
// 题目数据保证 horizontalCuts 中的所有元素各不相同
// 题目数据保证 verticalCuts 中的所有元素各不相同
//
//
// Related Topics 贪心 数组 排序 👍 82 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	h, w := 5, 4
	horizontalCuts := []int{3, 1}
	verticalCuts := []int{1}
	h, w = 7, 3
	horizontalCuts = []int{3, 4, 2}
	verticalCuts = []int{2}
	area := maxArea(h, w, horizontalCuts, verticalCuts)
	fmt.Println(area)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	// 贪心
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	const mod = 1_000_000_007
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	maxL := func(arr []int, l int) int {
		max, last := 0, 0
		for _, i := range arr {
			max = maxVal(max, i-last)
			last = i
		}
		return maxVal(max, l-arr[len(arr)-1])
	}
	return (maxL(horizontalCuts, h) * maxL(verticalCuts, w)) % mod

	// 超时
	//maxVal := func(a, b int) int {
	//	if b > a {
	//		return b
	//	}
	//	return a
	//}
	//const mod = 1_000_000_007
	//sort.Ints(horizontalCuts)
	//sort.Ints(verticalCuts)
	//area, l, m, n := 0, 0, len(horizontalCuts), len(verticalCuts)
	//for i := 0; i < m; i++ {
	//	d, last := horizontalCuts[i]-l, 0
	//	for j := 0; j < n; j++ {
	//		area = maxVal(area, d*(verticalCuts[j]-last))
	//		last = verticalCuts[j]
	//	}
	//	area = maxVal(area, d*(w-last))
	//	l = horizontalCuts[i]
	//}
	//d, last := h-l, 0
	//for j := 0; j < n; j++ {
	//	area = maxVal(area, d*(verticalCuts[j]-last))
	//	last = verticalCuts[j]
	//}
	//area = maxVal(area, d*(w-last))
	//return area % mod
}

//leetcode submit region end(Prohibit modification and deletion)
