

package main

import "slices"

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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ret, queue := make([][]int, 0), []*TreeNode{root}
	for len(queue) > 0 {
		l := make([]int, len(queue))
		for i := range l {
			n := queue[0]
			queue = queue[1:]
			l[i] = n.Val
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		ret = append(ret, l)
	}
	slices.Reverse(ret)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
