package main

import (
	"container/list"
	"fmt"
)

/*
LFU 论文
	http://dhruvbird.com/lfu.pdf
golang 实现
	https://cloud.tencent.com/developer/article/2306436

五种实现方式（五种数据结构）
	二层链表

补充
	所有解法都默认 head.next 优先级最低，tail.pre 优先级最高

*/

func main() {
	const N = 10
	lfu := Constructor(N)
	for i := 1; i <= N; i++ {
		lfu.Put(i, i*i)
	}
	//for i := 0; i < N; i++ {
	//	lfu.Get(rand.Intn(N) + 1)
	//}
	//for i := 1; i <= N; i++ {
	//	j := rand.Intn(N) + 10
	//	lfu.Put(j, j*j)
	//}
	lfu.Get(3)
	lfu.Get(2)

	for l := lfu.freqs.Front(); l != nil; l = l.Next() {
		frequencyItem := l.Value.(*FrequencyItem)
		//fmt.Println("frequencyItem:", frequencyItem.freq, frequencyItem)
		fmt.Printf("frequencyItem: %d, %p\n", frequencyItem.freq, l)
	}
	fmt.Println("==========================")
	for k, item := range lfu.bykey {
		fmt.Println(k, item)
		fmt.Printf("%p\n", item.frequencyParent.Value)
		//if item != nil {
		//	parent := item.frequencyParent
		//	fmt.Println(k, parent.Value)
		//}
	}
	//for l := lfu.freqs.Front(); l != nil; l = lfu.freqs.Front() {
	//	fmt.Println(l.Value)
	//	lfu.freqs.Remove(l)
	//}

	//temp := make(map[int]struct{})
	//for k, item := range lfu.bykey {
	//	fmt.Println(k, item.frequencyParent.Value.(*FrequencyItem).freq)
	//	//fmt.Println(k, item.key, item.value, item.frequencyParent.Value.(*FrequencyItem).freq)
	//	frequencyItem := item.frequencyParent.Value.(*FrequencyItem)
	//	//fmt.Println(last, frequencyItem.freq)
	//	if _, ok := temp[frequencyItem.freq]; !ok {
	//		temp[frequencyItem.freq] = struct{}{}
	//		m := frequencyItem.entries
	//		for key := range m {
	//			fmt.Println(key, frequencyItem.freq)
	//		}
	//	}
	//}
}

// ====================总 map + 链表 map（频率:链表）====================

type LFUCache struct {
	items   map[int]*ele
	lists   map[int]*eleList
	minFreq int
	cap     int
}
type eleList struct {
	head, tail *ele
}
type ele struct {
	key, val  int
	pre, next *ele
	f         int
}

func newEleList() *eleList {
	v := &ele{}
	v.next, v.pre = v, v
	return &eleList{
		head: v,
		tail: v,
	}
}
func (e *ele) remove() {
	e.pre.next, e.next.pre = e.next, e.pre
}
func (e *ele) move(at *ele) {
	at.next, at.next.pre, e.next, e.pre = e, e, at.next, at
}
func Constructor(capacity int) LFUCache {
	return LFUCache{
		items: make(map[int]*ele, capacity),
		lists: make(map[int]*eleList, capacity),
		cap:   capacity,
	}
}
func (this *LFUCache) Get(key int) int {
	if e, ok := this.items[key]; ok {
		this.update(e, false)
		return e.val
	}
	return -1
}
func (this *LFUCache) Put(key int, value int) {
	e, ok := this.items[key]
	if ok {
		e.val = value
		this.update(e, false)
		return
	}
	if len(this.items) == this.cap { // 复用结点
		list := this.lists[this.minFreq]
		e = list.head.next
		delete(this.items, e.key) // key 已变，必须删除
		e.key, e.val = key, value
		if this.minFreq == 1 { // 不用更新频率
			if list.tail.pre != e { // 调整位置
				e.remove()
				e.move(list.tail.pre)
			}
			goto add
		}
	} else {
		e = &ele{key: key, val: value}
	}
	this.update(e, true)
add:
	this.items[key] = e
}
func (this *LFUCache) update(e *ele, new bool) {
	if e.f > 0 { // ele 是旧结点
		e.remove()
		cur := this.lists[e.f]
		if cur.head.next == cur.tail {
			delete(this.lists, e.f)
			if this.minFreq == e.f { // 需要更新 minFreq
				this.minFreq++
			}
		}
	}
	if new { // 设置频率
		e.f = 1
		this.minFreq = 1 // 重置 minFreq
	} else {
		e.f++
	}
	next := this.lists[e.f]
	if next == nil { // 新的频率链表
		next = newEleList()
		this.lists[e.f] = next
	}
	e.move(next.tail.pre)
}

// ====================终版：自定义二维链表====================

type LFUCache struct {
	items      map[int]*ele
	head, tail *freqEle
	cap        int
}

//type Ele interface {
//	remove(e *doubly)
//	move(e, at *doubly)
//}
//
//type doubly struct { // // 需要根据 e.fe 找到 doubly 结点
//	pre, next *doubly
//}

type freqEle struct {
	//*doubly
	freq      int
	pre, next *freqEle
	h, t      *ele
}
type ele struct {
	//*doubly
	key, val  int
	pre, next *ele
	fe        *freqEle
}

//func (d *doubly) remove() {
//	d.pre.next, d.next.pre = d.next, d.pre
//}
//
//func (d *doubly) move(at *doubly) {
//	at.next, at.next.pre, d.next, d.pre = d, d, at.next, at
//}

func (fe *freqEle) remove() {
	fe.pre.next, fe.next.pre = fe.next, fe.pre
}
func (fe *freqEle) move(at *freqEle) {
	at.next, at.next.pre, fe.next, fe.pre = fe, fe, at.next, at
}
func (e *ele) remove() {
	e.pre.next, e.next.pre = e.next, e.pre
}
func (e *ele) move(at *ele) {
	at.next, at.next.pre, e.next, e.pre = e, e, at.next, at
}
func newFreqEle(freq int) *freqEle {
	e := &ele{}
	e.pre, e.next = e, e
	return &freqEle{
		freq: freq,
		h:    e,
		t:    e,
	}
}
func Constructor(capacity int) LFUCache {
	e := &freqEle{}
	e.pre, e.next = e, e
	return LFUCache{
		items: make(map[int]*ele, capacity),
		head:  e,
		tail:  e,
		cap:   capacity,
	}
}
func (this *LFUCache) Get(key int) int {
	if e, ok := this.items[key]; ok {
		this.update(e, false)
		return e.val
	}
	return -1
}
func (this *LFUCache) Put(key int, value int) {
	e, ok := this.items[key]
	if ok {
		e.val = value
		this.update(e, false)
		return
	}
	//defer func() {
	//	this.items[key] = e
	//}()
	if len(this.items) == this.cap {
		e = this.head.next.h.next
		delete(this.items, e.key) // key 已变，必须删除
		e.key, e.val = key, value
		if e.fe.freq == 1 { // 不用更新频率
			if t := this.head.next.t; t.pre != e { // 调整位置
				e.remove()
				e.move(t.pre)
			}
			goto add
		}
	} else {
		e = &ele{key: key, val: value}
	}
	this.update(e, true)
add:
	this.items[key] = e
}
func (this *LFUCache) update(e *ele, new bool) {
	var freqVal int
	var nextFe *freqEle
	if new { // 设置频率，尝试找到频率结点
		freqVal, nextFe = 1, this.head.next
	} else {
		freqVal, nextFe = e.fe.freq+1, e.fe.next
	}
	if nextFe == nil || nextFe.freq != freqVal { // nextFe.freq > freqVal
		nextFe = newFreqEle(freqVal)
		if new {
			nextFe.move(this.head) // 插入到 head
		} else {
			nextFe.move(e.fe) // 插入结点需要前缀结点 e.fe，所以“无法”复用 fe
		}
	}
	if e.fe != nil { // 删除结点
		e.remove()
		if e.fe.h.next == e.fe.t {
			e.fe.remove()
		}
	}
	e.fe = nextFe        // 更新频率
	e.move(nextFe.t.pre) // 插入到 tail
}

// ====================总 map + 总 链表 + 结点 map====================

type LFUCache struct {
	h, t  *ele
	items map[int]*ele
	freqs map[int]*ele
	cap   int
}

type ele struct {
	key, val  int
	pre, next *ele
	f         int
}

func (e *ele) remove() {
	e.pre.next, e.next.pre = e.next, e.pre
}
func (e *ele) move(at *ele) {
	at.next, at.next.pre, e.next, e.pre = e, e, at.next, at
}

func Constructor(capacity int) LFUCache {
	e := &ele{key: -1}
	e.pre, e.next = e, e
	return LFUCache{
		h:     e,
		t:     e,
		items: make(map[int]*ele, capacity),
		freqs: make(map[int]*ele, capacity),
		cap:   capacity,
	}
}
func (this *LFUCache) Get(key int) int {
	if e, ok := this.items[key]; ok {
		this.update(e, false)
		return e.val
	}
	return -1
}
func (this *LFUCache) Put(key int, value int) {
	e, ok := this.items[key]
	if ok {
		e.val = value
		this.update(e, false)
		return
	}
	if len(this.items) == this.cap {
		e = this.h.next
		delete(this.items, e.key) // key 已变，必须删除
		e.key, e.val = key, value
		if e.f == 1 { // 不用更新频率
			if t := this.freqs[1]; t != e { // 调整位置
				e.remove()
				e.move(t)
				this.freqs[1] = e // 注意：如果分别记录 head & tail，所以不能（或需要判断）复用对象
			}
			goto add
		}
	} else {
		e = &ele{key: key, val: value}
	}
	this.update(e, true)
add:
	this.items[key] = e
}
func (this *LFUCache) update(e *ele, new bool) {
	var (
		at *ele // 记录 e 要插入的位置（at.next = e）
		f  int
	)
	if e.f > 0 {
		at = this.freqs[e.f]  // 目标 freq 不存在时，用这里赋值的 at
		if at.f != at.pre.f { // 最后一个节点
			delete(this.freqs, at.f)
			at = at.pre
		} else if at == e { // 尾节点（即优先级最高）
			at = at.pre
			this.freqs[at.f] = at // 更新尾节点
		}
		e.remove()
	}
	if new {
		f = 1
	} else {
		f = e.f + 1
	}
	if t := this.freqs[f]; t == nil { // 补全 at 的赋值
		if f == 1 { // 新增结点，且 f == 1
			at = this.h
		}
	} else { // 目标 freq 存在
		at = t
	}
	e.f = f // 更新频率
	e.move(at)
	this.freqs[f] = e // 新频率的尾节点更新
}

// ====================list 包：二维链表====================

// LFUCache 频率越大，优先级越高；频率相等时，最近最久未使用，优先级越低
// 有多少个频率结点 *list.Element，就有多少个 freqVal
// 同频率的 item 的 freqItem 属性，指向的是同一个 *list.Element
// 每个 item 的 kvItem 属性，指向一个 *list.Element
type LFUCache struct {
	items    map[int]*item // 所有的 CacheItem 元素
	freqList *list.List    // 频率链表
	cap      int           // 最大容量
}

// item 每个 kvVal 对应一个 item
type item struct {
	freqItem *list.Element // 频率链表的结点
	kvItem   *list.Element // 相同频率元素的节点
}

// freqVal 频率链表中结点的 Value 值，可看做结点本身
// 每个 freqVal 的 freqItem 频率都相等
type freqVal struct {
	kvList *list.List // 频率相等的元素列表
	freq   int        // 频率
}

// kvVal k-v 数据
// 由于 list 包没有提供修改 Value 的接口，故保存 kvVal 的指针
type kvVal struct {
	key int
	val int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		items:    make(map[int]*item, capacity),
		freqList: list.New(),
		cap:      capacity,
	}
}
func (this *LFUCache) Get(key int) int {
	if it, ok := this.items[key]; ok {
		this.increment(it, it.kvItem.Value.(*kvVal))
		return it.kvItem.Value.(*kvVal).val
	}
	return -1
}
func (this *LFUCache) Put(key int, value int) {
	if it, ok := this.items[key]; ok {
		kv := it.kvItem.Value.(*kvVal)
		kv.val = value
		this.increment(it, kv)
	} else { // 新增元素
		it = new(item)
		if len(this.items) == this.cap { // 已满
			this.Remove(this.freqList.Front(), nil)
		}
		this.items[key] = it
		data := &kvVal{key, value}
		this.increment(it, data)
	}
}
func (this *LFUCache) increment(it *item, data *kvVal) {
	curFreqItem := it.freqItem // 当前结点
	var nextFreq int
	var nextFreqItem *list.Element

	if curFreqItem == nil { // 新增元素
		nextFreq = 1                         // 元素权重 =1
		nextFreqItem = this.freqList.Front() // 头部结点
	} else { // 更新
		nextFreq = curFreqItem.Value.(*freqVal).freq + 1 // 元素权重 +1
		nextFreqItem = curFreqItem.Next()                // 下一个结点
	}

	if nextFreqItem == nil || nextFreqItem.Value.(*freqVal).freq != nextFreq { // 频率不存在，新增 FrequencyItem 结点
		freqV := new(freqVal)
		freqV.freq = nextFreq
		freqV.kvList = list.New() // 结点的列表容器
		if curFreqItem == nil {
			nextFreqItem = this.freqList.PushFront(freqV) // 添加结点到头部
		} else {
			nextFreqItem = this.freqList.InsertAfter(freqV, curFreqItem) // 插入节点到 currentFrequency 后面
		}
	}

	it.freqItem = nextFreqItem // 所有元素都指向其所在的结点
	if curFreqItem != nil {
		this.Remove(curFreqItem, it.kvItem) // 从结点的列表中移除元素
	}
	//nextFrequency.Value.(*FrequencyItem).fl.PushFront(item.freqListItem)	// 错误使用
	it.kvItem = nextFreqItem.Value.(*freqVal).kvList.PushBack(data) // list 包的缺点
}
func (this *LFUCache) Remove(flItem *list.Element, kv *list.Element) {
	kl := flItem.Value.(*freqVal).kvList
	if kv == nil { // 淘汰一个元素
		delete(this.items, kl.Front().Value.(*kvVal).key)
		kl.Remove(kl.Front())
	} else { // 从 kvList 链表中删除一个元素
		kl.Remove(kv)
	}
	if kl.Len() == 0 { // 移除频率结点
		this.freqList.Remove(flItem)
	}
}

// ====================最初写法：“二维”计数器链表====================

type LFUCache struct {
	h, t *l         // 头 & 尾指针
	m    map[int]*l // 记录 k-v 元素
	lfu  map[int]*l // 记录计数器
	n    int        // 容量
}
type l struct {
	key, val  int // k-v 元素
	pre, next *l  // 维护元素的链表
	cnt       int // 计数器
	pl, nl    *l  // 维护计数器的链表
}

func Constructor(capacity int) LFUCache {
	p := &l{key: -1} // 0 <= key <= 10^5
	p.pre, p.next, p.pl, p.nl = p, p, p, p
	return LFUCache{
		h:   p,
		t:   p,
		m:   make(map[int]*l, capacity),
		lfu: make(map[int]*l, capacity),
		n:   capacity}
}
func (this *LFUCache) Get(key int) int {
	v, ok := this.m[key]
	if !ok { // key 不存在
		return -1
	}
	this.update(v, v.cnt+1) // key 存在，则更新 cnt
	return v.val
}
func (this *LFUCache) Put(key int, value int) {
	v, ok := this.m[key]
	if ok { // key 已存在
		v.val = value           // 更新 val
		this.update(v, v.cnt+1) // 更新 cnt
		return
	}
	if len(this.m) == this.n { // key 不存在，且 LFU 已满
		v = this.t.pre
		delete(this.m, v.key)     // this.m map 中删除结点
		v.key, v.val = key, value // 淘汰结点，并复用
		this.update(v, 1)
	} else { // key 不存在，且 LFU 未满
		v = &l{key: key, val: value, cnt: 1} // 新建结点
		this.push2LFU(v, nil)
	}
	this.m[key] = v
}
func (this *LFUCache) update(v *l, up int) {
	lv := this.lfu[v.cnt]
	if v == lv { // 计数相同的结点中，v 结点的优先级最低
		if v.pre.cnt == v.cnt { // 还有 v.cnt 的结点
			this.lfu[v.cnt] = v.pre                                         // lfu map 中更新 v.cnt 代表结点
			v.pre.pl, v.pre.nl, v.pl.nl, v.nl.pl = v.pl, v.nl, v.pre, v.pre // 计数链表中 v.pre 取代 v
		} else { // 最后一个 v.cnt 的结点
			delete(this.lfu, v.cnt)       // lfu map 中删除结点
			v.pl.nl, v.nl.pl = v.nl, v.pl // 计数链表中删除 v.cnt 结点
		}
	}
	delEle(v)
	v.cnt = up              // 更新计数
	this.push2LFU(v, lv.pl) // 更新计数器算法
}
func (this *LFUCache) push2LFU(v, p *l) { // p 为计数链表中 v.pl 结点
	if plv, ok := this.lfu[v.cnt]; ok { // cnt 存在
		insertEle(plv.pl, v)
	} else { // cnt 不存在
		if v.cnt == 1 { // 新插入节点
			p = this.t.pl
		}
		insertEle(p, v)
		this.lfu[v.cnt] = v                       // lfu map 中插入结点
		p.nl, p.nl.pl, v.pl, v.nl = v, v, p, p.nl // 计数链表中插入结点
	}
}
func delEle(v *l) { // 删除结点
	v.pre.next, v.next.pre = v.next, v.pre
}
func insertEle(h, v *l) { // 插入结点
	h.next, h.next.pre, v.pre, v.next = v, v, h, h.next
}

// ====================频率相等的元素的优先级是相同的，被淘汰的概率是随机的====================

// LFUCache 频率相等的元素的优先级是相同的，被淘汰的概率是随机的
// 有多少个结点 *list.Element，就有多少个 FrequencyItem
// 同频率的 CacheItem 的属性 frequencyParent，指向的是同一个 *list.Element
// 每个 *list.Element 的属性 Value，指向一个 FrequencyParent
type LFUCache struct {
	// 所有的 CacheItem 元素
	bykey map[int]*CacheItem // Hashmap containing *CacheItems for O(1) access
	// 存储 *list.Element 的链表
	freqs    *list.List // Linked list of frequencies
	capacity int        // Max number of items
	size     int        // Current size of cache
}

// CacheItem 存储 k-v 数据的元素
type CacheItem struct {
	key   int // Key of entry
	value int // Value of item
	// 指向元素所在的结点 *list.Element，即 FrequencyItem
	frequencyParent *list.Element // Pointer to parent in cacheList
}

// FrequencyItem 链表中结点元素的 Value，可看做结点本身
// 每个 FrequencyItem 的 CacheItem 频率都相等
type FrequencyItem struct {
	// 频率列表
	entries map[*CacheItem]byte // Set of entries
	// 频率
	freq int // Access frequency
}

func Constructor(capacity int) LFUCache {
	cache := LFUCache{}
	cache.bykey = make(map[int]*CacheItem)
	cache.freqs = list.New()
	cache.size = 0
	cache.capacity = capacity
	return cache
}
func (this *LFUCache) Get(key int) int {
	if e, ok := this.bykey[key]; ok {
		this.increment(e)
		return e.value
	}
	return -1
}
func (this *LFUCache) Put(key int, value int) {
	if item, ok := this.bykey[key]; ok {
		item.value = value
		this.increment(item)
	} else { // 新增元素
		item := new(CacheItem)
		item.key = key
		item.value = value
		this.bykey[key] = item
		this.size++
		if this.atCapacity() { // 已满
			//this.Evict(10) // 淘汰 10 个元素
			this.Evict(2)
		}
		this.increment(item)
	}
}
func (this *LFUCache) atCapacity() bool {
	//return this.size >= this.capacity
	return this.size > this.capacity
}
func (this *LFUCache) Evict(count int) {
	for i := 0; i < count; { // 删除 count 个元素
		if item := this.freqs.Front(); item != nil { // 从头部结点的列表开始删除
			for entry, _ := range item.Value.(*FrequencyItem).entries {
				if i < count {
					delete(this.bykey, entry.key) // 随机删除
					this.Remove(item, entry)      // 如果列表为空，则删除列表
					this.size--
					i++
				}
			}
		}
	}
}
func (this *LFUCache) increment(item *CacheItem) {
	currentFrequency := item.frequencyParent // 当前结点
	var nextFrequencyAmount int
	var nextFrequency *list.Element

	if currentFrequency == nil { // 新增元素
		nextFrequencyAmount = 1            // 元素权重 =1
		nextFrequency = this.freqs.Front() // 头部结点
	} else { // 更新
		nextFrequencyAmount = currentFrequency.Value.(*FrequencyItem).freq + 1 // 元素权重 +1
		nextFrequency = currentFrequency.Next()                                // 下一个结点
	}

	if nextFrequency == nil || nextFrequency.Value.(*FrequencyItem).freq != nextFrequencyAmount { // 频率不存在，新增 FrequencyItem 结点
		newFrequencyItem := new(FrequencyItem)
		newFrequencyItem.freq = nextFrequencyAmount
		newFrequencyItem.entries = make(map[*CacheItem]byte) // 结点的列表容器
		if currentFrequency == nil {
			nextFrequency = this.freqs.PushFront(newFrequencyItem) // 添加结点到头部
		} else {
			nextFrequency = this.freqs.InsertAfter(newFrequencyItem, currentFrequency) // 插入节点到 currentFrequency 后面
		}
	}

	item.frequencyParent = nextFrequency                   // 所有元素都指向其所在的结点
	nextFrequency.Value.(*FrequencyItem).entries[item] = 1 // ？
	if currentFrequency != nil {
		this.Remove(currentFrequency, item) // 从结点的列表中移除元素
	}
}
func (this *LFUCache) Remove(listItem *list.Element, item *CacheItem) {
	frequencyItem := listItem.Value.(*FrequencyItem)
	delete(frequencyItem.entries, item)
	if len(frequencyItem.entries) == 0 { // 列表已空
		this.freqs.Remove(listItem) // 移除结点
	}
}
