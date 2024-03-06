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
type FrontMiddleBackQueue struct {
	f, m, b *fmbq
	n       int
}
type fmbq struct {
	val       int
	pre, next *fmbq
}

func Constructor() FrontMiddleBackQueue {
	q := &fmbq{}
	q.pre, q.next = q, q
	return FrontMiddleBackQueue{q, q, q, 0}
}
func (this *FrontMiddleBackQueue) push(q *fmbq, val int) {
	node := &fmbq{val: val}
	node.pre, node.next, q.next, q.next.pre = q, q.next, node, node // 增加结点
	this.n++
}
func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.push(this.f, val)
	if this.n&1 == 0 || this.n == 1 { // 奇偶讨论
		this.m = this.m.pre
	}
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	if this.n&1 == 0 {
		this.push(this.m, val)
		this.m = this.m.next
	} else {
		this.push(this.m.pre, val)
		this.m = this.m.pre
	}
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.push(this.b.pre, val)
	if this.n&1 == 1 {
		this.m = this.m.next
	}
}

func (this *FrontMiddleBackQueue) pop(q *fmbq) bool {
	if this.n == 0 {
		return false
	}
	this.n--
	q.pre.next, q.next.pre = q.next, q.pre // 删除结点
	return true
}
func (this *FrontMiddleBackQueue) PopFront() int {
	q := this.f.next
	if !this.pop(q) {
		return -1
	}
	if this.n&1 == 1 || this.n == 0 { // 奇偶讨论
		this.m = this.m.next
	}
	return q.val
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	q := this.m
	if !this.pop(q) {
		return -1
	}
	if this.n&1 == 0 {
		this.m = this.m.pre
	} else {
		this.m = this.m.next
	}
	return q.val
}

func (this *FrontMiddleBackQueue) PopBack() int {
	q := this.b.pre
	if !this.pop(q) {
		return -1
	}
	if this.n&1 == 0 {
		this.m = this.m.pre
	}
	return q.val
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
