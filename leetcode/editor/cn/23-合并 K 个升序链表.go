//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
//
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 2463 👎 0

package main

func main() {

}

/*
思路：归并
	1.利用分治思想，将 n 条链表分为两个链表或一条链表（不用排序）
	2.两两排序后再归并，直至只剩最后一条链表
		这里排序和归并都用递归实现
思路：堆
	1.将 n 条链表的头节点在小顶堆中堆化，选出最小节点
	2.排序
		依次将最小节点取出组成新的链表
		取出堆顶元素时，如果它有next节点，则将next节点放入堆中
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {

}

//leetcode submit region end(Prohibit modification and deletion)

func mergeKLists(lists []*ListNode) *ListNode {
	// 堆
	//var pre ListNode
	//m := &mkl{}
	//for _, l := range lists {
	//	if l != nil {
	//		heap.Push(m, l) // 初始化堆数据
	//	}
	//}
	//for curr := &pre; m.Len() > 0; curr = curr.Next {
	//	next := heap.Pop(m).(*ListNode) // 取出最小节点
	//	if next.Next != nil {
	//		heap.Push(m, next.Next) // 放入下一个节点
	//	}
	//	curr.Next = next // 排序
	//}
	//return pre.Next

	// 归并
	if len(lists) == 0 {
		return nil
	}
	return mergeK(lists, 0, len(lists)-1)
}

type mkl []*ListNode

func (m mkl) Len() int           { return len(m) }
func (m mkl) Less(i, j int) bool { return m[i].Val < m[j].Val }
func (m mkl) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *mkl) Push(x any)        { *m = append(*m, x.(*ListNode)) }
func (m *mkl) Pop() any {
	ln := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return ln
}

func mergeK(lists []*ListNode, l, r int) *ListNode {
	switch r - l {
	case 0:
		return lists[l]
	case 1:
		return mergeList(lists[l], lists[r]) // 合并两个链表
	default:
		m := (r + l) >> 1 // 分治
		return mergeList(mergeK(lists, l, m), mergeK(lists, m+1, r))
	}
}
func mergeList(a, b *ListNode) *ListNode { // 递归写法
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val < b.Val {
		a.Next = mergeList(a.Next, b)
		return a
	} else {
		b.Next = mergeList(a, b.Next)
		return b
	}
}
