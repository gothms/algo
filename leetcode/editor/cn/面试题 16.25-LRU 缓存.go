//设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存
//被填满时，它应该删除最近最少使用的项目。
//
// 它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
// 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。 写入数据 put(key, value)
// - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
//
// 示例:
//
// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
//
//
// Related Topics 设计 哈希表 链表 双向链表 👍 184 👎 0

package main

func main() {
	//hash := make(map[int]int, 5)
	//hash[1] = 1
	//fmt.Println(len(hash))

	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)    // 返回 1
	cache.Put(3, 3) // 该操作会使得密钥 2 作废
	cache.Get(2)    // 返回 -1 (未找到)
	cache.Put(4, 4) // 该操作会使得密钥 1 作废
	cache.Get(1)    // 返回 -1 (未找到)
	cache.Get(3)    // 返回 3
	cache.Get(4)    // 返回 4
}

/*
思路：双向链表 + hash表
	1.hash表实现 O(1) 的查询，双向链表实现 O(1) 的删除、添加、更新
	2.Get 逻辑
		2.1.如果hash表中存在 key，则返回对应的 val，并将这个结点插入到头部
		2.2.如果不存在，则返回 -1
	3.Put 逻辑
		3.1.如果hash表中存在 key，则将 val 更新，并将这个结点插入到头部
		3.2.如果hash表中不存在，新建一个结点，并将这个结点插入到头部
			如果hash表中的结点总数超过预定 cap，则将尾结点从链表中删除，并从hash表中删除这个尾结点
*/

// leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	hash       map[int]*DoublyLinkedList // hash表
	head, tail *DoublyLinkedList         // 头、尾节点
	cap        int                       // 预定 cap 大小
}
type DoublyLinkedList struct { // 双向链表
	pre, next *DoublyLinkedList
	key, val  int
}

func Constructor(capacity int) LRUCache {
	h, t := &DoublyLinkedList{}, &DoublyLinkedList{}
	h.next, t.pre = t, h // 初始化头、尾节点
	return LRUCache{
		hash: make(map[int]*DoublyLinkedList, capacity),
		cap:  capacity,
		head: h,
		tail: t,
	}
}
func (this *LRUCache) Get(key int) int {
	if dll, ok := this.hash[key]; ok {
		this.insertHead(dll) // hash表中存在 key，将结点插入头部
		return dll.val
	}
	return -1
}
func (this *LRUCache) Put(key int, value int) {
	if dll, ok := this.hash[key]; ok { // 更新结点
		dll.val = value
		this.insertHead(dll) // hash表中存在 key，将结点插入头部
	} else { // 新增结点
		dll = &DoublyLinkedList{key: key, val: value}
		if len(this.hash) == this.cap { // 超过 cap，删除尾结点
			this.deleteTail()
		}
		this.hash[key] = dll // 将新增结点添加到hash表
		this.addHead(dll)    // 将新增结点添加到头部
	}
}
func (this *LRUCache) deleteTail() {
	del := this.tail.pre
	delete(this.hash, del.key)
	del.pre.next, del.next.pre, del.pre, del.next = del.next, del.pre, nil, nil
}
func (this *LRUCache) insertHead(dll *DoublyLinkedList) {
	if dll.pre == this.head {
		return
	}
	dll.pre.next, dll.next.pre = dll.next, dll.pre
	this.addHead(dll)
}
func (this *LRUCache) addHead(dll *DoublyLinkedList) {
	this.head.next, this.head.next.pre, dll.pre, dll.next = dll, dll, this.head, this.head.next
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
