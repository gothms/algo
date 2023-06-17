//给出二叉树的根节点 root，树上每个节点都有一个不同的值。
//
// 如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。
//
// 返回森林中的每棵树。你可以按任意顺序组织答案。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,2,3,4,5,6,7], to_delete = [3,5]
//输出：[[1,2,null,4],[6],[7]]
//
//
// 示例 2：
//
//
//输入：root = [1,2,4,null,3], to_delete = [3]
//输出：[[1,2,4]]
//
//
//
//
// 提示：
//
//
// 树中的节点数最大为 1000。
// 每个节点都有一个介于 1 到 1000 之间的值，且各不相同。
// to_delete.length <= 1000
// to_delete 包含一些从 1 到 1000、各不相同的值。
//
//
// Related Topics 树 深度优先搜索 数组 哈希表 二叉树 👍 208 👎 0

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路：dfs
	1.和删除二叉树的一个节点类似，只是不要求删除后再接回去
		那么可以把要删除的节点的 left right 添加到结果集中
	2.细节处理
		2.1.把要删除的节点放入map，实现 O(1) 的查询效率
		2.2.当前节点是否添加到结果集中，取决于父节点是否被删除
			用布尔值 del 标记父节点是否删除，并告知当前节点
		2.3.结合2.2.
			a)当前节点需要记录一个布尔值 nDel，标记当前节点是否被删除
			b)如果当前节点被删除，可以返回 nil 给父节点，省去大量 if 判断
				可以参考 二叉查找树 删除节点的一个做法：将父节点传给子节点
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	ts, memo := make([]*TreeNode, 0), make(map[int]bool)
	for _, v := range to_delete {
		memo[v] = true
	}
	var dfs func(*TreeNode, bool) *TreeNode
	dfs = func(r *TreeNode, del bool) *TreeNode {
		if r == nil {
			return nil
		}
		nDel := memo[r.Val]
		r.Left, r.Right = dfs(r.Left, nDel), dfs(r.Right, nDel)
		if nDel {
			return nil
		}
		if del {
			ts = append(ts, r)
		}
		return r
	}
	dfs(root, true)
	return ts
}
	// 优化
	//ts, memo := make([]*TreeNode, 0), make(map[int]bool)
	//for _, v := range to_delete {
	//	memo[v] = true
	//}
	//pivot := &TreeNode{Left: root}
	//var dfs func(*TreeNode, *TreeNode, bool)
	//dfs = func(r *TreeNode, p *TreeNode, del bool) {
	//	if r == nil {
	//		return
	//	}
	//	nDel := memo[r.Val]
	//	dfs(r.Left, r, nDel)
	//	dfs(r.Right, r, nDel)
	//	if del && !nDel {
	//		ts = append(ts, r)
	//	}
	//	if nDel {
	//		if p.Left == r {
	//			p.Left = nil
	//		} else {
	//			p.Right = nil
	//		}
	//	}
	//}
	//dfs(root, pivot, memo[root.Val])
	//if pivot.Left != nil {
	//	ts = append(ts, pivot.Left)
	//}
	//return ts

	//ts, memo := make([]*TreeNode, 0), make(map[int]bool)
	//for _, v := range to_delete {
	//	memo[v] = true
	//}
	//pivot := &TreeNode{Left: root}
	//var dfs func(*TreeNode, *TreeNode)
	//dfs = func(r *TreeNode, p *TreeNode) {
	//	if r == nil {
	//		return
	//	}
	//	dfs(r.Left, r)
	//	dfs(r.Right, r)
	//	if memo[r.Val] {
	//		//delete(memo, r.Val)
	//		if r.Left != nil {
	//			ts = append(ts, r.Left)
	//		}
	//		if r.Right != nil {
	//			ts = append(ts, r.Right)
	//		}
	//		if p.Left == r {
	//			p.Left = nil
	//		} else {
	//			p.Right = nil
	//		}
	//	}
	//}
	//dfs(root, pivot)
	//if pivot.Left != nil {
	//	ts = append(ts, pivot.Left)
	//}
	//return ts
}

//leetcode submit region end(Prohibit modification and deletion)
