//电子游戏“辐射4”中，任务 “通向自由” 要求玩家到达名为 “Freedom Trail Ring” 的金属表盘，并使用表盘拼写特定关键词才能开门。
//
// 给定一个字符串 ring ，表示刻在外环上的编码；给定另一个字符串 key ，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。
//
// 最初，ring 的第一个字符与 12:00 方向对齐。您需要顺时针或逆时针旋转 ring 以使 key 的一个字符在 12:00 方向对齐，然后按下中心按
//钮，以此逐个拼写完 key 中的所有字符。
//
// 旋转 ring 拼出 key 字符 key[i] 的阶段中：
//
//
// 您可以将 ring 顺时针或逆时针旋转 一个位置 ，计为1步。旋转的最终目的是将字符串 ring 的一个字符与 12:00 方向对齐，并且这个字符必须等于
//字符 key[i] 。
// 如果字符 key[i] 已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作 1 步。按完之后，您可以开始拼写 key 的下一个字符（下一阶段
//）, 直至完成所有拼写。
//
//
//
//
// 示例 1：
//
//
//
//
//
//
//
//
//输入: ring = "godding", key = "gd"
//输出: 4
//解释:
// 对于 key 的第一个字符 'g'，已经在正确的位置, 我们只需要1步来拼写这个字符。
// 对于 key 的第二个字符 'd'，我们需要逆时针旋转 ring "godding" 2步使它变成 "ddinggo"。
// 当然, 我们还需要1步进行拼写。
// 因此最终的输出是 4。
//
//
// 示例 2:
//
//
//输入: ring = "godding", key = "godding"
//输出: 13
//
//
//
//
// 提示：
//
//
// 1 <= ring.length, key.length <= 100
// ring 和 key 只包含小写英文字母
// 保证 字符串 key 一定可以由字符串 ring 旋转拼出
//
//
// Related Topics 深度优先搜索 广度优先搜索 字符串 动态规划 👍 277 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	ring := "godding"
	key := "gd"
	steps := findRotateSteps(ring, key)
	fmt.Println(steps)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findRotateSteps(ring string, key string) int {
	// dp：个人
	ret, m, n := math.MaxInt32, len(ring), len(key)
	dp, memo := make([][]int, n+1), [26][]int{}
	for i, c := range ring {
		ch := c - 'a'
		memo[ch] = append(memo[ch], i)
	}
	minV := func(a, b int) int { // 逆/顺时针
		if a > b {
			return min(a-b, m-a+b)
		} else {
			return min(b-a, m+a-b)
		}
	}
	dp[0] = make([]int, m)
	for i := 1; i < m; i++ {
		dp[0][i] = math.MaxInt32
	}
	for i, c := range key {
		dp[i+1] = make([]int, m)
		idx := c - 'a'
		for _, j := range memo[idx] {
			dp[i+1][j] = math.MaxInt32
			for k, v := range dp[i] {
				if i == 0 || v > 0 {
					dp[i+1][j] = min(dp[i+1][j], minV(k, j)+v+1)
				}
			}
		}
	}
	for _, v := range dp[n] {
		if v > 0 {
			ret = min(ret, v)
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
