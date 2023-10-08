//给定一个整数数组 A，坡是元组 (i, j)，其中 i < j 且 A[i] <= A[j]。这样的坡的宽度为 j - i。
//
// 找出 A 中的坡的最大宽度，如果不存在，返回 0 。
//
//
//
// 示例 1：
//
// 输入：[6,0,8,2,1,5]
//输出：4
//解释：
//最大宽度的坡为 (i, j) = (1, 5): A[1] = 0 且 A[5] = 5.
//
//
// 示例 2：
//
// 输入：[9,8,1,0,1,9,4,0,4,1]
//输出：7
//解释：
//最大宽度的坡为 (i, j) = (2, 9): A[2] = 1 且 A[9] = 1.
//
//
//
//
// 提示：
//
//
// 2 <= A.length <= 50000
// 0 <= A[i] <= 50000
//
//
//
//
// Related Topics 栈 数组 单调栈 👍 236 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxWidthRamp(nums []int) int {
	// stack：lc
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	stack := make([]int, 1)
	max, n := 0, len(nums)
	for i := 1; i < n; i++ { // 准备 stack 数据
		if nums[i] < nums[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}
	for i, last := n-1, len(stack)-1; i >= 0; i-- { // 重要逻辑
		for nums[i] >= nums[stack[last]] {
			max = maxVal(max, i-stack[last])
			if last == 0 { // stack 已遍历完
				return max
			}
			last--
		}
	}
	return max

	// stack
	//maxVal := func(a, b int) int {
	//	if b > a {
	//		return b
	//	}
	//	return a
	//}
	//max, stack := 0, make([]int, 0)
	//for i, v := range nums {
	//	n := len(stack)
	//	if n == 0 || v < nums[stack[n-1]] {
	//		stack = append(stack, i)
	//		continue
	//	}
	//	j := sort.Search(n, func(i int) bool {
	//		return nums[stack[i]] <= v
	//	})
	//	max = maxVal(max, i-stack[j])
	//}
	//return max

	// 排序：对索引进行排序，遍历并记录当前最小索引（后出现的元素，肯定大于先出现的）
}

//leetcode submit region end(Prohibit modification and deletion)
