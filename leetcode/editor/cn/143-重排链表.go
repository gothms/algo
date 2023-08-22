//给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
//
//L0 → L1 → … → Ln - 1 → Ln
//
//
// 请将其重新排列后变为：
//
//
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [1,2,3,4]
//输出：[1,4,2,3]
//
// 示例 2：
//
//
//
//
//输入：head = [1,2,3,4,5]
//输出：[1,5,2,4,3]
//
//
//
// 提示：
//
//
// 链表的长度范围为 [1, 5 * 10⁴]
// 1 <= node.val <= 1000
//
//
// Related Topics 栈 递归 链表 双指针 👍 1247 👎 0

package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
思路：dfs
	1.可以把链表对折，和头结点“对齐”的节点，就是需要插入头结点后面的结点
		先求出链表的中点，然后按中点将链表对折，那么怎么对折呢？
	2.dfs
		对中点进行dfs，当遍历到链表的尾时，同时返回“头结点”
		即dfs的参数节点是“递”，返回值的“头节点”是“归”
		对子结点的dfs返回节点进行处理，将“递”结点插入“归”结点的后面
思路：双指针
	1.将结点放入一个缓存列表中
	2.对列表首尾进行遍历，向中间缩进，并将尾插入头的后面
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	// dfs
	slow, fast := head, head
	for fast != nil && fast.Next != nil { // 找到中点 slow
		slow, fast = slow.Next, fast.Next.Next
	}
	var dfs func(*ListNode) *ListNode
	dfs = func(s *ListNode) *ListNode {
		if s.Next == nil {
			return head // 从头开始归
		}
		// 写法一：立即截断当前尾结点的尾巴
		//next := s.Next
		//s.Next = nil // s 为当前尾结点，需要先断开
		//t := dfs(next)
		//t.Next, next.Next = next, t.Next
		//return next.Next
		// 写法二：不截断
		h := dfs(s.Next)
		h.Next, s.Next.Next = s.Next, h.Next // 插入节点
		return s.Next.Next                   // 返回下个“头”
	}
	dfs(slow)
	slow.Next = nil // 截断尾巴

	// 双指针
	//nodes := make([]*ListNode, 0)
	//for ; head != nil; head = head.Next {
	//	nodes = append(nodes, head)	// 所有结点放入列表
	//}
	//for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
	//	nodes[i].Next, nodes[j].Next = nodes[j], nodes[i+1]
	//}	// 尾结点插入头结点的后面
	//nodes[len(nodes)>>1].Next = nil	// 掐断尾巴

	// 栈
	//s, f, i := head, head, 0
	//for f != nil && f.Next != nil {
	//	s, f = s.Next, f.Next.Next
	//	i++
	//}
	//stack := make([]*ListNode, 0)
	//for ; s != nil; s = s.Next {
	//	stack = append(stack, s)
	//}
	//for j := len(stack) - 1; i > 0; i, j = i-1, j-1 {
	//	head.Next, stack[j].Next, head = stack[j], head.Next, head.Next
	//}
	//stack[0].Next = nil
}

//leetcode submit region end(Prohibit modification and deletion)
