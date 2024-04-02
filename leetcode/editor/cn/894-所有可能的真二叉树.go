//给你一个整数 n ，请你找出所有可能含 n 个节点的 真二叉树 ，并以列表形式返回。答案中每棵树的每个节点都必须符合 Node.val == 0 。
//
// 答案的每个元素都是一棵真二叉树的根节点。你可以按 任意顺序 返回最终的真二叉树列表。
//
// 真二叉树 是一类二叉树，树中每个节点恰好有 0 或 2 个子节点。
//
//
//
// 示例 1：
//
//
//输入：n = 7
//输出：[[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0
//,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
//
//
// 示例 2：
//
//
//输入：n = 3
//输出：[[0,0,0]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 20
//
//
// Related Topics 树 递归 记忆化搜索 动态规划 二叉树 👍 325 👎 0

package main

import "fmt"

func main() {
	n := 7
	fbt := allPossibleFBT(n)
	bfs := func(r *TreeNode) []int {
		ret := make([]int, 0)
		q := []*TreeNode{r}
		for l := len(q); l > 0; l = len(q) {
			for i := 0; i < l; i++ {
				if q[i] == nil {
					ret = append(ret, -1)
				} else {
					ret = append(ret, 0)
					if q[i].Left != nil {
						q = append(q, q[i].Left, q[i].Right)
					} else {
						q = append(q, nil, nil)
					}
				}
			}
			q = q[l:]
		}
		return ret
	}
	for _, f := range fbt {
		arr := bfs(f)
		fmt.Println(arr)
	}
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
func allPossibleFBT(n int) []*TreeNode {
	// dp
	if n&1 == 0 {
		return nil
	}
	return memo894[n>>1]

	// dfs
	//if n&1 == 0 {
	//	return nil
	//}
	//const N = (20 - 2) >> 1
	//memo := make([][]*TreeNode, N)
	//memo[0], memo[1] = []*TreeNode{{0, nil, nil}},
	//	[]*TreeNode{{0, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}}}
	//var dfs func(i int) []*TreeNode
	//dfs = func(i int) []*TreeNode {
	//	if memo[i>>1] != nil {
	//		return memo[i>>1]
	//	}
	//	for j := 1; j < n; j += 2 {
	//
	//	}
	//}
	//dfs(n - 1)
	//return memo[N-1]
}

const n894 = 20 >> 1

var memo894 [][]*TreeNode

func init() {
	memo894 = make([][]*TreeNode, n894)
	memo894[0] = []*TreeNode{{}}
	for i := 1; i < n894; i++ { // dp
		for j := 0; j < i; j++ {
			for _, l := range memo894[j] {
				for _, r := range memo894[i-j-1] {
					memo894[i] = append(memo894[i], &TreeNode{0, l, r})
				}
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
