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

import "fmt"

func main() {
	titles := "CAB"
	possibilities := numTilePossibilities(titles)
	fmt.Println(possibilities)
}

//leetcode submit region begin(Prohibit modification and deletion)
func numTilePossibilities(tiles string) int {
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

//leetcode submit region end(Prohibit modification and deletion)
