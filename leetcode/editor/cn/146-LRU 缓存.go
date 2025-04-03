package main

import "fmt"

func main() {
	//f := func(lru LRUCache) {
	//	v := lru.h.next
	//	for v != lru.t {
	//		fmt.Print(v.val, "->")
	//		v = v.next
	//	}
	//	fmt.Println()
	//}

	//lru := Constructor(2)
	//lru.Put(1, 1)
	//lru.Put(2, 2)
	////f(lru)
	//get := lru.Get(1)
	//fmt.Println(get)
	//lru.Put(3, 3)
	////f(lru)
	////fmt.Println(lru.m)
	//get = lru.Get(2)
	//fmt.Println(get)
	//lru.Put(4, 4)
	////f(lru)
	//get = lru.Get(1)
	//fmt.Println(get)
	//get = lru.Get(3)
	//fmt.Println(get)
	//get = lru.Get(4)
	//fmt.Println(get)

	lru := Constructor(1)
	lru.Put(2, 1)
	//f(lru)
	get := lru.Get(2)
	fmt.Println(get)
	lru.Put(3, 2)
	//f(lru)
	//fmt.Println(lru.m)
	get = lru.Get(2)
	fmt.Println(get)
	get = lru.Get(3)
	fmt.Println(get)
}

/*
利用数组实现 LRU 缓存淘汰策略
	1. 哈希表存储 [k-index] 实现 O(1) 查询/插入/删除，即 key 和 key 在数组中存储的位置索引
	2. 数组存储 [k-v]，默认容量为两倍 capacity
	- 查询：通过哈希表实现，并在当前队首的数据之前插入
	- 插入：视为循环数组，在当前队首的数据之前插入
	- 删除：将数据标记为已删除，当累积到某个阈值时再批量删除
	- “批量删除”：1.数据总量即将超过 capacity，则从队尾开始找到一个有效数据删除，O(n)。2.在插入数据时队首的“前一个位置”已被占，此时将全部有效数据依次挪到队首数据的后面，然后再插入，O(n)。
	3. 细节
	- h 记录了当前的队首数据，在 lru 队列有变更时需要更新 h
	- 两倍的 capacity 数组长度，为标记已删除数据留了充足的空间
	- 在循环数组中更新数据时，要时刻保持 map 的数据与数组的数据的一致性
	- 将全部有效数据连续排列后，可以保证后 n 次操作都是 O(1) 的，所以均摊时间复杂度为 O(1)。极端情况为连续 x*n 次插入新数据，此为有待改进的地方。
	4. 优化：添加了队尾变量 t，耗时由 500 ms 变为 411 ms
*/

// leetcode submit region begin(Prohibit modification and deletion)
//type LRUCache struct {
//
//}
//type ele struct {
//
//}
//
//func Constructor(capacity int) LRUCache {
//
//}
//
//func (this *LRUCache) Get(key int) int {
//
//}
//
//func (this *LRUCache) Put(key int, value int) {
//
//}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)

//type LRUCache struct {
//	h, t *ele         // 最近最少使用缓存
//	m    map[int]*ele // 快速查询
//	n    int
//}
//type ele struct {
//	key, val  int
//	pre, next *ele
//}
//
//func Constructor(capacity int) LRUCache {
//	e := &ele{}
//	e.pre, e.next = e, e
//	return LRUCache{e, e, make(map[int]*ele, capacity), capacity}
//}
//func (this *LRUCache) Get(key int) int {
//	if e, ok := this.m[key]; ok {
//		this.update(e)
//		return e.val
//	}
//	return -1
//}
//func (this *LRUCache) Put(key int, value int) {
//	e, ok := this.m[key]
//	if ok { // 已存在
//		e.val = value // 更新 val
//	} else {
//		if len(this.m) == this.n { // 已满
//			e = this.h.next           // 淘汰的节点
//			delete(this.m, e.key)     // 从 cache 删除
//			e.key, e.val = key, value // 复用结点
//		} else { // 新建结点
//			e = &ele{key: key, val: value}
//		}
//		this.m[key] = e // 添加到 cache
//	}
//	this.update(e)
//}
//func (this *LRUCache) update(e *ele) {
//	if e.next == this.t { // 已经是优先级最高
//		return
//	}
//	if e.pre != nil { // 说明不是新建结点
//		e.pre.next, e.next.pre = e.next, e.pre // 删除结点
//	}
//	this.t.pre.next, this.t.pre, e.next, e.pre = e, e, this.t, this.t.pre // 插入结点
//}

// 数据结构与算法之美：06|链表（上），如何利用数组实现 LRU 缓存淘汰策略呢？
//const inf146 = -1
//
//type LRUCache struct {
//	arr  [][2]int    // [2]:k v
//	cap  int         // 总长度
//	h  int         // 当前队首
//	memo map[int]int // key-index，同时存储了总数
//}
//
//func Constructor(capacity int) LRUCache {
//	arr := make([][2]int, capacity<<1)
//	for i := range arr {
//		arr[i][0] = inf146
//	}
//	return LRUCache{arr, capacity << 1, 0, make(map[int]int, capacity)}
//}
//func (this *LRUCache) Get(key int) int {
//	if i, ok := this.memo[key]; ok { // 已存在
//		if i != this.h { // fast fail
//			this.update(i)
//			this.memo[key] = this.h
//		}
//		return this.arr[this.h][1]
//	}
//	return -1
//}
//func (this *LRUCache) Put(key int, value int) {
//	if i, ok := this.memo[key]; ok { // 已存在
//		if i != this.h { // fast fail
//			this.update(i)
//			this.memo[key] = this.h
//		}
//		this.arr[this.h][1] = value // 更新 value
//		return
//	}
//	j := (this.h + this.cap - 1) % this.cap
//	if len(this.memo) == this.cap>>1 { // 移除队尾的数据
//		b := j
//		for this.arr[b][0] == inf146 {
//			b = (b + this.cap - 1) % this.cap
//		}
//		delete(this.memo, this.arr[b][0])
//		this.arr[b][0] = inf146
//	} else if this.arr[j][0] != inf146 {
//		this.move(j, 2) // 队首 & 队尾
//	}
//	this.memo[key] = j
//	this.arr[j][0], this.arr[j][1] = key, value // 添加新数据
//	this.h = j
//}
//func (this *LRUCache) update(i int) {
//	j := (this.h + this.cap - 1) % this.cap // 在 h 之前插入
//	if i == j {
//		this.h = j
//		return
//	}
//	k := this.arr[i]        // 保存
//	this.arr[i][0] = inf146 // 清除标记
//	if this.arr[j][0] != inf146 {
//		this.move(j, 3) // 队首 & arr[i] & 队尾
//	}
//	this.arr[j] = k
//	this.h = j
//}
//func (this *LRUCache) move(i, d int) { // 目标位置已被占，把有效数据的挪动为连续的队列
//	left := (this.h + 1) % this.cap                                               // 跳过队首
//	for c, right := len(this.memo)-d, left; c > 0; right = (right + 1) % this.cap { // len(this.memo)-2：i、h、j 抹去
//		if this.arr[right][0] != inf146 {
//			if left != right { // 重要：l == r 时，不能 this.arr[r][0] = inf146
//				this.arr[left] = this.arr[right]     // 挪动
//				this.memo[this.arr[right][0]] = left // 更新 memo
//				this.arr[right][0] = inf146          // 清除标记
//			}
//			left = (left + 1) % this.cap
//			c--
//		}
//	}
//	this.arr[left] = this.arr[i] // 移动最后一个位置，即队尾且被占的位置
//	this.memo[this.arr[i][0]] = left
//}

const inf146 = -1

type LRUCache struct {
	arr  [][2]int    // [2]:k v
	cap  int         // 总长度
	h, t int         // 当前队首 / 队尾
	memo map[int]int // key-index，同时存储了总数
}

func Constructor(capacity int) LRUCache {
	arr := make([][2]int, capacity<<1)
	for i := range arr {
		arr[i][0] = inf146
	}
	return LRUCache{arr, capacity << 1, 0, 0, make(map[int]int, capacity)}
}
func (this *LRUCache) Get(key int) int {
	if i, ok := this.memo[key]; ok { // 已存在
		if i != this.h { // fast fail
			this.update(i)
			this.memo[key] = this.h
		}
		return this.arr[this.h][1]
	}
	return -1
}
func (this *LRUCache) Put(key int, value int) {
	if i, ok := this.memo[key]; ok { // 已存在
		if i != this.h { // fast fail
			this.update(i)
			this.memo[key] = this.h
		}
		this.arr[this.h][1] = value // 更新 value
		return
	}
	j := (this.h - 1 + this.cap) % this.cap
	if len(this.memo) == this.cap>>1 { // 数据总数达到 capacity
		delete(this.memo, this.arr[this.t][0])
		this.arr[this.t][0] = inf146 // 将队尾标记为删除
		this.tail()
	} else if this.arr[j][0] != inf146 {
		this.move(j, 2) // h & t
	}
	if len(this.memo) == 0 { // 第一个数据
		this.t = j
	}
	this.memo[key] = j
	this.arr[j][0], this.arr[j][1] = key, value // 添加新数据
	this.h = j
}
func (this *LRUCache) update(i int) {
	j := (this.h - 1 + this.cap) % this.cap // 在 h 之前插入
	if i == j {                             // 特殊情况 1 & 2
		this.h = j
		this.tail() // t 肯定是 j
		return
	}
	kv := this.arr[i] // 暂存
	this.arr[i][0] = inf146
	if this.arr[j][0] != inf146 { // 特殊情况 1
		this.move(j, 3) // h & arr[i] & t
	} else if i == this.t { // 特殊情况 2
		this.tail()
	}
	this.arr[j] = kv
	this.h = j
}
func (this *LRUCache) move(i, d int) { // 目标位置已被占，把有效数据的挪动为连续的队列
	left := (this.h + 1) % this.cap // 跳过队首
	for c, right := len(this.memo)-d, left; c > 0; right = (right + 1) % this.cap {
		if this.arr[right][0] != inf146 {
			if left != right { // 重要：l == r 时，不能 this.arr[r][0] = inf146
				this.arr[left] = this.arr[right]     // 挪动
				this.memo[this.arr[right][0]] = left // 更新 memo
				this.arr[right][0] = inf146          // 清除标记
			}
			left = (left + 1) % this.cap
			c--
		}
	}
	this.arr[left] = this.arr[i] // 跳过队尾：最后一个移动的数据，即在队尾的数据且位置已被占
	this.memo[this.arr[i][0]] = left
	this.t = left // this.arr[j][0] != inf146：j 肯定是队尾
}
func (this *LRUCache) tail() { // 更新队尾
	if len(this.memo) == 0 { // 防止 capacity = 1 导致无限循环
		return
	}
	b := (this.t - 1 + this.cap) % this.cap
	for this.arr[b][0] == inf146 {
		b = (b - 1 + this.cap) % this.cap
	}
	this.t = b
}
