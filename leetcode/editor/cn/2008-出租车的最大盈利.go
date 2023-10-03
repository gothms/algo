//你驾驶出租车行驶在一条有 n 个地点的路上。这 n 个地点从近到远编号为 1 到 n ，你想要从 1 开到 n ，通过接乘客订单盈利。你只能沿着编号递增的方
//向前进，不能改变方向。
//
// 乘客信息用一个下标从 0 开始的二维数组 rides 表示，其中 rides[i] = [starti, endi, tipi] 表示第 i 位乘客需要从
//地点 starti 前往 endi ，愿意支付 tipi 元的小费。
//
// 每一位 你选择接单的乘客 i ，你可以 盈利 endi - starti + tipi 元。你同时 最多 只能接一个订单。
//
// 给你 n 和 rides ，请你返回在最优接单方案下，你能盈利 最多 多少元。
//
// 注意：你可以在一个地点放下一位乘客，并在同一个地点接上另一位乘客。
//
//
//
// 示例 1：
//
// 输入：n = 5, rides = [[2,5,4],[1,5,1]]
//输出：7
//解释：我们可以接乘客 0 的订单，获得 5 - 2 + 4 = 7 元。
//
//
// 示例 2：
//
// 输入：n = 20, rides = [[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]
//
//输出：20
//解释：我们可以接以下乘客的订单：
//- 将乘客 1 从地点 3 送往地点 10 ，获得 10 - 3 + 2 = 9 元。
//- 将乘客 2 从地点 10 送往地点 12 ，获得 12 - 10 + 3 = 5 元。
//- 将乘客 5 从地点 13 送往地点 18 ，获得 18 - 13 + 1 = 6 元。
//我们总共获得 9 + 5 + 6 = 20 元。
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁵
// 1 <= rides.length <= 3 * 10⁴
// rides[i].length == 3
// 1 <= starti < endi <= n
// 1 <= tipi <= 10⁵
//
//
// Related Topics 数组 二分查找 动态规划 排序 👍 76 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	rides := [][]int{{1, 6, 1},
		{3, 10, 2},
		{10, 12, 3},
		{11, 12, 2},
		{12, 15, 2},
		{13, 18, 1}}
	n := 20
	rides = [][]int{{2, 3, 6}, // 33
		{8, 9, 8},
		{5, 9, 7},
		{8, 9, 1},
		{2, 9, 2},
		{9, 10, 6},
		{7, 10, 10},
		{6, 7, 9},
		{4, 9, 7},
		{2, 3, 1}}
	//rides = [][]int{{9, 10, 2}, // 22
	//	{4, 5, 6},
	//	{6, 8, 1},
	//	{1, 5, 5},
	//	{4, 9, 5},
	//	{1, 6, 5},
	//	{4, 8, 3},
	//	{4, 7, 10},
	//	{1, 9, 8},
	//	{2, 3, 5}}
	earnings := maxTaxiEarnings(n, rides)
	fmt.Println(earnings)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxTaxiEarnings(n int, rides [][]int) int64 {
	// dp
	maxVal := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}
	dp, path := make(map[int]int64, len(rides)), make([]int, 1, len(rides)+1)
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][1] < rides[j][1] // 按下车地点排序
		//return rides[i][0] < rides[j][0] ||
		//	rides[i][0] == rides[j][0] && rides[i][1] < rides[j][1]
	})
	for _, r := range rides {
		i := sort.SearchInts(path, r[0]) // path 记录下车地点
		last := len(path) - 1
		if i > last || path[i] != r[0] { // i 越界 / 上车点第一次出现
			i--
		}
		if r[1] != path[last] { // 避免重复记录下车地点
			path = append(path, r[1])
		}
		v := maxVal(dp[path[i]]+int64(r[1]-r[0]+r[2]), dp[path[last]]) // 载客 / 不载客
		dp[r[1]] = maxVal(dp[r[1]], v)                                 // 历史最大值
	}
	return dp[path[len(path)-1]]
}

//leetcode submit region end(Prohibit modification and deletion)
