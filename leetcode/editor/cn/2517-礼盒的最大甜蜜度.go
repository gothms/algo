//给你一个正整数数组 price ，其中 price[i] 表示第 i 类糖果的价格，另给你一个正整数 k 。
//
// 商店组合 k 类 不同 糖果打包成礼盒出售。礼盒的 甜蜜度 是礼盒中任意两种糖果 价格 绝对差的最小值。
//
// 返回礼盒的 最大 甜蜜度。
//
//
//
// 示例 1：
//
//
//输入：price = [13,5,1,8,21,2], k = 3
//输出：8
//解释：选出价格分别为 [13,5,21] 的三类糖果。
//礼盒的甜蜜度为 min(|13 - 5|, |13 - 21|, |5 - 21|) = min(8, 8, 16) = 8 。
//可以证明能够取得的最大甜蜜度就是 8 。
//
//
// 示例 2：
//
//
//输入：price = [1,3,1], k = 2
//输出：2
//解释：选出价格分别为 [1,3] 的两类糖果。
//礼盒的甜蜜度为 min(|1 - 3|) = min(2) = 2 。
//可以证明能够取得的最大甜蜜度就是 2 。
//
//
// 示例 3：
//
//
//输入：price = [7,7,7,7], k = 2
//输出：0
//解释：从现有的糖果中任选两类糖果，甜蜜度都会是 0 。
//
//
//
//
// 提示：
//
//
// 1 <= price.length <= 10⁵
// 1 <= price[i] <= 10⁹
// 2 <= k <= price.length
//
//
// Related Topics 数组 二分查找 排序 👍 36 👎 0

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	price := []int{13, 5, 1, 8, 21, 2}
	k := 3
	//price = []int{1, 3, 1}
	//k = 2
	//price = []int{7, 7, 7, 7}
	//k = 2
	//price = []int{34, 116, 83, 15, 150, 56, 69, 42, 26}
	//k = 6 // 19
	tastiness := maximumTastiness(price, k)
	fmt.Println("res:", tastiness)

	fmt.Println(int(math.Inf(-1)) >> 1)
}

/*
思路：二分
	1.price[i] 的 最大/最小值 分别是 max/min，那么两个糖果的
		1.1.最大差值，不会超过 max-min
		1.2.最小差值，直接假设为 0（也可以遍历一次 price，找出最小的两个数求出）
	2.那么可以在 [0,max-min] 这个区间内，使用 二分查找 得出 k 个糖果的 最大 甜蜜度
		2.1.将 price 排序
		2.2.第一次二分：找出 0 到 max-min 之间一个数 v，这个数是 k 个糖果 最大 甜蜜度
		2.3.第二次二分：遍历 price，找出满足差值为 v 的数一共有几个，返回给 第一次二分
代码注释：
	1.排序
	2.二分
		2.1.查找 最大 甜蜜度
		2.2.甜蜜度为 v 时，找出满足差值为 v 的 price[i] 总数
			a)查找条件是 price[n] >= v：最后一个元素 >= 下一次要查找的糖果值
			b)price[i] 为当前满足差值为 v 的糖果，所以二分时查找范围为 (i,n]
		2.3.修正2.1.查找到的值（采用标准库，也说明有时标准库不能满足需求，也可以自己写二分）
			需求是查找 最后一个小于等于 v 的数，而 go 标准库提供的二分是查找 大于等于 v 的数
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maximumTastiness(price []int, k int) int {
	//sort.Ints(price)
	//n := len(price) - 1
	//l, r := 0, price[n]-price[0]
	//for l < r {
	//	m, cnt := int(uint(l+r+1)>>1), 1
	//	//m, cnt, i := (l+r)>>1, 1, 0
	//	for i, v := 0, price[0]+m; price[n] >= v; v = price[i] + m {
	//		i += sort.Search(n-i, func(idx int) bool {
	//			return price[idx+i] >= v
	//		})
	//		cnt++
	//	}
	//	if cnt >= k {
	//		l = m
	//	} else {
	//		r = m - 1
	//	}
	//}
	//return l

	sort.Ints(price) // 1：快排空间 O(log n)
	n := len(price) - 1
	counter := func(m int) bool { // 2.2：O(k log n)
		cnt := 1
		for i, v := 0, price[0]+m; price[n] >= v; v = price[i] + m { // 2.2.a
			i += sort.Search(n-i, func(idx int) bool {
				return price[idx+i] >= v
			}) // 2.2.b
			cnt++
		}
		return cnt < k
	}
	max := sort.Search(price[n]-price[0], func(m int) bool { // 2
		return m > 0 && counter(m) // 2.1：O(log C) C=max-min
	})
	if max > 0 && counter(max) { // 2.3
		return max - 1
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
