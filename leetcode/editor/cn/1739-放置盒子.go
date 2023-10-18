//有一个立方体房间，其长度、宽度和高度都等于 n 个单位。请你在房间里放置 n 个盒子，每个盒子都是一个单位边长的立方体。放置规则如下：
//
//
// 你可以把盒子放在地板上的任何地方。
// 如果盒子 x 需要放置在盒子 y 的顶部，那么盒子 y 竖直的四个侧面都 必须 与另一个盒子或墙相邻。
//
//
// 给你一个整数 n ，返回接触地面的盒子的 最少 可能数量。
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 3
//输出：3
//解释：上图是 3 个盒子的摆放位置。
//这些盒子放在房间的一角，对应左侧位置。
//
//
// 示例 2：
//
//
//
//
//输入：n = 4
//输出：3
//解释：上图是 3 个盒子的摆放位置。
//这些盒子放在房间的一角，对应左侧位置。
//
//
// 示例 3：
//
//
//
//
//输入：n = 10
//输出：6
//解释：上图是 10 个盒子的摆放位置。
//这些盒子放在房间的一角，对应后方位置。
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁹
//
//
// Related Topics 贪心 数学 二分查找 👍 91 👎 0

package main

import (
	"fmt"
)

func main() {
	n := 10
	boxes := minimumBoxes(n)
	fmt.Println(boxes)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumBoxes(n int) int {
	// lc：https://leetcode.cn/problems/building-boxes/solutions/2031813/mei-xiang-ming-bai-yi-ge-dong-hua-miao-d-8vbe/
	ret, maxN := 0, 0 // maxN：盒子上限
	for i := 1; maxN+ret+i <= n; i++ {
		ret += i    // 1+2+3+⋯+i
		maxN += ret //  1+(1+2)+(1+2+3)+⋯+ i(i+1)/2
	}
	for j := 1; maxN < n; j++ {
		ret++
		maxN += j
	}
	return ret

	// O(1)
	//x := int(math.Cbrt(float64(6 * n))) // 开立方
	//ret := x * (x + 1) >> 1
	//maxN := x * (x + 1) * (x + 2) / 6
	//if maxN > n { // 修正 maxN
	//	maxN -= ret
	//	ret -= x
	//}
	//return ret + int(math.Ceil((-1+math.Sqrt(float64(1+8*(n-maxN))))/2))
}

//leetcode submit region end(Prohibit modification and deletion)
