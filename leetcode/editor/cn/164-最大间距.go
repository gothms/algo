//给定一个无序的数组 nums，返回 数组在排序之后，相邻元素之间最大的差值 。如果数组元素个数小于 2，则返回 0 。
//
// 您必须编写一个在「线性时间」内运行并使用「线性额外空间」的算法。
//
//
//
// 示例 1:
//
//
//输入: nums = [3,6,9,1]
//输出: 3
//解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。
//
// 示例 2:
//
//
//输入: nums = [10]
//输出: 0
//解释: 数组元素个数小于 2，因此返回 0。
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁵
// 0 <= nums[i] <= 10⁹
//
//
// Related Topics 数组 桶排序 基数排序 排序 👍 586 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{3, 6, 9, 1}
	i := maximumGap(nums)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumGap(nums []int) int {
	if len(nums) < 2 { // fast path
		return 0
	}
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	min, max, n := math.MaxInt32, 0, len(nums)
	for i := 0; i < n; i++ {
		min, max = minVal(min, nums[i]), maxVal(max, nums[i])
	}
	g := maxVal(1, (max-min)/(n-1)) // 在元素有不同的情况下，最大差肯定大于等于 g
	bucketSize := (max-min)/g + 1   // 桶的数量
	buckets := make([][2]int, bucketSize)
	for i := range buckets {
		buckets[i][0], buckets[i][1] = -1, -1 // 初始化：桶内最小值、最大值
	}
	for _, v := range nums { // 记录数据到桶
		i := (v - min) / g
		if buckets[i][0] == -1 {
			buckets[i][0] = v
			buckets[i][1] = v
		} else {
			buckets[i][0] = minVal(buckets[i][0], v)
			buckets[i][1] = maxVal(buckets[i][1], v)
		}
	}
	ret, last := 0, -1
	for _, b := range buckets { // cur.min - pre.max
		if b[0] == -1 {
			continue
		}
		if last != -1 {
			ret = maxVal(ret, b[0]-last)
		}
		last = b[1]
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
