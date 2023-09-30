//给定一个直方图(也称柱状图)，假设有人从上面源源不断地倒水，最后直方图能存多少水量?直方图的宽度为 1。
//
//
//
// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的直方图，在这种情况下，可以接 6 个单位的水（蓝色部分表示水）。 感谢
//Marcos 贡献此图。
//
// 示例:
//
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 289 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	t, n := 0, len(height)
	stack := make([]int, 0, n)
	for i, h := range height {
		last := len(stack) - 1
		for last >= 0 && h >= height[stack[last]] {
			if last-1 >= 0 { // 存在左边界
				left := stack[last-1] // 左边界
				// (左右边界更小值 - 中间柱子的高) * 宽度
				t += (minVal(h, height[left]) - height[stack[last]]) * (i - left - 1)
			}
			last--
		}
		stack = stack[:last+1]
		stack = append(stack, i)
	}
	return t
}

//leetcode submit region end(Prohibit modification and deletion)
