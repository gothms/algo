//
// 请你设计并实现一个满足
// LRU (最近最少使用) 缓存 约束的数据结构。
//
//
//
// 实现
// LRUCache 类：
//
//
//
//
//
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
//key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//
//
//
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
//
//
// 示例：
//
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
//
//
//
//
// 提示：
//
//
// 1 <= capacity <= 3000
// 0 <= key <= 10000
// 0 <= value <= 10⁵
// 最多调用 2 * 10⁵ 次 get 和 put
//
//
// Related Topics 设计 哈希表 链表 双向链表 👍 2868 👎 0

package main

import "fmt"

func main() {
	f := func(lru LRUCache) {
		v := lru.h.next
		for v != lru.t {
			fmt.Print(v.val, "->")
			v = v.next
		}
		fmt.Println()
	}
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	f(lru)
	get := lru.Get(1)
	fmt.Println(get)
	lru.Put(3, 3)
	f(lru)
	fmt.Println(lru.m)
	get = lru.Get(2)
	fmt.Println(get)
	lru.Put(4, 4)
	f(lru)
	get = lru.Get(1)
	fmt.Println(get)
	get = lru.Get(3)
	fmt.Println(get)
	get = lru.Get(4)
	fmt.Println(get)
}

// leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	h, t *ele         // 最近最少使用缓存
	m    map[int]*ele // 快速查询
	n    int
}
type ele struct {
	key, val  int
	pre, next *ele
}

func Constructor(capacity int) LRUCache {
	e := &ele{}
	e.pre, e.next = e, e
	return LRUCache{e, e, make(map[int]*ele, capacity), capacity}
}
func (this *LRUCache) Get(key int) int {
	if e, ok := this.m[key]; ok {
		this.update(e)
		return e.val
	}
	return -1
}
func (this *LRUCache) Put(key int, value int) {
	e, ok := this.m[key]
	if ok { // 已存在
		e.val = value // 更新 val
	} else {
		if len(this.m) == this.n { // 已满
			e = this.h.next           // 淘汰的节点
			delete(this.m, e.key)     // 从 cache 删除
			e.key, e.val = key, value // 复用结点
		} else { // 新建结点
			e = &ele{key: key, val: value}
		}
		this.m[key] = e // 添加到 cache
	}
	this.update(e)
}
func (this *LRUCache) update(e *ele) {
	if e.next == this.t { // 已经是优先级最高
		return
	}
	if e.pre != nil { // 说明不是新建结点
		e.pre.next, e.next.pre = e.next, e.pre // 删除结点
	}
	this.t.pre.next, this.t.pre, e.next, e.pre = e, e, this.t, this.t.pre // 插入结点
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
func (this *LRUCache) delVal(e *ele) {
	e.pre.next, e.next.pre = e.next, e.pre // 删除结点
}
func (this *LRUCache) pushHead(e *ele) {
	this.h.next, this.h.next.pre, e.pre, e.next = e, e, this.h, this.h.next // 插入结点
}
