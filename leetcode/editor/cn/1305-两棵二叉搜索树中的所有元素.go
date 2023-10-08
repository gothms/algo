//给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.
//
//
//
// 示例 1：
//
//
//
//
//输入：root1 = [2,1,4], root2 = [1,0,3]
//输出：[0,1,1,2,3,4]
//
//
// 示例 2：
//
//
//
//
//输入：root1 = [1,null,8], root2 = [8,1]
//输出：[1,1,8,8]
//
//
//
//
// 提示：
//
//
// 每棵树的节点数在 [0, 5000] 范围内
// -10⁵ <= Node.val <= 10⁵
//
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 排序 👍 171 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	// 中序遍历 + 归并
	arr, temp := make([]int, 0), make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Left)
		arr = append(arr, r.Val)
		dfs(r.Right)
	}
	dfs(root1)
	arr, temp = temp, arr
	dfs(root2)
	ret := make([]int, 0, len(arr)+len(temp))
	j := 0
	for i := 0; i < len(arr); i++ { // 归并
		for ; j < len(temp) && temp[j] < arr[i]; j++ {
			ret = append(ret, temp[j])
		}
		ret = append(ret, arr[i])
	}
	ret = append(ret, temp[j:]...)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
