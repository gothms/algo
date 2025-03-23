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
func findSecondMinimumValue(root *TreeNode) int {
	ans, v := -1, root.Val
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil || ans > 0 && root.Val >= ans {
			return
		}
		if root.Val > v {
			ans = root.Val
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return ans

	// 个人
	//if root.Left == nil {
	//	return -1
	//}
	//mn := root.Val
	//var f func(*TreeNode) *TreeNode
	//f = func(root *TreeNode) *TreeNode {
	//	if root == nil {
	//		return nil
	//	}
	//	if v := root.Val; v == mn {
	//		l, r := f(root.Left), f(root.Right)
	//		if l == nil {
	//			return r
	//		} else if r == nil {
	//			return l
	//		} else if l.Val < r.Val {
	//			return l
	//		} else {
	//			return r
	//		}
	//	}
	//	return root
	//}
	//l, r := f(root.Left), f(root.Right)
	//if l == nil && r == nil {
	//	return -1
	//} else if r == nil || l != nil && l.Val < r.Val {
	//	return l.Val
	//} else {
	//	return r.Val
	//}
}

//leetcode submit region end(Prohibit modification and deletion)

func findSecondMinimumValue_(root *TreeNode) int {
	//v := root.Val
	//var dfs func(*TreeNode) int
	//dfs = func(root *TreeNode) int {
	//	if root.Left == nil {
	//		return root.Val
	//	}
	//	l, r := root.Left.Val, root.Right.Val
	//	var ans int
	//	switch {
	//	case l < r:
	//		ans = dfs(root.Left)
	//		if ans == v {
	//			return root.Right.Val
	//		}
	//		return min(ans, root.Right.Val)
	//	case l > r:
	//		ans = dfs(root.Right)
	//		if ans == v {
	//			return root.Left.Val
	//		}
	//		return min(ans, root.Left.Val)
	//	default:
	//		lv, rv := dfs(root.Left), dfs(root.Right)
	//		if lv == v {
	//			return rv
	//		} else if rv == v {
	//			return lv
	//		}
	//		return min(lv, rv)
	//	}
	//}
	//ans := dfs(root)
	//if ans == v {
	//	return -1
	//}
	//return ans

	// lc
	ans := -1
	v := root.Val
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil || ans > 0 && r.Val > ans {
			return
		}
		if r.Val > v {
			ans = r.Val
		}
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	return ans
}
