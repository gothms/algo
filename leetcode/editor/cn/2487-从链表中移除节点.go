//给你一个链表的头节点 head 。
//
// 移除每个右侧有一个更大数值的节点。
//
// 返回修改后链表的头节点 head 。
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [5,2,13,3,8]
//输出：[13,8]
//解释：需要移除的节点是 5 ，2 和 3 。
//- 节点 13 在节点 5 右侧。
//- 节点 13 在节点 2 右侧。
//- 节点 8 在节点 3 右侧。
//
//
// 示例 2：
//
//
//输入：head = [1,1,1,1]
//输出：[1,1,1,1]
//解释：每个节点的值都是 1 ，所以没有需要移除的节点。
//
//
//
//
// 提示：
//
//
// 给定列表中的节点数目在范围 [1, 10⁵] 内
// 1 <= Node.val <= 10⁵
//
//
// Related Topics 栈 递归 链表 单调栈 👍 35 👎 0

package main

import "math"

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
func removeNodes(head *ListNode) *ListNode {
	// 单调栈
	pre := &ListNode{math.MaxInt32, head}
	st := []*ListNode{pre}
	for ; head != nil; head = head.Next {
		last := len(st) - 1
		for last > 0 && head.Val > st[last].Val {
			last--
		}
		st[last].Next = head
		st = append(st[:last+1], head)
	}
	return pre.Next

	// 单调栈：递减
	//pre := &ListNode{math.MaxInt32, head}
	//stack := []*ListNode{pre} // 哨兵
	//for ; head != nil; head = head.Next {
	//	last := len(stack) - 1
	//	for last > 0 && head.Val > stack[last].Val {
	//		last--
	//	}
	//	stack[last].Next = head
	//	stack = stack[:last+1]
	//	stack = append(stack, head)
	//}
	//return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
