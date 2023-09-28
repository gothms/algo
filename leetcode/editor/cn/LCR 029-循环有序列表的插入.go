//给定循环单调非递减列表中的一个点，写一个函数向这个列表中插入一个新元素 insertVal ，使这个列表仍然是循环升序的。
//
// 给定的可以是这个列表中任意一个顶点的指针，并不一定是这个列表中最小元素的指针。
//
// 如果有多个满足条件的插入位置，可以选择任意一个位置插入新的值，插入后整个列表仍然保持有序。
//
// 如果列表为空（给定的节点是 null），需要创建一个循环有序列表并返回这个节点。否则。请返回原先给定的节点。
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [3,4,1], insertVal = 2
//输出：[3,4,1,2]
//解释：在上图中，有一个包含三个元素的循环有序列表，你获得值为 3 的节点的指针，我们需要向表中插入元素 2 。新插入的节点应该在 1 和 3 之间，插入之后
//，整个列表如上图所示，最后返回节点 3 。
//
//
//
//
// 示例 2：
//
//
//输入：head = [], insertVal = 1
//输出：[1]
//解释：列表为空（给定的节点是 null），创建一个循环有序列表并返回这个节点。
//
//
// 示例 3：
//
//
//输入：head = [1], insertVal = 0
//输出：[1,0]
//
//
//
//
// 提示：
//
//
// 0 <= Number of Nodes <= 5 * 10^4
// -10^6 <= Node.val <= 10^6
// -10^6 <= insertVal <= 10^6
//
//
//
//
//
// 注意：本题与主站 708 题相同： https://leetcode-cn.com/problems/insert-into-a-sorted-
//circular-linked-list/
//
// Related Topics 链表 👍 156 👎 0

package main

import "fmt"

func main() {
	node := &Node{3, &Node{5, &Node{1, nil}}}
	x := 0
	n := insert(node, x)
	fmt.Println(n)
}

type Node struct {
	Val  int
	Next *Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(aNode *Node, x int) *Node {
	// lc
	xn := &Node{Val: x}
	if aNode == nil {
		xn.Next = xn
		return xn
	}
	//if aNode.Next == aNode {
	//	aNode.Next, xn.Next = xn, aNode
	//	return aNode
	//}
	pre, cur := aNode, aNode.Next
	for ; cur != aNode; pre, cur = cur, cur.Next {
		if pre.Val <= x && x <= cur.Val || // 插入点
			pre.Val > cur.Val && (pre.Val < x || x < cur.Val) { // 翻转点
			break
		}
	}
	pre.Next, xn.Next = xn, cur // cur == aNode 是最后一个没查找的情况
	return aNode

	// 个人
	//xn := &Node{Val: x}
	//if aNode == nil {
	//	xn.Next = xn
	//	return xn
	//}
	//pre := aNode
	//for cur := aNode; ; {
	//	if x >= cur.Val && x <= cur.Next.Val { // 插入点
	//		cur.Next, xn.Next = xn, cur.Next
	//		return aNode
	//	} else if cur.Val > cur.Next.Val { // 翻转点
	//		pre = cur
	//	}
	//	cur = cur.Next    // ugly
	//	if cur == aNode { // 一圈
	//		break
	//	}
	//}
	//pre.Next, xn.Next = xn, pre.Next // 插入翻转点
	//return aNode
}

//leetcode submit region end(Prohibit modification and deletion)
