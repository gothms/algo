//给定单个链表的头
// head ，使用 插入排序 对链表进行排序，并返回 排序后链表的头 。
//
// 插入排序 算法的步骤:
//
//
// 插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
// 每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
// 重复直到所有输入数据插入完为止。
//
//
// 下面是插入排序算法的一个图形示例。部分排序的列表(黑色)最初只包含列表中的第一个元素。每次迭代时，从输入数据中删除一个元素(红色)，并就地插入已排序的列表
//中。
//
// 对链表进行插入排序。
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
//输入: head = [4,2,1,3]
//输出: [1,2,3,4]
//
// 示例 2：
//
//
//
//
//输入: head = [-1,5,3,4,0]
//输出: [-1,0,3,4,5]
//
//
//
// 提示：
//
//
//
//
//
// 列表中的节点数在 [1, 5000]范围内
// -5000 <= Node.val <= 5000
//
//
// Related Topics 链表 排序 👍 618 👎 0

package main

func main() {

}

/*
思路：迭代
	1.哨兵 pre 为新的头结点，避免插入点是在头节点之前
	2.使用插入排序的思路，不同的是，链表需要从头开始找插入点
		2.1.preV = pre，遍历 preV 的 Next 结点
			直到 preV.Next.Val >= 当前结点 curr.Val
		2.2.在 preV 和 preV.Next 结点之间插入 curr 结点
			由于要删除 curr 结点，所以要记录 curr 的前驱结点，这里就用 head
	3.链表的操作关键在于指针的变更
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre := &ListNode{0, head}
	// preV 是要插入位置的前驱结点，curr 是当前遍历的节点，head 是 curr 的前驱结点
	for curr, preV := head.Next, pre; curr != nil; curr, preV = head.Next, pre {
		for preV.Next.Val < curr.Val { // 寻找插入点
			preV = preV.Next
		}
		if preV.Next != curr { // 变更顺序，preV 后面插入 curr 结点，head 后面删除 curr 结点
			head.Next, preV.Next, curr.Next = curr.Next, curr, preV.Next
		} else { // 不需要变更顺序
			head = curr
		}
	}
	return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
