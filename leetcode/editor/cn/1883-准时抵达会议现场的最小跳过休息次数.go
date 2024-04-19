//给你一个整数 hoursBefore ，表示你要前往会议所剩下的可用小时数。要想成功抵达会议现场，你必须途经 n 条道路。道路的长度用一个长度为 n 的整数
//数组 dist 表示，其中 dist[i] 表示第 i 条道路的长度（单位：千米）。另给你一个整数 speed ，表示你在道路上前进的速度（单位：千米每小时）。
//
//
// 当你通过第 i 条路之后，就必须休息并等待，直到 下一个整数小时 才能开始继续通过下一条道路。注意：你不需要在通过最后一条道路后休息，因为那时你已经抵达会
//议现场。
//
//
// 例如，如果你通过一条道路用去 1.4 小时，那你必须停下来等待，到 2 小时才可以继续通过下一条道路。如果通过一条道路恰好用去 2 小时，就无需等待，可以
//直接继续。
//
//
// 然而，为了能准时到达，你可以选择 跳过 一些路的休息时间，这意味着你不必等待下一个整数小时。注意，这意味着与不跳过任何休息时间相比，你可能在不同时刻到达接
//下来的道路。
//
//
// 例如，假设通过第 1 条道路用去 1.4 小时，且通过第 2 条道路用去 0.6 小时。跳过第 1 条道路的休息时间意味着你将会在恰好 2 小时完成通过第
// 2 条道路，且你能够立即开始通过第 3 条道路。
//
//
// 返回准时抵达会议现场所需要的 最小跳过次数 ，如果 无法准时参会 ，返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：dist = [1,3,2], speed = 4, hoursBefore = 2
//输出：1
//解释：
//不跳过任何休息时间，你将用 (1/4 + 3/4) + (3/4 + 1/4) + (2/4) = 2.5 小时才能抵达会议现场。
//可以跳过第 1 次休息时间，共用 ((1/4 + 0) + (3/4 + 0)) + (2/4) = 1.5 小时抵达会议现场。
//注意，第 2 次休息时间缩短为 0 ，由于跳过第 1 次休息时间，你是在整数小时处完成通过第 2 条道路。
//
//
// 示例 2：
//
//
//输入：dist = [7,3,5,5], speed = 2, hoursBefore = 10
//输出：2
//解释：
//不跳过任何休息时间，你将用 (7/2 + 1/2) + (3/2 + 1/2) + (5/2 + 1/2) + (5/2) = 11.5 小时才能抵达会议现
//场。
//可以跳过第 1 次和第 3 次休息时间，共用 ((7/2 + 0) + (3/2 + 0)) + ((5/2 + 0) + (5/2)) = 10 小时抵达
//会议现场。
//
//
// 示例 3：
//
//
//输入：dist = [7,3,5,5], speed = 1, hoursBefore = 10
//输出：-1
//解释：即使跳过所有的休息时间，也无法准时参加会议。
//
//
//
//
// 提示：
//
//
// n == dist.length
// 1 <= n <= 1000
// 1 <= dist[i] <= 10⁵
// 1 <= speed <= 10⁶
// 1 <= hoursBefore <= 10⁷
//
//
// Related Topics 数组 动态规划 👍 43 👎 0

package main

import (
	"fmt"
)

func main() {
	dist := []int{1, 3, 2}
	speed := 4
	hoursBefore := 2 // 1
	dist = []int{7, 3, 5, 5}
	hoursBefore = 10 // 2
	speed = 2
	//dist = []int{1, 2, 3, 4, 5}
	//speed = 1000000
	//hoursBefore = 1 // 4
	skips := minSkips(dist, speed, hoursBefore)
	fmt.Println(skips)

	//x := 8
	//fmt.Println(x, ^x+1, ^x&x)
	//fmt.Println(x>>31+1, (-x)>>31+1)
	//fmt.Println(0>>31 + 1)
	//
	//fmt.Println(math.MinInt32 >> 31)
	//fmt.Println(-8 >> 31)
	//fmt.Println(-1 >> 31)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minSkips(dist []int, speed int, hoursBefore int) int {
	// dp：个人
	sum, n := 0, len(dist)
	for _, v := range dist {
		sum += v
	}
	h := (sum + speed - 1) / speed
	if h > hoursBefore { // 时间不够
		return -1
	} else if h == hoursBefore && sum%speed == 0 { // 特殊情况
		ret, s := 0, 0
		for _, v := range dist {
			if s += v; s%speed != 0 {
				ret++
				s %= speed
			}
		}
		return ret
	}
	// 时间够
	if speed == 1 || h-hoursBefore >= n-1 { // 跳不跳都一样 || 有充裕时间
		return 0
	}
	ret := 1
	//dp := make([][][2]int, n+1) // [2]：不休息、（不休息的）剩余长度、休息、（休息的）剩余长度
	//dp[0] = make([][2]int, n)
	//for i := 1; i <= n; i++ {
	//	dp[i] = make([][2]int, n)
	//	for j := ret; j < n; j++ {
	//		temp := [2]int{dp[i-1][j][0] + dist[i-1]/speed, dist[i-1] % speed} // 同次数 j
	//		if dp[i-1][j][1] > 0 {
	//			temp[0]++
	//		}
	//		if j > 1 {
	//			s := dp[i-1][j-1][1] + dist[i-1]
	//			//fmt.Println(i, j, s)
	//			a, b := dp[i-1][j-1][0]+s/speed, s%speed // 次数 j-1
	//			if a < temp[0] || a == temp[0] && b < temp[1] {
	//				temp[0], temp[1] = a, b
	//			}
	//		}
	//		if temp[0] > hoursBefore || temp[0] == hoursBefore && temp[1] > 0 {
	//			ret = j + 1
	//		}
	//		dp[i][j] = temp
	//	}
	//}
	// 简化
	dp := make([][2]int, n+1) // [2]：耗时、剩余长度
	for i := 1; i <= n; i++ {
		for j := n - 1; j >= ret; j-- {
			temp := [2]int{dp[j][0] + dist[i-1]/speed, dist[i-1] % speed} // 同次数 j
			if dp[j][1] > 0 {
				temp[0]++
			}
			if j > 1 { // 至少跳过 1 次
				s := dp[j-1][1] + dist[i-1]
				a, b := dp[j-1][0]+s/speed, s%speed // 次数 j-1
				if a < temp[0] || a == temp[0] && b < temp[1] {
					temp[0], temp[1] = a, b
				}
			}
			//if temp[0] > hoursBefore || temp[0] == hoursBefore && temp[1] > 0 {
			//	ret = j + 1 // 跳过的次数太少，导致超时
			//}
			if temp[0]+(temp[1]-1)>>31+1 > hoursBefore { // (temp[1]-1)>>31+1：正数则 1，0 则 0
				ret = j + 1 // 跳过的次数太少，导致超时
			}
			dp[j] = temp
		}
	}
	return ret - 1

	// lc
	//n, s, total := len(dist), int64(speed), int64(speed)*int64(hoursBefore)
	//dp := make([]int64, n+1)  // 记录“总里程” = 时间 * 速度
	//for i := 1; i <= n; i++ { // dist 遍历
	//	dp[i] = dp[i-1] + int64(dist[i-1]) // 存 dist[i+1] 的“原始总数”，用于下一次跳过
	//	for j := i - 1; j > 0; j-- {
	//		// 不跳过，次数 +0 / 跳过，次数 +1
	//		// 重点：(dp[j]+int64(dist[i-1])+s-1)/s*s
	//		// 因为不跳时：dp[j-1]+int64(dist[i-1])，还没跳
	//		/*
	//			[8 7 0 0 0]
	//			[12 10 10 0 0]
	//			[18 16 15 15 0]
	//			[24 22 20 20 20]
	//		*/
	//		dp[j] = min((dp[j]+int64(dist[i-1])+s-1)/s*s, dp[j-1]+int64(dist[i-1]))
	//		//dp[j] = min((dp[j]+s-1)/s*s, dp[j-1]) + int64(dist[i-1])	// 错误
	//	}
	//	dp[0] += int64(dist[i-1]+speed-1) / s * s // 永远不跳过
	//}
	//for i, v := range dp {
	//	if v <= total {
	//		return i
	//	}
	//}
	//return -1
}

//leetcode submit region end(Prohibit modification and deletion)
