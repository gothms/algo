package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 3}
	nums = []int{1, 1, 1, 1, 1}
	nums = []int{-6, -1, 5, 4, 1, -8, 6, 7, -3, 6, 0, -6, -7, 8, -8, -4, 1}
	elements := nextGreaterElements(nums)
	fmt.Println(elements)
}

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElements(nums []int) []int {

}

//leetcode submit region end(Prohibit modification and deletion)

func nextGreaterElements_(nums []int) []int {
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
