//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
//输入：l1 = [], l2 = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
//
// Related Topics 递归 链表 👍 3205 👎 0

package main

func main() {

}

/*
思路：
	总是将当前最小的结点，作为当前的头结点，然后进行下一次比较排序
	若一条链尾 nil 时，将另外一条链作为尾巴添加到结尾
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 递归
	//if list1 == nil {
	//	return list2 // 尾巴
	//}
	//if list2 == nil {
	//	return list1
	//}
	//if list1.Val < list2.Val {
	//	list1.Next = mergeTwoLists(list1.Next, list2) // 排序 list1
	//	return list1                                  // 返回当前最小结点
	//} else {
	//	list2.Next = mergeTwoLists(list1, list2.Next)
	//	return list2
	//}

	// 迭代
	pre := &ListNode{} // 哨兵
	curr := pre
	for list1 != nil && list2 != nil { // 排序
		if list1.Val < list2.Val {
			curr.Next, curr = list1, list1
			list1 = list1.Next
		} else {
			curr.Next, curr = list2, list2
			list2 = list2.Next
		}
	}
	if list1 != nil { // 收尾
		curr.Next = list1
	} else {
		curr.Next = list2
	}
	return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
