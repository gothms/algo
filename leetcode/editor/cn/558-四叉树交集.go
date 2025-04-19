package main

func main() {

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

func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
	// 个人
	if quadTree1.IsLeaf {
		if quadTree1.Val {
			return quadTree1
		} else {
			return quadTree2
		}
	}
	if quadTree2.IsLeaf {
		if quadTree2.Val {
			return quadTree2
		} else {
			return quadTree1
		}
	}
	tl, tr, bl, br :=
		intersect(quadTree1.TopLeft, quadTree2.TopLeft),
		intersect(quadTree1.TopRight, quadTree2.TopRight),
		intersect(quadTree1.BottomLeft, quadTree2.BottomLeft),
		intersect(quadTree1.BottomRight, quadTree2.BottomRight)
	if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf &&
		//(tl.Val && tr.Val && bl.Val && br.Val || !tl.Val && !tr.Val && !bl.Val && !br.Val) {
		tl.Val == tr.Val && tr.Val == bl.Val && bl.Val == br.Val {
		return &Node{Val: tl.Val, IsLeaf: true}
	}
	return &Node{TopLeft: tl, TopRight: tr, BottomLeft: bl, BottomRight: br}
}

//leetcode submit region end(Prohibit modification and deletion)
