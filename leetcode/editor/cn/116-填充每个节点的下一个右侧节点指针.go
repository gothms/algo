//给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
//
//
//struct Node {
//  int val;
//  Node *left;
//  Node *right;
//  Node *next;
//}
//
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//
// 初始状态下，所有 next 指针都被设置为 NULL。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,2,3,4,5,6,7]
//输出：[1,#,2,3,#,4,5,6,7,#]
//解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由
//next 指针连接，'#' 标志着每一层的结束。
//
//
//
//
//
// 示例 2:
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点的数量在
// [0, 2¹² - 1] 范围内
// -1000 <= node.val <= 1000
//
//
//
//
// 进阶：
//
//
// 你只能使用常量级额外空间。
// 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 链表 二叉树 👍 1031 👎 0

package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
思路：dfs
	1.节点的 Next 节点为同层的“右边”节点，反过来思考
		可以记录同层的“左边”节点到一个数组 pre 中，用“层” level 作为id标识
		当前节点为 r，那么 pre[level].Next = r，再将 r 放入 pre[level] 位置备下一个 Next 使用
	2.可以在dfs中完成上面的逻辑
思路：bfs
	1.根据dfs算法的启发，我们可以将同层的节点先构造成链表
		然后遍历链表，链表节点的左右子节点，将它们串成链表
		即：当前层已经是链表，遍历当前层，把下一层也串成链表，这个过程就像层序遍历
	2.怎么把下一层串成链表？
		遍历当前层时，先将节点的左右节点串起来：curr.Left.Next = curr.Right
		再将 curr 的右节点，与curr.Next的左节点串起来：curr.Right.Next = curr.Next.Left
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	// bfs：迭代
	//if root == nil {
	//	return nil
	//}
	//for leftNode := root; leftNode.Left != nil; leftNode = leftNode.Left {
	//	for curr := leftNode; curr != nil; curr = curr.Next {
	//		curr.Left.Next = curr.Right
	//		if curr.Next != nil {
	//			curr.Right.Next = curr.Next.Left
	//		}
	//	}
	//}
	//return root

	// bfs：递归
	if root == nil {
		return nil
	}
	var bfs func(*Node)
	bfs = func(r *Node) {
		if r.Left == nil {
			return
		}
		for curr := r; curr != nil; curr = curr.Next {
			curr.Left.Next = curr.Right // 串联节点的左右节点
			if curr.Next != nil {       // 串联节点的右节点，和 节点.Next 的左节点
				curr.Right.Next = curr.Next.Left
			}
		}
		bfs(r.Left) // 层序遍历的“层头”
	}
	bfs(root)
	return root

	// dfs
	//pre := make([]*Node, 0) // 每层的节点
	//var dfs func(*Node, int)
	//dfs = func(r *Node, i int) {
	//	if r == nil {
	//		return
	//	}
	//	if i == len(pre) { // 层的第一个节点（最左节点），它没有pre节点
	//		pre = append(pre, r)
	//	} else { // 当前节点的 pre 节点为同层的上一个节点
	//		pre[i].Next, pre[i] = r, r
	//	}
	//	dfs(r.Left, i+1) // dfs逻辑
	//	dfs(r.Right, i+1)
	//}
	//dfs(root, 0)
	//return root
}

//leetcode submit region end(Prohibit modification and deletion)
