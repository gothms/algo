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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// lc
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a

	// 个人
	//if headA == nil || headB == nil {
	//	return nil
	//}
	//cur := headB
	//for cur.Next != nil {
	//	cur = cur.Next
	//}
	//cur.Next = headB
	//defer func() { cur.Next = nil }()
	//
	//fast, slow := headA, headA
	//for fast != nil && fast.Next != nil { // fast == slow：headA 找到环
	//	fast, slow = fast.Next.Next, slow.Next
	//	if fast == slow {
	//		break
	//	}
	//}
	//if fast == nil || fast.Next == nil {
	//	return nil
	//}
	//for fast = headA; fast != slow; {
	//	fast, slow = fast.Next, slow.Next
	//}
	//return fast
}

//leetcode submit region end(Prohibit modification and deletion)
