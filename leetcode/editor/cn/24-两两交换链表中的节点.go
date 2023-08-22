//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
//
// Related Topics 递归 链表 👍 1915 👎 0

package main

func main() {

}

/*
思路：迭代 / 递归
	如果能有 2 个结点，则交换这两个节点的位置
	直到没有结点了或者只有 1 个结点（不处理）
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// 迭代
	//pre := &ListNode{0, head} // 哨兵
	//curr := pre               // 前驱结点
	//for head != nil && head.Next != nil {
	//	curr.Next, curr, head.Next, head.Next.Next, head =
	//		head.Next, head, head.Next.Next, head, head.Next.Next
	//} // 挪动前驱结点，并交换结点
	//return pre.Next

	// 递归
	if head == nil || head.Next == nil {
		return head
	} // 没有结点/只有 1 个结点，不处理
	next := head.Next
	head.Next, next.Next = swapPairs(head.Next.Next), head // 交换结点
	return next                                            // 返回新的头结点
}

//leetcode submit region end(Prohibit modification and deletion)
