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
func getDecimalValue(head *ListNode) int {
	v := 0
	for head != nil {
		v = v<<1 + head.Val
		head = head.Next
	}
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
