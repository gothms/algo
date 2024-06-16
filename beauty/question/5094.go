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

// ====================循环链表====================

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
//	if as.cnt[0] == 0 || as.cnt[1] > 0 && h.next.idx > as.ht[1][0].next.idx {	// 重点
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

// ====================循环数组====================

type AnimalShelf struct {
	q   [2][][2]int // dog / cat : animals : id / c
	ids [2][2]int   // dog / cat : h / t
	k   int         // cap+1
	c   int
}

func NewAnimalShelf(shelfSize int) *AnimalShelf {
	shelfSize++
	q := [2][][2]int{make([][2]int, shelfSize), make([][2]int, shelfSize)}
	return &AnimalShelf{q: q, k: shelfSize}
}

func (as *AnimalShelf) EnqueueDog(dogID int) bool {
	return as.Enqueue(dogID, 0)
}

func (as *AnimalShelf) EnqueueCat(catID int) bool {
	return as.Enqueue(catID, 1)
}

func (as *AnimalShelf) Enqueue(id, i int) bool {
	is := as.ids[i]
	if ot := as.ids[i^1]; (is[1]-is[0]+ot[1]-ot[0]+as.k<<1)%as.k == as.k-1 { // 重点
		return false // 会错题意的代价！
	}
	as.c++
	as.q[i][is[1]] = [2]int{id, as.c} // append
	as.ids[i][1] = (is[1] + 1) % as.k
	return true
}

func (as *AnimalShelf) DequeueAny() int {
	isDog, isCat := as.ids[0], as.ids[1]
	if isDog[0] == isDog[1] && isCat[0] == isCat[1] { // empty
		return -1
	}
	i, is := 0, isDog
	if isDog[0] == isDog[1] ||
		isCat[0] != isCat[1] && as.q[0][isDog[0]][1] > as.q[1][isCat[0]][1] {
		i, is = 1, isCat
	}
	v := as.q[i][is[0]][0]
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
	v := as.q[i][is[0]][0]
	as.ids[i][0] = (is[0] + 1) % as.k
	return v
}
