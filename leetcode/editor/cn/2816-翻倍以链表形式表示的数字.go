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
func doubleIt(head *ListNode) *ListNode {
	// lc：迭代，计算 cur 时，查看 cur.Next.Val 是否 >=5

	// 个人
	var dfs func(*ListNode) int
	dfs = func(cur *ListNode) int {
		if cur == nil {
			return 0
		}
		v := cur.Val<<1 + dfs(cur.Next)
		cur.Val = v % 10
		return v / 10
	}
	if dfs(head) == 1 {
		return &ListNode{1, head}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
