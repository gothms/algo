//随机产生数字并传递给一个方法。你能否完成这个方法，在每次产生新值时，寻找当前所有值的中间值（中位数）并保存。
//
// 中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
//
// 例如，
//
// [2,3,4] 的中位数是 3
//
// [2,3] 的中位数是 (2 + 3) / 2 = 2.5
//
// 设计一个支持以下两种操作的数据结构：
//
//
// void addNum(int num) - 从数据流中添加一个整数到数据结构中。
// double findMedian() - 返回目前所有元素的中位数。
//
//
// 示例：
//
// addNum(1)
//addNum(2)
//findMedian() -> 1.5
//addNum(3)
//findMedian() -> 2
//
//
// Related Topics 设计 双指针 数据流 排序 堆（优先队列） 👍 37 👎 0

package main

import (
	"container/heap"
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
type MedianFinder struct {
	maxHp, minHp *mfHp
}
type mfHp struct {
	sort.IntSlice
}

func (m *mfHp) Push(x any) {
	m.IntSlice = append(m.IntSlice, x.(int))
}
func (m *mfHp) Pop() any {
	v := m.IntSlice[len(m.IntSlice)-1]
	m.IntSlice = m.IntSlice[:len(m.IntSlice)-1]
	return v
}
func Constructor() MedianFinder {
	return MedianFinder{&mfHp{}, &mfHp{}}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minHp.Len() == 0 || num >= this.minHp.IntSlice[0] { // 优先小顶堆
		heap.Push(this.minHp, num)
		if this.minHp.Len() > this.maxHp.Len()+1 { // 调整平衡
			heap.Push(this.maxHp, -heap.Pop(this.minHp).(int))
		}
	} else {
		heap.Push(this.maxHp, -num) // 大顶堆：存负值
		if this.maxHp.Len() > this.minHp.Len() {
			heap.Push(this.minHp, -heap.Pop(this.maxHp).(int))
		}
	}
}
func (this *MedianFinder) FindMedian() float64 {
	if this.minHp.Len() == this.maxHp.Len() { // 偶数
		return float64(this.minHp.IntSlice[0]-this.maxHp.IntSlice[0]) / 2.0
	} else {
		return float64(this.minHp.IntSlice[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
//leetcode submit region end(Prohibit modification and deletion)
