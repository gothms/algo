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
func nodesBetweenCriticalPoints(head *ListNode) []int {
	i, l, r, mn := 1, 0, 0, 0
	for pre, cur := head.Val, head.Next; cur.Next != nil; pre, cur = cur.Val, cur.Next {
		if cur.Val > pre && cur.Val > cur.Next.Val ||
			cur.Val < pre && cur.Val < cur.Next.Val {
			if l == 0 {
				l = i
			} else if r == 0 {
				r = i
				mn = r - l
			} else {
				mn = min(mn, i-r)
				r = i
			}
		}
		i++
	}
	if r > 0 {
		return []int{mn, r - l}
	}
	return []int{-1, -1}
}

//leetcode submit region end(Prohibit modification and deletion)
