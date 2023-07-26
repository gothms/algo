//给你一个头结点为 head 的单链表和一个整数 k ，请你设计一个算法将链表分隔为 k 个连续的部分。
//
// 每部分的长度应该尽可能的相等：任意两部分的长度差距不能超过 1 。这可能会导致有些部分为 null 。
//
// 这 k 个部分应该按照在链表中出现的顺序排列，并且排在前面的部分的长度应该大于或等于排在后面的长度。
//
// 返回一个由上述 k 部分组成的数组。
//
// 示例 1：
//
//
//输入：head = [1,2,3], k = 5
//输出：[[1],[2],[3],[],[]]
//解释：
//第一个元素 output[0] 为 output[0].val = 1 ，output[0].next = null 。
//最后一个元素 output[4] 为 null ，但它作为 ListNode 的字符串表示是 [] 。
//
//
// 示例 2：
//
//
//输入：head = [1,2,3,4,5,6,7,8,9,10], k = 3
//输出：[[1,2,3,4],[5,6,7],[8,9,10]]
//解释：
//输入被分成了几个连续的部分，并且每部分的长度相差不超过 1 。前面部分的长度大于等于后面部分的长度。
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 1000]
// 0 <= Node.val <= 1000
// 1 <= k <= 50
//
//
// Related Topics 链表 👍 286 👎 0

package main

import "fmt"

func main() {
	tail := &ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, &ListNode{10, nil}}}}}
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, tail}}}}}
	k := 3
	parts := splitListToParts(head, k)
	fmt.Println(parts)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
思路：
	1.每个子链的长度大于等于 n/k
	2.贪心
		多余了 m=n%k 个节点，均分给前 m 个子链，所以前 m 个子链长度为 n/k+1
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	// 代码简化
	split, n := make([]*ListNode, k), 0
	for curr := head; curr != nil; curr = curr.Next {
		n++ // 计算链表长度
	}
	for length, l, m, i := n/k, 0, n%k, 0; head != nil; i++ {
		split[i], l = head, length // 存储子链的头
		if m > 0 {                 // 靠前的子链长度是 length+1，靠后的子链长度是 length
			l++
			m-- // 一共有 m 个靠前的子链
		}
		for j := 1; j < l; j++ { // 少遍历一个节点，指针停在每个子链的尾节点
			head = head.Next
		}
		head, head.Next = head.Next, nil // 截断链表
	}
	return split

	//split, n := make([]*ListNode, k), 0
	//if head == nil {
	//	return split
	//}
	//split[0] = head
	//for curr := head; curr != nil; curr = curr.Next {
	//	n++
	//}
	//kPre, M, i := n/k, n%k, 0
	//curr := head
	//for m := M; m > 0; m-- {
	//	split[i], i = curr, i+1
	//	for j := 0; j < kPre; j++ {
	//		curr = curr.Next
	//	}
	//	curr, curr.Next = curr.Next, nil
	//}
	//kTail := kPre - 1
	//if kTail < 0 {
	//	return split
	//}
	//for m := k - M - 1; m > 0; m-- {
	//	split[i], i = curr, i+1
	//	for j := 0; j < kTail; j++ {
	//		curr = curr.Next
	//	}
	//	curr, curr.Next = curr.Next, nil
	//}
	//split[i] = curr
	//return split
}

//leetcode submit region end(Prohibit modification and deletion)
