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
func swapNodes(head *ListNode, k int) *ListNode {
	fast := head
	for i := 1; i < k; i++ {
		fast = fast.Next
	}
	f, slow := fast, head
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	f.Val, slow.Val = slow.Val, f.Val
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
