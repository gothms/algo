package main

import (
	"algo/sort"
	"fmt"
	gosys "sort"
)

func main() {
	// sort
	arr := []int{3, 2, 6, 13, 54, 32, 654, 889, 44, 0, 34, 987, 45, 765, 23, 76}
	arr = []int{3, 2, 6, 13, 54, 32, 3, 654, 44, 889, 0, 987, 44, 0, 34, 987, 45, 765, 23, 76}
	//sort.BubbleSort(arr)
	//sort.InsertionSort(arr)
	//sort.SelectionSort(arr)
	//sort.QuickSort(arr, 0, len(arr)-1)
	//sort.HeapSort(arr)
	sort.MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
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
}
