//输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
//
// 示例1：
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
// 限制：
//
// 0 <= 链表长度 <= 1000
//
// 注意：本题与主站 21 题相同：https://leetcode-cn.com/problems/merge-two-sorted-lists/
//
// Related Topics 递归 链表 👍 347 👎 0

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{0, nil}
	curr := pre
	for ; l1 != nil && l2 != nil; curr = curr.Next {
		if l1.Val < l2.Val {
			curr.Next, l1 = l1, l1.Next
		} else {
			curr.Next, l2 = l2, l2.Next
		}
	}
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}
	return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
