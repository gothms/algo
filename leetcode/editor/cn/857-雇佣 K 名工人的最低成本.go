//有 n 名工人。 给定两个数组 quality 和 wage ，其中，quality[i] 表示第 i 名工人的工作质量，其最低期望工资为 wage[i]
//。
//
// 现在我们想雇佣 k 名工人组成一个工资组。在雇佣 一组 k 名工人时，我们必须按照下述规则向他们支付工资：
//
//
// 对工资组中的每名工人，应当按其工作质量与同组其他工人的工作质量的比例来支付工资。
// 工资组中的每名工人至少应当得到他们的最低期望工资。
//
//
// 给定整数 k ，返回 组成满足上述条件的付费群体所需的最小金额 。在实际答案的 10⁻⁵ 以内的答案将被接受。。
//
//
//
//
//
//
// 示例 1：
//
//
//输入： quality = [10,20,5], wage = [70,50,30], k = 2
//输出： 105.00000
//解释： 我们向 0 号工人支付 70，向 2 号工人支付 35。
//
// 示例 2：
//
//
//输入： quality = [3,1,10,10,1], wage = [4,8,2,2,7], k = 3
//输出： 30.66667
//解释： 我们向 0 号工人支付 4，向 2 号和 3 号分别支付 13.33333。
//
//
//
// 提示：
//
//
// n == quality.length == wage.length
// 1 <= k <= n <= 10⁴
// 1 <= quality[i], wage[i] <= 10⁴
//
//
// Related Topics 贪心 数组 排序 堆（优先队列） 👍 314 👎 0

package main

import (
	"container/heap"
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	// 堆 + 排序
	// 总金额 = k 名工人中最高的工资倍率 * k 名工人的总质量
	// k 名工人中最高的工资倍率：按性价比升序遍历工人，当前工人为 i，则工资倍率就是工人 i 的倍率
	// k 名工人的总质量：维护 k 大小的大顶堆，存储的是工人的质量，记录堆中 k 名工人的总质量
	type s857 struct {
		i   int     // 工人编号
		per float64 // 性价比：工资/质量
	}
	n := len(quality)
	h, idx := &hp857{}, make([]*s857, n)
	for i := range idx {
		idx[i] = &s857{i, float64(wage[i]) / float64(quality[i])}
	}
	sort.Slice(idx, func(i, j int) bool { // 按 性价比 排序（升序）
		return idx[i].per < idx[j].per
	})
	s := 0
	for i := 0; i < k; i++ {
		id := idx[i].i
		s += quality[id]
		heap.Push(h, quality[id])
	}
	ans := idx[k-1].per * float64(s)
	for i := k; i < n; i++ { // 工人 i 的工资倍率，是当前最高
		id := idx[i].i
		if quality[id] < (*h)[0] { // 大顶堆
			s += quality[id] - (*h)[0] // 已遍历工人中选 k 名 quality 最低的工人，并记录总质量
			(*h)[0] = quality[id]
			heap.Fix(h, 0) // 更新堆顶元素
			ans = min(ans, idx[i].per*float64(s))
		}
	}
	return ans
}

type hp857 []int

func (h hp857) Len() int {
	return len(h)
}
func (h hp857) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h hp857) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp857) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *hp857) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)

//func mincostToHireWorkers(quality, wage []int, k int) float64 {
//	// lc
//	type pair struct{ q, w int }
//	pairs := make([]pair, len(quality))
//	for i, q := range quality {
//		pairs[i] = pair{q, wage[i]}
//	}
//	slices.SortFunc(pairs, func(a, b pair) int { return a.w*b.q - b.w*a.q }) // 按照 r 值排序
//
//	h := hp{make([]int, k)}
//	sumQ := 0
//	for i, p := range pairs[:k] {
//		h.IntSlice[i] = p.q
//		sumQ += p.q
//	}
//	heap.Init(&h)
//
//	ans := float64(sumQ*pairs[k-1].w) / float64(pairs[k-1].q) // 选 r 值最小的 k 名工人
//
//	for _, p := range pairs[k:] { // 后面的工人 r 值更大
//		if p.q < h.IntSlice[0] { // 但是 sumQ 可以变小，从而可能得到更优的答案
//			sumQ -= h.IntSlice[0] - p.q
//			h.IntSlice[0] = p.q
//			heap.Fix(&h, 0) // 更新堆顶
//			ans = min(ans, float64(sumQ*p.w)/float64(p.q))
//		}
//	}
//	return ans
//}
//
//type hp struct{ sort.IntSlice }
//
//func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
//func (hp) Push(any)             {}
//func (hp) Pop() (_ any)         { return }
