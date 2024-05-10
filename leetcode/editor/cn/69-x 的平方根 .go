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

import (
	"fmt"
	"math"
)

func main() {
	for i := 100; i < 199; i++ {
		fmt.Println(i, mySqrt(i))
	}

	var x = 7.34680654489
	k := sqrtK(x) // 2.7104993166201794 / 2.710499
	fmt.Println(k)
}

/*
袖珍计算器算法：
	根号 x = x 的 1/2 次方 = e 的 lnx 次方的 1/2 次方 = e 的 1/2乘lnx 次方
	https://leetcode-cn.com/problems/sqrtx/solution/x-de-ping-fang-gen-by-leetcode-solution/
	https://baike.baidu.com/item/IEEE%20754
二分法：
牛顿迭代法：
	http://www.matrix67.com/blog/archives/361
	http://www.voidcn.com/article/p-eudisdmk-zm.html
Magic：
	http://lomont.org/papers/2003/InvSqrt.pdf
	https://handwiki.org/wiki/Fast_inverse_square_root
*/
//leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {

}

// leetcode submit region end(Prohibit modification and deletion)

// 牛顿迭代法，精确到小数点后 6 位
func sqrtK(x float64) float64 {
	g := x
	d := 1e-6
	for g*g > x+d {
		g = (g + x/g) / 2 // 二分
	}
	return g

	// 二分
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

	// api
	//return math.Trunc(g*1e6+0.5) * d

	//return int(InvSqrt(float32(x)))	// 太精准，提交不成功
}

func InvSqrt(x float32) float32 {
	var xhalf float32 = 0.5 * x // get bits for floating VALUE
	i := math.Float32bits(x)    // gives initial guess y0
	i = 0x5f375a86 - (i >> 1)   // convert bits BACK to float
	x = math.Float32frombits(i) // Newton step, repeating increases accuracy
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)
	return 1 / x
}

/*
16 -> 4.0000005
16 -> 4.006779
*/
func invSqrt(x float32) float32 { // 使用 float64 会运算成 0
	xhalf := 0.5 * x            // get bits for floating VALUE
	i := math.Float32bits(x)    // gives initial guess y0
	i = 0x5f375a86 - (i >> 1)   // convert bits BACK to float	(or i = 0x5f3759df - (i >> 1))
	x = math.Float32frombits(i) // Newton step, repeating increases accuracy
	x = x * (1.5 - xhalf*x*x)   // 运行3次，会非常精确
	return 1 / x
}
