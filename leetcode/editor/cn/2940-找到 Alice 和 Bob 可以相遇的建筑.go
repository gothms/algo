package main

import "container/heap"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i := range ans {
		ans[i] = -1
	}
	qs := make([][]pair2940, len(heights))
	for i, q := range queries {
		a, b := q[0], q[1]
		if a > b {
			a, b = b, a // 保证 a <= b
		}
		if a == b || heights[a] < heights[b] {
			ans[i] = b // a 直接跳到 b
		} else {
			qs[b] = append(qs[b], pair2940{heights[a], i}) // 离线询问
		}
	}

	h := hp2940{}
	for i, x := range heights {
		for h.Len() > 0 && h[0].h < x {
			// 堆顶的 heights[a] 可以跳到 heights[i]
			ans[heap.Pop(&h).(pair2940).i] = i
		}
		for _, p := range qs[i] {
			heap.Push(&h, p) // 后面再回答
		}
	}
	return ans
}

type pair2940 struct{ h, i int }
type hp2940 []pair2940

func (h hp2940) Len() int           { return len(h) }
func (h hp2940) Less(i, j int) bool { return h[i].h < h[j].h }
func (h hp2940) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp2940) Push(v any)        { *h = append(*h, v.(pair2940)) }
func (h *hp2940) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

//leetcode submit region end(Prohibit modification and deletion)
