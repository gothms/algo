//给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
// 你不需要 保留 每个分区中各节点的初始相对位置。
//
//
//
// 示例 1：
//
//
//输入：head = [1,4,3,2,5,2], x = 3
//输出：[1,2,2,4,3,5]
//
//
// 示例 2：
//
//
//输入：head = [2,1], x = 2
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 200] 内
// -100 <= Node.val <= 100
// -200 <= x <= 200
//
//
// Related Topics 链表 双指针 👍 145 👎 0

package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, nil}}
	x := 2
	node := partition(head, x)
	fmt.Println(node)
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	// 两条链
	var lp, rp ListNode
	l, r := &lp, &rp
	for ; head != nil; head = head.Next {
		if head.Val < x {
			l.Next, l = head, head
		} else {
			r.Next, r = head, head
		}
	}
	l.Next, r.Next = rp.Next, nil // 连接起来，并把右链的尾巴断开
	return lp.Next

	// 一条链
	//lPre := &ListNode{}
	//var rPre, r *ListNode
	//l := lPre
	//for head != nil {
	//	if head.Val < x {
	//		l.Next, l, head, head.Next = head, head, head.Next, rPre
	//	} else {
	//		if rPre == nil {
	//			rPre = head
	//		} else {
	//			r.Next = head
	//		}
	//		r, head, head.Next = head, head.Next, nil
	//	}
	//}
	//if l.Next == nil { // 连接起来
	//	l.Next = rPre
	//}
	//return lPre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
