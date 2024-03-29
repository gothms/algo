//给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,2,1]
//输出：true
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：false
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围[1, 10⁵] 内
// 0 <= Node.val <= 9
//
//
//
//
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
// Related Topics 栈 递归 链表 双指针 👍 1786 👎 0

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
func isPalindrome(head *ListNode) bool {
	var dfs func(*ListNode) bool
	dfs = func(cur *ListNode) bool { // 相对于 LCR 027 的解法，重复比对了一遍
		if cur != nil { // 继续遍历结点
			if !dfs(cur.Next) || cur.Val != head.Val { // 已经不匹配，或当前结点不匹配
				return false
			}
			head = head.Next // dfs 从尾往头，head 从头往尾
		}
		return true
	}
	return dfs(head)
}

//leetcode submit region end(Prohibit modification and deletion)
