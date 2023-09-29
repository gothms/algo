//编写一个函数，检查输入的链表是否是回文的。
//
//
//
// 示例 1：
//
// 输入： 1->2
//输出： false
//
//
// 示例 2：
//
// 输入： 1->2->2->1
//输出： true
//
//
//
//
// 进阶： 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
// Related Topics 栈 递归 链表 双指针 👍 136 👎 0

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
	// 同 LCR 027
	var down *ListNode
	var dfs func(*ListNode, *ListNode) bool
	dfs = func(slow, fast *ListNode) bool {
		if fast == nil || fast.Next == nil {
			return true
		}
		fast = fast.Next.Next
		if !dfs(slow.Next, fast) {
			return false
		}
		if down == nil {
			if fast == nil { // 偶数
				down = slow.Next
			} else { // 奇数
				down = slow.Next.Next
			}
		}
		if slow.Val != down.Val {
			return false
		}
		down = down.Next
		return true
	}
	return dfs(head, head)
}

//leetcode submit region end(Prohibit modification and deletion)
