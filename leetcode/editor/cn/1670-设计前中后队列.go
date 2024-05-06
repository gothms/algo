//请你设计一个队列，支持在前，中，后三个位置的 push 和 pop 操作。
//
// 请你完成 FrontMiddleBack 类：
//
//
// FrontMiddleBack() 初始化队列。
// void pushFront(int val) 将 val 添加到队列的 最前面 。
// void pushMiddle(int val) 将 val 添加到队列的 正中间 。
// void pushBack(int val) 将 val 添加到队里的 最后面 。
// int popFront() 将 最前面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
// int popMiddle() 将 正中间 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
// int popBack() 将 最后面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
//
//
// 请注意当有 两个 中间位置的时候，选择靠前面的位置进行操作。比方说：
//
//
// 将 6 添加到 [1, 2, 3, 4, 5] 的中间位置，结果数组为 [1, 2, 6, 3, 4, 5] 。
// 从 [1, 2, 3, 4, 5, 6] 的中间位置弹出元素，返回 3 ，数组变为 [1, 2, 4, 5, 6] 。
//
//
//
//
// 示例 1：
//
//
//输入：
//["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle",
//"popFront", "popMiddle", "popMiddle", "popBack", "popFront"]
//[[], [1], [2], [3], [4], [], [], [], [], []]
//输出：
//[null, null, null, null, null, 1, 3, 4, 2, -1]
//
//解释：
//FrontMiddleBackQueue q = new FrontMiddleBackQueue();
//q.pushFront(1);   // [1]
//q.pushBack(2);    // [1, 2]
//q.pushMiddle(3);  // [1, 3, 2]
//q.pushMiddle(4);  // [1, 4, 3, 2]
//q.popFront();     // 返回 1 -> [4, 3, 2]
//q.popMiddle();    // 返回 3 -> [4, 2]
//q.popMiddle();    // 返回 4 -> [2]
//q.popBack();      // 返回 2 -> []
//q.popFront();     // 返回 -1 -> [] （队列为空）
//
//
//
//
// 提示：
//
//
// 1 <= val <= 10⁹
// 最多调用 1000 次 pushFront， pushMiddle， pushBack， popFront， popMiddle 和 popBack 。
//
//
//
// Related Topics 设计 队列 数组 链表 数据流 👍 89 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)

type fmbq struct {
	pre, next *fmbq
	val       int
}
type FrontMiddleBackQueue struct {
	h, m, t *fmbq
	n       int
}

func Constructor() FrontMiddleBackQueue {
	f := &fmbq{}
	f.pre, f.next = f, f
	return FrontMiddleBackQueue{f, f, f, 0}
}
func (this *FrontMiddleBackQueue) PushFront(val int) {
	insertAfter(this.h, &fmbq{val: val})
	this.n++
	if this.n == 1 || this.n&1 == 0 {
		this.m = this.m.pre
	}
}
func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	f := &fmbq{val: val}
	if this.n&1 == 0 { // 往后插入
		insertAfter(this.m, f)
	} else { // 往前
		insertBefore(this.m, f)
	}
	this.n++
	this.m = f
}
func (this *FrontMiddleBackQueue) PushBack(val int) {
	insertBefore(this.t, &fmbq{val: val})
	this.n++
	if this.n&1 == 1 {
		this.m = this.m.next
	}
}
func (this *FrontMiddleBackQueue) PopFront() int {
	return this.pop(this.h.next, func() {
		if this.n == 0 || this.n&1 == 1 {
			this.m = this.m.next
		}
	})
}
func (this *FrontMiddleBackQueue) PopMiddle() int {
	return this.pop(this.m, func() {
		if this.n&1 == 0 {
			this.m = this.m.pre
		} else {
			this.m = this.m.next
		}
	})
}
func (this *FrontMiddleBackQueue) PopBack() int {
	return this.pop(this.t.pre, func() {
		if this.n&1 == 0 {
			this.m = this.m.pre
		}
	})
}
func insertAfter(pre, c *fmbq) {
	pre.next, c.pre, c.next, pre.next.pre = c, pre, pre.next, c
}
func insertBefore(next, c *fmbq) {
	next.pre, c.next, c.pre, next.pre.next = c, next, next.pre, c
}
func (this *FrontMiddleBackQueue) pop(c *fmbq, f func()) int {
	if this.n == 0 {
		return -1
	}
	c.pre.next, c.next.pre = c.next, c.pre // 删除结点
	this.n--
	f()                      // 调整中点
	c.pre, c.next = nil, nil // 回收
	return c.val
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
//leetcode submit region end(Prohibit modification and deletion)
