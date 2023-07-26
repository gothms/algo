//给你一个整数数组，返回它的某个 非空 子数组（连续元素）在执行一次可选的删除操作后，所能得到的最大元素总和。换句话说，你可以从原数组中选出一个子数组，并可以
//决定要不要从中删除一个元素（只能删一次哦），（删除后）子数组中至少应当有一个元素，然后该子数组（剩下）的元素总和是所有子数组之中最大的。
//
// 注意，删除一个元素后，子数组 不能为空。
//
//
//
// 示例 1：
//
//
//输入：arr = [1,-2,0,3]
//输出：4
//解释：我们可以选出 [1, -2, 0, 3]，然后删掉 -2，这样得到 [1, 0, 3]，和最大。
//
// 示例 2：
//
//
//输入：arr = [1,-2,-2,3]
//输出：3
//解释：我们直接选出 [3]，这就是最大和。
//
//
// 示例 3：
//
//
//输入：arr = [-1,-1,-1,-1]
//输出：-1
//解释：最后得到的子数组不能为空，所以我们不能选择 [-1] 并从中删去 -1 来得到 0。
//     我们应该直接选择 [-1]，或者选择 [-1, -1] 再从中删去一个 -1。
//
//
//
//
// 提示：
//
//
//
// 1 <= arr.length <= 10⁵
// -10⁴ <= arr[i] <= 10⁴
//
//
// Related Topics 数组 动态规划 👍 168 👎 0

package main

import "fmt"

func main() {
	arr := []int{1, -2, 0, 3}
	arr = []int{-1, -1, -2, -1}
	arr = []int{1, -2, -2, 3}
	arr = []int{1, -2, -2, 3, 6, -4, 2, 5}
	sum := maximumSum(arr)
	fmt.Println(sum)
}

/*
思路：dp
	1.对于任意 arr[i]，当前最大和 = max(不删除元素的前缀和, 删除一个元数的前缀和)
	2.分情况讨论状态转移方程
		2.1.当前最大和 = 不删除元素的前缀和：
			dp[i] = arr[i]，dp[i-1]<0
			dp[i] = dp[i-1] + arr[i]，dp[i-1]>=0
		2.2.当前最大和 = 删除一个元数的前缀和：
			dp[i] = dp[i-1]，删除 arr[i]
			dp[i] = dp[i-1] + arr[i]，含删除某个元素
*/
// leetcode submit region begin(Prohibit modification and deletion)
func maximumSum(arr []int) int {
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//max, sum, dp, n := arr[0], arr[0], 0, len(arr)
	//for i := 1; i < n; i++ {
	//	dp = maxVal(sum, dp+arr[i])
	//	if sum < 0 {
	//		sum = arr[i]
	//	} else {
	//		sum += arr[i]
	//	}
	//	max = maxVal(max, dp)
	//}
	//return maxVal(max, sum)
	// dp
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 初始化，防止 arr[0]<0
	max, sum, dp, n := arr[0], arr[0], 0, len(arr)
	for i := 1; i < n; i++ {
		dp = maxVal(sum, dp+arr[i]) // 初始化当前最大和
		if sum < 0 {                // 当前最大前缀和
			sum = arr[i]
		} else {
			sum += arr[i]
		}
		dp = maxVal(dp, sum)  // 当前最大和
		max = maxVal(max, dp) // 当前最大最大和
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
