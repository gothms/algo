//中位数是有序序列最中间的那个数。如果序列的长度是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。
//
// 例如：
//
//
// [2,3,4]，中位数是 3
// [2,3]，中位数是 (2 + 3) / 2 = 2.5
//
//
// 给你一个数组 nums，有一个长度为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗
//口中元素的中位数，并输出由它们组成的数组。
//
//
//
// 示例：
//
// 给出 nums = [1,3,-1,-3,5,3,6,7]，以及 k = 3。
//
//
//窗口位置                      中位数
//---------------               -----
//[1  3  -1] -3  5  3  6  7       1
// 1 [3  -1  -3] 5  3  6  7      -1
// 1  3 [-1  -3  5] 3  6  7      -1
// 1  3  -1 [-3  5  3] 6  7       3
// 1  3  -1  -3 [5  3  6] 7       5
// 1  3  -1  -3  5 [3  6  7]      6
//
//
// 因此，返回该滑动窗口的中位数数组 [1,-1,-1,3,5,6]。
//
//
//
// 提示：
//
//
// 你可以假设 k 始终有效，即：k 始终小于等于输入的非空数组的元素个数。
// 与真实值误差在 10 ^ -5 以内的答案将被视作正确答案。
//
//
// Related Topics 数组 哈希表 滑动窗口 堆（优先队列） 👍 435 👎 0

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	//nums = []int{1, 2, 3, 4, 2, 3, 1, 4, 2}
	k := 3 // [1.00000,-1.00000,-1.00000,3.00000,5.00000,6.00000]
	//nums = []int{7, 9, 3, 8, 0, 2, 4, 8, 3, 9}
	//nums = []int{1, 2}
	//k = 1
	window := medianSlidingWindow(nums, k)
	fmt.Println(window)

	//for i := range nums[:len(nums)-1] {
	//	fmt.Println(nums[i+1])
	//}
}

/*
堆 + 延迟删除
重点：
	1.保持堆顶元素是有效元素，否则 添加/删除 数据时，会误判堆顶元素
		1.1.push 后
		1.2.remove 后
		1.3.remove 的是堆顶元素，还要先 lazy
	2.正确计算 size 值
	3.lazy 删除已被移除的元素
		go API 中，优先队列的缺点正是不能“正确”删除堆中的元素
		此 lazy 方案提供了一种“正确“删除堆中元素的方式
	4.delay map 记录了被淘汰元素，以及其数量
		如果还有记录，则可以从堆顶删除
		否则说明堆顶元素在窗口中
*/
// leetcode submit region begin(Prohibit modification and deletion)
func medianSlidingWindow(nums []int, k int) []float64 {
	// 执行耗时:56 ms,击败了97.67% 的Go用户
	// 内存消耗:13.1 MB,击败了34.88% 的Go用户
	n := len(nums)
	ret := make([]float64, 0, n-k+1)
	switch k {
	case 1: // fast path
		for _, v := range nums {
			ret = append(ret, float64(v))
		}
		return ret
	case 2:
		for i := 1; i < n; i++ {
			ret = append(ret, float64(nums[i]+nums[i-1])/2.0)
		}
		return ret
	}
	hMin, hMax := &hp480{}, &hp480{}
	minCnt, maxCnt := (k+1)>>1, k>>1
	median := func() {
		if k&1 == 0 {
			ret = append(ret, float64((*hMin)[0][0]-(*hMax)[0][0])/2.0)
		} else {
			if minCnt > maxCnt {
				ret = append(ret, float64((*hMin)[0][0]))
			} else {
				ret = append(ret, float64(-(*hMax)[0][0]))
			}
		}
	}
	moveHeap := func(i int, f1, f2 func() bool) { // minCnt maxCnt 两者差的绝对值 > 1，则需要调整
		var (
			ma, mi *hp480
			ad     int
		)
		if f1() {
			ma, mi, ad = hMin, hMax, 1
		} else if f2() {
			ma, mi, ad = hMax, hMin, -1
		}
		for f1() || f2() {
			v := heap.Pop(ma).([2]int)
			if v[1] > i {
				minCnt -= ad
				maxCnt += ad
				v[0] = -v[0]
				heap.Push(mi, v)
			}
		}
	}
	balance := func(i int) {
		if k&1 == 0 { // 偶数：元素多的堆弹出
			moveHeap(i, func() bool {
				return minCnt > maxCnt
			}, func() bool {
				return minCnt < maxCnt
			})
		} else { // 奇数：元素”相对多“的堆弹出
			moveHeap(i, func() bool {
				return minCnt > maxCnt+1
			}, func() bool {
				return minCnt+1 < maxCnt
			})
		}
		for hMin.Len() > 0 && (*hMin)[0][1] <= i { // 平衡后，滑出淘汰的数据
			heap.Pop(hMin)
		}
		for hMax.Len() > 0 && (*hMax)[0][1] <= i {
			heap.Pop(hMax)
		}
	}
	for i := 0; i < k; i++ { // 先处理前 k 个元素
		heap.Push(hMin, [2]int{nums[i], i})
	}
	for i := k >> 1; i > 0; i-- {
		v := heap.Pop(hMin).([2]int)
		v[0] = -v[0]
		heap.Push(hMax, v)
	}
	median() // 第一个中位数
	for i := k; i < n; i++ {
		d := i - k
		out, in := nums[d], nums[i]
		if out > (*hMin)[0][0] || out == (*hMin)[0][0] && d >= (*hMin)[0][1] { // 重要：滑出的元素在哪个堆
			minCnt--
		} else {
			maxCnt--
		}
		if in > (*hMin)[0][0] { // 滑入的元素在哪个堆
			heap.Push(hMin, [2]int{in, i})
			minCnt++
		} else {
			heap.Push(hMax, [2]int{-in, i})
			maxCnt++
		}
		balance(d)
		median()
	}
	return ret

}

type hp480 [][2]int

func (h hp480) Len() int { return len(h) }
func (h hp480) Less(i, j int) bool {
	return h[i][0] < h[j][0] || h[i][0] == h[j][0] && h[i][1] < h[j][1]
}
func (h hp480) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp480) Push(x any)   { *h = append(*h, x.([2]int)) }
func (h *hp480) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

// leetcode submit region end(Prohibit modification and deletion)

func medianSlidingWindow_(nums []int, k int) []float64 {
	// 对顶堆：个人
	// 执行耗时:60 ms,击败了97.67% 的Go用户
	// 内存消耗:12.6 MB,击败了37.21% 的Go用户
	n := len(nums)
	ret := make([]float64, 0, n-k+1)
	switch k {
	case 1:
		for _, v := range nums {
			ret = append(ret, float64(v))
		}
		return ret
	case 2:
		for i := 1; i < n; i++ {
			ret = append(ret, float64(nums[i]+nums[i-1])/2.0)
		}
		return ret
	}
	hMin, hMax := &hp480{}, &hp480{}
	minCnt, maxCnt := (k+1)>>1, k>>1
	median := func() {
		if k&1 == 0 {
			ret = append(ret, float64((*hMin)[0][0]-(*hMax)[0][0])/2.0)
		} else {
			if minCnt > maxCnt {
				ret = append(ret, float64((*hMin)[0][0]))
			} else {
				ret = append(ret, float64(-(*hMax)[0][0]))
			}
		}
	}
	balance := func(i int) {
		var (
			ma, mi *hp480
			ad     int
		)
		if k&1 == 0 { // 偶数：元素多的堆弹出
			if minCnt > maxCnt {
				ma, mi, ad = hMin, hMax, 1
			} else if minCnt < maxCnt {
				ma, mi, ad = hMax, hMin, -1
			}
			for minCnt != maxCnt { // 平衡对顶堆
				v := heap.Pop(ma).([2]int)
				if v[1] > i {
					minCnt -= ad
					maxCnt += ad
					v[0] = -v[0]
					heap.Push(mi, v)
				}
			}
		} else { // 奇数：元素”相对多“的堆弹出
			if minCnt > maxCnt+1 {
				ma, mi, ad = hMin, hMax, 1
			} else if minCnt+1 < maxCnt {
				ma, mi, ad = hMax, hMin, -1
			}
			for minCnt != maxCnt+1 && minCnt+1 != maxCnt {
				v := heap.Pop(ma).([2]int)
				if v[1] > i {
					minCnt -= ad
					maxCnt += ad
					v[0] = -v[0]
					heap.Push(mi, v)
				}
			}
		}
		for hMin.Len() > 0 && (*hMin)[0][1] <= i { // 平衡后，滑出淘汰的数据
			heap.Pop(hMin)
		}
		for hMax.Len() > 0 && (*hMax)[0][1] <= i {
			heap.Pop(hMax)
		}
	}
	for i := 0; i < k; i++ { // 先处理前 k 个元素
		heap.Push(hMin, [2]int{nums[i], i})
	}
	for i := k >> 1; i > 0; i-- {
		v := heap.Pop(hMin).([2]int)
		v[0] = -v[0]
		heap.Push(hMax, v)
	}
	median() // 第一个中位数
	for i := k; i < n; i++ {
		d := i - k
		out, in := nums[d], nums[i]
		if out > (*hMin)[0][0] || out == (*hMin)[0][0] && d >= (*hMin)[0][1] { // 重要：滑出的元素在哪个堆
			minCnt--
		} else {
			maxCnt--
		}
		if in > (*hMin)[0][0] { // 滑入的元素在哪个堆
			heap.Push(hMin, [2]int{in, i})
			minCnt++
		} else {
			heap.Push(hMax, [2]int{-in, i})
			maxCnt++
		}
		balance(d)
		median()
	}
	return ret
}

func medianSlidingWindow_lc(nums []int, k int) []float64 {
	// lc：堆 + 延迟删除，很复杂
	n := len(nums)
	ret := make([]float64, 0, n-k+1)
	if k == 1 { // fast path
		for _, v := range nums {
			ret = append(ret, float64(v))
		}
		return ret
	}
	var (
		maxHp, minHp = &mswHp{}, &mswHp{} // 大 & 小顶堆，优先小顶堆
		delay        = make(map[int]int)  // k：已删除元素，v：个数
	)
	lazy := func(h *mswHp) { // 由于 maxHp, minHp 是局部的，所以函数也设为局部
		var v int
		for h.Len() > 0 {
			v = h.IntSlice[0]
			if h == maxHp {
				v = -v // 大顶堆
			}
			if c, ok := delay[v]; !ok { // 尝试 Pop 已出列的元素
				break
			} else {
				if c == 1 { // 记录了元素的个数
					delete(delay, v)
				} else {
					delay[v]--
				}
				heap.Pop(h) // Pop
			}
		}
	}
	balance := func() { // 维护两个堆元素的”平衡“，优先小顶堆
		if minHp.size > maxHp.size+1 {
			heap.Push(maxHp, -heap.Pop(minHp).(int))
			minHp.size-- // 维护 size
			maxHp.size++
			lazy(minHp)
		} else if maxHp.size > minHp.size {
			heap.Push(minHp, -heap.Pop(maxHp).(int))
			minHp.size++
			maxHp.size--
			lazy(maxHp) // 重要：每一次 balance 后，都要维护堆顶元素的有效性
		}
	}
	push := func(v int) { // 往两个堆中添加元素
		if minHp.Len() == 0 || v >= minHp.IntSlice[0] {
			heap.Push(minHp, v)
			minHp.size++
		} else {
			heap.Push(maxHp, -v)
			maxHp.size++
		}
		balance()
	}
	remove := func(v int) { // “删除”出列的元素
		delay[v]++
		if v >= minHp.IntSlice[0] {
			minHp.size--
			if v == minHp.IntSlice[0] {
				lazy(minHp) // 如果是堆顶元素，则先 lazy
				//balance()
			}
		} else {
			maxHp.size--
			if v == -maxHp.IntSlice[0] {
				lazy(maxHp)
				//balance()
			}
		}
		balance()
	}
	getMedian := func() float64 {
		if k&1 == 0 {
			return float64(minHp.IntSlice[0]-maxHp.IntSlice[0]) / 2.0
		} else {
			return float64(minHp.IntSlice[0])
		}
	}
	for i := 0; i < k; i++ {
		push(nums[i])
	}
	ret = append(ret, getMedian())
	for i := k; i < n; i++ {
		push(nums[i])
		remove(nums[i-k])
		ret = append(ret, getMedian())
	}
	return ret
}

type mswHp struct {
	sort.IntSlice
	size int
}

func (m *mswHp) Push(x any) { m.IntSlice = append(m.IntSlice, x.(int)) }
func (m *mswHp) Pop() any {
	a := m.IntSlice
	v := a[len(a)-1]
	m.IntSlice = a[:len(a)-1]
	return v
}
