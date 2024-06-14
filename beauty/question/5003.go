package main

type MyArray struct {
	arr []int
	k   int
}

func newMyArray() *MyArray {
	return &MyArray{make([]int, 4), 0}
}

func (this *MyArray) Insert(index int, value int) bool {
	if index > this.k { // index == this.k 即末尾追加
		return false
	}
	if this.k == cap(this.arr) {
		arr := make([]int, cap(this.arr)<<1)
		copy(arr, this.arr[:index])
		copy(arr[index+1:], this.arr[index:])
		this.arr = arr
	} else {
		copy(this.arr[index+1:], this.arr[index:])
	}
	this.arr[index] = value
	this.k++
	return true
}

func (this *MyArray) Delete(index int) int {
	if index < 0 || index >= this.k {
		return -1
	}
	v := this.arr[index]
	copy(this.arr[index:], this.arr[index+1:])
	this.k--
	return v
}

func (this *MyArray) Get(index int) int {
	if index < 0 || index >= this.k {
		return -1
	}
	return this.arr[index]
}
