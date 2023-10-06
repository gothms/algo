package main

import (
	"container/heap"
	"fmt"
)

func main() {
	vm := Constructor()
	vm.AddItem(0, 1, "A", 4, 3)
	vm.AddItem(1, 3, "A", 4, 2)
	vm.Sell(2, "M", "A", 2)
	for k, hp := range vm.m {
		fmt.Println(hp, k)
	}
	for k, hp := range vm.tm {
		for _, item := range hp.h {
			fmt.Println(*item, k, hp.cnt)
		}
	}
	vm.AddItem(2, 1, "B", 2, 5)
	vm.Sell(4, "J", "B", 2)
	vm.Sell(4, "M", "B", 1)
	vm.Sell(4, "M", "A", 1)
	vm.AddItem(6, 200, "A", 2, 5)
	vm.Sell(6, "J", "A", 100)
	for _, hp := range vm.m {
		fmt.Println(hp)
	}
	for k, hp := range vm.tm {
		for _, item := range hp.h {
			fmt.Println(*item, k, hp.cnt)
		}
	}
	sell := vm.Sell(7, "M", "A", 100)
	fmt.Println(sell)
}

/*
细节
	1.优先队列作为 k-v 的 val 时，使用指针，否则修改的是副本
	2.优先队列中的元素在多处被使用时，最好也使用指针，则修改优先队列时也修改了原值
*/

type VendingMachine struct {
	m  map[string]*vmHp
	tm map[string]*vmPHp
	cm map[string]int
}
type vmItem struct {
	num      int
	price    int
	duration int
}
type vmHp []*vmItem

func (v vmHp) Len() int { return len(v) }
func (v vmHp) Less(i, j int) bool {
	return v[i].price < v[j].price || v[i].price == v[j].price && v[i].duration < v[j].duration
}
func (v vmHp) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v *vmHp) Push(x any)   { *v = append(*v, x.(*vmItem)) }
func (v *vmHp) Pop() any {
	top := (*v)[len(*v)-1]
	*v = (*v)[:len(*v)-1]
	return top
}

type vmPHp struct {
	h   []*vmItem
	cnt int
}

func (v vmPHp) Len() int { return len(v.h) }
func (v vmPHp) Less(i, j int) bool {
	return (v.h)[i].duration < (v.h)[j].duration
}
func (v vmPHp) Swap(i, j int) { (v.h)[i], (v.h)[j] = (v.h)[j], (v.h)[i] }
func (v *vmPHp) Push(x any)   { v.h = append(v.h, x.(*vmItem)) }
func (v *vmPHp) Pop() any {
	top := (v.h)[len(v.h)-1]
	v.h = (v.h)[:len(v.h)-1]
	return top
}

func Constructor() VendingMachine {
	return VendingMachine{
		m:  make(map[string]*vmHp),
		tm: make(map[string]*vmPHp),
		cm: make(map[string]int),
	}
}
func (this *VendingMachine) AddItem(time int, number int, item string, price int, duration int) {
	vm := vmItem{number, price, time + duration}
	vHp, vPHp := this.m[item], this.tm[item]
	if vHp == nil {
		vHp = &vmHp{}
		this.m[item] = vHp
		vPHp = &vmPHp{}
		this.tm[item] = vPHp
	}
	heap.Push(vHp, &vm)
	heap.Push(vPHp, &vm)
	vPHp.cnt += number
	for vPHp.Len() > 0 && (vPHp.h)[0].duration < time { // 过期
		top := heap.Pop(vPHp).(*vmItem)
		vPHp.cnt -= top.num
		top.num = 0
	}
}
func (this *VendingMachine) Sell(time int, customer string, item string, number int) int64 {
	vHp, vPHp := this.m[item], this.tm[item]
	if vHp == nil {
		return -1
	}
	for vPHp.Len() > 0 && (vPHp.h)[0].duration < time { // 过期
		top := heap.Pop(vPHp).(*vmItem)
		vPHp.cnt -= top.num
		top.num = 0
	}
	if vPHp.cnt < number {
		return -1
	}
	curPrice := this.cm[customer]
	if curPrice == 0 {
		curPrice = 100
	}
	vPHp.cnt -= number // 修正总数量
	v := 0
	for number > 0 {
		top := (*vHp)[0] // 修改指针
		number -= top.num
		if number < 0 {
			v += (top.num + number) * top.price
			top.num = -number
			heap.Fix(vHp, 0)
		} else {
			v += top.num * top.price
			top.num = 0
			heap.Pop(vHp)
		}
	}
	this.cm[customer] = curPrice - 1
	return int64(v*curPrice+99) / 100 // 向上取整的技巧
}
