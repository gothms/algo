//给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
//
// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
//
// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
//
//
//
// 示例 1：
//
//
//输入：x = 4
//输出：2
//
//
// 示例 2：
//
//
//输入：x = 8
//输出：2
//解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
//
//
//
//
// 提示：
//
//
// 0 <= x <= 2³¹ - 1
//
//
// Related Topics 数学 二分查找 👍 1343 👎 0

package main

import "fmt"

func main() {
	for i := 100; i < 199; i++ {
		fmt.Println(i, mySqrt(i))
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	g := x
	for g*g > x {
		g = (g + x/g) / 2
	}
	return g

	// 二分
	//l, r := 0, x
	//for l <= r {
	//	mid := (l + r) >> 1
	//	if v := mid * mid; v == x {
	//		return mid
	//	} else if v < x {
	//		l = mid + 1
	//	} else {
	//		r = mid - 1
	//	}
	//}
	//return r

	//l, r, sqrt := 0, x, 1
	//for l <= r {
	//	mid := (l + r) >> 1
	//	if mid*mid <= x {
	//		sqrt, l = mid, mid+1
	//	} else {
	//		r = mid - 1
	//	}
	//}
	//return sqrt
}

//leetcode submit region end(Prohibit modification and deletion)
