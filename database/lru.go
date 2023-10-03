package main

import (
	"container/list"
	"fmt"
)

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

// ====================自定义 List====================

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

// ====================API====================

type LRUCacheApi struct {
	l   list.List
	m   map[int]*list.Element
	cap int
}

func ConstructorApi(capacity int) LRUCacheApi {
	return LRUCacheApi{m: make(map[int]*list.Element, capacity), cap: capacity}
}
func (this *LRUCacheApi) Get(key int) int {
	if e, ok := this.m[key]; ok {
		this.update(e, nil) // Get 不用放入 map
		return e.Value.([]int)[1]
	}
	return -1
}
func (this *LRUCacheApi) Put(key int, value int) {
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
func (this *LRUCacheApi) update(e *list.Element, val []int) {
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
