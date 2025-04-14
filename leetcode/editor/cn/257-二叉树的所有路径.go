package main

import (
	"strconv"
	"strings"
)

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
func binaryTreePaths_(root *TreeNode) []string {
	
}

//leetcode submit region end(Prohibit modification and deletion)

func binaryTreePaths(root *TreeNode) []string {
	// lc
	ans, path := make([]string, 0), make([]string, 0)
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, strconv.Itoa(root.Val))
		if root.Left == root.Right {
			ans = append(ans, strings.Join(path, "->"))
		} else {
			dfs(root.Left)
			dfs(root.Right)
		}
		path = path[:len(path)-1]
	}
	dfs(root)
	return ans

	// 个人
	//ans, path := make([]string, 0), make([]byte, 0)
	//var dfs func(root *TreeNode)
	//dfs = func(root *TreeNode) {
	//	i := len(path)
	//	v := strconv.Itoa(root.Val)
	//	path = append(path, v...)
	//	defer func() { path = path[:i] }() // 回溯
	//	if root.Left == nil && root.Right == nil {
	//		ans = append(ans, string(path))
	//		return
	//	}
	//	path = append(path, '-', '>')
	//	if root.Left != nil {
	//		dfs(root.Left)
	//	}
	//	if root.Right != nil {
	//		dfs(root.Right)
	//	}
	//}
	//dfs(root)
	//return ans
}
