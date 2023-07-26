//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
//
// Related Topics 递归 链表 数学 👍 9696 👎 0

package main

import "fmt"

func main() {
	l1 := &ListNode{9, &ListNode{9, &ListNode{9, nil}}}
	l2 := &ListNode{9, &ListNode{9, nil}}
	l1 = &ListNode{5, nil}
	l2 = &ListNode{5, nil}
	numbers := addTwoNumbers(l1, l2)
	fmt.Println(numbers)
}

/*
思路：遍历
	1.创建哨兵节点 pre，新的链表的当前节点为 curr = pre
	2.同时遍历两条链表，计算两个当前节点的和 v
		v >= 10：curr.Next = &ListNode{v - 10, nil}
			v = 1，保留到下一次的和计算
		v < 10：curr.Next = &ListNode{v, nil}
	3.当 l1 == nil && l2 == nil 后，如果 v = 1
		补位：curr.Next = &ListNode{1, nil}
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
	pre := &ListNode{0, nil}
	curr, v := pre, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{v % 10, nil}
		curr = curr.Next
		v /= 10
	}
	if v > 0 {
		curr.Next = &ListNode{1, nil}
	}
	return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
