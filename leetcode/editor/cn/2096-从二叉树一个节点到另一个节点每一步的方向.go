package main

import (
	"fmt"
)

func main() {
	s, d := 2, 1 // U
	root := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	s, d = 2, 1 // UL
	root = &TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}
	s, d = 2, 1 // UUUU
	root = &TreeNode{1, &TreeNode{3, &TreeNode{7, &TreeNode{6, &TreeNode{2, nil, nil}, nil}, nil}, nil},
		&TreeNode{8, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}}
	s, d = 6, 15 // UURR
	root = &TreeNode{1, nil, &TreeNode{10, &TreeNode{12, &TreeNode{4, nil, nil}, &TreeNode{6, nil, nil}},
		&TreeNode{13, nil, &TreeNode{15, nil, nil}}}}
	directions := getDirections(root, s, d)
	fmt.Println(directions)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
func getDirections(root *TreeNode, startValue int, destValue int) string {
	// lc：dfs+bfs

	// lc
	//path := make([]byte, 0)
	//var dfs func(*TreeNode, int) bool
	//dfs = func(cur *TreeNode, t int) bool {
	//	if cur == nil {
	//		return false
	//	}
	//	if cur.Val == t {
	//		return true
	//	}
	//	path = append(path, 'L')
	//	if dfs(cur.Left, t) {
	//		return true
	//	}
	//	path[len(path)-1] = 'R'
	//	if dfs(cur.Right, t) {
	//		return true
	//	}
	//	path = path[:len(path)-1]
	//	return false
	//}
	//dfs(root, startValue)
	//up := path
	//path = nil
	//dfs(root, destValue)
	//for len(up) > 0 && len(path) > 0 && up[0] == path[0] {
	//	up, path = up[1:], path[1:]
	//}
	//return strings.Repeat("U", len(up)) + string(path)

	//path := make([]byte, 0)
	//var dfs func(*TreeNode) *TreeNode
	//dfs = func(cur *TreeNode) *TreeNode {
	//	if cur == nil || cur.Val == startValue || cur.Val == destValue {
	//		return cur
	//	}
	//	l, r := dfs(cur.Left), dfs(cur.Right)
	//	if l == nil {
	//		return r
	//	} else if r == nil {
	//		return l
	//	}
	//	return cur
	//}
	//var down func(*TreeNode) bool
	//down = func(cur *TreeNode) bool {
	//	if cur == nil {
	//		return false
	//	}
	//	if cur.Val == destValue {
	//		return true
	//	}
	//	path = append(path, 'L')
	//	if down(cur.Left) {
	//		return true
	//	}
	//	path[len(path)-1] = 'R'
	//	if down(cur.Right) {
	//		return true
	//	}
	//	path = path[:len(path)-1]
	//	return false
	//}
	//var up func(*TreeNode) bool
	//up = func(cur *TreeNode) bool {
	//	if cur == nil {
	//		return false
	//	}
	//	if cur.Val == startValue {
	//		return true
	//	}
	//	path = append(path, 'U')
	//	if up(cur.Left) || up(cur.Right) {
	//		return true
	//	}
	//	path = path[:len(path)-1]
	//	return false
	//}
	//
	//parent := dfs(root)           // 1.最近公共祖先，lc 236
	//if parent.Val == startValue { // 2.分情况讨论
	//	path = append(path, 'L')
	//	if !down(parent.Left) {
	//		path[len(path)-1] = 'R'
	//		down(parent.Right)
	//	}
	//} else if parent.Val == destValue {
	//	path = append(path, 'U')
	//	if !up(parent.Left) {
	//		up(parent.Right)
	//	}
	//} else {
	//	path = append(path, 'U')
	//	if up(parent.Left) {
	//		path = append(path, 'R')
	//		down(parent.Right)
	//	} else {
	//		up(parent.Right)
	//		path = append(path, 'L')
	//		down(parent.Left)
	//	}
	//}
	//return string(path)

	// 个人
	//var dfs func(*TreeNode) (*TreeNode, bool)
	//dfs = func(cur *TreeNode) (*TreeNode, bool) {
	//	if cur == nil {
	//		return nil, false
	//	}
	//	if cur.Val == startValue || cur.Val == destValue {
	//		return cur, cur.Val == startValue
	//	}
	//	l, ls := dfs(cur.Left)
	//	r, rs := dfs(cur.Right)
	//	if l != nil && r != nil {
	//		return cur, ls
	//	} else if l != nil {
	//		return l, ls
	//	} else {
	//		return r, rs // 此刻是右子树，up 后也可以是左子树
	//	}
	//}
	//parent, s := dfs(root) // 1.最近公共祖先，lc 236
	//
	//path := make([]byte, 0)
	//var down func(*TreeNode) bool
	//down = func(cur *TreeNode) bool {
	//	if cur == nil {
	//		return false
	//	}
	//	if cur.Val == destValue {
	//		return true
	//	}
	//	path = append(path, 'L')
	//	if down(cur.Left) {
	//		return true
	//	}
	//	path[len(path)-1] = 'R'
	//	if down(cur.Right) {
	//		return true
	//	}
	//	path = path[:len(path)-1]
	//	return false
	//}
	//var up func(*TreeNode) bool
	//up = func(cur *TreeNode) bool {
	//	if cur == nil {
	//		return false
	//	}
	//	if cur.Val == startValue {
	//		return true
	//	}
	//	path = append(path, 'U')
	//	if up(cur.Left) || up(cur.Right) {
	//		return true
	//	}
	//	path = path[:len(path)-1]
	//	return false
	//}
	//if parent.Val == startValue { // 2.分情况讨论
	//	path = append(path, 'L')
	//	if !down(parent.Left) {
	//		path[len(path)-1] = 'R'
	//		down(parent.Right)
	//	}
	//} else if parent.Val == destValue {
	//	path = append(path, 'U')
	//	if !up(parent.Left) {
	//		up(parent.Right)
	//	}
	//} else {
	//	path = append(path, 'U')
	//	if s {
	//		up(parent.Left)
	//		path = append(path, 'R')
	//		down(parent.Right)
	//	} else {
	//		up(parent.Right)
	//		path = append(path, 'L')
	//		down(parent.Left)
	//	}
	//}
	//return string(path)
}

//leetcode submit region end(Prohibit modification and deletion)
