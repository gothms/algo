//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//
//
// 示例 2：
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
// Related Topics 链表 双指针 分治 排序 归并排序 👍 2011 👎 0

package main

import "fmt"

func main() {
	head := &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{}}}}}
	//head = &ListNode{3, &ListNode{2, &ListNode{4, nil}}}
	list := sortList(head)
	for ; list != nil; list = list.Next {
		fmt.Print(list.Val, " ")
	}
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

// leetcode submit region end(Prohibit modification and deletion)

func sortList(head *ListNode) *ListNode {
	// 进阶：在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序
	// heap 依赖索引，快排 O(log n) 空间复杂度，归并呢？已知合并两条有序链表是 O(1) 空间复杂度
	// 自顶向下：O(log n) 空间复杂度，即递归的深度。类似数组的归并排序，但采用“快慢指针”寻找中点
	// 自底向上：O(1) 空间复杂度，“快慢指针”数结点没有更快，反而增加了繁琐度
	pre := &ListNode{Next: head}
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	merge2List := func(t, l, r *ListNode) *ListNode {
		for l != nil && r != nil {
			if l.Val < r.Val {
				t.Next, t, l = l, l, l.Next
			} else {
				t.Next, t, r = r, r, r.Next
			}
		}
		if l == nil {
			t.Next = r
		} else {
			t.Next = l
		}
		for t.Next != nil { // 找到尾结点
			t = t.Next
		}
		return t
	}
	for i := 1; i < n; i <<= 1 {
		// O(log n)
		for tail, cur := pre, pre.Next; cur != nil; { // tail：前驱 & 尾结点，cur 用于遍历
			// O(n)
			l := cur
			for j := 1; j < i && cur.Next != nil; j++ { // cur.Next != nil
				cur = cur.Next
			}
			cur.Next, cur = nil, cur.Next // cur != nil
			r := cur
			for j := 1; j < i && cur != nil; j++ { // cur != nil，从 1 开始
				cur = cur.Next
			}
			if cur != nil { // 可能 cur == nil
				cur.Next, cur = nil, cur.Next
			}
			// O(n)
			tail = merge2List(tail, l, r) // 返回合并后的尾结点
		}
	}
	return pre.Next
}

//func sortList(head *ListNode) *ListNode {
//	// 自底向上：已 AC，但“快慢指针”数结点没有更快，反而增加了繁琐度
//	if head == nil || head.Next == nil {
//		return head
//	}
//	pre := &ListNode{Next: head}
//	n := 0
//	for h, l, r := pre, head, head.Next; l != nil; {
//		n++
//		if r = l.Next; r == nil {
//			break
//		}
//		n++
//		if l.Val > r.Val {
//			h.Next, l.Next, r.Next, h, l = r, r.Next, l, l, r.Next // 交换位置，并准备下一轮
//		} else {
//			h, l = r, r.Next // 下一轮
//		}
//	}
//	mergeTwoSortedList := func(pre, l, lt, r, rt *ListNode) *ListNode {
//		for l != nil && r != nil {
//			if l.Val < r.Val {
//				pre.Next, pre, l = l, l, l.Next
//			} else {
//				pre.Next, pre, r = r, r, r.Next
//			}
//		}
//		if l == nil {
//			pre.Next = r
//			return rt
//		} else {
//			pre.Next = l
//			return lt
//		}
//	}
//out:
//	for i := 2; i < n; i <<= 1 {
//		cur, slow, fast := pre, pre.Next, pre.Next.Next // 提前跳一“步”
//		for slow != nil {
//			l := slow
//			for j := 1; j < i && slow != nil; j++ {
//				slow = slow.Next
//				if fast != nil {
//					if fast = fast.Next; fast != nil {
//						fast = fast.Next
//					}
//				}
//			}
//			if slow == nil {
//				continue out
//			}
//			lt, r, rt := slow, slow.Next, fast // 左尾、右头、右尾
//			lt.Next = nil                      // 掐断
//			if r == nil {
//				continue out
//			}
//			if fast == nil {
//				slow = nil // 最后的合并了
//			} else {
//				slow, fast = fast.Next, fast.Next // 提前跳一“步”
//				if fast != nil {
//					fast = fast.Next
//				}
//				rt.Next = nil // 掐断
//			}
//			cur = mergeTwoSortedList(cur, l, lt, r, rt) // 返回尾结点
//			if cur != nil {
//				cur.Next = slow // 连接被掐断的
//			}
//		}
//	}
//	return pre.Next
//}
