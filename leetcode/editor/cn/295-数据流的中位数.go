//中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。
//
//
// 例如 arr = [2,3,4] 的中位数是 3 。
// 例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。
//
//
// 实现 MedianFinder 类:
//
//
// MedianFinder() 初始化 MedianFinder 对象。
// void addNum(int num) 将数据流中的整数 num 添加到数据结构中。
// double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10⁻⁵ 以内的答案将被接受。
//
//
// 示例 1：
//
//
//输入
//["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
//[[], [1], [2], [], [3], []]
//输出
//[null, null, null, 1.5, null, 2.0]
//
//解释
//MedianFinder medianFinder = new MedianFinder();
//medianFinder.addNum(1);    // arr = [1]
//medianFinder.addNum(2);    // arr = [1, 2]
//medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
//medianFinder.addNum(3);    // arr[1, 2, 3]
//medianFinder.findMedian(); // return 2.0
//
// 提示:
//
//
// -10⁵ <= num <= 10⁵
// 在调用 findMedian 之前，数据结构中至少有一个元素
// 最多 5 * 10⁴ 次调用 addNum 和 findMedian
//
//
// Related Topics 设计 双指针 数据流 排序 堆（优先队列） 👍 881 👎 0

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	mf.AddNum(3)
	mf.AddNum(2)
	mf.AddNum(4)
	mf.AddNum(6)
	median := mf.FindMedian()
	fmt.Println(median)
	fmt.Println(mf.maxHp)
	fmt.Println(mf.minHp)
}

// leetcode submit region begin(Prohibit modification and deletion)

// MedianFinder 一种堆
type MedianFinder struct {
	maxHp *mfHp // 更小的数
	minHp *mfHp // 更大的数，长度 >= minHp
}
type mfHp struct{ sort.IntSlice } // 默认小顶堆

func (m *mfHp) Push(x any) { m.IntSlice = append(m.IntSlice, x.(int)) }
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
		if this.minHp.Len()-1 > this.maxHp.Len() {
			heap.Push(this.maxHp, -heap.Pop(this.minHp).(int))
		}
	} else {
		heap.Push(this.maxHp, -num) // 存“负值”
		if this.maxHp.Len() > this.minHp.Len() {
			heap.Push(this.minHp, -heap.Pop(this.maxHp).(int))
		}
	}
}
func (this *MedianFinder) FindMedian() float64 {
	if this.minHp.Len() > this.maxHp.Len() {
		return float64(this.minHp.IntSlice[0])
	}
	return float64(this.minHp.IntSlice[0]-this.maxHp.IntSlice[0]) / 2.0
}

// leetcode submit region end(Prohibit modification and deletion)

// MedianFinder 笨方法：两种堆
//type MedianFinder struct {
//	maxHp *maxMfHp // 更小的数，长度 >= minHp
//	minHp *mfHp    // 更大的数
//}
//type mfHp struct{ sort.IntSlice } // 默认小顶堆
//type maxMfHp struct{ mfHp }
//
//func (mm maxMfHp) Less(i, j int) bool { return mm.IntSlice[i] > mm.IntSlice[j] }
//func (m *mfHp) Push(x any)            { m.IntSlice = append(m.IntSlice, x.(int)) }
//func (m *mfHp) Pop() any {
//	v := m.IntSlice[len(m.IntSlice)-1]
//	m.IntSlice = m.IntSlice[:len(m.IntSlice)-1]
//	return v
//}
//
//func Constructor() MedianFinder {
//	return MedianFinder{&maxMfHp{}, &mfHp{}}
//}
//func (this *MedianFinder) AddNum(num int) {
//	if this.maxHp.Len() == 0 || num <= this.maxHp.IntSlice[0] { // 优先大顶堆
//		heap.Push(this.maxHp, num)
//		if this.maxHp.Len()-1 > this.minHp.Len() {
//			heap.Push(this.minHp, heap.Pop(this.maxHp))
//		}
//	} else {
//		heap.Push(this.minHp, num)
//		if this.minHp.Len() > this.maxHp.Len() {
//			heap.Push(this.maxHp, heap.Pop(this.minHp))
//		}
//	}
//}
//func (this *MedianFinder) FindMedian() float64 {
//	//if this.maxHp.Len() == 0 {	// 题意：肯定会先添加一个元素
//	//	return 0
//	//}
//	if this.maxHp.Len() > this.minHp.Len() {
//		return float64(this.maxHp.IntSlice[0])
//	}
//	return float64(this.maxHp.IntSlice[0]+this.minHp.IntSlice[0]) / 2.0
//}

// lc：红黑树？二叉平衡树？
//type MedianFinder struct {
//	nums        *redblacktree.Tree
//	total       int
//	left, right iterator
//}
//
//func Constructor() MedianFinder {
//	return MedianFinder{nums: redblacktree.NewWithIntComparator()}
//}
//
//func (mf *MedianFinder) AddNum(num int) {
//	if count, has := mf.nums.Get(num); has {
//		mf.nums.Put(num, count.(int)+1)
//	} else {
//		mf.nums.Put(num, 1)
//	}
//	if mf.total == 0 {
//		it := mf.nums.Iterator()
//		it.Next()
//		mf.left = iterator{it, 1}
//		mf.right = mf.left
//	} else if mf.total%2 == 1 {
//		if num < mf.left.Key().(int) {
//			mf.left.prev()
//		} else {
//			mf.right.next()
//		}
//	} else {
//		if mf.left.Key().(int) < num && num < mf.right.Key().(int) {
//			mf.left.next()
//			mf.right.prev()
//		} else if num >= mf.right.Key().(int) {
//			mf.left.next()
//		} else {
//			mf.right.prev()
//			mf.left = mf.right
//		}
//	}
//	mf.total++
//}
//
//func (mf *MedianFinder) FindMedian() float64 {
//	return float64(mf.left.Key().(int)+mf.right.Key().(int)) / 2
//}
//
//type iterator struct {
//	redblacktree.Iterator
//	count int
//}
//
//func (it *iterator) prev() {
//	if it.count > 1 {
//		it.count--
//	} else {
//		it.Prev()
//		it.count = it.Value().(int)
//	}
//}
//
//func (it *iterator) next() {
//	if it.count < it.Value().(int) {
//		it.count++
//	} else {
//		it.Next()
//		it.count = 1
//	}
//}
