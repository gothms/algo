//将一个 二叉搜索树 就地转化为一个 已排序的双向循环链表 。
//
// 对于双向循环列表，你可以将左右孩子指针作为双向循环链表的前驱和后继指针，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。
//
// 特别地，我们希望可以 就地 完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中最小元素的指针。
//
//
//
// 示例 1：
//
//
//输入：root = [4,2,5,1,3]
//
//
//输出：[1,2,3,4,5]
//
//解释：下图显示了转化后的二叉搜索树，实线表示后继关系，虚线表示前驱关系。
//
//
//
// 示例 2：
//
//
//输入：root = [2,1,3]
//输出：[1,2,3]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//解释：输入是空树，所以输出也是空链表。
//
//
// 示例 4：
//
//
//输入：root = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// -1000 <= Node.val <= 1000
// Node.left.val < Node.val < Node.right.val
// Node.val 的所有值都是独一无二的
// 0 <= Number of Nodes <= 2000
//
//
// 注意：本题与主站 426 题相同：https://leetcode-cn.com/problems/convert-binary-search-tree-
//to-sorted-doubly-linked-list/
//
// Related Topics 栈 树 深度优先搜索 二叉搜索树 链表 二叉树 双向链表 👍 734 👎 0

package main

func main() {

}

// There is no code of Go type for this problem
// 牛客：https://www.nowcoder.com/practice/947f6eb80d944a84850b0538bf0ec3a5?tpId=13&tqId=11179&ru=/exam/oj
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	list := &TreeNode{}
	for cur, head := root, list; cur != nil; {
		if cur.Left == nil {
			head.Right, cur.Left, head, cur = cur, head, cur, cur.Right
		} else {
			pre := cur.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			cur, cur.Left, pre.Right = cur.Left, nil, cur
		}
	}
	ans := list.Right
	list.Right, ans.Left = nil, nil
	return ans
}

//func treeToDoublyList(root *TreeNode) *TreeNode {
//	var head, pre *TreeNode          // 新的头节点，pre 节点
//	head.Left, pre.Right = pre, head // 环形链表：error 空指针
//	var dfs func(*TreeNode)
//	dfs = func(cur *TreeNode) {
//		if cur == nil {
//			return
//		}
//		dfs(cur.Left)
//		if pre == nil {
//			head = cur // 最小值
//		} else {
//			pre.Right = cur
//		}
//		cur.Left, pre = pre, cur
//		dfs(cur.Right)
//	}
//	dfs(root)
//	return head
//}
