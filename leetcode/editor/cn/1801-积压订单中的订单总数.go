package main

import (
	"container/heap"
	"fmt"
)

func main() {
	orders := [][]int{{7, 1000000000, 1},
		{15, 3, 0},
		{5, 999999995, 0},
		{5, 1, 1}} // 999999984
	backlogOrders := getNumberOfBacklogOrders(orders)
	fmt.Println(backlogOrders)
}

// leetcode submit region begin(Prohibit modification and deletion)
func getNumberOfBacklogOrders(orders [][]int) int {
	const mod = 1_000_000_007
	buy, sell := &hp1801{data: make([][2]int, 0), cmp: func(x, y int) bool {
		return x >= y
	}}, &hp1801{data: make([][2]int, 0), cmp: func(x, y int) bool {
		return x <= y
	}} // buy 大顶堆，sell 小顶堆
	var b, s int
	for _, order := range orders {
		p, a := order[0], order[1]
		switch order[2] {
		case 0:
			for sell.Len() > 0 && (*sell).data[0][0] <= p {
				c := (*sell).data[0][1]
				a -= c
				if a < 0 { // 已匹配完，a == 0 也匹配完，但是这里优先淘汰 sell[0]
					(*sell).data[0][1] = -a
					a = 0
					break
				} else {
					heap.Pop(sell)
				}
			}
			s = (s - order[1] + a + mod) % mod // 如果 a = 0，也要修正
			if a > 0 {
				b = (b + a) % mod
				heap.Push(buy, [2]int{p, a}) // 放入负数，变大顶堆
			}
		case 1:
			for buy.Len() > 0 && (*buy).data[0][0] >= p { // 修正放入的是负数
				c := (*buy).data[0][1]
				a -= c
				if a < 0 { // 已匹配完，a == 0 也匹配完，但是这里优先淘汰 sell[0]
					(*buy).data[0][1] = -a
					a = 0
					break
				} else {
					heap.Pop(buy)
				}
			}
			b = (b - order[1] + a + mod) % mod
			if a > 0 {
				s = (s + a) % mod
				heap.Push(sell, [2]int{p, a})
			}
		}
	}
	return (b + s) % mod
}

type hp1801 struct {
	data [][2]int
	cmp  func(int, int) bool
}

func (h hp1801) Len() int { return len(h.data) }

func (h hp1801) Less(i, j int) bool { return h.cmp(h.data[i][0], h.data[j][0]) }
func (h hp1801) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp1801) Push(x any)        { (*h).data = append((*h).data, x.([2]int)) }
func (h *hp1801) Pop() any {
	v := (*h).data[h.Len()-1]
	(*h).data = (*h).data[:h.Len()-1]
	return v
}

//var _ heap.Interface = (*hp1801)(nil)

//leetcode submit region end(Prohibit modification and deletion)

//func getNumberOfBacklogOrders_(orders [][]int) int {
//	const mod = 1_000_000_007
//	buy, sell := &hp1801{}, &hp1801{} // buy 大顶堆，sell 小顶堆
//	var b, s int
//	for _, order := range orders {
//		p, a := order[0], order[1]
//		switch order[2] {
//		case 0:
//			for sell.Len() > 0 && (*sell)[0][0] <= p {
//				c := (*sell)[0][1]
//				a -= c
//				if a < 0 { // 已匹配完，a == 0 也匹配完，但是这里优先淘汰 sell[0]
//					(*sell)[0][1] = -a
//					a = 0
//					break
//				} else {
//					heap.Pop(sell)
//				}
//			}
//			s = (s - order[1] + a + mod) % mod // 如果 a = 0，也要修正
//			if a > 0 {
//				b = (b + a) % mod
//				heap.Push(buy, [2]int{-p, a}) // 放入负数，变大顶堆
//			}
//		case 1:
//			for buy.Len() > 0 && -(*buy)[0][0] >= p { // 修正放入的是负数
//				c := (*buy)[0][1]
//				a -= c
//				if a < 0 { // 已匹配完，a == 0 也匹配完，但是这里优先淘汰 sell[0]
//					(*buy)[0][1] = -a
//					a = 0
//					break
//				} else {
//					heap.Pop(buy)
//				}
//			}
//			b = (b - order[1] + a + mod) % mod
//			if a > 0 {
//				s = (s + a) % mod
//				heap.Push(sell, [2]int{p, a})
//			}
//		}
//	}
//	return (b + s) % mod
//}
//
//type hp1801 [][2]int
//
//func (h hp1801) Len() int           { return len(h) }
//func (h hp1801) Less(i, j int) bool { return h[i][0] <= h[j][0] }
//func (h hp1801) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//func (h *hp1801) Push(x any)        { *h = append(*h, x.([2]int)) }
//func (h *hp1801) Pop() any {
//	v := (*h)[h.Len()-1]
//	*h = (*h)[:h.Len()-1]
//	return v
//}
//
//var _ heap.Interface = (*hp1801)(nil)
