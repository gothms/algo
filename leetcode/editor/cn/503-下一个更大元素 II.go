//给定一个循环数组 nums （ nums[nums.length - 1] 的下一个元素是 nums[0] ），返回 nums 中每个元素的 下一个更大元素
// 。
//
// 数字 x 的 下一个更大的元素 是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1
//。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,2,1]
//输出: [2,-1,2]
//解释: 第一个 1 的下一个更大的数是 2；
//数字 2 找不到下一个更大的数；
//第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
//
//
// 示例 2:
//
//
//输入: nums = [1,2,3,4,3]
//输出: [2,3,4,-1,4]
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics 栈 数组 单调栈 👍 867 👎 0

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 3}
	//nums = []int{2, 2, 3, 4, 3}
	elements := nextGreaterElements(nums)
	fmt.Println(elements)
}

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	nge, stack := make([]int, n), make([]int, 0, n+1)
	for i := range nge {
		nge[i] = -1 // 标记
	}
	for i := 0; i < n<<1-1; i++ { // 循环两圈
		last := len(stack) - 1
		for last >= 0 && nums[stack[last]] < nums[i%n] { // 找到更大的数 nums[i%n]
			if nge[stack[last]] == -1 {
				nge[stack[last]] = nums[i%n] // 赋值
			}
			last--
		}
		stack = stack[:last+1]
		stack = append(stack, i%n)
	}
	return nge

	//n := len(nums)
	//nge, stack := make([]int, n), make([]int, 0, n+1)
	////stack[0] = nums[0]
	//for i, v := range nums {
	//	last := len(stack) - 1
	//	if last < 0 || v > nums[stack[last]] {
	//		stack = append(stack, i)
	//	} else {
	//		stack = append(stack, stack[last])
	//	}
	//}
	//fmt.Println(stack)
	//return nge
}

//leetcode submit region end(Prohibit modification and deletion)
