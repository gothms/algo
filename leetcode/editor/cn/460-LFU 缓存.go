//请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。
//
// 实现 LFUCache 类：
//
//
// LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
// int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
// void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量
//capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
//
//
// 为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。
//
// 当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
//
//
// 示例：
//
//
//输入：
//["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get",
//"get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
//输出：
//[null, null, null, 1, null, -1, 3, null, -1, 3, 4]
//
//解释：
//// cnt(x) = 键 x 的使用计数
//// cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
//LFUCache lfu = new LFUCache(2);
//lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
//lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
//lfu.get(1);      // 返回 1
//                 // cache=[1,2], cnt(2)=1, cnt(1)=2
//lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
//                 // cache=[3,1], cnt(3)=1, cnt(1)=2
//lfu.get(2);      // 返回 -1（未找到）
//lfu.get(3);      // 返回 3
//                 // cache=[3,1], cnt(3)=2, cnt(1)=2
//lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
//                 // cache=[4,3], cnt(4)=1, cnt(3)=2
//lfu.get(1);      // 返回 -1（未找到）
//lfu.get(3);      // 返回 3
//                 // cache=[3,4], cnt(4)=1, cnt(3)=3
//lfu.get(4);      // 返回 4
//                 // cache=[3,4], cnt(4)=2, cnt(3)=3
//
//
//
// 提示：
//
//
// 1 <= capacity <= 10⁴
// 0 <= key <= 10⁵
// 0 <= value <= 10⁹
// 最多调用 2 * 10⁵ 次 get 和 put 方法
//
//
// Related Topics 设计 哈希表 链表 双向链表 👍 689 👎 0

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// map test
	//m := map[int]*ele{}
	//e1 := &ele{key: 1, val: 2}
	//e2 := &ele{key: 3, val: 4}
	//m[1] = e1
	//m[2] = e2
	//fmt.Println(m[2])
	//e2.val = 99
	//fmt.Println(m[2])

	lfu := Constructor(2)
	lfu.Put(1, 11)
	lfu.Put(2, 22)
	fmt.Println(len(lfu.items))
	for l := lfu.head.next; l != lfu.tail; l = l.next {
		fmt.Println(l, ", ", l.freq)
		for k := l.h.next; k != l.t; k = k.next {
			fmt.Print(k.key, k.val, ", ")
		}
		fmt.Println()
	}
	v := lfu.Get(1)
	fmt.Println(v)
	//fmt.Println(len(lfu.items))
	//fmt.Println(lfu.items)
	//for l := lfu.head.next; l != lfu.tail; l = l.next {
	//	fmt.Println(l, ", ", l.freq)
	//	for k := l.h.next; k != l.t; k = k.next {
	//		fmt.Print(k.key, k.val, ", ")
	//	}
	//	fmt.Println()
	//}

	//fmt.Println(lfu.freqList.Len())

	//for l := lfu.freqList.Front(); l != nil; l = l.Next() {
	//	frequencyItem := l.Value.(*freqVal)
	//	//fmt.Println("frequencyItem:", frequencyItem.freq, frequencyItem)
	//	fmt.Printf("frequencyItem: %d, %d, %p\n", frequencyItem.freq, frequencyItem.kvList.Len(), l)
	//}
	//fmt.Println("==========================")
	//for k, it := range lfu.items {
	//	fmt.Println(k, it)
	//	fmt.Printf("%p\n", it.freqItem.Value)
	//	//if item != nil {
	//	//	parent := item.frequencyParent
	//	//	fmt.Println(k, parent.Value)
	//	//}
	//}
	//fmt.Println("==========================")

	lfu.Put(3, 33)
	//fmt.Println(lfu.items)
	v = lfu.Get(2)
	fmt.Println(v)
	v = lfu.Get(3)
	fmt.Println(v)
	lfu.Put(4, 44)
	//rangeLFU(lfu)

	v = lfu.Get(1)
	fmt.Println(v)
	v = lfu.Get(3)
	fmt.Println(v)
	v = lfu.Get(4)
	fmt.Println(v)
}

//func rangeLFU(lfu LFUCache) {
//	for k, v := range lfu.m {
//		fmt.Print(k, v, ", ")
//	}
//	fmt.Println()
//	for k, v := range lfu.lfu {
//		fmt.Print(k, v, ", ")
//	}
//	fmt.Println()
//}

/*
LFU 论文
	http://dhruvbird.com/lfu.pdf
golang 实现
	https://cloud.tencent.com/developer/article/2306436

list 包的缺点
	1.不能修改 Value，除非用指针
	2.不同链表直接的结点不能互相操作，需要解析出 Value 再添加 Value
*/

// leetcode submit region begin(Prohibit modification and deletion)

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

//leetcode submit region end(Prohibit modification and deletion)

// ====================list 包====================

// LFUCache 频率越大，优先级越高；频率相等时，最近最久未使用，优先级越低
// 有多少个频率结点 *list.Element，就有多少个 freqVal
// 同频率的 item 的 freqItem 属性，指向的是同一个 *list.Element
// 每个 item 的 kvItem 属性，指向一个 *list.Element
//type LFUCache struct {
//	items    map[int]*item // 所有的 CacheItem 元素
//	freqList *list.List    // 频率链表
//	cap      int           // 最大容量
//}
//
//// item 每个 kvVal 对应一个 item
//type item struct {
//	freqItem *list.Element // 频率链表的结点
//	kvItem   *list.Element // 相同频率元素的节点
//}
//
//// freqVal 频率链表中结点的 Value 值，可看做结点本身
//// 每个 freqVal 的 freqItem 频率都相等
//type freqVal struct {
//	kvList *list.List // 频率相等的元素列表
//	freq   int        // 频率
//}
//
//// kvVal k-v 数据
//// 由于 list 包没有提供修改 Value 的接口，故保存 kvVal 的指针
//type kvVal struct {
//	key int
//	val int
//}
//
//func Constructor(capacity int) LFUCache {
//	return LFUCache{
//		items:    make(map[int]*item, capacity),
//		freqList: list.New(),
//		cap:      capacity,
//	}
//}
//func (this *LFUCache) Get(key int) int {
//	if it, ok := this.items[key]; ok {
//		this.increment(it, it.kvItem.Value.(*kvVal))
//		return it.kvItem.Value.(*kvVal).val
//	}
//	return -1
//}
//func (this *LFUCache) Put(key int, value int) {
//	if it, ok := this.items[key]; ok {
//		kv := it.kvItem.Value.(*kvVal)
//		kv.val = value
//		this.increment(it, kv)
//	} else { // 新增元素
//		it = new(item)
//		if len(this.items) == this.cap { // 已满
//			this.Remove(this.freqList.Front(), nil)
//		}
//		this.items[key] = it
//		data := &kvVal{key, value}
//		this.increment(it, data)
//	}
//}
//func (this *LFUCache) increment(it *item, data *kvVal) {
//	curFreqItem := it.freqItem // 当前结点
//	var nextFreq int
//	var nextFreqItem *list.Element
//
//	if curFreqItem == nil { // 新增元素
//		nextFreq = 1                         // 元素权重 =1
//		nextFreqItem = this.freqList.Front() // 头部结点
//	} else { // 更新
//		nextFreq = curFreqItem.Value.(*freqVal).freq + 1 // 元素权重 +1
//		nextFreqItem = curFreqItem.Next()                // 下一个结点
//	}
//
//	if nextFreqItem == nil || nextFreqItem.Value.(*freqVal).freq != nextFreq { // 频率不存在，新增 FrequencyItem 结点
//		freqV := new(freqVal)
//		freqV.freq = nextFreq
//		freqV.kvList = list.New() // 结点的列表容器
//		if curFreqItem == nil {
//			nextFreqItem = this.freqList.PushFront(freqV) // 添加结点到头部
//		} else {
//			nextFreqItem = this.freqList.InsertAfter(freqV, curFreqItem) // 插入节点到 currentFrequency 后面
//		}
//	}
//
//	it.freqItem = nextFreqItem // 所有元素都指向其所在的结点
//	if curFreqItem != nil {
//		this.Remove(curFreqItem, it.kvItem) // 从结点的列表中移除元素
//	}
//	//nextFrequency.Value.(*FrequencyItem).fl.PushFront(item.freqListItem)	// 错误使用
//	it.kvItem = nextFreqItem.Value.(*freqVal).kvList.PushBack(data) // list 包的缺点
//}
//func (this *LFUCache) Remove(flItem *list.Element, kv *list.Element) {
//	kl := flItem.Value.(*freqVal).kvList
//	if kv == nil { // 淘汰一个元素
//		delete(this.items, kl.Front().Value.(*kvVal).key)
//		kl.Remove(kl.Front())
//	} else { // 从 kvList 链表中删除一个元素
//		kl.Remove(kv)
//	}
//	if kl.Len() == 0 { // 移除频率结点
//		this.freqList.Remove(flItem)
//	}
//}

// ====================个人写法====================

// LFUCache 个人写法
//type LFUCache struct {
//	h, t *l         // 头 & 尾指针
//	m    map[int]*l // 记录 k-v 元素
//	lfu  map[int]*l // 记录计数器
//	n    int        // 容量
//}
//type l struct {
//	key, val  int // k-v 元素
//	pre, next *l  // 维护元素的链表
//	cnt       int // 计数器
//	pl, nl    *l  // 维护计数器的链表
//}
//
//func Constructor(capacity int) LFUCache {
//	p := &l{key: -1} // 0 <= key <= 10^5
//	p.pre, p.next, p.pl, p.nl = p, p, p, p
//	return LFUCache{
//		h:   p,
//		t:   p,
//		m:   make(map[int]*l, capacity),
//		lfu: make(map[int]*l, capacity),
//		n:   capacity}
//}
//func (this *LFUCache) Get(key int) int {
//	v, ok := this.m[key]
//	if !ok { // key 不存在
//		return -1
//	}
//	this.update(v, v.cnt+1) // key 存在，则更新 cnt
//	return v.val
//}
//func (this *LFUCache) Put(key int, value int) {
//	v, ok := this.m[key]
//	if ok { // key 已存在
//		v.val = value           // 更新 val
//		this.update(v, v.cnt+1) // 更新 cnt
//		return
//	}
//	if len(this.m) == this.n { // key 不存在，且 LFU 已满
//		v = this.t.pre
//		delete(this.m, v.key)     // this.m map 中删除结点
//		v.key, v.val = key, value // 淘汰结点，并复用
//		this.update(v, 1)
//	} else { // key 不存在，且 LFU 未满
//		v = &l{key: key, val: value, cnt: 1} // 新建结点
//		this.push2LFU(v, nil)
//	}
//	this.m[key] = v
//}
//func (this *LFUCache) update(v *l, up int) {
//	lv := this.lfu[v.cnt]
//	if v == lv { // 计数相同的结点中，v 结点的优先级最低
//		if v.pre.cnt == v.cnt { // 还有 v.cnt 的结点
//			this.lfu[v.cnt] = v.pre                                         // lfu map 中更新 v.cnt 代表结点
//			v.pre.pl, v.pre.nl, v.pl.nl, v.nl.pl = v.pl, v.nl, v.pre, v.pre // 计数链表中 v.pre 取代 v
//		} else { // 最后一个 v.cnt 的结点
//			delete(this.lfu, v.cnt)       // lfu map 中删除结点
//			v.pl.nl, v.nl.pl = v.nl, v.pl // 计数链表中删除 v.cnt 结点
//		}
//	}
//	delEle(v)
//	v.cnt = up              // 更新计数
//	this.push2LFU(v, lv.pl) // 更新计数器算法
//}
//func (this *LFUCache) push2LFU(v, p *l) { // p 为计数链表中 v.pl 结点
//	if plv, ok := this.lfu[v.cnt]; ok { // cnt 存在
//		insertEle(plv.pl, v)
//	} else { // cnt 不存在
//		if v.cnt == 1 { // 新插入节点
//			p = this.t.pl
//		}
//		insertEle(p, v)
//		this.lfu[v.cnt] = v                       // lfu map 中插入结点
//		p.nl, p.nl.pl, v.pl, v.nl = v, v, p, p.nl // 计数链表中插入结点
//	}
//}
//func delEle(v *l) { // 删除结点
//	v.pre.next, v.next.pre = v.next, v.pre
//}
//func insertEle(h, v *l) { // 插入结点
//	h.next, h.next.pre, v.pre, v.next = v, v, h, h.next
//}

// ====================lc====================

//type entry struct {
//	key, value, freq int // freq 表示这本书被看的次数
//}
//
//type LFUCache struct {
//	capacity   int
//	minFreq    int
//	keyToNode  map[int]*list.Element
//	freqToList map[int]*list.List
//}
//
//func Constructor(capacity int) LFUCache {
//	return LFUCache{
//		capacity:   capacity,
//		keyToNode:  map[int]*list.Element{},
//		freqToList: map[int]*list.List{},
//	}
//}
//
//func (c *LFUCache) pushFront(e *entry) {
//	if _, ok := c.freqToList[e.freq]; !ok {
//		c.freqToList[e.freq] = list.New() // 双向链表
//	}
//	c.keyToNode[e.key] = c.freqToList[e.freq].PushFront(e)
//}
//
//func (c *LFUCache) getEntry(key int) *entry {
//	node := c.keyToNode[key]
//	if node == nil { // 没有这本书
//		return nil
//	}
//	e := node.Value.(*entry)
//	lst := c.freqToList[e.freq]
//	lst.Remove(node)    // 把这本书抽出来
//	if lst.Len() == 0 { // 抽出来后，这摞书是空的
//		delete(c.freqToList, e.freq) // 移除空链表
//		if c.minFreq == e.freq {     // 这摞书是最左边的
//			c.minFreq++
//		}
//	}
//	e.freq++       // 看书次数 +1
//	c.pushFront(e) // 放在右边这摞书的最上面
//	return e
//}
//
//func (c *LFUCache) Get(key int) int {
//	if e := c.getEntry(key); e != nil { // 有这本书
//		return e.value
//	}
//	return -1 // 没有这本书
//}
//
//func (c *LFUCache) Put(key, value int) {
//	if e := c.getEntry(key); e != nil { // 有这本书
//		e.value = value // 更新 value
//		return
//	}
//	if len(c.keyToNode) == c.capacity { // 书太多了
//		lst := c.freqToList[c.minFreq]                           // 最左边那摞书
//		delete(c.keyToNode, lst.Remove(lst.Back()).(*entry).key) // 移除这摞书的最下面的书
//		if lst.Len() == 0 {                                      // 这摞书是空的
//			delete(c.freqToList, c.minFreq) // 移除空链表
//		}
//	}
//	c.pushFront(&entry{key, value, 1}) // 新书放在「看过 1 次」的最上面
//	c.minFreq = 1
//}
