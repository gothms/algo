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
func countPairs(root *TreeNode, distance int) int {
	// dfs：个人
	if distance == 1 { // fast path
		return 0
	}
	ans := 0
	var dfs func(*TreeNode) []int
	dfs = func(r *TreeNode) []int {
		if r == nil {
			return nil
		}
		ls, rs := dfs(r.Left), dfs(r.Right)
		if ls == nil && rs == nil {
			ret := make([]int, distance)
			ret[1] = 1 // 叶子
			return ret
		} else if ls == nil {
			return append([]int{0}, rs...)[:distance]
		} else if rs == nil {
			return append([]int{0}, ls...)[:distance]
		}
		ret := make([]int, distance) // ret[i] 代表当前深度为 i 的叶子节点有 ret[i] 个，ret[0] 肯定为 0
		for i := 2; i < distance; i++ {
			ret[i] = ls[i-1] + rs[i-1] // 舍弃深度为 distance 的叶子节点
		}
		for i := 2; i < distance; i++ { // 前缀和
			rs[i] += rs[i-1]
		}
		for i := 1; i < distance; i++ { // 统计路径 <= distance 的叶子节点对数
			ans += ls[i] * rs[distance-i]
		}
		return ret
	}
	dfs(root)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
