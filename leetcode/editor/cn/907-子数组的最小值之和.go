//给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。
//
// 由于答案可能很大，因此 返回答案模 10^9 + 7 。
//
//
//
// 示例 1：
//
//
//输入：arr = [3,1,2,4]
//输出：17
//解释：
//子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。
//最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。
//
// 示例 2：
//
//
//输入：arr = [11,81,94,43,3]
//输出：444
//
//
//
//
// 提示：
//
//
// 1 <= arr.length <= 3 * 10⁴
// 1 <= arr[i] <= 3 * 10⁴
//
//
//
//
// Related Topics 栈 数组 动态规划 单调栈 👍 683 👎 0

package main

import "fmt"

func main() {
	arr := []int{11, 81, 94, 43, 3} // 444
	mins := sumSubarrayMins(arr)
	fmt.Println(mins)
}

// leetcode submit region begin(Prohibit modification and deletion)
func sumSubarrayMins(arr []int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func sumSubarrayMins_(arr []int) int {
	// 个人
	const mod = 1_000_000_007
	st := [][2]int{{-1, 0}}
	ans := 0
	for i, v := range arr {
		for len(st) > 1 && v < arr[st[len(st)-1][0]] {
			st = st[:len(st)-1]
		}
		last := st[len(st)-1]
		cur := last[1] + (i-last[0])*v
		ans += cur
		st = append(st, [2]int{i, cur})
	}
	return ans % mod

	// lc
	//ans := 0
	//arr = append(arr, -1)
	//st := []int{-1} // 哨兵
	//for r, x := range arr {
	//	for len(st) > 1 && arr[st[len(st)-1]] >= x {
	//		i := st[len(st)-1]
	//		st = st[:len(st)-1]
	//		ans += arr[i] * (i - st[len(st)-1]) * (r - i) // 累加贡献
	//		fmt.Println(r, ans)
	//	}
	//	st = append(st, r)
	//}
	//return ans % (1e9 + 7)
}
