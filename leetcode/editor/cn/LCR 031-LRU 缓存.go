//
// 运用所掌握的数据结构，设计和实现一个 LRU (Least Recently Used，最近最少使用) 缓存机制 。
//
//
// 实现 LRUCache 类：
//
//
// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上
//限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
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
//
//
// 进阶：是否可以在 O(1) 时间复杂度内完成这两种操作？
//
//
//
//
// 注意：本题与主站 146 题相同：https://leetcode-cn.com/problems/lru-cache/
//
// Related Topics 设计 哈希表 链表 双向链表 👍 102 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	//f := func(lru LRUCache) {
	//	v := lru.h.next
	//	for v != lru.t {
	//		fmt.Print(v.val, "->")
	//		v = v.next
	//	}
	//	fmt.Println()
	//}
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	get := lru.Get(1)
	fmt.Println(get)
	lru.Put(3, 3)
	fmt.Println(lru.m)
	get = lru.Get(2)
	fmt.Println(get)
	lru.Put(4, 4)
	get = lru.Get(1)
	fmt.Println(get)
	get = lru.Get(3)
	fmt.Println(get)
	get = lru.Get(4)
	fmt.Println(get)
}

// leetcode submit region begin(Prohibit modification and deletion)

type LRUCache struct {
	l   list.List
	m   map[int]*list.Element
	cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{m: make(map[int]*list.Element, capacity), cap: capacity}
}
func (this *LRUCache) Get(key int) int {
	if e, ok := this.m[key]; ok {
		this.update(e, nil) // Get 不用放入 map
		return e.Value.([]int)[1]
	}
	return -1
}
func (this *LRUCache) Put(key int, value int) {
	e, ok := this.m[key] // 三种情况
	val := []int{key, value}
	if !ok && len(this.m) == this.cap {
		e = this.l.Front()
		delete(this.m, e.Value.([]int)[0])
		this.m[key] = e // 复用结点，key 已变
	}
	if e != nil { // 修改结点 / 复用结点
		e.Value = val
	}
	this.update(e, val)
}
func (this *LRUCache) update(e *list.Element, val []int) {
	if e != nil {
		if this.l.Back() == e {
			return
		}
		this.l.MoveToBack(e)
	} else { // 只有新建结点时，e == nil
		back := this.l.PushBack(val)
		this.m[val[0]] = back
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
