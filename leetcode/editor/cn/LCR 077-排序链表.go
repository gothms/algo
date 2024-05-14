//给定链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
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
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//
//
// 示例 2：
//
//
//
//
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
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
// 链表中节点的数目在范围 [0, 5 * 10⁴] 内
// -10⁵ <= Node.val <= 10⁵
//
//
//
//
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
//
//
//
//
// 注意：本题与主站 148 题相同：https://leetcode-cn.com/problems/sort-list/
//
// Related Topics 链表 双指针 分治 排序 归并排序 👍 131 👎 0

package main

import "fmt"

func main() {
	head := &ListNode{2, &ListNode{1, &ListNode{-3, nil}}}
	list := sortList(head)
	fmt.Println(list)
	//for list != nil {
	//	fmt.Println(list.Val)
	//}
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
func sortList(head *ListNode) *ListNode {

}

//leetcode submit region end(Prohibit modification and deletion)

//func sortList(head *ListNode) *ListNode {
//	// 归并
//	if head == nil || head.Next == nil {
//		return head
//	}
//	n := 1
//	for cur := head.Next; cur != nil; cur = cur.Next {
//		n++ // 统计结点数
//	}
//	mergeList := func(l, r *ListNode) *ListNode { // 合并链表
//		pre := &ListNode{}
//		cur := pre
//		for r != nil {
//			for l != nil && l.Val < r.Val {
//				cur.Next, cur, l = l, l, l.Next
//			}
//			cur.Next, cur, r = r, r, r.Next
//		}
//		cur.Next = l // 收尾
//		return pre.Next
//	}
//	var dfs func(*ListNode, int) *ListNode
//	dfs = func(ln *ListNode, n int) *ListNode {
//		if n <= 1 {
//			return ln
//		}
//		m := (n + 1) >> 1 // 左边链表的结点数
//		cur := ln
//		for i := 1; i < m; i++ { // 左边的链表
//			cur = cur.Next
//		}
//		cur, cur.Next = cur.Next, nil               // 断开
//		return mergeList(dfs(ln, m), dfs(cur, n-m)) // 分治 & 归并
//	}
//	return dfs(head, n)
//}
