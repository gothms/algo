package main

import (
	"container/heap"
	"fmt"
	"slices"
)

func main() {
	w := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 3 // ["i","love","coding"]
	frequent := topKFrequent(w, k)
	fmt.Println(frequent)
}

// leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(words []string, k int) []string {
	memo := make(map[string]int)
	for _, w := range words {
		memo[w]++
	}
	arr := make(hp, 0, k)
	h := &arr
	for w, c := range memo {
		heap.Push(h, pair{c, w})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, 0, k)
	for h.Len() > 0 {
		ans = append(ans, heap.Pop(h).(pair).s)
	}
	slices.Reverse(ans)
	return ans
}

type pair struct {
	fre int
	s   string
}
type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	return h[i].fre < h[j].fre || h[i].fre == h[j].fre && h[i].s > h[j].s
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)   { *h = append(*h, x.(pair)) }
func (h *hp) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

var _ heap.Interface = (*hp)(nil)

//leetcode submit region end(Prohibit modification and deletion)
