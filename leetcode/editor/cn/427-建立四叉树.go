package main

import "fmt"

func main() {
	grid := [][]int{ // [[0,1],[1,1],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0, 0}}
	n := construct(grid)
	fmt.Printf("%t", *n)
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	// lc：矩阵前缀和

	// 个人：分治
	var dfs func(int, int, int, int) *Node
	dfs = func(l, r, t, b int) *Node {
		if l == r {
			return &Node{Val: grid[l][t] == 1, IsLeaf: true}
		}
		mt, mb := (l+r)>>1, (t+b)>>1
		tl, tr, bl, br :=
			dfs(l, mt, t, mb), dfs(l, mt, mb+1, b), dfs(mt+1, r, t, mb), dfs(mt+1, r, mb+1, b)
		if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf &&
			(tl.Val && tr.Val && bl.Val && br.Val || !tl.Val && !tr.Val && !bl.Val && !br.Val) {
			return &Node{Val: tl.Val, IsLeaf: true}
		}
		return &Node{TopLeft: tl, TopRight: tr, BottomLeft: bl, BottomRight: br}
	}
	return dfs(0, len(grid)-1, 0, len(grid)-1)
}

//leetcode submit region end(Prohibit modification and deletion)
