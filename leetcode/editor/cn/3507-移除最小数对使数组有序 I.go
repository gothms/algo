package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{5, 2, 3, 1}                             // 2
	nums = []int{2, 2, -1, 3, -2, 2, 1, 1, 1, 0, -1}      // 9
	nums = []int{-2, 1, 2, -1, -1, -2, -2, -1, -1, 1, 1}  // 10
	nums = []int{3, 6, 4, -6, 2, -4, 5, -7, -3, 6, 3, -4} // 10
	nums = []int{-7, -2, -4, 4, 8, -6, 0, 0, 4, 5, 1, -8} // 11
	removal := minimumPairRemoval(nums)
	fmt.Println(removal)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumPairRemoval(nums []int) int {
	n := len(nums)
	idxPre, idxNext := make([]int, n+1), make([]int, n+1)
	for i := 0; i <= n; i++ {
		idxPre[i], idxNext[i] = i-1, i
	}
	var findPre func(int) int
	findPre = func(i int) int {
		if i != idxPre[i+1] {
			idxPre[i+1] = findPre(idxPre[i+1])
		}
		return idxPre[i+1]
	}
	var findNext func(int) int
	findNext = func(i int) int {
		if i != idxNext[i] {
			idxNext[i] = findNext(idxNext[i])
		}
		return idxNext[i]
	}

	arr := make(hp, 0, n-1)
	h := &arr
	ans, cnt, pre := 0, 0, nums[0]
	for i, v := range nums[1:] {
		*h = append(*h, [4]int{i, i + 1, pre, v})
		if pre > v {
			cnt++
		}
		pre = v
	}
	heap.Init(h)
	for cnt > 0 {
		top := heap.Pop(h).([4]int)
		i, j, iv, jv := top[0], top[1], top[2], top[3]
		if j != findNext(j) || i != findNext(i) || jv != nums[j] {
			continue
		}
		if iv > jv {
			cnt--
		}
		ans++
		nums[i] += jv
		p, s := findPre(i-1), findNext(j+1)
		idxNext[j] = s
		if s < n {
			idxPre[j+1] = i
			heap.Push(h, [4]int{i, s, nums[i], nums[s]})
			if jv > nums[s] {
				cnt--
			}
			if nums[i] > nums[s] {
				cnt++
			}
		}
		if p >= 0 {
			heap.Push(h, [4]int{p, i, nums[p], nums[i]})
			if nums[p] > iv {
				cnt--
			}
			if nums[p] > nums[i] {
				cnt++
			}
		}
	}
	return ans
}

type hp [][4]int

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	s1, s2 := h[i][2]+h[i][3], h[j][2]+h[j][3]
	return s1 < s2 || s1 == s2 && h[i][0] < h[j][0]
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)   { *h = append(*h, x.([4]int)) }
func (h *hp) Pop() any {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}

var _ heap.Interface = (*hp)(nil)

//leetcode submit region end(Prohibit modification and deletion)
