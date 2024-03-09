//给你一个整数数组 nums 和一个 正 整数 k 。你可以选择数组的任一 子序列 并且对其全部元素求和。
//
// 数组的 第 k 大和 定义为：可以获得的第 k 个 最大 子序列和（子序列和允许出现重复）
//
// 返回数组的 第 k 大和 。
//
// 子序列是一个可以由其他数组删除某些或不删除元素排生而来的数组，且派生过程不改变剩余元素的顺序。
//
// 注意：空子序列的和视作 0 。
//
//
//
// 示例 1：
//
// 输入：nums = [2,4,-2], k = 5
//输出：2
//解释：所有可能获得的子序列和列出如下，按递减顺序排列：
//- 6、4、4、2、2、0、0、-2
//数组的第 5 大和是 2 。
//
//
// 示例 2：
//
// 输入：nums = [1,-2,3,4,-10,12], k = 16
//输出：10
//解释：数组的第 16 大和是 10 。
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// 1 <= k <= min(2000, 2ⁿ)
//
//
// Related Topics 数组 排序 堆（优先队列） 👍 84 👎 0

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 5, 9, -15}
	k := 15
	i := kSum(nums, k)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func kSum(nums []int, k int) int64 {
	sum, n := 0, len(nums)
	for i, v := range nums {
		if v >= 0 {
			sum += v
		} else {
			nums[i] = -v
		}
	}
	sort.Ints(nums)
	h := &ksHp{{}}
	for i := k; i > 1; i-- {
		c := heap.Pop(h).([2]int)
		if c[0] < n { // 当前要处理的索引
			heap.Push(h, [2]int{c[0] + 1, c[1] + nums[c[0]]}) // 新增元素和
			if c[0] > 0 {                                     // 新增元素和：替换元素
				heap.Push(h, [2]int{c[0] + 1, c[1] - nums[c[0]-1] + nums[c[0]]})
			}
		}
	}
	return int64(sum - (*h)[0][1])
}

type ksHp [][2]int

func (k ksHp) Len() int           { return len(k) }
func (k ksHp) Less(i, j int) bool { return k[i][1] < k[j][1] }
func (k ksHp) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k *ksHp) Push(x any)        { *k = append(*k, x.([2]int)) }
func (k *ksHp) Pop() any {
	v := (*k)[len(*k)-1]
	*k = (*k)[:len(*k)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
