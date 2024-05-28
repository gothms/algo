package main

import (
	"container/heap"
	"fmt"
)

func main() {
	left, right := 0., 1.
	fmt.Println(left, right)
	fmt.Printf("%T,%T\n", left, right) // float64,float64

	arr := []int{1, 2, 3, 5}
	k := 3
	fraction := kthSmallestPrimeFraction(arr, k)
	fmt.Println(fraction)
}

// leetcode submit region begin(Prohibit modification and deletion)
func kthSmallestPrimeFraction(arr []int, k int) []int {

}

//leetcode submit region end(Prohibit modification and deletion)

func kthSmallestPrimeFraction_(arr []int, k int) []int {
	// 二分 + 双指针
	n := len(arr)
	l, r := 0., 1.
	for {
		m := (l + r) / 2
		cnt, i := 0, 0 // i 的巧妙运用
		x, y := 0, 1
		for j := 1; j < n; j++ {
			for float64(arr[i])/float64(arr[j]) <= m { // 符合的分数
				if arr[i]*y > arr[j]*x { // 符合的分数中的最大值
					x, y = arr[i], arr[j]
				}
				i++
			}
			cnt += i // 统计个数
		}
		if cnt == k {
			return []int{x, y}
		} else if cnt < k {
			l = m
		} else {
			r = m
		}
	}

	// heap
	//n := len(arr)
	//h := &hp786{make([][2]int, 0), arr}
	//for i := 1; i < n; i++ {
	//	heap.Push(h, [2]int{0, i})
	//}
	//for i := 1; i < k; i++ {
	//	if h.hp[0][0]+1 < h.hp[0][1] { // i < j
	//		h.hp[0][0]++
	//		heap.Fix(h, 0)
	//	} else {
	//		heap.Pop(h)
	//	}
	//}
	//return []int{arr[h.hp[0][0]], arr[h.hp[0][1]]}

	// 排序
}

type hp786 struct {
	hp  [][2]int
	arr []int
}

func (h hp786) Len() int { return len(h.hp) }
func (h hp786) Less(i, j int) bool {
	return h.arr[h.hp[i][0]]*h.arr[h.hp[j][1]] < h.arr[h.hp[j][0]]*h.arr[h.hp[i][1]]
}
func (h hp786) Swap(i, j int) { h.hp[i], h.hp[j] = h.hp[j], h.hp[i] }
func (h *hp786) Push(x any)   { h.hp = append(h.hp, x.([2]int)) }
func (h *hp786) Pop() any {
	v := h.hp[len(h.hp)-1]
	h.hp = h.hp[:len(h.hp)-1]
	return v
}

var _ heap.Interface = (*hp786)(nil)
