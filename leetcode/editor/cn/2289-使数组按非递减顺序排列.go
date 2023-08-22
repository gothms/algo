//给你一个下标从 0 开始的整数数组 nums 。在一步操作中，移除所有满足 nums[i - 1] > nums[i] 的 nums[i] ，其中 0 <
//i < nums.length 。
//
// 重复执行步骤，直到 nums 变为 非递减 数组，返回所需执行的操作数。
//
//
//
// 示例 1：
//
//
//输入：nums = [5,3,4,4,7,3,6,11,8,5,11]
//输出：3
//解释：执行下述几个步骤：
//- 步骤 1 ：[5,3,4,4,7,3,6,11,8,5,11] 变为 [5,4,4,7,6,11,11]
//- 步骤 2 ：[5,4,4,7,6,11,11] 变为 [5,4,7,11,11]
//- 步骤 3 ：[5,4,7,11,11] 变为 [5,7,11,11]
//[5,7,11,11] 是一个非递减数组，因此，返回 3 。
//
//
// 示例 2：
//
//
//输入：nums = [4,5,7,7,13]
//输出：0
//解释：nums 已经是一个非递减数组，因此，返回 0 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁹
//
//
// Related Topics 栈 数组 链表 单调栈 👍 122 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11}
	//nums = []int{11, 12, 5, 3, 14, 4, 1, 4}
	//nums = []int{5, 11, 7, 8, 11} // 2
	//nums = []int{10, 6, 5, 10, 15} // 1
	steps := totalSteps(nums)
	fmt.Println(steps)
}

/*
注意：
	[11,8,5] = 1
	[11,8,9] = 2
思路：并查集（超时）
思路：链表（超时）
思路：链
思路：单调栈 + 线段树
	单调栈求左边最近更大元素位置，线段树维护区间最大值
思路：栈
	1.
思路：逆序对
*/
// leetcode submit region begin(Prohibit modification and deletion)
type list struct {
	val  int
	next *list
}

func totalSteps(nums []int) int {
	// 栈
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//type vs struct {
	//	v    int // nums[i] 的值
	//	step int // nums[i] 在第 step 次被移除
	//}
	//totalStep, n := 0, len(nums)
	//stack := []vs{{nums[0], 0}}
	//for i := 1; i < n; i++ {
	//	l, currStep := len(stack)-1, 0
	//	for ; l >= 0 && nums[i] >= stack[l].v; l-- { // nums[i] 后面的元素，优先被 nums[i] 移除
	//		currStep = maxVal(currStep, stack[l].step) // 遍历元素被移除的最大 step
	//	} // nums[i-1] > nums[i] 那么 i 会被 i-1 移除，且移除次数 f(i) = f(i-1)+1
	//	if l < 0 { // nums[i] 是最大元素，不会被移除，所以它的次数为 0
	//		currStep = 0
	//	} else { // stack 中还有更大的元素，那么 nums[i] 会被这个元素移除，次数为 currStep+1
	//		currStep++
	//		totalStep = maxVal(totalStep, currStep) // 记录最大次数
	//	}
	//	stack = append(stack[:l+1], vs{nums[i], currStep}) // nums[i] 和 它的次数 入栈
	//}
	//return totalStep

	// 链/并查集"+模拟
	nums = append(nums, math.MaxInt32) // 防止越界
	cnt, n := 0, len(nums)
	next, visited := make([]int, n), make([]bool, n)
	for i := 0; i < n; i++ {
		next[i] = i + 1 // i 要删除的第一个元素为 i+1
	}
	del := make([]int, 0)
	for i := n - 1; i > 0; i-- { // 需要倒序遍历
		if nums[i-1] > nums[i] {
			//visited[i] = true
			//next[i]++
			del = append(del, i-1) // 标记 i-1 元素后面的数将被删除
		}
	}
	for {
		nDel := make([]int, 0)
		for _, i := range del {
			if !visited[i] && nums[i] > nums[next[i]] {
				visited[next[i]] = true // 已删除标记
				next[i] = next[next[i]] // 下一个可能被删除的元素为 next[i]
				nDel = append(nDel, i)  // 标记 i 元素后面的数将被删除
			}
		}
		if len(nDel) == 0 { // 已删除完
			break
		}
		cnt++
		del = nDel // 新的数组
	}
	return cnt

	// 链+模拟：lc
	//nums = append(nums, math.MaxInt32)
	//cnt, n := 0, len(nums)
	//next, visited := make([]int, n), make([]bool, n)
	//for i := 0; i < n; i++ {
	//	next[i] = i + 1
	//}
	//v1 := make([]int, 0)
	//for i := n - 1; i > 0; i-- {
	//	if nums[i-1] > nums[i] {
	//		v1 = append(v1, i-1)
	//	}
	//}
	//if len(v1) == 0 {
	//	return 0
	//}
	//fmt.Println("v1", v1)
	//// nums := []int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11}
	//for {
	//	v2 := make([]int, 0)
	//	//for i := len(v1) - 1; i >= 0; i-- {
	//	for _, i := range v1 {
	//		if !visited[i] && nums[i] > nums[next[i]] {
	//			visited[next[i]] = true
	//			next[i] = next[next[i]]
	//			v2 = append(v2, i)
	//		}
	//	}
	//	fmt.Println(v2)
	//	fmt.Println(next)
	//	if len(v2) > 0 {
	//		v1 = v2
	//	} else {
	//		return cnt
	//	}
	//	cnt++
	//}
	//return cnt

	// 链表：超时
	//cnt, n := -1, len(nums)
	//head := &list{nums[0], nil}
	//for i, curr := 1, head; i < n; i++ {
	//	curr.next = &list{nums[i], nil}
	//	curr = curr.next
	//}
	//for hasNext := true; hasNext; cnt++ {
	//	hasNext = false
	//	for curr := head; curr != nil; curr = curr.next {
	//		pre := curr
	//		for pre.next != nil && pre.val > pre.next.val { // 连续的降区间
	//			pre = pre.next
	//		}
	//		if pre != curr {
	//			curr.next = pre.next
	//			//hasNext = true	// 有最后一次“空跑”
	//			if curr.val > pre.val { // 没有最后一次“空跑”
	//				hasNext = true
	//			}
	//		}
	//	}
	//}
	//return cnt

	// 并查集：最后一个用例超时
	//cnt, n := -1, len(nums)
	//uni := make([]int, n)
	//for i := 1; i < n; i++ {
	//	uni[i] = i
	//}
	//var find func(int) int
	//find = func(p int) int {
	//	//if uni[p] == p {
	//	//	return p
	//	//}
	//	//for uni[p] != uni[uni[p]] {
	//	//	uni[p] = uni[uni[p]]
	//	//}
	//	if p != uni[p] {
	//		uni[p] = find(uni[p])
	//	}
	//	return uni[p]
	//}
	//for next := true; next; cnt++ {
	//	next = false
	//	for i, pre := n-1, 0; i > 0; i = pre { // 此题必须倒序遍历
	//		if pre = find(i - 1); i == uni[i] && nums[i] < nums[pre] { // 跳过已经淘汰的数
	//			uni[i] = pre // 跳过的数都并到上一个（索引最小）更大的数
	//			next = true  // 是否有执行操作
	//		}
	//	}
	//}
	//return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
