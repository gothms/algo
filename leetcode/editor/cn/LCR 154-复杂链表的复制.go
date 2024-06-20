package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//leetcode submit region begin(Prohibit modification and deletion)

func copyRandomList(head *Node) *Node {
	// O(1)空间
	if head == nil {
		return nil
	}
	for cur := head; cur != nil; {
		copy := &Node{Val: cur.Val}
		cur.Next, copy.Next = copy, cur.Next
		cur = copy.Next
	}
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	newHead := head.Next
	for cur := head; cur != nil; cur = cur.Next {
		copy := cur.Next
		cur.Next = copy.Next
		if copy.Next != nil {
			copy.Next = copy.Next.Next
		}
	}
	return newHead

	//memo := make(map[*Node]*Node)
	//pre := &Node{}
	//for p, cur := pre, head; cur != nil; cur = cur.Next {
	//	next := &Node{Val: cur.Val}
	//	p.Next, p = next, next
	//	memo[cur] = next
	//}
	//for p, cur := pre.Next, head; cur != nil; cur = cur.Next {
	//	p.Random, p = memo[cur.Random], p.Next
	//}
	//return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
