//给你一个下标从 0 开始的整数数组 nums ，它表示英雄的能力值。如果我们选出一部分英雄，这组英雄的 力量 定义为：
//
//
// i0 ，i1 ，... ik 表示这组英雄在数组中的下标。那么这组英雄的力量为 max(nums[i0],nums[i1] ... nums[ik])²
//* min(nums[i0],nums[i1] ... nums[ik]) 。
//
//
// 请你返回所有可能的 非空 英雄组的 力量 之和。由于答案可能非常大，请你将结果对 109 + 7 取余。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,1,4]
//输出：141
//解释：
//第 1 组：[2] 的力量为 2² * 2 = 8 。
//第 2 组：[1] 的力量为 1² * 1 = 1 。
//第 3 组：[4] 的力量为 4² * 4 = 64 。
//第 4 组：[2,1] 的力量为 2² * 1 = 4 。
//第 5 组：[2,4] 的力量为 4² * 2 = 32 。
//第 6 组：[1,4] 的力量为 4² * 1 = 16 。
//第​ ​​​​​​7 组：[2,1,4] 的力量为 4²​​​​​​​ * 1 = 16 。
//所有英雄组的力量之和为 8 + 1 + 64 + 4 + 32 + 16 + 16 = 141 。
//
//
// 示例 2：
//
//
//输入：nums = [1,1,1]
//输出：7
//解释：总共有 7 个英雄组，每一组的力量都是 1 。所以所有英雄组的力量之和为 7 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁹
//
//
// Related Topics 数组 数学 前缀和 排序 👍 71 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 1, 4}
	// 567530242
	nums = []int{658, 489, 777, 2418, 1893, 130, 2448, 178, 1128, 2149, 1059, 1495, 1166, 608, 2006, 713, 1906, 2108, 680, 1348, 860, 1620, 146, 2447, 1895, 1083, 1465, 2351, 1359, 1187, 906, 533, 1943, 1814, 1808, 2065, 1744, 254, 1988, 1889, 1206}
	// 297094082
	nums = []int{8672, 8685, 7500, 3692, 9877, 1057, 6851, 226, 4371, 9418, 8597, 7707, 8706, 1801, 816, 1875, 4628, 8471, 6966, 9095, 8561, 7782, 4463, 4316, 6540, 5297, 4206, 8532, 202, 4783, 9277, 8976, 9770, 5404, 4034, 6134, 9585, 1250, 685, 2731, 7337, 2432, 8281, 5573, 4508}
	// 100279694
	nums = []int{1705, 7581, 5816, 8226, 3173, 2381, 9562, 5447, 8770, 2247, 2106, 6925, 4347, 4651, 3785, 7013, 4826, 1260, 9151, 5321, 6521, 2984, 2553, 6035, 4095, 9021, 8296, 7682, 4071, 2830, 4182, 1994, 9222, 5343, 6826, 330, 8214, 5657, 425, 7030, 9074, 4016, 1393, 4598, 9563, 2811, 9156, 9399, 140, 5627, 6590, 9980, 6620, 2461, 6213, 3966, 301, 6254, 8024, 3508, 2979, 1707, 4293, 3789, 7349, 728, 5802, 4223, 2195, 9914, 2294, 2551, 8221, 3807, 3860, 4967, 2506, 3086, 7531, 6922, 8546, 5636, 6249}
	// 711909510
	nums = []int{625006, 846432, 764290, 653039}
	power := sumOfPower(nums)
	fmt.Println(power)
	//arr := []int{1, 2, 3, 4, 5}

}

/*
思路：dp
	1.有序
		求所有英雄组的力量和，也就是求 nums 的所有子集（除去空子集）的力量和
		假设 nums 是有序的，它的任意子集的最小值是 sub[0]，最大值是 sub[-1]
	2.子集
		假设 nums = [1,2,3,4]，可以列出它所有子集，根据子集的数学公式可知：
		max = 4 的子集有 8 个，这 8 个子集中，最小值情况：
			min = 1 的子集有 4 个
			min = 2：2
			min = 3：1
			min = 4：1（先不考虑这种单个元素的子集）
	3.公式思路
		由上可得，对于所有以 nums[i] 为最大值的子集，它们的力量和：
		power[i] = nums[i]*nums[i] * (1<<(i-1)*nums[0] + 1<<(i-2)*nums[1] + ... + 1*nums[i-1])
		设 (1<<(i-1)*nums[0] + 1<<(i-2)*nums[1] + ... + 1*nums[i-1]) = times[i-1]
		则 power[i+1] = nums[i+1]*nums[i+1] * (times[i-1]<<1 + nums[i])
	4.状态转移方程
		补上：min = 4：1（先不考虑这种单个元素的子集）
		dp[i] = (times[i-1]+nums[i]) * nums[i]*nums[i]
		times[i] = times[i-1]<<1 + nums[i]
*/
// leetcode submit region begin(Prohibit modification and deletion)
func sumOfPower(nums []int) int {
	// dp：终版
	const mod = 1000000007
	dp, times, n := 0, 0, len(nums)

	//power := make([]int, n)
	sort.Ints(nums) // 排序
	//for i := 0; i < n; i++ {
	//	power[i] = nums[i] * nums[i] % mod
	//}
	for i := 0; i < n; i++ {
		//dp = ((times+nums[i])*power[i] + dp) % mod
		dp = (nums[i]*nums[i]%mod*(times+nums[i]) + dp) % mod // 当前力量和
		times = (times<<1 + nums[i]) % mod                    // 下一轮 dp 的因子
	}
	return dp

	// dp：得和lc流程一样，才能保证 mod 顺序
	//const mod = 1000000007
	////sum, n := 0, len(nums)
	//n := len(nums)
	//power := make([]int, n)
	//sort.Ints(nums)
	//for i := 0; i < n; i++ {
	//	power[i] = nums[i] * nums[i] % mod
	//	//sum = (power[i]*nums[i]%mod + sum) % mod
	//}
	////sum -= power[0]
	//dp, times := power[0], 0
	//for i := 1; i < n; i++ {
	//	times = (times<<1%mod + nums[i-1]) % mod
	//	dp = (times*power[i]%mod + dp) % mod
	//}
	//for i := 1; i < n; i++ {
	//	dp = (power[i]*nums[i]%mod + dp) % mod
	//}
	////return (dp + sum) % mod
	//return dp

	// lc
	//n := len(nums)
	//sort.Ints(nums)
	//dp := 0
	//preSum := 0
	//res := 0
	//mod := int(1e9 + 7)
	//for i := 0; i < n; i++ {
	//	dp = (nums[i] + preSum) % mod
	//	preSum = (preSum + dp) % mod
	//	res = (res + (nums[i]*nums[i]%mod*dp)%mod) % mod
	//}
	//return res

	// math：个人
	//const mod = 1000000007
	//sum, n := 0, len(nums)
	//cnt, power := make([]int, n), make([]int, n)
	//sort.Ints(nums)
	//for i := 0; i < n; i++ {
	//	cnt[i] = 1 << i % mod
	//	power[i] = nums[i] * nums[i] % mod
	//}
	//for i := n - 1; i >= 0; i-- {
	//	for j, k := 0, i-1; j < i; j++ {
	//		//sum = (power[i]*nums[j]%mod*cnt[k]%mod + sum) % mod
	//		sum = (power[i]*nums[j]%mod*cnt[k]%mod + sum) % mod
	//		//sum %= mod
	//		k--
	//	}
	//	sum = (power[i]*nums[i]%mod + sum) % mod
	//	//sum %= mod
	//}
	//return sum
}

/*
5：1<<i
	1：8
	2：4
	3：2
	4：1
	5：1
*/
//leetcode submit region end(Prohibit modification and deletion)
/*
[5 4 3 2 1]
[5 4 3 2]
[5 4 3 1]
[5 4 3]
[5 4 2 1]
[5 4 2]
[5 4 1]
[5 4]
[5 3 2 1]
[5 3 2]
[5 3 1]
[5 3]
[5 2 1]
[5 2]
[5 1]
[5]
[4 3 2 1]
[4 3 2]
[4 3 1]
[4 3]
[4 2 1]
[4 2]
[4 1]
[4]
[3 2 1]
[3 2]
[3 1]
[3]
[2 1]
[2]
[1]
*/
