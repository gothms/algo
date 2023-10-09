//给你两个数组 arr1 和 arr2 ，它们一开始都是空的。你需要往它们中添加正整数，使它们满足以下条件：
//
//
// arr1 包含 uniqueCnt1 个 互不相同 的正整数，每个整数都 不能 被 divisor1 整除 。
// arr2 包含 uniqueCnt2 个 互不相同 的正整数，每个整数都 不能 被 divisor2 整除 。
// arr1 和 arr2 中的元素 互不相同 。
//
//
// 给你 divisor1 ，divisor2 ，uniqueCnt1 和 uniqueCnt2 ，请你返回两个数组中 最大元素 的 最小值 。
//
//
//
// 示例 1：
//
//
//输入：divisor1 = 2, divisor2 = 7, uniqueCnt1 = 1, uniqueCnt2 = 3
//输出：4
//解释：
//我们可以把前 4 个自然数划分到 arr1 和 arr2 中。
//arr1 = [1] 和 arr2 = [2,3,4] 。
//可以看出两个数组都满足条件。
//最大值是 4 ，所以返回 4 。
//
//
// 示例 2：
//
//
//输入：divisor1 = 3, divisor2 = 5, uniqueCnt1 = 2, uniqueCnt2 = 1
//输出：3
//解释：
//arr1 = [1,2] 和 arr2 = [3] 满足所有条件。
//最大值是 3 ，所以返回 3 。
//
// 示例 3：
//
//
//输入：divisor1 = 2, divisor2 = 4, uniqueCnt1 = 8, uniqueCnt2 = 2
//输出：15
//解释：
//最终数组为 arr1 = [1,3,5,7,9,11,13,15] 和 arr2 = [2,6] 。
//上述方案是满足所有条件的最优解。
//
//
//
//
// 提示：
//
//
// 2 <= divisor1, divisor2 <= 10⁵
// 1 <= uniqueCnt1, uniqueCnt2 < 10⁹
// 2 <= uniqueCnt1 + uniqueCnt2 <= 10⁹
//
//
// Related Topics 数学 二分查找 数论 👍 39 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	divisor1, divisor2 := 2, 4
	uniqueCnt1, uniqueCnt2 := 8, 2 // 15
	divisor1, divisor2 = 6, 4
	uniqueCnt1, uniqueCnt2 = 18, 27 // 49
	i := minimizeSet(divisor1, divisor2, uniqueCnt1, uniqueCnt2)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	// O(1)：lc
	// [0,m] 内有 m/d 个数被 d 整除，则有 u = m+1-m/d，有 m = (u-1)*d / (d-1)
	// 那么从 1 开始，所求 m = f(d,k) = (u-1)*d / (d-1) + 1
	//v1, v2 := divisor1, divisor2
	//for v1 != 0 { // 最大公约数
	//	v1, v2 = v2%v1, v1
	//}
	//div := divisor2 / v2 * divisor1 // 最小公倍数
	//return maxVal(div*(uniqueCnt1+uniqueCnt2-1)/(div-1)+1,
	//	maxVal(divisor1*(uniqueCnt1-1)/(divisor1-1)+1, divisor2*(uniqueCnt2-1)/(divisor2-1)+1))

	// 二分：https://leetcode.cn/problems/minimize-the-maximum-of-two-arrays/solutions/2031827/er-fen-da-an-by-endlesscheng-y8fp/
	v1, v2 := divisor1, divisor2
	for v1 != 0 { // 最大公约数
		v1, v2 = v2%v1, v1
	}
	lcm := divisor1 / v2 * divisor2 // 最小公倍数
	return sort.Search((uniqueCnt1+uniqueCnt2)<<1-1, func(i int) bool {
		v := i / lcm
		left1 := maxVal(uniqueCnt1-i/divisor2+v, 0) // i/divisor2 - v：被 left1 独享的数
		left2 := maxVal(uniqueCnt2-i/divisor1+v, 0)
		common := i - i/divisor1 - i/divisor2 + v // 区间 i 内不能被 d1 和 d2 整除的数
		return common >= left1+left2
	})

	// 个人写法：O(1)
	//v1, v2 := divisor1, divisor2
	//for v1 != 0 { // 最大公约数
	//	v1, v2 = v2%v1, v1
	//}
	//div := divisor2 / v2 * divisor1 // 最小公倍数
	//f := func(d, u int) int {       // [1,t] 区间有 u 个数，不是 d 的倍数
	//	t := u / (d - 1) * d          // u / (d - 1)：多少次轮询 [1,d)，其中，此时 t 能整除 d
	//	if v := u % (d - 1); v == 0 { // 修正 t：余数为 0
	//		return t - 1
	//	} else {
	//		return u/(d-1)*d + v // 修正 t：加上余数
	//	}
	//}
	//return maxVal(f(div, uniqueCnt1+uniqueCnt2), maxVal(f(divisor1, uniqueCnt1), f(divisor2, uniqueCnt2)))

	// 以下求的是：既不被 divisor1 整除，也不被 divisor2 整除的数，最大元素的最小值
	//var d int
	//if divisor1 < divisor2 {
	//	d = divisor1
	//} else {
	//	d = divisor2
	//}
	//k := uniqueCnt1 + uniqueCnt2
	//n := (k/(d-1)*d + k%(d-1)) << 1
	//fmt.Println(n)
	//return sort.Search(n, func(i int) bool {
	//	return i-i/divisor1-i/divisor2+i/(divisor1*divisor2) >= k
	//})
}

//leetcode submit region end(Prohibit modification and deletion)
