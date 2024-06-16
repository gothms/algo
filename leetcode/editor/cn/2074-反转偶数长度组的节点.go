package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{1, &ListNode{0, &ListNode{6, &ListNode{5, nil}}}}}
	// [4,3,0,5,1,2,7,8,6]
	// [4,0,3,5,1,2,7,8,6]
	head = &ListNode{4, &ListNode{3, &ListNode{0, &ListNode{5, &ListNode{1,
		&ListNode{2, &ListNode{7, &ListNode{8, &ListNode{6, nil}}}}}}}}}
	groups := reverseEvenLengthGroups(head)
	for groups != nil {
		fmt.Println(groups.Val)
		groups = groups.Next
	}
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//leetcode submit region begin(Prohibit modification and deletion)

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	i, j := 1, 0
	pre := head
	for cur := head; cur != nil; cur = cur.Next {
		if i&1 == 0 {
			for j = 1; j < i && cur.Next != nil; j++ { // 只需要反转 i-1 个结点
				pre.Next, cur.Next.Next, cur.Next = cur.Next, pre.Next, cur.Next.Next
			}
		} else {
			for j = 1; j < i && cur.Next != nil; j++ { // 留一个结点做 pre
				cur = cur.Next
			}
		}
		if j == i { // 循环正常完成
			pre = cur
			i++
		}
	}
	if (i^j)&1 == 1 && pre.Next != nil { // (i^j)&1：i 和 j 的奇偶性相反，则末尾的序列需要反一下
		for cur := pre.Next; cur.Next != nil; {
			pre.Next, cur.Next.Next, cur.Next = cur.Next, pre.Next, cur.Next.Next
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
