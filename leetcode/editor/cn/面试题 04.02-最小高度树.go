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
func sortedArrayToBST(nums []int) *TreeNode {
	// 分治
	var div func(int, int) *TreeNode
	div = func(l, r int) *TreeNode {
		if l > r {
			return nil
		} else if l == r {
			return &TreeNode{Val: nums[l]}
		}
		mid := (l + r) >> 1
		left, right := div(l, mid-1), div(mid+1, r)
		return &TreeNode{Val: nums[mid], Left: left, Right: right}
	}
	return div(0, len(nums)-1)
}

//leetcode submit region end(Prohibit modification and deletion)
