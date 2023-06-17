//动物收容所。有家动物收容所只收容狗与猫，且严格遵守“先进先出”的原则。在收养该收容所的动物时，收养人只能收养所有动物中“最老”（由其进入收容所的时间长短而定
//）的动物，或者可以挑选猫或狗（同时必须收养此类动物中“最老”的）。换言之，收养人不能自由挑选想收养的对象。请创建适用于这个系统的数据结构，实现各种操作方法，比如
//enqueue、dequeueAny、dequeueDog和dequeueCat。允许使用Java内置的LinkedList数据结构。
//
// enqueue方法有一个animal参数，animal[0]代表动物编号，animal[1]代表动物种类，其中 0 代表猫，1 代表狗。
//
// dequeue*方法返回一个列表[动物编号, 动物种类]，若没有可以收养的动物，则返回[-1,-1]。
//
// 示例1:
//
//  输入：
//["AnimalShelf", "enqueue", "enqueue", "dequeueCat", "dequeueDog",
//"dequeueAny"]
//[[], [[0, 0]], [[1, 0]], [], [], []]
// 输出：
//[null,null,null,[0,0],[-1,-1],[1,0]]
//
//
// 示例2:
//
//  输入：
//["AnimalShelf", "enqueue", "enqueue", "enqueue", "dequeueDog", "dequeueCat",
//"dequeueAny"]
//[[], [[0, 0]], [[1, 0]], [[2, 1]], [], [], []]
// 输出：
//[null,null,null,null,[2,1],[0,0],[1,0]]
//
//
// 说明:
//
//
// 收纳所的最大容量为20000
//
//
// Related Topics 设计 队列 👍 52 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type AnimalShelf struct {
	ch, ct *Node
	dh, dt *Node
	idx    int
}
type Node struct {
	next, pre *Node
	v         []int
	i         int
}

func Constructor() AnimalShelf {
	c, d := &Node{}, &Node{}
	return AnimalShelf{ch: c, ct: c, dh: d, dt: d, idx: 1}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	n := &Node{v: animal, i: this.idx}
	if animal[1] == 0 {
		this.ct.next, this.ct, n.pre = n, n, this.ct
	} else {
		this.dt.next, this.dt, n.pre = n, n, this.dt
	}
	this.idx++
}

func (this *AnimalShelf) DequeueAny() []int {
	c, d := this.ch.next, this.dh.next
	if c == nil && d == nil {
		this.idx = 1
		return []int{-1, -1}
	} else if c == nil || d != nil && c.i > d.i {
		if d == this.dt {
			this.dt = d.pre
		}
		this.dh.next, d.pre = d.next, nil
		if d.next != nil {
			d.next.pre, d.next = this.dh, nil
		}
		return d.v
	} else {
		if c == this.ct {
			this.ct = c.pre
		}
		this.ch.next, c.pre = c.next, nil
		if c.next != nil {
			c.next.pre, c.next = this.ch, nil
		}
		return c.v
	}
}

func (this *AnimalShelf) DequeueDog() []int {
	d := this.dh.next
	if d == nil {
		return []int{-1, -1}
	}
	if d == this.dt {
		this.dt = d.pre
	}
	this.dh.next, d.pre = d.next, nil
	if d.next != nil {
		d.next.pre, d.next = this.dh, nil
	}
	return d.v
}

func (this *AnimalShelf) DequeueCat() []int {
	c := this.ch.next
	if c == nil {
		return []int{-1, -1}
	}
	if c == this.ct {
		this.ct = c.pre
	}
	this.ch.next, c.pre = c.next, nil
	if c.next != nil {
		c.next.pre, c.next = this.ch, nil
	}
	return c.v
}

//["AnimalShelf", "enqueue", "enqueue", "enqueue", "dequeueDog", "dequeueCat", "dequeueAny","dequeueAny"]
//[[], [[0, 0]], [[1, 0]], [[2, 1]], [], [], [],[]]
/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */
//leetcode submit region end(Prohibit modification and deletion)
