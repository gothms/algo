//给定链表头结点 head，该链表上的每个结点都有一个 唯一的整型值 。同时给定列表 nums，该列表是上述链表中整型值的一个子集。
//
// 返回列表 nums 中组件的个数，这里对组件的定义为：链表中一段最长连续结点的值（该值必须在列表 nums 中）构成的集合。
//
//
//
// 示例 1：
//
//
//
//
//输入: head = [0,1,2,3], nums = [0,1,3]
//输出: 2
//解释: 链表中,0 和 1 是相连接的，且 nums 中不包含 2，所以 [0, 1] 是 nums 的一个组件，同理 [3] 也是一个组件，故返回 2。
//
//
// 示例 2：
//
//
//
//
//输入: head = [0,1,2,3,4], nums = [0,3,1,4]
//输出: 2
//解释: 链表中，0 和 1 是相连接的，3 和 4 是相连接的，所以 [0, 1] 和 [3, 4] 是两个组件，故返回 2。
//
//
//
// 提示：
//
//
// 链表中节点数为n
// 1 <= n <= 10⁴
// 0 <= Node.val < n
// Node.val 中所有值 不同
// 1 <= nums.length <= n
// 0 <= nums[i] < n
// nums 中所有值 不同
//
//
// Related Topics 数组 哈希表 链表 👍 194 👎 0

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
func numComponents(head *ListNode, nums []int) int {
	// 思路是可行的，但是“没法”统计结尾有多少个结点不在集合中
	//memo := make(map[int]bool, len(nums))
	//for _, v := range nums {
	//	memo[v] = true
	//}
	//cnt, n := 1, len(nums)
	//if !memo[head.Val] {
	//	n++
	//}
	//for ; head.Next != nil; head = head.Next {
	//	cnt++
	//}
	//if !memo[head.Val] {
	//	n++
	//}
	//return cnt - n + 1

	memo := make(map[int]bool, len(nums))
	for _, v := range nums {
		memo[v] = true
	}
	nc, cnt := 0, false
	for ; head != nil; head = head.Next {
		if memo[head.Val] {
			cnt = true
		} else {
			if cnt { // 求的是组件的个数
				nc++
			}
			cnt = false
		}
	}
	if cnt {
		return nc + 1
	}
	return nc
}

//leetcode submit region end(Prohibit modification and deletion)
