//你有一套活字字模 tiles，其中每个字模上都刻有一个字母 tiles[i]。返回你可以印出的非空字母序列的数目。
//
// 注意：本题中，每个活字字模只能使用一次。
//
//
//
// 示例 1：
//
//
//输入："AAB"
//输出：8
//解释：可能的序列为 "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA"。
//
//
// 示例 2：
//
//
//输入："AAABBC"
//输出：188
//
//
// 示例 3：
//
//
//输入："V"
//输出：1
//
//
//
// 提示：
//
//
// 1 <= tiles.length <= 7
// tiles 由大写英文字母组成
//
//
// Related Topics 哈希表 字符串 回溯 计数 👍 170 👎 0

package main

import (
	"fmt"
)

func main() {
	titles := "AAB"
	titles = "AAABBC"
	possibilities := numTilePossibilities(titles)
	fmt.Println(possibilities)
}

// leetcode submit region begin(Prohibit modification and deletion)

const mx = 8

var combineNTP [mx][mx]int

func init() {
	for i := 0; i < mx; i++ {
		combineNTP[i][0], combineNTP[i][i] = 1, 1
		for j := 1; j < i; j++ {
			combineNTP[i][j] = combineNTP[i-1][j-1] + combineNTP[i-1][j] // 预处理组合数
		}
	}
}

func numTilePossibilities(tiles string) int {
	// NTT：O(n log2N)
	// https://leetcode.cn/problems/letter-tile-possibilities/solutions/2275356/on2-ji-shu-dppythonjavacgo-by-endlessche-hmez/

	//dp：终版
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	counts := make(map[rune]int)
	for _, v := range tiles {
		counts[v]++ // 统计每个字母的出现次数
	}
	dp := make([]int, len(tiles)+1)
	dp[0] = 1 // 构造空序列的方案数
	ret, sum := 0, 0
	for _, cnt := range counts { // 枚举第 i 种字母
		sum += cnt
		for j := sum; j > 0; j-- { // 枚举序列长度 j
			for k := 1; k <= minVal(j, cnt); k++ { // 枚举第 i 种字母选了 k 个
				dp[j] += dp[j-k] * combineNTP[j][k] // 长度为 j，选出 k 个位置放置字母 "i"
			}
		}
	}
	for _, v := range dp[1:] {
		ret += v
	}
	return ret

	// dp：test
	//minVal := func(a, b int) int {
	//	if b < a {
	//		return b
	//	}
	//	return a
	//}
	//counts := map[rune]int{}
	//for _, ch := range tiles {
	//	counts[ch]++ // 统计每个字母的出现次数
	//}
	//ret, n, m := 0, len(tiles), len(counts)
	//f := make([][]int, m+1)
	//f[0] = make([]int, n+1)
	//f[0][0] = 1 // 构造空序列的方案数
	//i := 1
	////for _, cnt := range counts { // 枚举第 i 种字母
	//cnter := []int{3, 2, 1}
	//for _, cnt := range cnter { // 枚举第 i 种字母
	//	f[i] = make([]int, n+1)
	//	for j := 0; j <= n; j++ { // 枚举序列长度 j
	//		max := minVal(j, cnt)
	//		for k := 0; k <= max; k++ { // 枚举第 i 种字母选了 k 个
	//			//f[i][j] += f[i-1][j-k] * c[j][k] // 长度为 j，选出 k 个位置放置字母 "i"
	//			f[i][j] += f[i-1][j-k] * c[max][k] // 会出现重复组合，如：AAABB，j=2，选择了两次 AB
	//			if j == 2 {
	//				fmt.Println(i, j, k, ",", f[i-1][j-k], c[max][k])
	//			}
	//		}
	//	}
	//	i++
	//}
	//for _, ints := range f {
	//	fmt.Println(ints)
	//}
	//for _, x := range f[m][1:] {
	//	ret += x
	//}
	//return ret

	// dp
	//counts := map[rune]int{}
	//for _, ch := range tiles {
	//	counts[ch]++ // 统计每个字母的出现次数
	//}
	//ret, n, m := 0, len(tiles), len(counts)
	//f := make([][]int, m+1)
	//f[0] = make([]int, n+1)
	//f[0][0] = 1 // 构造空序列的方案数
	//i := 1
	//for _, cnt := range counts { // 枚举第 i 种字母
	//	f[i] = make([]int, n+1)
	//	for j := 0; j <= n; j++ { // 枚举序列长度 j
	//		for k := 0; k <= j && k <= cnt; k++ { // 枚举第 i 种字母选了 k 个
	//			f[i][j] += f[i-1][j-k] * c[j][k] // 长度为 j，选出 k 个位置放置字母 "i"
	//		}
	//	}
	//	i++
	//}
	////for _, ints := range f {
	////	fmt.Println(ints)
	////}
	//for _, x := range f[m][1:] {
	//	ret += x
	//}
	//return ret

	// 回溯：排序
	//arr := []byte(tiles)
	//sort.Slice(arr, func(i, j int) bool {
	//	return arr[i] < arr[j]
	//})
	//n := len(arr)
	//visited := make([]bool, n) // 标识某位是否已选择
	//var dfs func()
	//cnt := 0
	//dfs = func() {
	//	for i := 0; i < n; i++ {
	//		if visited[i] || i > 0 && !visited[i-1] && arr[i] == arr[i-1] {
	//			continue // 相同数字：第一次出现被标识选择了，后面才继续选择
	//		}
	//		cnt++
	//		visited[i] = true
	//		dfs()
	//		visited[i] = false // 回溯
	//	}
	//}
	//dfs()
	//return cnt

	// 回溯：map
	//memo, n := make(map[uint8]int), len(tiles)
	//for i := 0; i < n; i++ {
	//	memo[tiles[i]]++
	//}
	//var dfs func(int, map[uint8]int) int
	//dfs = func(i int, memo map[uint8]int) (c int) {
	//	if i == 0 {
	//		return 1
	//	}
	//	c++ // 1 到 n-1 的情况
	//	for k, v := range memo {
	//		if v > 0 {
	//			memo[k]--
	//			c += dfs(i-1, memo)
	//			memo[k]++
	//		}
	//	}
	//	return
	//}
	//return dfs(len(tiles), memo) - 1
}

//leetcode submit region end(Prohibit modification and deletion)
