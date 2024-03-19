//给你一个整数数组 ranks ，表示一些机械工的 能力值 。ranksi 是第 i 位机械工的能力值。能力值为 r 的机械工可以在 r * n² 分钟内修好
// n 辆车。
//
// 同时给你一个整数 cars ，表示总共需要修理的汽车数目。
//
// 请你返回修理所有汽车 最少 需要多少时间。
//
// 注意：所有机械工可以同时修理汽车。
//
//
//
// 示例 1：
//
//
//输入：ranks = [4,2,3,1], cars = 10
//输出：16
//解释：
//- 第一位机械工修 2 辆车，需要 4 * 2 * 2 = 16 分钟。
//- 第二位机械工修 2 辆车，需要 2 * 2 * 2 = 8 分钟。
//- 第三位机械工修 2 辆车，需要 3 * 2 * 2 = 12 分钟。
//- 第四位机械工修 4 辆车，需要 1 * 4 * 4 = 16 分钟。
//16 分钟是修理完所有车需要的最少时间。
//
//
// 示例 2：
//
//
//输入：ranks = [5,1,8], cars = 6
//输出：16
//解释：
//- 第一位机械工修 1 辆车，需要 5 * 1 * 1 = 5 分钟。
//- 第二位机械工修 4 辆车，需要 1 * 4 * 4 = 16 分钟。
//- 第三位机械工修 1 辆车，需要 8 * 1 * 1 = 8 分钟。
//16 分钟时修理完所有车需要的最少时间。
//
//
//
//
// 提示：
//
//
// 1 <= ranks.length <= 10⁵
// 1 <= ranks[i] <= 100
// 1 <= cars <= 10⁶
//
//
// Related Topics 数组 二分查找 👍 29 👎 0

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	ranks := []int{4, 2, 3, 1}
	cars := 10
	i := repairCars(ranks, cars)
	fmt.Println(i)
}

/*
思路：优先队列
	1.优先队列的元素定义为数组 arr = [4]int{}
		arr[0]：工人 i 的能力值
		arr[1]：修理汽车的数目
		arr[2]：修理 arr[1] 辆汽车，最少用时，此项为优先队列的比较元素
			根据时间（arr[2]）来判断这个用时，最多能修理多少汽车
		arr[3]：能力值为 arr[0] 的工人数目
			由题意 1 <= ranks.length <= 10^5，1 <= ranks[i] <= 100，可知：
			能力值相同的工人数的期望均值为 1000，所以统计相同能力值的工人数，以提高效率
	2.最少用时
		每次取出堆顶元素，即用时最少（arr[2]）的工人，统计能修理的最大汽车数
		2.1.当修理的汽车数 >= 目标值 cars，则返回这个用时，即为最少用时
		2.2.否则，将堆顶元素增加一辆汽车（同时更新它的用时），再堆化对顶元素
思路：二分
	二分查找的数据准备
		二分的上限：可以找出能力值最大的工人，他独自修理所有汽车的用时，为二分的上限
		二分的条件：找到在用时 i 的情况下，统计所有工人能修理的最大汽车数 >= cars
			所以这里要遍历每个工人，计算他用时 i 能修理多少汽车，然后累加所有汽车
		统计汽车：由于工人的能力值有大量的相同值，所以统计每个能力值的工人数，提高统计效率
*/
// leetcode submit region begin(Prohibit modification and deletion)
func repairCars(ranks []int, cars int) int64 {
	// 二分
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	minRank, arr := 101, [101]int{} // 二分用数组来统计工人数
	for _, v := range ranks {
		arr[v]++                     // 统计相同能力值的工人
		minRank = minVal(minRank, v) // 能力值最大的工人
	}
	mTime := minRank * cars * cars // 二分的上限时间：假设由能力值最大的工人单独修理 cars
	return int64(sort.Search(mTime, func(i int) bool {
		cnt := 0
		for r, c := range arr { // 统计用时 i，能修理的最大汽车数
			if c > 0 {
				cnt += int(math.Sqrt(float64(i)/float64(r))) * c
			}
		}
		return cnt >= cars // 二分的目标值
	}))

	// 优先队列
	//h, memo := &rcHp{}, make(map[int]int)
	//for _, v := range ranks {
	//	memo[v]++ // 统计相同能力值的工人
	//}
	//for k, v := range memo {
	//	heap.Push(h, [4]int{k, 1, k, v}) // 下一个目标是修理一辆汽车
	//}
	//for cnt := 0; ; heap.Fix(h, 0) { // 堆化堆顶元素
	//	t := &(*h)[0]    // 堆顶元素
	//	cnt += t[3]      // 统计修理的汽车
	//	if cnt >= cars { // 找到最少用时
	//		return int64(t[2])
	//	}
	//	t[2] += t[0] * (t[1]<<1 + 1) // 更新堆顶元素的用时
	//	t[1]++                       // 增加一辆汽车
	//}
}

type rcHp [][4]int

func (r rcHp) Len() int           { return len(r) }
func (r rcHp) Less(i, j int) bool { return r[i][2] < r[j][2] }
func (r rcHp) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r *rcHp) Push(x any)        { *r = append(*r, x.([4]int)) }
func (r *rcHp) Pop() any {
	v := (*r)[len(*r)-1]
	*r = (*r)[:len(*r)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
