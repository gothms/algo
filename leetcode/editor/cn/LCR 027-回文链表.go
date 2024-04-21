//给定一个链表的 头节点 head ，请判断其是否为回文链表。
//
// 如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。
//
//
//
// 示例 1：
//
//
//
//
//输入: head = [1,2,3,3,2,1]
//输出: true
//
// 示例 2：
//
//
//
//
//输入: head = [1,2]
//输出: false
//
//
//
//
// 提示：
//
//
// 链表 L 的长度范围为 [1, 10⁵]
// 0 <= node.val <= 9
//
//
//
//
// 进阶：能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
//
//
//
// 注意：本题与主站 234 题相同：https://leetcode-cn.com/problems/palindrome-linked-list/
//
// Related Topics 栈 递归 链表 双指针 👍 124 👎 0

package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	head = &ListNode{1, &ListNode{0, &ListNode{1, nil}}}
	palindrome := isPalindrome(head)
	fmt.Println(palindrome)
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
func isPalindrome(head *ListNode) bool {
	var mid *ListNode
	var dfs func(*ListNode, *ListNode) bool
	dfs = func(slow, fast *ListNode) (ret bool) {
		if fast == nil || fast.Next == nil { // 找到中点就终止 dfs
			return true
		}
		fast = fast.Next.Next
		if !dfs(slow.Next, fast) { // 保证 slow 慢一拍，才能回溯到 slow.Val
			return false
		}
		if mid == nil { // mid 只赋值一次
			if fast == nil {
				mid = slow.Next // 偶数个结点
			} else {
				mid = slow.Next.Next // 奇数个结点
			}
		}
		ret, mid = mid.Val == slow.Val, mid.Next // 结点是否回文：mid 往右，slow 往左
		return
	}
	return dfs(head, head)
}

//leetcode submit region end(Prohibit modification and deletion)
