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
	nums = []int{1, 2, 3, 4, 2, 3, 1, 4, 2}
	k := 3
	nums = []int{7, 9, 3, 8, 0, 2, 4, 8, 3, 9}
	nums = []int{1, 2}
	k = 1
	window := medianSlidingWindow(nums, k)
	fmt.Println(window)
}

/*
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
*/
// leetcode submit region begin(Prohibit modification and deletion)
func medianSlidingWindow(nums []int, k int) []float64 {
	// 堆 + 延迟删除
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

//leetcode submit region end(Prohibit modification and deletion)
