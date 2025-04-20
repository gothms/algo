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
	removal := minimumPairRemoval(nums)
	fmt.Println(removal)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumPairRemoval(nums []int) int {
	// lc：懒删除
	// map[pair]int 记录点对 pair 需要从堆中删除的次数
	// 当堆顶元素 for map[pair] > 0 时，直接丢弃元素

	// 个人
	n := len(nums)
	pre, suf := make([]int, n+1), make([]int, n+1) // 哨兵
	for i := range pre {                           // [-1 0 1 2 ... n-1]
		pre[i] = i - 1
	}
	for i := range suf { // [0 1 2 3 ... n-1 n]
		suf[i] = i
	}
	var findPre func(int) int   // 找到前驱
	findPre = func(i int) int { // i+1
		if i != pre[i+1] {
			pre[i+1] = findPre(pre[i+1])
		}
		return pre[i+1]
	}
	var findSuf func(int) int // 找到后驱
	findSuf = func(i int) int {
		if i != suf[i] {
			suf[i] = findSuf(suf[i])
		}
		return suf[i]
	}

	arr := make(hp, 0, n-1)
	h := &arr
	ans, preV, cnt := 0, nums[0], 0
	for i, v := range nums[1:] {
		heap.Push(h, [4]int{i, i + 1, preV, v})
		if preV > v {
			cnt++ // 相邻的逆序对
		}
		preV = v
	}
	for cnt > 0 /*&& (*h).Len() > 0*/ {
		top := (*h)[0]
		i, j, iv, jv := top[0], top[1], top[2], top[3]
		heap.Pop(h)                             // 无论如何都要弹出
		if findSuf(j) == j && findSuf(i) == i { // find(i) == i，说明 [i-1,i] 还没被操作
			if nums[j] != jv { // nums[j] 已被修改
				continue
			}
			// 1.删除 j
			// 找到 i 的 前缀 和 后缀 p s
			// j 的后缀修改为 s，如果 s 合法时 j 的前缀修改为 i
			// 2.作废旧的对，新增新的对
			// 作废时尝试 cnt--，新增时尝试 cnt++
			if iv > jv { // 逆序对减少
				cnt--
			}
			ans++         // 操作一次
			nums[i] += jv // 合并
			p, s := findPre(i-1), findSuf(j+1)
			suf[j] = s
			if s < n { // 新对
				pre[j+1] = i
				heap.Push(h, [4]int{i, s, nums[i], nums[s]})
				if jv > nums[s] { // 作废的对
					cnt--
				}
				if nums[i] > nums[s] { // 新增的对
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
	}
	return ans
}

type hp [][4]int // i,j,sum

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	s1, s2 := h[i][2]+h[i][3], h[j][2]+h[j][3]
	return s1 < s2 || s1 == s2 && h[i][0] < h[j][0]
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)   { *h = append(*h, x.([4]int)) }
func (h *hp) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

var _ heap.Interface = (*hp)(nil)

//leetcode submit region end(Prohibit modification and deletion)
