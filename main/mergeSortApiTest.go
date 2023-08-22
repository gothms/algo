package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// rotate Test
	//ro := []int{4, 5, 6, 7, 8, 9, 1, 2, 3, 4}
	//rotate(ro, 0, 6, 10)
	//fmt.Println(ro)

	n, max := 10, 10
	arr := makeSortSlice(n, max)
	brr := makeSortSlice(n, max-2)
	arr = append(arr, brr...)
	//arr = []int{1, 2, 3, 4, 5, 4, 4, 5, 6, 7, 8}
	fmt.Println(arr)

	//arr := []int{3, 3, 5, 7, 9, 0, 1, 1, 8, 10}
	a, m, b := 0, len(arr)>>1, len(arr)
	symMerge(arr, a, m, b)
	fmt.Println(arr)
	boo := check(arr)
	if !boo {
		fmt.Println("===========有 bug=========")
	}

	// 结论：将 merge 两个数组中，相等的元素提前交换，比较好，即：第一个大于等于目标的元素
	data := []int{1, 2, 3, 4, 4, 5, 6}
	start, r := 0, len(data)
	tar := 4
	for start < r {
		c := int(uint(start+r) >> 1)
		//if !data.Less(p-c, c) {
		//if data[p-c] >= data[c] {
		//if tar >= data[c] {	// 第一个大于目标的元素
		if tar > data[c] { // 第一个大于等于目标的元素
			start = c + 1
		} else {
			r = c
		}
	}
	fmt.Println(start, r)
}

func symMerge(data []int, a, m, b int) {
	// Avoid unnecessary recursions of symMerge
	// by direct insertion of data[a] into data[m:b]
	// if data[a:m] only contains one element.
	if m-a == 1 {
		// Use binary search to find the lowest index i
		// such that data[i] >= data[a] for m <= i < b.
		// Exit the search loop with i == b in case no such index exists.
		i := m
		j := b
		for i < j {
			h := int(uint(i+j) >> 1)
			//if data.Less(h, a) {
			if data[h] < data[a] {
				i = h + 1
			} else {
				j = h
			}
		}
		// Swap values until data[a] reaches the position before i.
		for k := a; k < i-1; k++ {
			//data.Swap(k, k+1)
			data[k], data[k+1] = data[k+1], data[k]
		}
		return
	}

	// Avoid unnecessary recursions of symMerge
	// by direct insertion of data[m] into data[a:m]
	// if data[m:b] only contains one element.
	if b-m == 1 {
		// Use binary search to find the lowest index i
		// such that data[i] > data[m] for a <= i < m.
		// Exit the search loop with i == m in case no such index exists.
		i := a
		j := m
		for i < j {
			h := int(uint(i+j) >> 1)
			//if !data.Less(m, h) {
			if data[m] >= data[h] {
				i = h + 1
			} else {
				j = h
			}
		}
		// Swap values until data[m] reaches the position i.
		for k := m; k > i; k-- {
			//data.Swap(k, k-1)
			data[k], data[k-1] = data[k-1], data[k]
		}
		return
	}

	mid := int(uint(a+b) >> 1)
	n := mid + m
	var start, r int
	if m > mid {
		start = n - b
		r = mid
	} else {
		start = a
		r = m
	}
	p := n - 1

	//fmt.Println("start:", start, r, p)
	for start < r {
		c := int(uint(start+r) >> 1)
		//if !data.Less(p-c, c) {
		if data[p-c] > data[c] { // 优于 >=
			start = c + 1
		} else {
			r = c
		}
	}
	//fmt.Println(start, r, p)
	// [start,m-1] 有序区间最小值 > [m,end) 有序区间最大值
	end := n - start
	if start < m && m < end {
		//fmt.Println("rotate:", start, m, end)
		rotate(data, start, m, end) // 交换
	}
	fmt.Println("data:", data)
	if a < start && start < mid {
		symMerge(data, a, start, mid)
	}
	if mid < end && end < b {
		symMerge(data, mid, end, b)
	}
}

// rotate rotates two consecutive blocks u = data[a:m] and v = data[m:b] in data:
// Data of the form 'x u v y' is changed to 'x v u y'.
// rotate performs at most b-a many calls to data.Swap,
// and it assumes non-degenerate arguments: a < m && m < b.
func rotate(data []int, a, m, b int) {
	fmt.Println(a, m, b)
	i := m - a
	j := b - m

	for i != j {
		if i > j {
			swapRange(data, m-i, m, j)
			i -= j
		} else {
			swapRange(data, m-i, m+j-i, i)
			j -= i
		}
		//fmt.Println(data)
	}
	// i == j
	swapRange(data, m-i, m, i)
}
func swapRange(data []int, a, b, n int) {
	for i := 0; i < n; i++ {
		//data.Swap(a+i, b+i)
		data[a+i], data[b+i] = data[b+i], data[a+i]
	}
}
func makeSortSlice(n, max int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		v := rand.Intn(max)
		ret[i] = v
	}
	sort.Ints(ret)
	return ret
}
func check(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
