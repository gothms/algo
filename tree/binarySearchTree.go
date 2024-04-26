package tree

/*
二叉查找树的其他操作
	查找最大节点和最小节点、前驱节点和后继节点
	中序遍历二叉查找树，可以输出有序的数据序列，时间复杂度是 O(n)，非常高效
*/

// FindBinarySearchTree 二叉查找树的查找操作
func FindBinarySearchTree(root *TreeNode, v int) *TreeNode {
	for cur := root; cur != nil; {
		switch {
		case cur.Val > v:
			cur = cur.Left
		case cur.Val < v:
			cur = cur.Right
		case cur.Val == v:
			return cur
		}
	}
	return nil // 不存在
}

// InsertBinarySearchTree 二叉查找树的插入操作
// 新插入的数据一般都是在叶子节点上
func InsertBinarySearchTree(root *TreeNode, v int) {
	if root == nil {
		root = &TreeNode{Val: v}
		return
	}
	for cur := root; cur != nil; {
		switch {
		case cur.Val > v:
			if cur.Left == nil {
				cur.Left = &TreeNode{Val: v}
				return
			}
			cur = cur.Left
		case cur.Val < v:
			if cur.Right == nil {
				cur.Right = &TreeNode{Val: v}
				return
			}
			cur = cur.Right
		case cur.Val == v: // 已存在
			return
		}
	}
}

// DeleteBinarySearchTree 二叉查找树的删除操作
// 1.如果要删除的节点没有子节点
// 只需要直接将父节点中，指向要删除节点的指针置为 nil
// 2.如果要删除的节点只有一个子节点（只有左子节点或者右子节点）
// 只需要更新父节点中，指向要删除节点的指针，让它指向要删除节点的子节点
// 3.如果要删除的节点有两个子节点
// 需要找到这个节点的右子树中的最小节点，把它替换到要删除的节点上
// 再删除掉这个最小节点，因为最小节点肯定没有左子节点，应用上面两条规则来删除这个最小节点
// 还有个非常简单、取巧的方法，就是单纯将要删除的节点标记为“已删除”，但是并不真正从树中将这个节点去掉
// 这样原本删除的节点还需要存储在内存中，比较浪费内存空间，但是删除操作就变得简单了很多
func DeleteBinarySearchTree(root *TreeNode, v int) bool {
	//var (
	//	pre *TreeNode // 记录 cur 的父节点
	//	cur = root    // 当前节点，指向要删除的节点
	//)
	//for cur != nil && cur.Val != v { // 查询 v 节点
	//	pre = cur
	//	if v < cur.Val {
	//		cur = cur.Left
	//	} else {
	//		cur = cur.Right
	//	}
	//}

	var pre *TreeNode                    // 记录 cur 的父节点
	cur := FindBinarySearchTree(root, v) // 当前节点，指向要删除的节点
	if cur == nil {                      // 不存在
		return false
	}

	if cur.Left != nil && cur.Right != nil { // 要删除的节点有两个子节点
		p, c := cur, cur.Right // 查找右子树中最小节点 c
		for c.Left != nil {
			p, c = c, c.Left
		}
		cur.Val, pre, cur = c.Val, p, c // 将 c 的数据替换到 cur 中，变成删除 c
	}
	// 要删除的节点是叶子节点，或只有一个子节点
	var child *TreeNode   // pre 的子节点
	if cur.Right != nil { // 有两个子节点时，只可能有右节点/没有子节点（所以先判断 cur.Right）
		child = cur.Right
	} else if cur.Left != nil {
		child = cur.Left
	} // else 没有子节点

	if pre == nil { // 删除目标是根节点，且根节点只有一个子节点/没有子节点
		root = child
	} else if pre.Left == cur {
		pre.Left = child
	} else {
		pre.Right = child
	}
	return true
}
