//给你一个下标从 0 开始的二维整数数组 flowers ，其中 flowers[i] = [starti, endi] 表示第 i 朵花的 花期 从
//starti 到 endi （都 包含）。同时给你一个下标从 0 开始大小为 n 的整数数组 people ，people[i] 是第 i 个人来看花的时间。
//
// 请你返回一个大小为 n 的整数数组 answer ，其中 answer[i]是第 i 个人到达时在花期内花的 数目 。
//
//
//
// 示例 1：
//
//
//
//
//输入：flowers = [[1,6],[3,7],[9,12],[4,13]], people = [2,3,7,11]
//输出：[1,2,2,2]
//解释：上图展示了每朵花的花期时间，和每个人的到达时间。
//对每个人，我们返回他们到达时在花期内花的数目。
//
//
// 示例 2：
//
//
//
//
//输入：flowers = [[1,10],[3,3]], people = [3,3,2]
//输出：[2,2,1]
//解释：上图展示了每朵花的花期时间，和每个人的到达时间。
//对每个人，我们返回他们到达时在花期内花的数目。
//
//
//
//
// 提示：
//
//
// 1 <= flowers.length <= 5 * 10⁴
// flowers[i].length == 2
// 1 <= starti <= endi <= 10⁹
// 1 <= people.length <= 5 * 10⁴
// 1 <= people[i] <= 10⁹
//
//
// Related Topics 数组 哈希表 二分查找 有序集合 前缀和 排序 👍 60 👎 0

package main

import (
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func fullBloomFlowers(flowers [][]int, people []int) []int {
	// 差分
	//n, m := len(flowers), len(people)
	//memo := make(map[int]int, n<<1) // 统计每个时刻的“开谢”
	//for _, f := range flowers {
	//	memo[f[0]]++
	//	memo[f[1]+1]--
	//}
	//times := make([]int, 0, len(memo)) // 花期的时刻
	//for k := range memo {
	//	times = append(times, k)
	//}
	//sort.Ints(times)
	//fbf, queue := make([]int, m), make([]int, m)
	//for i := 1; i < m; i++ {
	//	queue[i] = i
	//}
	//sort.Slice(queue, func(i, j int) bool { // 前缀和需要依次查询
	//	return people[queue[i]] < people[queue[j]]
	//})
	//j, k, sum := 0, len(times), 0
	//for _, i := range queue {
	//	for ; j < k && times[j] <= people[i]; j++ {
	//		sum += memo[times[j]] // 差分
	//	}
	//	fbf[i] = sum
	//}
	//return fbf

	// 二分
	n, m := len(flowers), len(people)
	fbf, start, end := make([]int, m), make([]int, n), make([]int, n)
	for i, f := range flowers {
		//start[i], end[i] = f[0], f[1]+1 // 花期开始和结束时间
		start[i], end[i] = f[0], f[1] // 花期开始和结束时间
	}
	sort.Ints(start)
	sort.Ints(end)
	for i, v := range people {
		//fbf[i] = sort.Search(n, func(i int) bool { // 已开花的数目
		//	return start[i] > v
		//}) - sort.Search(n, func(i int) bool { // 已凋谢的数目
		//	return end[i] > v
		//})
		fbf[i] = sort.SearchInts(start, v+1) - sort.SearchInts(end, v) // 已开花 - 已凋谢
	}
	return fbf
}

//leetcode submit region end(Prohibit modification and deletion)
