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
	h, t *l         // 最近最少使用缓存
	m    map[int]*l // 快速查询
	n    int
}
type l struct {
	key, val  int
	pre, next *l
}

func Constructor(capacity int) LRUCache {
	v := &l{}
	v.pre, v.next = v, v
	return LRUCache{v, v, make(map[int]*l, capacity), capacity}
}
func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.delVal(v)
		this.pushHead(v)
		return v.val
	}
	return -1
}
func (this *LRUCache) Put(key int, value int) {
	v, ok := this.m[key]
	if ok { // 已存在
		v.val = value // 更新 val
		if v.pre == this.h {
			return
		}
		this.delVal(v)
	} else {
		if len(this.m) == this.n { // 已满
			del := this.t.pre       // 淘汰的节点
			delete(this.m, del.key) // 删除
			this.delVal(del)
		}
		v = &l{key: key, val: value}
		this.m[key] = v // 添加
	}
	this.pushHead(v)
}
func (this *LRUCache) delVal(v *l) {
	v.pre.next, v.next.pre = v.next, v.pre // 删除结点
}
func (this *LRUCache) pushHead(v *l) {
	this.h.next, this.h.next.pre, v.pre, v.next = v, v, this.h, this.h.next // 插入结点
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
