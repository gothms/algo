//你有两个果篮，每个果篮中有 n 个水果。给你两个下标从 0 开始的整数数组 basket1 和 basket2 ，用以表示两个果篮中每个水果的成本。
//
// 你希望两个果篮相等。为此，可以根据需要多次执行下述操作：
//
//
// 选中两个下标 i 和 j ，并交换 basket1 中的第 i 个水果和 basket2 中的第 j 个水果。
// 交换的成本是 min(basket1i,basket2j) 。
//
//
// 根据果篮中水果的成本进行排序，如果排序后结果完全相同，则认为两个果篮相等。
//
// 返回使两个果篮相等的最小交换成本，如果无法使两个果篮相等，则返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：basket1 = [4,2,2,2], basket2 = [1,4,1,2]
//输出：1
//解释：交换 basket1 中下标为 1 的水果和 basket2 中下标为 0 的水果，交换的成本为 1 。此时，basket1 = [4,1,2,2]
//且 basket2 = [2,4,1,2] 。重排两个数组，发现二者相等。
//
//
// 示例 2：
//
//
//输入：basket1 = [2,3,4,1], basket2 = [3,2,5,1]
//输出：-1
//解释：可以证明无法使两个果篮相等。
//
//
//
//
// 提示：
//
//
// basket1.length == bakste2.length
// 1 <= basket1.length <= 10⁵
// 1 <= basket1i,basket2i <= 10⁹
//
//
// Related Topics 贪心 数组 哈希表 👍 26 👎 0

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Println(-3 & 1)
	basket1 := []int{84, 80, 43, 8, 80, 88, 43, 14, 100, 88} // 48
	basket2 := []int{32, 32, 42, 68, 68, 100, 42, 84, 14, 8}
	cost := minCost(basket1, basket2)
	fmt.Println(cost)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minCost(basket1 []int, basket2 []int) int64 {
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	absVal := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	mc, in, cnt, n := int64(0), math.MaxInt32, 0, len(basket1)
	memo := make(map[int]int, n<<1)
	for i := 0; i < n; i++ {
		memo[basket1[i]]++ // 统计次数
		memo[basket2[i]]--
	}
	temp := make([][2]int, 0, len(memo))
	for k, v := range memo {
		in = minVal(in, k) // 记录最小值
		if v == 0 {        // 不需要交换
			continue
		}
		if v&1 == 1 { // 奇数无法平分
			return -1
		}
		c := absVal(v >> 1)
		cnt += c                          // 记录不相等的值的总数
		temp = append(temp, [2]int{k, c}) // 不相等的数集
	}
	sort.Slice(temp, func(i, j int) bool { // 排序
		return temp[i][0] < temp[j][0]
	})
	cnt >>= 1 // 交换次数为总数 or 总数的一半
	for i := 0; i < len(temp); i++ {
		c := minVal(cnt, temp[i][1])               // 以防超过交换次数
		mc += int64(minVal(in<<1, temp[i][0]) * c) // 使用最小值为交换媒介 or 直接交换
		cnt -= c
		if cnt == 0 { // 完成交换
			break
		}
	}
	return mc
}
	// 个人写法：但是还没考虑“已经不用交换的数，但是可以作为交换的媒介”
	//mc, n := int64(0), len(basket1)
	//minVal := func(a, b int) int {
	//	if b < a {
	//		return b
	//	}
	//	return a
	//}
	//memo := make(map[int]int, 2*n)
	//for i := 0; i < n; i++ {
	//	memo[basket1[i]]++
	//	memo[basket2[i]]--
	//}
	//l, r := make([][2]int, 0, len(memo)>>1), make([][2]int, 0, len(memo)>>1)
	//for k, v := range memo {
	//	if v&1 == 1 {
	//		return -1
	//	}
	//	if v > 0 {
	//		l = append(l, [2]int{k, v >> 1})
	//	} else if v < 0 {
	//		r = append(r, [2]int{k, -v >> 1})
	//	}
	//}
	//sort.Slice(l, func(i, j int) bool {
	//	return l[i][0] > l[j][0]
	//})
	//sort.Slice(r, func(i, j int) bool {
	//	return r[i][0] < r[j][0]
	//})
	//for il, ir := len(l)-1, len(r)-1; il >= 0; {
	//	in := minVal(l[il][0], r[ir][0])
	//	d := l[il][1] - r[ir][1]
	//	switch {
	//	case d == 0:
	//		mc += int64(in * l[il][1])
	//		il--
	//		ir--
	//	case d < 0:
	//		mc += int64(in * l[il][1])
	//		il--
	//	case d > 0:
	//		mc += int64(in * r[ir][1])
	//		ir--
	//	}
	//}
	//return mc
}

//leetcode submit region end(Prohibit modification and deletion)
