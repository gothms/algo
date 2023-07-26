//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
// Related Topics 递归 链表 👍 3188 👎 0

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
/*
LeetCode-206
思路：迭代
	1.设置标兵节点 pre，pre的next节点永远是反转后链表的头
	2.遍历链表的每个节点，同时修改一下几个指针
		temp = curr.Next
		curr.Next = pre.Next
		pre.Next = curr
		curr = temp
*/
func reverseList(head *ListNode) *ListNode {
	pre := &ListNode{0, nil}
	for head != nil {
		pre.Next, head, head.Next = head, head.Next, pre.Next
	}
	return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
