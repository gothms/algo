//给你两个 正 整数 n 和 x 。
//
// 请你返回将 n 表示成一些 互不相同 正整数的 x 次幂之和的方案数。换句话说，你需要返回互不相同整数 [n1, n2, ..., nk] 的集合数目，满
//足 n = n1ˣ + n2ˣ + ... + nkˣ 。
//
// 由于答案可能非常大，请你将它对 10⁹ + 7 取余后返回。
//
// 比方说，n = 160 且 x = 3 ，一个表示 n 的方法是 n = 2³ + 3³ + 5³ 。
//
//
//
// 示例 1：
//
// 输入：n = 10, x = 2
//输出：1
//解释：我们可以将 n 表示为：n = 3² + 1² = 10 。
//这是唯一将 10 表达成不同整数 2 次方之和的方案。
//
//
// 示例 2：
//
// 输入：n = 4, x = 1
//输出：2
//解释：我们可以将 n 按以下方案表示：
//- n = 4¹ = 4 。
//- n = 3¹ + 1¹ = 4 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 300
// 1 <= x <= 5
//
//
// Related Topics 动态规划 👍 7 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	n, x := 10, 2
	n, x = 10, 1 // map[1:1 2:1 3:2 4:2 5:3 6:4 7:5 8:6 9:8 10:10]
	ofWays := numberOfWays(n, x)
	fmt.Println(ofWays)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfWays(n int, x int) int {
	// 查表
	const mod = 1_000_000_007
	if x == 1 {
		return nowArr[n] % mod
	}
	return now[x-2][n]

	// lc：优化
	//minVal := func(a, b int) int {
	//	if b < a {
	//		return b
	//	}
	//	return a
	//}
	//const mod = 1_000_000_007
	//f := make([]int, n+1)
	//f[0] = 1
	//for i, sum := 1, 0; ; i++ {
	//	v := int(math.Pow(float64(i), float64(x)))
	//	sum += v // 累加值，即能达到的和
	//	if v > n {
	//		break
	//	}
	//	for s := minVal(sum, n); s > v; s-- { // 倒序遍历
	//		f[s] += f[s-v] // 状态转移方程：包含 v，即包含 i
	//	}
	//	f[v]++ // 自身
	//}
	//return f[n] % mod

	// lc
	//const mod = 1_000_000_007
	//f := make([]int, n+1)
	//f[0] = 1
	//for i := 1; ; i++ {
	//	v := int(math.Pow(float64(i), float64(x)))
	//	if v > n {
	//		break
	//	}
	//	for s := n; s >= v; s-- {
	//		f[s] += f[s-v] // 包含 v，即包含 i
	//		fmt.Println(f, s)
	//	}
	//}
	//return f[n] % mod
}

var (
	nowArr []int          // x = 1
	now    [4]map[int]int // x = [2,5]
)

func init() {
	const M = 300
	pow := func(x, n int) int { // 幂运算：x^n
		ret := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				ret *= x
			}
			x *= x
		}
		return ret
	}
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	nowArr = make([]int, M+1)
	//nowArr[0] = 1 // 哨兵
	for i, sum := 1, 0; i <= M; i++ {
		sum += i                              // 累加值，即能达到的和
		for m := minVal(sum, M); m > i; m-- { // 倒序遍历
			nowArr[m] += nowArr[m-i] // 状态转移方程：包含 v，即包含 i
		}
		nowArr[i]++ // 自身
	}
	var dfs func(int, int, int, []int) // x = [2,5]
	dfs = func(i, j, sum int, t []int) {
		if j == len(t) { // 已遍历完
			return
		}
		dfs(i, j+1, sum, t)        // 不选 j，则没有新的和
		if sum += t[j]; sum <= M { // 合法
			now[i][sum]++ // 新的和
			dfs(i, j+1, sum, t)
		}
	}
	for i := 2; i <= 5; i++ {
		now[i-2] = make(map[int]int)
		temp := []int{1}
		for j := 2; ; j++ {
			v := pow(j, i) // 幂运算
			if v > M {     // 终止循环
				break
			}
			temp = append(temp, v) // <= M 的幂值
		}
		dfs(i-2, 0, 0, temp) // 可达的幂和
	}
	//fmt.Println(nowArr)
	//for _, v := range now {
	//	fmt.Println(v)
	//}
}

// leetcode submit region end(Prohibit modification and deletion)
func init_old() {
	const M = 300
	// 严重超时
	var dfsX func(int, int, int, []int)
	dfsX = func(i, j, sum int, t []int) {
		if j == len(t) {
			return
		}
		dfsX(i, j+1, sum, t)
		if sum += t[j]; sum <= M {
			now[i][sum]++
			dfsX(i, j+1, sum, t)
		}
	}
	temp := make([]int, M)
	for i := 0; i < M; i++ {
		temp[i] = i + 1
	}
	now[1] = make(map[int]int)
	dfsX(1, 0, 0, temp)
	var dfs func(int, int, int, []int) // x = [2,5]
	dfs = func(i, j, sum int, t []int) {
		if j == len(t) {
			return
		}
		dfs(i, j+1, sum, t)
		if sum += t[j]; sum <= M {
			now[i][sum]++
			dfs(i, j+1, sum, t)
		}
	}
	for i := 2; i <= 5; i++ {
		now[i] = make(map[int]int)
		temp := []int{1}
		for j := 2; ; j++ {
			v := math.Pow(float64(j), float64(i))
			if v > M {
				break
			}
			temp = append(temp, int(v))
		}
		dfs(i, 0, 0, temp)
	}
}
