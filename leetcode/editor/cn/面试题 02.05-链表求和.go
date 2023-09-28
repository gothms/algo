//给定两个用链表表示的整数，每个节点包含一个数位。
//
// 这些数位是反向存放的，也就是个位排在链表首部。
//
// 编写函数对这两个整数求和，并用链表形式返回结果。
//
//
//
// 示例：
//
// 输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
//输出：2 -> 1 -> 9，即912
//
//
// 进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?
//
// 示例：
//
// 输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
//输出：9 -> 1 -> 2，即912
//
//
// Related Topics 递归 链表 数学 👍 188 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 老写法：简洁
	pre, v := &ListNode{}, 0
	l := pre
	for ; l1 != nil || l2 != nil; v /= 10 {
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		l.Next = &ListNode{Val: v % 10}
		l = l.Next
	}
	if v != 0 {
		l.Next = &ListNode{Val: 1}
	}
	return pre.Next

	// 老写法
	//ln, v := &ListNode{}, 0
	//pre, l := ln, ln
	//for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
	//	v += l1.Val + l2.Val
	//	l1.Val, v = v%10, v/10
	//	l.Next, l = l1, l1
	//}
	//if l2 != nil {
	//	l.Next = l2
	//}
	//for ; v != 0 && l.Next != nil; l = l.Next {
	//	v += l.Next.Val
	//	l.Next.Val, v = v%10, v/10
	//}
	//if v != 0 {
	//	l.Next = &ListNode{Val: 1}
	//}
	//return pre.Next

	// 新写法
	//var (
	//	up int
	//	ln ListNode
	//)
	//pre, l := &ln, &ln
	//tail := func(ln *ListNode) *ListNode {
	//	for l.Next = ln; up != 0 && l.Next != nil; l = l.Next {
	//		v := l.Next.Val + 1 // 参数 ln 有可能为 nil，所以要记录 l，操作 l.Next
	//		l.Next.Val, up = v%10, v/10
	//	}
	//	if up != 0 {
	//		l.Next = &ListNode{Val: 1}
	//	}
	//	return pre.Next
	//}
	//for {
	//	if l1 == nil {
	//		return tail(l2)
	//	} else if l2 == nil {
	//		return tail(l1)
	//	}
	//	v := l1.Val + l2.Val + up
	//	node := &ListNode{Val: v % 10}
	//	l.Next, l, l1, l2 = node, node, l1.Next, l2.Next
	//	up = v / 10
	//}
}

//leetcode submit region end(Prohibit modification and deletion)
