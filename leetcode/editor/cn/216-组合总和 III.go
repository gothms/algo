//找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
//
//
// 只使用数字1到9
// 每个数字 最多使用一次
//
//
// 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
//
//
//
// 示例 1:
//
//
//输入: k = 3, n = 7
//输出: [[1,2,4]]
//解释:
//1 + 2 + 4 = 7
//没有其他符合的组合了。
//
// 示例 2:
//
//
//输入: k = 3, n = 9
//输出: [[1,2,6], [1,3,5], [2,3,4]]
//解释:
//1 + 2 + 6 = 9
//1 + 3 + 5 = 9
//2 + 3 + 4 = 9
//没有其他符合的组合了。
//
// 示例 3:
//
//
//输入: k = 4, n = 1
//输出: []
//解释: 不存在有效的组合。
//在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。
//
//
//
//
// 提示:
//
//
// 2 <= k <= 9
// 1 <= n <= 60
//
//
// Related Topics 数组 回溯 👍 821 👎 0

package main

import (
	"fmt"
	"slices"
)

func main() {
	k, n := 3, 7
	sum3 := combinationSum3(k, n)
	fmt.Println(sum3)
}

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum3(k int, n int) [][]int {
	// 回溯问题的时间复杂度，通用公式：
	// 路径长度×搜索树的叶子数
	// 如 lc-216：O(k C(9,k))，时间/长度为 k，组合数为 C(9,k)
	mi, ma := (1+k)*k>>1, 45
	if n < mi || n > ma { // fast fail
		return nil
	}
	ret, path := make([][]int, 0), make([]int, 0, k)
	var dfs func(i, t int)
	dfs = func(i, t int) {
		if len(path) == k { // 终止：成功 / 失败
			if t == 0 {
				ret = append(ret, slices.Clone(path))
			}
			return
		}
		for i++; i <= min(9, t); i++ { // 剪枝：t<0 或 t<i 都将失败
			path = append(path, i)
			dfs(i, t-i)
			path = path[:len(path)-1] // 回溯
		}
	}
	dfs(0, n)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
