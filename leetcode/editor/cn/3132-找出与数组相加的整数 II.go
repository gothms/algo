//给你两个整数数组 nums1 和 nums2。
//
// 从 nums1 中移除两个元素，并且所有其他元素都与变量 x 所表示的整数相加。如果 x 为负数，则表现为元素值的减少。
//
// 执行上述操作后，nums1 和 nums2 相等 。当两个数组中包含相同的整数，并且这些整数出现的频次相同时，两个数组 相等 。
//
// 返回能够实现数组相等的 最小 整数 x 。
//
//
//
// 示例 1:
//
//
// 输入：nums1 = [4,20,16,12,8], nums2 = [14,18,10]
//
//
// 输出：-2
//
// 解释：
//
// 移除 nums1 中下标为 [0,4] 的两个元素，并且每个元素与 -2 相加后，nums1 变为 [18,14,10] ，与 nums2 相等。
//
// 示例 2:
//
//
// 输入：nums1 = [3,5,5,3], nums2 = [7,7]
//
//
// 输出：2
//
// 解释：
//
// 移除 nums1 中下标为 [0,3] 的两个元素，并且每个元素与 2 相加后，nums1 变为 [7,7] ，与 nums2 相等。
//
//
//
// 提示：
//
//
// 3 <= nums1.length <= 200
// nums2.length == nums1.length - 2
// 0 <= nums1[i], nums2[i] <= 1000
// 测试用例以这样的方式生成：存在一个整数 x，nums1 中的每个元素都与 x 相加后，再移除两个元素，nums1 可以与 nums2 相等。
//
//
// Related Topics 数组 双指针 枚举 排序 👍 2 👎 0

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums1 := []int{3, 5, 5, 3}
	nums2 := []int{7, 7}
	integer := minimumAddedInteger(nums1, nums2)
	fmt.Println(integer)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumAddedInteger(nums1 []int, nums2 []int) int {
	ans, minV := math.MaxInt32, nums2[len(nums2)-1]
	arr := [3]int(nums1[:3])
	sort.Ints(arr[:])
	for i, v := range nums1[3:] {
		if arr[2] > v {
			arr[2] = v
			sort.Ints(arr[:])
		}
		minV = min(minV, nums2[i])
	}
	for _, val := range arr {
		memo := make(map[int]int)
		cnt, d := 0, minV-val
		for _, v := range nums2 {
			memo[v]++
		}
		for _, v := range nums1 {
			if memo[v+d] == 0 {
				cnt++
			} else {
				memo[v+d]--
			}
		}
		if cnt == 2 {
			ans = min(ans, d)
		}
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
//func minimumAddedInteger(nums1 []int, nums2 []int) int {
//	// hash
//	// 其他方法：排序后，取 nums1[2] nums1[1] nums1[0] 尝试
//	hp, m := &hp3132{}, math.MaxInt32
//	for _, v := range nums1 {
//		if hp.Len() < 3 {
//			heap.Push(hp, v)
//		} else if v < hp.IntSlice[0] {
//			hp.IntSlice[0] = v
//			heap.Fix(hp, 0)
//		}
//	}
//
//	ans := math.MaxInt32
//	for _, mv := range hp.IntSlice {
//		memo := make(map[int]int, len(nums2))
//		for _, v := range nums2 {
//			m = min(m, v)
//			memo[v]++
//		}
//		d, cnt := m-mv, 0
//		for _, v := range nums1 {
//			if memo[v+d] <= 0 {
//				cnt++
//			}
//			memo[v+d]--
//		}
//		if cnt == 2 {
//			ans = min(ans, d)
//		}
//	}
//	return ans
//}
//
//type hp3132 struct {
//	sort.IntSlice
//}
//
//func (h hp3132) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
//func (h *hp3132) Push(x any)        { h.IntSlice = append(h.IntSlice, x.(int)) }
//func (h *hp3132) Pop() any {
//	v := h.IntSlice[len(h.IntSlice)-1]
//	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
//	return v
//}
