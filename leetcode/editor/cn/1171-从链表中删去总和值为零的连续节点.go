//给你一个链表的头节点 head，请你编写代码，反复删去链表中由 总和 值为 0 的连续节点组成的序列，直到不存在这样的序列为止。
//
// 删除完毕后，请你返回最终结果链表的头节点。
//
//
//
// 你可以返回任何满足题目要求的答案。
//
// （注意，下面示例中的所有序列，都是对 ListNode 对象序列化的表示。）
//
// 示例 1：
//
// 输入：head = [1,2,-3,3,1]
//输出：[3,1]
//提示：答案 [1,2,1] 也是正确的。
//
//
// 示例 2：
//
// 输入：head = [1,2,3,-3,4]
//输出：[1,2,4]
//
//
// 示例 3：
//
// 输入：head = [1,2,3,-3,-2]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 给你的链表中可能有 1 到 1000 个节点。
// 对于链表中的每个节点，节点的值：-1000 <= node.val <= 1000.
//
//
// Related Topics 哈希表 链表 👍 194 👎 0

package main

func main() {

}

/*
思路：hash
	1.如果链表的某一段 s-t 和为 0，则应该把这段截掉
		s节点的前驱节点.Next = t节点.Next
		所以在删除时，同时要记录的是 前驱节点
	2.s-t 这段节点的和为 0，那么从 head 到 s 节点的前驱节点和 sum(s.pre) == sum(t)
		若有很多个 t1 t2 ... 满足 sum(t) == sum(s.pre)
		可以总是 贪心 的选择最后一个 t
	3.综上
		3.1.一次遍历
			使用map记录链表的 前缀和：key = sum，value = 前驱节点的指针
			sum相等时，后出现的节点，会覆盖先出现的节点
		3.2.再次遍历
			如果 map[sum] 记录的节点，不是当前节点 curr
			说明 s.pre 到 t 的节点的和 == 0，删除操作 curr.Next = map[sum].Next
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	pre, preSum := &ListNode{0, head}, make(map[int]*ListNode)
	for curr, sum := pre, 0; curr != nil; curr = curr.Next {
		sum += curr.Val
		preSum[sum] = curr
	}
	for curr, sum := pre, 0; curr != nil; curr = curr.Next {
		sum += curr.Val
		//curr.Next = preSum[sum].Next
		if node := preSum[sum]; node != curr {
			curr.Next = node.Next
		}
	}
	return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
