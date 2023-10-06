//给你 2 枚相同 的鸡蛋，和一栋从第 1 层到第 n 层共有 n 层楼的建筑。
//
// 已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都 会碎 ，从 f 楼层或比它低 的楼层落下的鸡蛋都 不会碎 。
//
// 每次操作，你可以取一枚 没有碎 的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有
//摔碎，则可以在之后的操作中 重复使用 这枚鸡蛋。
//
// 请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出：2
//解释：我们可以将第一枚鸡蛋从 1 楼扔下，然后将第二枚从 2 楼扔下。
//如果第一枚鸡蛋碎了，可知 f = 0；
//如果第二枚鸡蛋碎了，但第一枚没碎，可知 f = 1；
//否则，当两个鸡蛋都没碎时，可知 f = 2。
//
//
// 示例 2：
//
//
//输入：n = 100
//输出：14
//解释：
//一种最优的策略是：
//- 将第一枚鸡蛋从 9 楼扔下。如果碎了，那么 f 在 0 和 8 之间。将第二枚从 1 楼扔下，然后每扔一次上一层楼，在 8 次内找到 f 。总操作次数
//= 1 + 8 = 9 。
//- 如果第一枚鸡蛋没有碎，那么再把第一枚鸡蛋从 22 层扔下。如果碎了，那么 f 在 9 和 21 之间。将第二枚鸡蛋从 10 楼扔下，然后每扔一次上一层楼
//，在 12 次内找到 f 。总操作次数 = 2 + 12 = 14 。
//- 如果第一枚鸡蛋没有再次碎掉，则按照类似的方法从 34, 45, 55, 64, 72, 79, 85, 90, 94, 97, 99 和 100 楼分别
//扔下第一枚鸡蛋。
//不管结果如何，最多需要扔 14 次来确定 f 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 1000
//
//
// Related Topics 数学 动态规划 👍 69 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 100
	drop := twoEggDrop(n)
	fmt.Println(drop)
}

// leetcode submit region begin(Prohibit modification and deletion)
func twoEggDrop(n int) int {
	// dp
	//return tedDP[n]

	// 记忆化搜索
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	k := 2 // k 个鸡蛋
	memo := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, k+1)
	}
	var dfs func(int, int) int
	dfs = func(n, k int) int {
		if n == 0 || k == 1 { // 到顶层 / 最后一个蛋
			return n
		}
		v := &memo[n][k]
		if *v > 0 {
			return *v
		}
		ret := math.MaxInt32
		for i := 1; i <= n; i++ {
			ret = minVal(ret, 1+maxVal(dfs(i-1, k-1), dfs(n-i, k))) // 碎 / 没碎
		}
		*v = ret
		return ret
	}
	return dfs(n, k)

	// 数学
	// 让公式 1 + 2 + 3 + 4 + .... + x >= n 成立的最小的x，就是 f 的最小次数
	//return int(math.Ceil((math.Sqrt(float64(n*8+1)) - 1.0) / 2.0))
}

var tedDP []int

func init() {
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	const N = 1000
	tedDP = make([]int, N+1)
	//tedDP[0], tedDP[1], tedDP[2] = 0, 1, 2
	//for i := 3; i <= N; i++ {
	for i := 1; i <= N; i++ {
		min := math.MaxInt32
		for j := 1; j <= i; j++ {
			min = minVal(min, maxVal(tedDP[i-j]+1, j)) // 没碎 / 碎
		}
		tedDP[i] = min
	}
}

//leetcode submit region end(Prohibit modification and deletion)
