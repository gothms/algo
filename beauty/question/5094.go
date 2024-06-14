package main

import "fmt"

func main() {
	var ret int
	as := NewAnimalShelf(1)
	as.EnqueueDog(1)
	as.EnqueueCat(2)
	as.EnqueueDog(3)
	ret = as.DequeueAny()
	fmt.Println(ret) // 1
	ret = as.DequeueCat()
	fmt.Println(ret) // 2
	as.EnqueueCat(4)
	ret = as.DequeueAny()
	fmt.Println(ret) // 4
	ret = as.DequeueDog()
	fmt.Println(ret) // -1
	as.EnqueueCat(5)
	ret = as.DequeueAny()
	fmt.Println(ret) // 5
}

//type ListNode struct {
//	val  int
//	idx  int
//	next *ListNode
//}
//type AnimalShelf struct {
//	ht  [2][2]*ListNode // dh, dt, ch, ct *ListNode
//	cnt [2]int
//	k   int
//	id  int
//}
//
//func NewAnimalShelf(shelfSize int) *AnimalShelf {
//	d, c := &ListNode{}, &ListNode{}
//	d.next, c.next = d, c
//	return &AnimalShelf{ht: [2][2]*ListNode{{d, d}, {c, c}}, k: shelfSize}
//}
//
//func (as *AnimalShelf) EnqueueDog(dogID int) bool {
//	return as.Enqueue(dogID, 0)
//}
//
//func (as *AnimalShelf) EnqueueCat(catID int) bool {
//	return as.Enqueue(catID, 1)
//}
//
//func (as *AnimalShelf) Enqueue(id, i int) bool {
//	if as.cnt[0]+as.cnt[1] == as.k {
//		return false
//	}
//	as.id++
//	cur := &ListNode{val: id, idx: as.id}
//	h, t := as.ht[i][0], as.ht[i][1]
//	t.next, cur.next = cur, h // 插入结点
//	as.ht[i][1] = cur
//	as.cnt[i]++
//	return true
//}
//
//func (as *AnimalShelf) DequeueAny() int {
//	if as.cnt[0] == 0 && as.cnt[1] == 0 {
//		return -1
//	}
//	i, h := 0, as.ht[0][0]
//	if as.cnt[0] == 0 || h.next.idx > as.ht[1][0].next.idx {
//		i, h = 1, as.ht[1][0]
//	}
//	return as.delete(h, i)
//}
//
//func (as *AnimalShelf) DequeueDog() int {
//	return as.Dequeue(0)
//}
//
//func (as *AnimalShelf) DequeueCat() int {
//	return as.Dequeue(1)
//}
//
//func (as *AnimalShelf) Dequeue(i int) int {
//	if as.cnt[i] == 0 {
//		return -1
//	}
//	return as.delete(as.ht[i][0], i)
//}
//
//func (as *AnimalShelf) delete(h *ListNode, i int) int {
//	v := h.next.val
//	h.next, h.next.next = h.next.next, nil // 删除结点
//	as.cnt[i]--
//	if as.cnt[i] == 0 { // 修正 tail
//		as.ht[i][1] = h
//	}
//	return v
//}

type AnimalShelf struct {
	q   [][2]int
	ids [2][2]int
	k   int
	c   int
}

func NewAnimalShelf(shelfSize int) *AnimalShelf {
	shelfSize++
	return &AnimalShelf{q: make([][2]int, shelfSize<<1), k: shelfSize}
}

func (as *AnimalShelf) EnqueueDog(dogID int) bool {
	return as.Enqueue(dogID, 0)
}

func (as *AnimalShelf) EnqueueCat(catID int) bool {
	return as.Enqueue(catID, 1)
}

func (as *AnimalShelf) Enqueue(id, i int) bool {
	is := as.ids[i]
	if (is[1]+1)%as.k == is[0] { // full
		return false
	}
	as.c++
	as.q[is[1]<<1+i] = [2]int{id, as.c}
	as.ids[i][1] = (is[1] + 1) % as.k
	return true
}

func (as *AnimalShelf) DequeueAny() int {
	isDog, isCat := as.ids[0], as.ids[1]
	if isDog[0] == isDog[1] && isCat[0] == isCat[1] { // empty
		return -1
	}
	i, is := 0, isDog
	if isDog[0] == isDog[1] || as.q[isDog[0]<<1][1] > as.q[isCat[0]<<1+1][1] {
		i, is = 1, isCat
	}
	v := as.q[is[0]<<1+i][0]
	as.ids[i][0] = (is[0] + 1) % as.k
	return v
}

func (as *AnimalShelf) DequeueDog() int {
	return as.Dequeue(0)
}

func (as *AnimalShelf) DequeueCat() int {
	return as.Dequeue(1)
}

func (as *AnimalShelf) Dequeue(i int) int {
	is := as.ids[i]
	if is[0] == is[1] { // empty
		return -1
	}
	v := as.q[is[0]<<1+i][0]
	as.ids[i][0] = (is[0] + 1) % as.k
	return v
}

//type ListNode struct {
//	val       int
//	pre, next *ListNode
//}
//
//type AnimalShelf struct {
//	q    []*ListNode
//	ids  [2][2]int
//	k    int
//	h, t *ListNode
//}
//
//func NewAnimalShelf(shelfSize int) *AnimalShelf {
//	l := &ListNode{}
//	l.pre, l.next = l, l
//	shelfSize++
//	return &AnimalShelf{make([]*ListNode, shelfSize<<1), [2][2]int{}, shelfSize, l, l}
//}
//
//func (as *AnimalShelf) EnqueueDog(dogID int) bool {
//	return as.Enqueue(dogID, 0)
//}
//
//func (as *AnimalShelf) EnqueueCat(catID int) bool {
//	return as.Enqueue(catID, 1)
//}
//
//func (as *AnimalShelf) Enqueue(id, i int) bool {
//	is := as.ids[i]
//	if (is[1]+1)%as.k == is[0] { // full
//		return false
//	}
//	cur := &ListNode{val: id<<1 | i}
//	as.t.pre, as.t.pre.next, cur.pre, cur.next = cur, cur, as.t.pre, as.t // 插入结点
//	as.q[is[1]<<1+i] = cur
//	as.ids[i][1] = (is[1] + 1) % as.k
//	return true
//}
//
//func (as *AnimalShelf) DequeueAny() int {
//	if as.h.next == as.t { // empty
//		return -1
//	}
//	cur := as.h.next
//	as.h.next, cur.next.pre, cur.next, cur.pre = cur.next, as.h, nil, nil // 删除结点
//	i := cur.val & 1                                                      // dog / cat
//	as.ids[i][0] = (as.ids[i][0] + 1) % as.k
//	return cur.val >> 1
//}
//
//func (as *AnimalShelf) DequeueDog() int {
//	return as.Dequeue(0)
//}
//
//func (as *AnimalShelf) DequeueCat() int {
//	return as.Dequeue(1)
//}
//
//func (as *AnimalShelf) Dequeue(i int) int {
//	is := as.ids[i]
//	if is[0] == is[1] { // empty
//		return -1
//	}
//	cur := as.q[is[0]<<1+i]
//	cur.pre.next, cur.next.pre, cur.pre, cur.next = cur.next, cur.pre, nil, nil // 删除节点
//	as.ids[i][0] = (is[0] + 1) % as.k
//	return cur.val >> 1
//}
