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

/*
思路
	用两个返回值来分别记录，sick 感染源的深度、d 子树的深度
	所求值 = MAX（sick节点在root中的深度 + 与sick不同侧的root子树的最大深度, MAX（sick左子树的最大深度，sick右子树的最大深度）)
	1.与sick不同侧的root子树的最大深度：变量 d 记录，即每遇到 sick 节点的子树，叶子结点返回深度为 0，保留最大深度即可
	2.MAX（sick左子树的最大深度，sick右子树的最大深度）：当递归在“归”时，遇到 sick 节点就从 d 变量中求得该值。后续“归”中返回的 d 值都默认为 0
	3.sick节点在root中的深度：当递归在“归”时，遇到 sick 节点，则返回 sick = 1。后续“归”中，只要 sick 返回值大于 0，则表示该子树是 sick 所在的子树，sick+1 后返回
		以上root指的是任意节点为根
	总结：sick 和 d 两个返回值，只会有一个是有效的（另一个无效且为 0）。优先级为 sick > d，即当 sick 大于 0，考虑的是 sick 节点的处理，否则才是考虑无感染子树的处理
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
func amountOfTime(root *TreeNode, start int) int {
	// dfs：个人
	//ret := 0
	//var dfs func(*TreeNode) (int, int)
	//dfs = func(r *TreeNode) (int, int) {
	//	if r == nil {
	//		return 0, 0
	//	}
	//	s1, d1 := dfs(r.Left)
	//	s2, d2 := dfs(r.Right)
	//	var sick, d int               // 感染源的深度，子树的深度
	//	if sick = s1 + s2; sick > 0 { // 左/右子树是感染源
	//		// 必须在此处结算，因为往后的“归”中 d = 0
	//		ret = max(ret, d1+d2+sick) // 感染源深度 + 无感染子树 + 0
	//		sick++
	//	} else if r.Val == start { // 感染源
	//		ret = max(d1, d2) // 肯定是 ret 第一次赋值
	//		sick = 1          // 标识感染源的深度
	//	} else {
	//		d = max(d1, d2) + 1 // 无感染的子树
	//	}
	//	return sick, d // 两者至少一个是 0：是感染源所在子树，则 d 置为 0
	//}
	//dfs(root)
	//return ret

	// lc 优秀题解：https://leetcode.cn/problems/amount-of-time-for-binary-tree-to-be-infected/solutions/2753470/cong-liang-ci-bian-li-dao-yi-ci-bian-li-tmt0x/

	// lc：将 start “置为” root，再求扩散时间
	var ans int
	var dfs func(*TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, false
		}
		lLen, lFound := dfs(node.Left)
		rLen, rFound := dfs(node.Right)
		if node.Val == start {
			// 计算子树 start 的最大深度
			// 注意这里和方法一的区别，max 后面没有 +1，所以算出的也是最大深度
			ans = max(lLen, rLen)
			return 1, true // 找到了 start
		}
		if lFound || rFound {
			// 只有在左子树或右子树包含 start 时，才能更新答案
			ans = max(ans, lLen+rLen) // 两条链拼成直径
			// 保证 start 是直径端点
			if lFound {
				return lLen + 1, true
			}
			return rLen + 1, true
		}
		return max(lLen, rLen) + 1, false
	}
	dfs(root)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
