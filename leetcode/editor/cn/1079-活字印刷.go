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
	titles := "AAB"   // 8
	titles = "AAABBC" // 188
	possibilities := numTilePossibilities(titles)
	fmt.Println(possibilities)
}

// leetcode submit region begin(Prohibit modification and deletion)

func numTilePossibilities(tiles string) int {

}

//leetcode submit region end(Prohibit modification and deletion)

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

func numTilePossibilities_(tiles string) int {
	// NTT：O(n log2N)
	// https://leetcode.cn/problems/letter-tile-possibilities/solutions/2275356/on2-ji-shu-dppythonjavacgo-by-endlessche-hmez/

	//dp：终版
	//counts := make(map[rune]int)
	//for _, v := range tiles {
	//	counts[v]++ // 统计每个字母的出现次数
	//}
	//dp := make([]int, len(tiles)+1)
	//dp[0] = 1 // 构造空序列的方案数
	//ret, sum := 0, 0
	//for _, cnt := range counts { // 枚举第 i 种字母
	//	sum += cnt
	//	for j := sum; j > 0; j-- { // 枚举序列长度 j
	//		for k := 1; k <= min(j, cnt); k++ { // 枚举第 i 种字母选了 k 个
	//			dp[j] += dp[j-k] * combineNTP[j][k] // 长度为 j，选出 k 个位置放置字母 "i"
	//		}
	//	}
	//}
	//for _, v := range dp[1:] {
	//	ret += v
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
	memo, n := make(map[uint8]int), len(tiles)
	for i := 0; i < n; i++ {
		memo[tiles[i]]++
	}
	var dfs func(int, map[uint8]int) int
	dfs = func(i int, memo map[uint8]int) (c int) {
		if i == 0 {
			return 1
		}
		c++ // 1 到 n-1 的情况
		for k, v := range memo {
			if v > 0 {
				memo[k]--
				c += dfs(i-1, memo)
				memo[k]++
			}
		}
		return
	}
	return dfs(len(tiles), memo) - 1
}
