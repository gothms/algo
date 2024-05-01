//给你一个下标从 0 开始的整数数组 costs ，其中 costs[i] 是雇佣第 i 位工人的代价。
//
// 同时给你两个整数 k 和 candidates 。我们想根据以下规则恰好雇佣 k 位工人：
//
//
// 总共进行 k 轮雇佣，且每一轮恰好雇佣一位工人。
// 在每一轮雇佣中，从最前面 candidates 和最后面 candidates 人中选出代价最小的一位工人，如果有多位代价相同且最小的工人，选择下标更小的
//一位工人。
//
// 比方说，costs = [3,2,7,7,1,2] 且 candidates = 2 ，第一轮雇佣中，我们选择第 4 位工人，因为他的代价最小 [3,2,
//7,7,1,2] 。
// 第二轮雇佣，我们选择第 1 位工人，因为他们的代价与第 4 位工人一样都是最小代价，而且下标更小，[3,2,7,7,2] 。注意每一轮雇佣后，剩余工人的下
//标可能会发生变化。
//
// 如果剩余员工数目不足 candidates 人，那么下一轮雇佣他们中代价最小的一人，如果有多位代价相同且最小的工人，选择下标更小的一位工人。
// 一位工人只能被选择一次。
//
//
// 返回雇佣恰好 k 位工人的总代价。
//
//
//
// 示例 1：
//
// 输入：costs = [17,12,10,2,7,2,11,20,8], k = 3, candidates = 4
//输出：11
//解释：我们总共雇佣 3 位工人。总代价一开始为 0 。
//- 第一轮雇佣，我们从 [17,12,10,2,7,2,11,20,8] 中选择。最小代价是 2 ，有两位工人，我们选择下标更小的一位工人，即第 3 位工人
//。总代价是 0 + 2 = 2 。
//- 第二轮雇佣，我们从 [17,12,10,7,2,11,20,8] 中选择。最小代价是 2 ，下标为 4 ，总代价是 2 + 2 = 4 。
//- 第三轮雇佣，我们从 [17,12,10,7,11,20,8] 中选择，最小代价是 7 ，下标为 3 ，总代价是 4 + 7 = 11 。注意下标为 3
//的工人同时在最前面和最后面 4 位工人中。
//总雇佣代价是 11 。
//
//
// 示例 2：
//
// 输入：costs = [1,2,4,1], k = 3, candidates = 3
//输出：4
//解释：我们总共雇佣 3 位工人。总代价一开始为 0 。
//- 第一轮雇佣，我们从 [1,2,4,1] 中选择。最小代价为 1 ，有两位工人，我们选择下标更小的一位工人，即第 0 位工人，总代价是 0 + 1 = 1
// 。注意，下标为 1 和 2 的工人同时在最前面和最后面 3 位工人中。
//- 第二轮雇佣，我们从 [2,4,1] 中选择。最小代价为 1 ，下标为 2 ，总代价是 1 + 1 = 2 。
//- 第三轮雇佣，少于 3 位工人，我们从剩余工人 [2,4] 中选择。最小代价是 2 ，下标为 0 。总代价为 2 + 2 = 4 。
//总雇佣代价是 4 。
//
//
//
//
// 提示：
//
//
// 1 <= costs.length <= 105
// 1 <= costs[i] <= 10⁵
// 1 <= k, candidates <= costs.length
//
//
// Related Topics 数组 双指针 模拟 堆（优先队列） 👍 62 👎 0

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	costs := []int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}
	k, candidates := 11, 2
	cost := totalCost(costs, k, candidates)
	fmt.Println(cost)
}

// leetcode submit region begin(Prohibit modification and deletion)
func totalCost(costs []int, k int, candidates int) int64 {
	// 优化
	ans, n := 0, len(costs)
	if candidates<<1+k > n { // 随便选
		sort.Ints(costs) // or TopK
		for _, v := range costs[:k] {
			ans += v
		}
		return int64(ans)
	}
	h := &hp2462{}
	for i := 0; i < candidates; i++ { // 初始化堆
		heap.Push(h, [2]int{costs[i], 0})
		heap.Push(h, [2]int{costs[n-i-1], 1})
	}
	idx := [2]int{candidates, n - candidates - 1} // 方便操作和简化代码
	for i := 0; i < k; i++ {
		v := (*h)[0]
		(*h)[0][0] = costs[idx[v[1]]] // 直接修改堆顶元素，再堆化
		heap.Fix(h, 0)
		idx[v[1]] += v[1]<<1 ^ 2 - 1 // 0 -> +1 或 1 -> -1
		ans += v[0]
	}
	return int64(ans)

	// 堆：也可以用两个堆
	//ans, n := 0, len(costs)
	//l, r := candidates, n-candidates-1 // 标记左/右下一次选择的位置，防止超选
	//h := &hp2462{}
	//for i := 0; i < min(n, l); i++ { // 初始化堆
	//	heap.Push(h, [2]int{costs[i], 0})
	//}
	//for i := n - 1; i >= max(r+1, candidates); i-- {
	//	heap.Push(h, [2]int{costs[i], 1})
	//}
	//for i := 0; i < k; i++ {
	//	v := heap.Pop(h).([2]int)
	//	if l <= r { // 还有元素可选
	//		if v[1] == 0 { // 左
	//			heap.Push(h, [2]int{costs[l], 0})
	//			l++
	//		} else { // 右
	//			heap.Push(h, [2]int{costs[r], 1})
	//			r--
	//		}
	//	}
	//	ans += v[0]
	//}
	//return int64(ans)
}

type hp2462 [][2]int

func (h hp2462) Len() int { return len(h) }
func (h hp2462) Less(i, j int) bool {
	return h[i][0] < h[j][0] || h[i][0] == h[j][0] && h[i][1] < h[j][1]
}
func (h hp2462) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp2462) Push(x any)   { *h = append(*h, x.([2]int)) }
func (h *hp2462) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
