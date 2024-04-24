package main

import (
	"algo/sort"
	"fmt"
	"math/rand"
	gosys "sort"
)

func main() {
	// sort
	arr := []int{3, 2, 6, 13, 54, 32, 654, 889, 44, 0, 34, 987, 45, 765, 23, 76}
	arr = []int{3, 2, 6, 13, 54, 32, 3, 654, 44, 889, 0, 987, 44, 0, 34, 987, 45, 765, 23, 76}
	n, max := 27, 100
	arr = makeSortSlice(n, max)
	brr := make([]int, 0, n)
	brr = append(brr, arr...)

	//sort.BubbleSort(arr)
	//sort.InsertionSort(arr)
	//sort.SelectionSort(arr)
	//sort.QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	sort.HeapSort(arr)
	b := check(arr)
	fmt.Println("check:", b)
	sort.MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	sort.QuickSort(brr, 0, len(brr)-1)
	fmt.Println(brr)

	arr = []int{3, 2, 6, 0, 6, 0, 5, 4, 8, 1, 7, 7, 3, 9, 2, 2, 2, 5}
	//sort.CountingSortGeekBang(arr, 9)
	sort.CountingSort(arr, 9)
	sort.CountingSortGeekBang(arr, 9)
	fmt.Println("CountingSort", arr)
	fmt.Println("CountingSortGeekBang", arr)
	arr = []int{3, 2, 6, 13, 54, 32, 65, 89, 44, 0, 34, 98, 45, 75, 23, 76, 13, 72, 51, 40}
	sort.BucketSort(arr, 5)
	fmt.Println(arr)
	arr = []int{3, 2, 6, 13, 54, 0, 32, 654, 0, 0, 889, 44, 0, 34, 987, 45, 765, 23, 76}
	sort.RadixSort(arr)
	fmt.Println(arr)

	sortArr := gosys.IntSlice{1, 3, 5, 9, 7}
	gosys.Sort(gosys.Reverse(sortArr)) // 只是逆序排序了
	fmt.Println("sort.Reverse:", sortArr)

	// rand
	//l, r := 3, 7
	//for i := 0; i < 20; i++ {
	//	pivot := l + rand.Intn(r-l+1)
	//	fmt.Println(pivot)
	//}

	arr = []int{3, 3, 5, 7, 9, 0, 6, 7, 8, 10}
	start, r, p := 0, len(arr)>>1, len(arr)-1
	for start < r {
		c := int(uint(start+r) >> 1)
		//if !data.Less(p-c, c) {
		if arr[p-c] >= arr[c] {
			start = c + 1
		} else {
			r = c
		}
	}
	fmt.Println(start, r, p)
}
func makeSortSlice(n, max int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		v := rand.Intn(max)
		ret[i] = v
	}
	//sort.Ints(ret)
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
