//给你一个数组 seats 表示一排座位，其中 seats[i] = 1 代表有人坐在第 i 个座位上，seats[i] = 0 代表座位 i 上是空的（下标
//从 0 开始）。
//
// 至少有一个空座位，且至少有一人已经坐在座位上。
//
// 亚历克斯希望坐在一个能够使他与离他最近的人之间的距离达到最大化的座位上。
//
// 返回他到离他最近的人的最大距离。
//
//
//
// 示例 1：
//
//
//输入：seats = [1,0,0,0,1,0,1]
//输出：2
//解释：
//如果亚历克斯坐在第二个空位（seats[2]）上，他到离他最近的人的距离为 2 。
//如果亚历克斯坐在其它任何一个空位上，他到离他最近的人的距离为 1 。
//因此，他到离他最近的人的最大距离是 2 。
//
//
// 示例 2：
//
//
//输入：seats = [1,0,0,0]
//输出：3
//解释：
//如果亚历克斯坐在最后一个座位上，他离最近的人有 3 个座位远。
//这是可能的最大距离，所以答案是 3 。
//
//
// 示例 3：
//
//
//输入：seats = [0,1]
//输出：1
//
//
//
//
// 提示：
//
//
// 2 <= seats.length <= 2 * 10⁴
// seats[i] 为 0 或 1
// 至少有一个 空座位
// 至少有一个 座位上有人
//
//
// Related Topics 数组 👍 200 👎 0

package main

func main() {

}

/*
思路：双指针
	1.选择的座位分两种情况：
		坐中间时：距离 = (连续0长度 + 1) / 2
		靠两边墙：距离 = 连续0长度
	2.双指针 (l,i) 表示连续 0 的区间
		l：上一个 1 的位置
		i：找到下一个 1 的位置
*/
// leetcode submit region begin(Prohibit modification and deletion)
func maxDistToClosest(seats []int) int {
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxDTC, l, n := 0, -1, len(seats)
	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			if l < 0 {
				maxDTC = i // 贴左墙
			} else {
				maxDTC = maxVal(maxDTC, (i-l)>>1) // 中间
			}
			l = i
		}
	}
	return maxVal(maxDTC, n-l-1) // 贴右墙
}
	//maxDTC, i, j := 0, 0, len(seats)-1
	//for seats[i] == 0 {
	//	i++
	//}
	//for seats[j] == 0 {
	//	j--
	//}
	//maxDTC = maxVal(i, len(seats)-j-1) << 1
	//for next := i + 1; next < j; next++ {
	//	for seats[next] == 0 {
	//		next++
	//	}
	//	maxDTC, i = maxVal(maxDTC, next-i-1), next
	//}
	//return (maxDTC + 1) >> 1
}

//leetcode submit region end(Prohibit modification and deletion)
