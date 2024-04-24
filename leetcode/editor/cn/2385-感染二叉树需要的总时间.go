//给你一棵二叉树的根节点 root ，二叉树中节点的值 互不相同 。另给你一个整数 start 。在第 0 分钟，感染 将会从值为 start 的节点开始爆发
//。
//
// 每分钟，如果节点满足以下全部条件，就会被感染：
//
//
// 节点此前还没有感染。
// 节点与一个已感染节点相邻。
//
//
// 返回感染整棵树需要的分钟数。
//
//
//
// 示例 1：
// 输入：root = [1,5,3,null,4,10,6,9,2], start = 3
//输出：4
//解释：节点按以下过程被感染：
//- 第 0 分钟：节点 3
//- 第 1 分钟：节点 1、10、6
//- 第 2 分钟：节点5
//- 第 3 分钟：节点 4
//- 第 4 分钟：节点 9 和 2
//感染整棵树需要 4 分钟，所以返回 4 。
//
//
// 示例 2：
// 输入：root = [1], start = 1
//输出：0
//解释：第 0 分钟，树中唯一一个节点处于感染状态，返回 0 。
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [1, 10⁵] 内
// 1 <= Node.val <= 10⁵
// 每个节点的值 互不相同
// 树中必定存在值为 start 的节点
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 44 👎 0

package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	start := 1 // 0
	root = &TreeNode{1, &TreeNode{2, &TreeNode{3, &TreeNode{4, &TreeNode{Val: 5}, nil}, nil}, nil}, nil}
	start = 3 // 2
	root = &TreeNode{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, &TreeNode{Val: 5}}}}}
	start = 2 // 3
	root = &TreeNode{2, &TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{1, nil, &TreeNode{Val: 5}}}, nil}
	root = &TreeNode{Left: root}
	start = 1 // 2
	time := amountOfTime(root, start)
	fmt.Println(time)
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
func amountOfTime(root *TreeNode, start int) int {
	// dfs：个人
	ret := 0
	var dfs func(*TreeNode) (int, int)
	dfs = func(r *TreeNode) (int, int) {
		if r == nil {
			return 0, 0
		}
		s1, d1 := dfs(r.Left)
		s2, d2 := dfs(r.Right)
		var sick, d int               // 感染源的深度，子树的深度
		if sick = s1 + s2; sick > 0 { // 左/右子树是感染源
			// 必须在此处结算，因为往后的“归”中 d = 0
			ret = max(ret, d1+d2+sick) // 感染源深度 + 无感染子树 + 0
			sick++
		} else if r.Val == start { // 感染源
			ret = max(d1, d2) // 肯定是 ret 第一次赋值
			sick = 1          // 标识感染源的深度
		} else {
			d = max(d1, d2) + 1 // 无感染的子树
		}
		return sick, d // 两者至少一个是 0：是感染源所在子树，则 d 置为 0
	}
	dfs(root)
	return ret

	// lc：将 start “置为” root，再求扩散时间

	// lc：不推荐
	//ret, d := 0, -1
	//var dfs func(*TreeNode, int) int
	//dfs = func(r *TreeNode, depth int) int {
	//	if r == nil {
	//		return 0
	//	}
	//	if r.Val == start {
	//		d = depth
	//	}
	//	ld := dfs(r.Left, depth+1)
	//	inLeft := d > 0 // 左子树已遍历完，记录 inLeft
	//	rd := dfs(r.Right, depth+1)
	//	if r.Val == start {
	//		ret = max(ret, max(ld, rd))
	//	}
	//	if inLeft {
	//		// 最大值一定产生在感染源的祖先的计算，即 root 是 r 的祖先，r 是 感染源的祖先
	//		ret = max(ret, d-depth+rd) // 算法：感染源深度 - r 的深度 + 右子树最大深度
	//	} else {
	//		ret = max(ret, d-depth+ld)
	//	}
	//	return max(ld, rd) + 1
	//}
	//dfs(root, 0)
	//return ret
}

//leetcode submit region end(Prohibit modification and deletion)
