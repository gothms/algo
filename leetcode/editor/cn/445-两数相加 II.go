//给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//
//
// 示例1：
//
//
//
//
//输入：l1 = [7,2,4,3], l2 = [5,6,4]
//输出：[7,8,0,7]
//
//
// 示例2：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[8,0,7]
//
//
// 示例3：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 链表的长度范围为 [1, 100]
// 0 <= node.val <= 9
// 输入数据保证链表代表的数字无前导 0
//
//
//
//
// 进阶：如果输入链表不能翻转该如何解决？
//
// Related Topics 栈 链表 数学 👍 606 👎 0

package main

import "fmt"

func main() {
	l1 := &ListNode{5, nil}
	l2 := &ListNode{5, nil}
	numbers := addTwoNumbers(l1, l2)
	fmt.Println(numbers)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
思路：栈
	1.使用两个栈，分别记录 l1 l2 的结点的 Val 值
	2.根据两个栈，计算新链表的结点的 Val 值
思路：回溯
	1.由于两数之和的高位先遍历到，所以可以提供一种回溯写法
	2.记录其中一条的链表的 Val
	3.另一条链表通过回溯，返回低位计算的结果，从而计算高位值
	4.如果记录的 Val 没有计算完整，最后再补位
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 回溯
	cache := make([]int, 0)
	for ; l2 != nil; l2 = l2.Next { // 记录 l2 的 Val
		cache = append(cache, l2.Val)
	}
	var dfs func(*ListNode) (int, int)
	dfs = func(r *ListNode) (int, int) {
		if r == nil { // 开始回溯
			return len(cache) - 1, 0
		}
		i, v := dfs(r.Next) // 低位计算的结果
		v += r.Val
		if i >= 0 { // cache 还没遍历完
			v += cache[i]
		}
		r.Val = v % 10       // 新结点的 Val
		return i - 1, v / 10 // 返回下一次计算的 cache 索引，和是否进位
	}
	i, v := dfs(l1)
	for ; i >= 0; i-- { // cache 没有遍历完，需要补位
		v += cache[i]
		l1 = &ListNode{v % 10, l1}
		v /= 10
	}
	if v > 0 { // 需要进位
		return &ListNode{1, l1}
	}
	return l1

	// stack
	//s1, s2 := make([]int, 0), make([]int, 0)
	//for ; l1 != nil; l1 = l1.Next {
	//	s1 = append(s1, l1.Val) // 记录 l1 的 Val
	//}
	//for ; l2 != nil; l2 = l2.Next {
	//	s2 = append(s2, l2.Val) // 记录 l2 的 Val
	//}
	//var l *ListNode // 新链表的head
	//v := 0
	//for i, j := len(s1)-1, len(s2)-1; i >= 0 || j >= 0; {
	//	if i >= 0 { // 开始计算新结点的 Val
	//		v += s1[i]
	//		i--
	//	}
	//	if j >= 0 {
	//		v += s2[j]
	//		j--
	//	}
	//	l = &ListNode{v % 10, l} // 新结点为新的head
	//	v /= 10
	//}
	//if v > 0 { // 补head
	//	return &ListNode{v, l}
	//}
	//return l
}

//leetcode submit region end(Prohibit modification and deletion)
