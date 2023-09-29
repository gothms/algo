//编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
//
// 示例1:
//
//
// 输入：[1, 2, 3, 3, 2, 1]
// 输出：[1, 2, 3]
//
//
// 示例2:
//
//
// 输入：[1, 1, 1, 1, 2]
// 输出：[1, 2]
//
//
// 提示：
//
//
// 链表长度在[0, 20000]范围内。
// 链表元素在[0, 20000]范围内。
//
//
// 进阶：
//
// 如果不得使用临时缓冲区，该怎么解决？
//
// Related Topics 哈希表 链表 双指针 👍 191 👎 0

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
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	memo := map[int]bool{head.Val: true}
	for cur, next := head, head.Next; next != nil; next = next.Next {
		if memo[next.Val] {
			cur.Next = next.Next // 已存在，删除，且 cur 不动
		} else {
			memo[next.Val] = true // 不存在，记录
			cur = next
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
