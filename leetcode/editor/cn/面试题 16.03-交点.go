//给定两条线段（表示为起点start = {X1, Y1}和终点end = {X2, Y2}），如果它们有交点，请计算其交点，没有交点则返回空值。
//
// 要求浮点型误差不超过10^-6。若有多个交点（线段重叠）则返回 X 值最小的点，X 坐标相同则返回 Y 值最小的点。
//
//
//
// 示例 1：
//
// 输入：
//line1 = {0, 0}, {1, 0}
//line2 = {1, 1}, {0, -1}
//输出： {0.5, 0}
//
//
// 示例 2：
//
// 输入：
//line1 = {0, 0}, {3, 3}
//line2 = {1, 1}, {2, 2}
//输出： {1, 1}
//
//
// 示例 3：
//
// 输入：
//line1 = {0, 0}, {1, 1}
//line2 = {1, 0}, {2, 1}
//输出： {}，两条线段没有交点
//
//
//
//
// 提示：
//
//
// 坐标绝对值不会超过 2^7
// 输入的坐标均是有效的二维坐标
//
//
// Related Topics 几何 数学 👍 84 👎 0

package main

import "fmt"

func main() {
	//fmt.Println(-1 ^ 3)
	//fmt.Println(-1 ^ -3)
	//fmt.Println(1 ^ -3)
	//fmt.Println(1 ^ 3)
	//fmt.Println(0 ^ -3)

	start1 := []int{0, 0}
	end1 := []int{-1, 1}
	start2 := []int{0, 0}
	end2 := []int{1, -1}
	float64s := intersection(start1, end1, start2, end2)
	fmt.Println(float64s)
}

// leetcode submit region begin(Prohibit modification and deletion)
func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	// 数学
	// 平行：判断平行、判断重叠、判断有交点
	// 相交：求交点（斜截式、截距式、参数方程式）、判断交点在线段上
	x1, y1 := start1[0], start1[1]
	x2, y2 := end1[0], end1[1]
	x3, y3 := start2[0], start2[1]
	x4, y4 := end2[0], end2[1]
	ret := make([]float64, 0, 2)
	in := func(x1, y1, x2, y2, xt, yt int) bool {
		// 线段与 x、y 轴平行 or xt, yt 在线段上
		return (x1 == x2 || (x1-xt)*(x2-xt) <= 0) && (y1 == y2 || (y1-yt)*(y2-yt) <= 0)
	}
	up := func(x, y float64) { // 多个交点（线段重叠）则返回 X 值最小的点，X 坐标相同则返回 Y 值最小的点
		if len(ret) == 0 || x < ret[0] || x == ret[0] && y < ret[1] {
			ret = append(ret[:0], x, y)
		}
	}
	if (y2-y1)*(x4-x3) == (y4-y3)*(x2-x1) { // 平行
		if (y3-y1)*(x2-x1) == (y2-y1)*(x3-x1) { // 重叠
			if in(x1, y1, x2, y2, x3, y3) {
				ret = append(ret, float64(x3), float64(y3))
			}
			if in(x1, y1, x2, y2, x4, y4) {
				up(float64(x4), float64(y4))
			}
			if in(x3, y3, x4, y4, x1, y1) {
				up(float64(x1), float64(y1))
			}
			if in(x3, y3, x4, y4, x2, y2) {
				up(float64(x2), float64(y2))
			}
		}
	} else { // 非平行
		// 求交点：x = λx1 + (1-λ)x2，且 y = λy1 + (1-λ)y2
		// x1 + t1(x2 - x1) = x3 + t2(x4 - x3)
		// y1 + t1(y2 - y1) = y3 + t2(y4 - y3)
		t1 := float64(x3*(y4-y3)+y1*(x4-x3)-y3*(x4-x3)-x1*(y4-y3)) / float64((x2-x1)*(y4-y3)-(x4-x3)*(y2-y1))
		t2 := float64(x1*(y2-y1)+y3*(x2-x1)-y1*(x2-x1)-x3*(y2-y1)) / float64((x4-x3)*(y2-y1)-(x2-x1)*(y4-y3))
		// 交点是否在线段上
		if t1 >= 0.0 && t1 <= 1.0 && t2 >= 0.0 && t2 <= 1.0 {
			ret = append(ret, float64(x1)+t1*float64(x2-x1), float64(y1)+t1*float64(y2-y1))
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
