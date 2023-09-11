//有一只跳蚤的家在数轴上的位置 x 处。请你帮助它从位置 0 出发，到达它的家。
//
// 跳蚤跳跃的规则如下：
//
//
// 它可以 往前 跳恰好 a 个位置（即往右跳）。
// 它可以 往后 跳恰好 b 个位置（即往左跳）。
// 它不能 连续 往后跳 2 次。
// 它不能跳到任何 forbidden 数组中的位置。
//
//
// 跳蚤可以往前跳 超过 它的家的位置，但是它 不能跳到负整数 的位置。
//
// 给你一个整数数组 forbidden ，其中 forbidden[i] 是跳蚤不能跳到的位置，同时给你整数 a， b 和 x ，请你返回跳蚤到家的最少跳跃
//次数。如果没有恰好到达 x 的可行方案，请你返回 -1 。
//
//
//
// 示例 1：
//
//
//输入：forbidden = [14,4,18,1,15], a = 3, b = 15, x = 9
//输出：3
//解释：往前跳 3 次（0 -> 3 -> 6 -> 9），跳蚤就到家了。
//
//
// 示例 2：
//
//
//输入：forbidden = [8,3,16,6,12,20], a = 15, b = 13, x = 11
//输出：-1
//
//
// 示例 3：
//
//
//输入：forbidden = [1,6,2,14,5,17,4], a = 16, b = 9, x = 7
//输出：2
//解释：往前跳一次（0 -> 16），然后往回跳一次（16 -> 7），跳蚤就到家了。
//
//
//
//
// 提示：
//
//
// 1 <= forbidden.length <= 1000
// 1 <= a, b, forbidden[i] <= 2000
// 0 <= x <= 2000
// forbidden 中所有位置互不相同。
// 位置 x 不在 forbidden 中。
//
//
// Related Topics 广度优先搜索 数组 动态规划 👍 81 👎 0

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	forbidden := []int{162, 118, 178, 152, 167, 100, 40, 74, 199, 186, 26, 73, 200, 127, 30, 124, 193, 84, 184, 36, 103, 149, 153, 9, 54, 154, 133, 95, 45, 198, 79, 157, 64, 122, 59, 71, 48, 177, 82, 35, 14, 176, 16, 108, 111, 6, 168, 31, 134, 164, 136, 72, 98}
	a, b, x := 29, 98, 80
	jumps := minimumJumps(forbidden, a, b, x)
	fmt.Println(jumps)
}

/*
思路：bfs
	1.通过搜索，我们可以确定能不能跳到 x 位置，搜索的左区间是 0，右区间是多少呢？
		前跳次数必然是 >= 后跳次数的，可以分3种情况讨论右区间
		当 a>b：此时跳蚤只会前进（正方向），不会后退，所以右区间是 x+b（超过 x 后，可以回跳一次）
		当 a=b：此时后跳已经没有意义，右区间是 x
		当 a<b：右区间为 max(max(forbidden)+a+b, x)，证明过程较复杂，请参考LeetCode
	2.Dijkstra
		cache 记录障碍点，visited 记录已经跳过的点，mj 记录当前跳的方向、点、次数
		每次取出当前跳的最小次数的点，然后判断是否能右跳/左跳
		最先跳到 x 的情况，用的次数最小
*/
// leetcode submit region begin(Prohibit modification and deletion)
func minimumJumps(forbidden []int, a int, b int, x int) int {
	// Dijkstra
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	max, visited, cache := 0, make(map[int]bool), make(map[int]struct{}, len(forbidden))
	for _, v := range forbidden {
		max = maxVal(max, v)
		cache[v] = struct{}{} // forbidden 的位置
	}
	max = maxVal(max+a, x) + b // 右边界：最核心逻辑
	h := &mj{{}}               // 初始位置起跳
	for h.Len() > 0 {
		p := heap.Pop(h).([3]int) // 堆顶
		if p[1] == x {            // 成功
			return p[2]
		}
		nrp := p[1] + a // 右跳
		if _, ok := cache[nrp]; !ok && nrp <= max && !visited[nrp] {
			heap.Push(h, [3]int{1, nrp, p[2] + 1})
			visited[nrp] = true // 已向右跳过
		}
		if p[0] == 0 { // 不能连续左跳
			continue
		}
		nlp := p[1] - b // 左跳
		if _, ok := cache[nlp]; !ok && nlp > 0 && !visited[nlp] {
			heap.Push(h, [3]int{0, nlp, p[2] + 1})
			visited[-nlp] = true // 已向左跳过
		}
	}
	return -1
}

type mj [][3]int

func (m mj) Len() int           { return len(m) }
func (m mj) Less(i, j int) bool { return m[i][2] < m[j][2] }
func (m mj) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *mj) Push(x any)        { *m = append(*m, x.([3]int)) }
func (m *mj) Pop() any {
	v := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
