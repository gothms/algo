//给你一个正整数数组 nums 。每一次操作中，你可以从 nums 中选择 任意 一个数并将它减小到 恰好 一半。（注意，在后续操作中你可以对减半过的数继续执
//行操作）
//
// 请你返回将 nums 数组和 至少 减少一半的 最少 操作数。
//
//
//
// 示例 1：
//
// 输入：nums = [5,19,8,1]
//输出：3
//解释：初始 nums 的和为 5 + 19 + 8 + 1 = 33 。
//以下是将数组和减少至少一半的一种方法：
//选择数字 19 并减小为 9.5 。
//选择数字 9.5 并减小为 4.75 。
//选择数字 8 并减小为 4 。
//最终数组为 [5, 4.75, 4, 1] ，和为 5 + 4.75 + 4 + 1 = 14.75 。
//nums 的和减小了 33 - 14.75 = 18.25 ，减小的部分超过了初始数组和的一半，18.25 >= 33/2 = 16.5 。
//我们需要 3 个操作实现题目要求，所以返回 3 。
//可以证明，无法通过少于 3 个操作使数组和减少至少一半。
//
//
// 示例 2：
//
// 输入：nums = [3,8,20]
//输出：3
//解释：初始 nums 的和为 3 + 8 + 20 = 31 。
//以下是将数组和减少至少一半的一种方法：
//选择数字 20 并减小为 10 。
//选择数字 10 并减小为 5 。
//选择数字 3 并减小为 1.5 。
//最终数组为 [1.5, 8, 5] ，和为 1.5 + 8 + 5 = 14.5 。
//nums 的和减小了 31 - 14.5 = 16.5 ，减小的部分超过了初始数组和的一半， 16.5 >= 31/2 = 16.5 。
//我们需要 3 个操作实现题目要求，所以返回 3 。
//可以证明，无法通过少于 3 个操作使数组和减少至少一半。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁷
//
//
// Related Topics 贪心 数组 堆（优先队列） 👍 12 👎 0

package main

import "container/heap"

func main() {

}

/*
思路：优先队列
	1.大顶堆记录 nums 的元素，nums 的总和为 sum
	2.每次取出堆顶元素并折半
		2.1.sum -= 折半值
		2.2.折半后，再次放入优先队列中堆化
		2.3.重复上面操作，直到 sum <= 初始 值的一半
*/
// leetcode submit region begin(Prohibit modification and deletion)
func halveArray(nums []int) int {
	// 优先队列
	sum, n := 0, len(nums)
	h := ha2208{}
	for i := 0; i < n; i++ {
		sum += nums[i]
		heap.Push(&h, float64(nums[i]))
	}
	var (
		cnt = 0
		s   = float64(sum)
		v   = float64(sum) / 2.0
	)
	for ; s > v; cnt++ {
		h[0] /= 2.0
		s -= h[0]
		heap.Fix(&h, 0)
	}
	return cnt
}

type ha2208 []float64

func (h ha2208) Len() int {
	return len(h)
}
func (h ha2208) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h ha2208) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *ha2208) Push(x any) {
	*h = append(*h, x.(float64))
}
func (h *ha2208) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
