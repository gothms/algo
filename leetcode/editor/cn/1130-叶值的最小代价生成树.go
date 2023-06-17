//给你一个正整数数组 arr，考虑所有满足以下条件的二叉树：
//
//
// 每个节点都有 0 个或是 2 个子节点。
// 数组 arr 中的值与树的中序遍历中每个叶节点的值一一对应。
// 每个非叶节点的值等于其左子树和右子树中叶节点的最大值的乘积。
//
//
// 在所有这样的二叉树中，返回每个非叶节点的值的最小可能总和。这个和的值是一个 32 位整数。
//
// 如果一个节点有 0 个子节点，那么该节点为叶节点。
//
//
//
// 示例 1：
//
//
//输入：arr = [6,2,4]
//输出：32
//解释：有两种可能的树，第一种的非叶节点的总和为 36 ，第二种非叶节点的总和为 32 。
//
//
// 示例 2：
//
//
//输入：arr = [4,11]
//输出：44
//
//
//
//
// 提示：
//
//
// 2 <= arr.length <= 40
// 1 <= arr[i] <= 15
// 答案保证是一个 32 位带符号整数，即小于 2³¹ 。
//
//
// Related Topics 栈 贪心 数组 动态规划 单调栈 👍 265 👎 0

package main

import (
	"fmt"
)

func main() {
	arr := []int{6, 2, 4}
	//arr = []int{7, 12, 8, 10, 9, 11}
	arr = []int{15, 13, 5, 3, 15} // 500
	//arr = []int{1, 2, 3, 4, 5}
	values := mctFromLeafValues(arr)
	fmt.Println(values)
}

/*
思路：动态规划
	1.由题意：
		1.1.arr[i] 都是叶子节点
		1.2.arr[i] 要满足中序遍历排列，则不能排序
			只能 原地 划分为 左/右 子树
		1.3.len(arr)=n，则有 n-1 个乘积
	2.dp
		2.1.假设 arr 是升序排列，arr=[1,2,3,4,5]
			那么最小乘积和为：1*2 + 2*3 + 3*4 + 4*5
			可见，一个数总是选择 比自己更大的数 “归边”
		2.2.再举个详细例子 arr=[7,12,8,10,9,11]，索引为 i
			i=0：0
			i=1：7*12
			i=2：7*12 + 12*8
				8 会选择比自己更大的 12 归边，但是 12 之前的数，先于 8 和 12 归边
				一个父节点的是否是这样建立的？8 会一直往左找：
					如果不比自己大，则选择和自己归边（right子树）
					如果比自己大，则停止（找到了left子树）
					最后和这个比自己大的数相乘，得到两者的父节点
			i=3：7*12 + 8*10 + 12*10
				10 找到 8，比自己小，和自己归边：8*10
				10 找到 12，比自己大，10和12相乘，得到二者父节点：12*10
			i=4：7*12 + 8*10 + 10*12 + 10*9
				9 找到 10，比自己大，9 直接归边于 10：10*9
					可见，此时直接用 i=3 的 最小乘积和 加上 10*9 即可
			i=5：可以自己递推下
				注意：会有一个 11*10，即 比自己小的最大数 * 自己，思考下为什么有 11*10
		2.3.由上面例子可知，我们只用保留 比自己索引小且比自己值大 的元素
			比如 i=3 时，arr[2]=8 在后续的递推中，已经不需要保留下来
			那么我们可以用一个 单调递减栈(stack) 来保存之前的元素
	3.状态转移方程
		3.1.arr[i] < stack[-1]：
			dp[i] = dp[i-1] + stack[-1]*arr[i]
		3.2.arr[i] >= stack[-1]：从后往前遍历 stack，直到 stack[j]>arr[i]
			dp[i] = dp[i-1] + stack[j]*arr[i] +
				stack[j-1]*arr[i] - stack[j-1]*stack[j]
			stack=stack[:j]：小于等于 arr[i] 的数都不保留
		3.3.计算 dp[i] 后，将 arr[i] 追加到 stack 尾
代码注释：
	1.初始化 dp 和 stack
	2.dp
		2.1.arr[i] < stack[-1]
		2.2.arr[i] >= stack[-1]
			a)寻找 stack[j]>arr[i]，也可以 二分查找
			b)计算 dp
			c)淘汰不大于 arr[i] 的元素
		2.3.追加 arr[i]
*/
//leetcode submit region begin(Prohibit modification and deletion)
func mctFromLeafValues(arr []int) int {
	//dp, stack, n := 0, []int{arr[0]}, len(arr)
	//for i, last := 1, len(stack)-1; i < n; i, last = i+1, len(stack)-1 {
	//	j, v := last, stack[last]
	//	if arr[i] >= stack[j] {
	//		j--
	//		for l := 0; l <= j; {
	//			m := (l + j) >> 1
	//			if stack[m] > arr[i] {
	//				l = m + 1
	//			} else {
	//				j = m - 1
	//			}
	//		}
	//		j++
	//		if j > 0 {
	//			dp += (arr[i] - stack[j]) * stack[j-1]
	//		}
	//		v, stack = stack[j], stack[:j]
	//	}
	//	dp += v * arr[i]
	//	stack = append(stack, arr[i])
	//}
	//return dp

	// dp+单调栈
	//dp, stack, n := arr[0]*arr[1], make([]int, 0), len(arr)
	//if arr[0] >= arr[1] {
	//	stack = append(stack, arr[0])
	//}
	//stack = append(stack, arr[1]) // i 从 2 开始
	dp, stack, n := 0, []int{arr[0]}, len(arr)                     // 1
	for i, l := 1, len(stack)-1; i < n; i, l = i+1, len(stack)-1 { // 2
		j, v := l, stack[l]
		for j >= 0 && arr[i] >= stack[j] { // 2.2
			j-- // 2.2.a
		}
		if j < l {
			if j++; j > 0 {
				dp += (arr[i] - stack[j]) * stack[j-1] // 2.2.b
			}
			v, stack = stack[j], stack[:j] // 2.2.c
		}
		dp += v * arr[i]              // 2.1 & 2.2.b
		stack = append(stack, arr[i]) // 2.3
	}
	return dp

	// 错误：中序遍历
	//mct, n := 0, len(arr)
	//sort.Ints(arr)
	//for i := 1; i < n; i++ {
	//	mct += arr[i] * arr[i-1]
	//}
	//return mct
}

//leetcode submit region end(Prohibit modification and deletion)
