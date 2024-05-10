//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回 滑动窗口中的最大值 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 2543 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	//nums = []int{1, 3, 1, 2, 0, 5}
	//k = 3 // [3,3,2,5]
	window := maxSlidingWindow(nums, k)
	fmt.Println(window)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {

}

// leetcode submit region end(Prohibit modification and deletion)

//func maxSlidingWindow(nums []int, k int) []int {
//	// 左右“滑动”窗体
//	// lWin, rWin 两者的算法相同，最后“错位”求最大值
//	// [1 3 3 -3 5 5 6 7]
//	//     [3 3 -1 5 5 3 7 7]
//	n := len(nums)
//	ans, lWin, rWin := make([]int, n-k+1), make([]int, n), make([]int, n)
//	for i, j, lv, rv := 0, n-1, 0, 0; i < n; i, j = i+1, j-1 {
//		if i%k == 0 {
//			lv = nums[i]
//		} else {
//			lv = max(lv, nums[i]) // k 个元素中的最大值
//		}
//		lWin[i] = lv
//		if (j+1)%k == 0 { // rWin 的算法和 lWin 相同
//			rv = nums[j]
//		} else {
//			rv = max(rv, nums[j])
//		}
//		rWin[j] = rv
//	}
//	for i := 0; i <= n-k; i++ {
//		ans[i] = max(lWin[i+k-1], rWin[i]) // “错位”最大值
//	}
//	return ans
//
//	// 双端队列：单调递减
//	//n := len(nums)
//	//ans, deque := make([]int, 0, n-k+1), []int{-k} // 哨兵
//	//for i, v := range nums {
//	//	if deque[0] == i-k { // 哨兵应对 i==0
//	//		deque = deque[1:]
//	//	}
//	//	for len(deque) > 0 && v >= nums[deque[len(deque)-1]] { // 单调递减
//	//		deque = deque[:len(deque)-1]
//	//	}
//	//	deque = append(deque, i)
//	//	if i >= k-1 {
//	//		ans = append(ans, nums[deque[0]])
//	//	}
//	//}
//	//return ans
//}
