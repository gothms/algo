//给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的
// 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
//
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
//
// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
//
//
//
// 示例 1：
//
//
//输入：candidates = [2,3,6,7], target = 7
//输出：[[2,2,3],[7]]
//解释：
//2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
//7 也是一个候选， 7 = 7 。
//仅有这两种组合。
//
// 示例 2：
//
//
//输入: candidates = [2,3,5], target = 8
//输出: [[2,2,2,2],[2,3,3],[3,5]]
//
// 示例 3：
//
//
//输入: candidates = [2], target = 1
//输出: []
//
//
//
//
// 提示：
//
//
// 1 <= candidates.length <= 30
// 2 <= candidates[i] <= 40
// candidates 的所有元素 互不相同
// 1 <= target <= 40
//
//
// Related Topics 数组 回溯 👍 2763 👎 0

package main

import (
	"fmt"
	"slices"
)

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	i := combinationSum(candidates, target)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	// 回溯
	//n := len(candidates)
	//ret, path := make([][]int, 0), make([]int, 0, target/candidates[0]) // path 的最大长度
	//var dfs func(int, int)
	//dfs = func(i, s int) {
	//	if s == target {
	//		ret = append(ret, append([]int{}, path...))
	//		return
	//	}
	//	if i == n || s > target {
	//		return
	//	}
	//	//dfs(i+1, s)	// 已经回溯了，这里会导致重复计算
	//	for j := i; j < n; j++ {
	//		path = append(path, candidates[j])
	//		dfs(j, s+candidates[j])
	//		path = path[:len(path)-1] // 回溯
	//	}
	//}
	//dfs(0, 0)
	//return ret

	// lc：完全背包预处理 + 可行性剪枝
	slices.Sort(candidates)
	ans, n := make([][]int, 0), len(candidates)
	// 完全背包
	f := make([][]bool, n+1)
	f[0] = make([]bool, target+1)
	f[0][0] = true
	for i, x := range candidates {
		f[i+1] = make([]bool, target+1)
		for j, b := range f[i] {
			f[i+1][j] = b || j >= x && f[i+1][j-x]
		}
	}
	//for i, b := range f {
	//	fmt.Println(i, b)
	//}

	path := []int{}
	var dfs func(int, int)
	dfs = func(i, left int) {
		if left == 0 {
			// 找到一个合法组合
			ans = append(ans, slices.Clone(path))
			return
		}

		// 无法用下标在 [0, i] 中的数字组合出 left
		if left < 0 || !f[i+1][left] {
			return
		}

		// 不选
		dfs(i-1, left)

		// 选
		path = append(path, candidates[i])
		dfs(i, left-candidates[i])
		path = path[:len(path)-1]
	}

	// 倒着递归，这样参数符合 f 数组的定义
	dfs(n-1, target)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
