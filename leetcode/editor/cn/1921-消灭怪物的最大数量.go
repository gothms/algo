//你正在玩一款电子游戏，在游戏中你需要保护城市免受怪物侵袭。给你一个 下标从 0 开始 且长度为 n 的整数数组 dist ，其中 dist[i] 是第 i
//个怪物与城市的 初始距离（单位：米）。
//
// 怪物以 恒定 的速度走向城市。给你一个长度为 n 的整数数组 speed 表示每个怪物的速度，其中 speed[i] 是第 i 个怪物的速度（单位：米/分
//）。
//
// 怪物从 第 0 分钟 时开始移动。你有一把武器，并可以 选择 在每一分钟的开始时使用，包括第 0 分钟。但是你无法在一分钟的中间使用武器。这种武器威力惊人
//，一次可以消灭任一还活着的怪物。
//
// 一旦任一怪物到达城市，你就输掉了这场游戏。如果某个怪物 恰 在某一分钟开始时到达城市，这会被视为 输掉 游戏，在你可以使用武器之前，游戏就会结束。
//
// 返回在你输掉游戏前可以消灭的怪物的 最大 数量。如果你可以在所有怪物到达城市前将它们全部消灭，返回 n 。
//
//
//
// 示例 1：
//
//
//输入：dist = [1,3,4], speed = [1,1,1]
//输出：3
//解释：
//第 0 分钟开始时，怪物的距离是 [1,3,4]，你消灭了第一个怪物。
//第 1 分钟开始时，怪物的距离是 [X,2,3]，你没有消灭任何怪物。
//第 2 分钟开始时，怪物的距离是 [X,1,2]，你消灭了第二个怪物。
//第 3 分钟开始时，怪物的距离是 [X,X,1]，你消灭了第三个怪物。
//所有 3 个怪物都可以被消灭。
//
// 示例 2：
//
//
//输入：dist = [1,1,2,3], speed = [1,1,1,1]
//输出：1
//解释：
//第 0 分钟开始时，怪物的距离是 [1,1,2,3]，你消灭了第一个怪物。
//第 1 分钟开始时，怪物的距离是 [X,0,1,2]，你输掉了游戏。
//你只能消灭 1 个怪物。
//
//
// 示例 3：
//
//
//输入：dist = [3,2,4], speed = [5,3,2]
//输出：1
//解释：
//第 0 分钟开始时，怪物的距离是 [3,2,4]，你消灭了第一个怪物。
//第 1 分钟开始时，怪物的距离是 [X,0,2]，你输掉了游戏。
//你只能消灭 1 个怪物。
//
//
//
//
// 提示：
//
//
// n == dist.length == speed.length
// 1 <= n <= 10⁵
// 1 <= dist[i], speed[i] <= 10⁵
//
//
// Related Topics 贪心 数组 排序 👍 31 👎 0

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	dist := []int{4, 3, 4}
	speed := []int{1, 1, 2} // 3
	maximum := eliminateMaximum(dist, speed)
	fmt.Println(maximum)
}

/*
思路：排序+贪心
	1.贪心：每次选择消灭最先到的僵尸
		怎么选出最先到的僵尸呢？
	2.排序
		计算 dist[i] / speed[i]，即僵尸到达的时间，并排序
		然后依次遍历僵尸到达的时间，如果在 i 时刻已达到，则游戏失败
*/
// leetcode submit region begin(Prohibit modification and deletion)
func eliminateMaximum(dist []int, speed []int) int {
	// 排序+贪心
	n := len(dist)
	arr := make([]int, n)
	for i := 0; i < n; i++ { // 计算到达的时间
		arr[i] = int(math.Ceil(float64(dist[i]) / float64(speed[i])))
	}
	sort.Ints(arr) // 排序
	for i := 0; i < n; i++ {
		if arr[i] <= i { // 僵尸入城
			return i
		}
	}
	return n

	// 优先队列
	//h := NewAr(len(dist))
	//for i := range dist {
	//	h[i] = float64(dist[i]) / float64(speed[i])
	//}
	//heap.Init(&h)
	//for i := float64(0); h.Len() > 0; i++ {
	//	top := heap.Pop(&h).(float64)
	//	if top <= i {
	//		return int(i)
	//	}
	//}
	//return len(dist)
}

type ar []float64

func (a ar) Len() int           { return len(a) }
func (a ar) Less(i, j int) bool { return a[i] < a[j] }
func (a ar) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a *ar) Push(x any)        { *a = append(*a, x.(float64)) }
func (a *ar) Pop() any {
	v := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return v
}
func NewAr(n int) ar {
	return make([]float64, n)
}

//func (a *ar) Len() int           { return len(*a) }
//func (a *ar) Less(i, j int) bool { return (*a)[i] < (*a)[j] }
//func (a *ar) Swap(i, j int)      { (*a)[i], (*a)[j] = (*a)[j], (*a)[i] }

//leetcode submit region end(Prohibit modification and deletion)
