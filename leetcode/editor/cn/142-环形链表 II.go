//给定一个链表的头节点 head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到
//链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//
// 不允许修改 链表。
//
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [3,2,0,-4], pos = 1
//输出：返回索引为 1 的链表节点
//解释：链表中有一个环，其尾部连接到第二个节点。
//
//
// 示例 2：
//
//
//
//
//输入：head = [1,2], pos = 0
//输出：返回索引为 0 的链表节点
//解释：链表中有一个环，其尾部连接到第一个节点。
//
//
// 示例 3：
//
//
//
//
//输入：head = [1], pos = -1
//输出：返回 null
//解释：链表中没有环。
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围在范围 [0, 10⁴] 内
// -10⁵ <= Node.val <= 10⁵
// pos 的值为 -1 或者链表中的一个有效索引
//
//
//
//
// 进阶：你是否可以使用 O(1) 空间解决此题？
//
// Related Topics 哈希表 链表 双指针 👍 2191 👎 0

package main

func main() {

}

/*
思路：快慢指针
    1.通过 LeetCode-141 的方式，判断是否有环
	2.假设有环，环的长度为 y（即有 y 个结点），头的长度为 x（即非环的长度）
		fast 走了 m 圈环，slow 走了 n 圈环，它们在环的第 a 个结点相遇
		也就是说，此时 fast和slow 距离环的起点有 y-a 个结点
	3.计算
		f(fast) = m*y+x+a
		f(slow) = n*y+x+a
		又因为：f(fast) = 2*f(slow)
		两式相减：(m-2n)*y-x-a = 0
		化简为：x = (m-2n)y - a
		m-2n 肯定为 >=1 的正整数，我们设 m-2n = k + 1
		则可化简为 x = k*y + y-a，而 y-a 恰好是现在 slow 移动到 “环的起点“ 的距离
		结论：
			当 fast和slow 相遇时，slow 再走 k*y + y-a 个结点，恰好会来到环的起点
			而这个距离，也是从 head 走到环起点的距离 x
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// 快慢指针
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil { // 无环
			return nil
		}
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow { // 有环
			break
		}
	}
	for fast = head; slow != fast; { // 此时 fast 也是每次走一步
		fast, slow = fast.Next, slow.Next // 双慢指针
	}
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)
