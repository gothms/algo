package test

import (
	"algo/sort"
	"algo/stack"
	"fmt"
	"testing"
)

var arr = []int{7, 2, 3, 4, 1, 0, 9, 6, 8, 5}

func TestBubbleSortTest(t *testing.T) {
	sort.BubbleSort(arr)
	fmt.Println(arr)
}
func TestInsertSortTest(t *testing.T) {
	sort.InsertionSort(arr)
	fmt.Println(arr)
}
func TestSelectSortTest(t *testing.T) {
	sort.SelectionSort(arr)
	fmt.Println(arr)
}

//func TestQuickSortTest(t *testing.T) {
//	basic.QuickSortTest(arr, 0, len(arr)-1)
//	fmt.Println(arr)
//}
//func TestMergeSortTest(t *testing.T) {
//	basic.MergeSortTest(arr, 0, len(arr)-1)
//	fmt.Println(arr)
//}
//func TestHeapSortTest(t *testing.T) {
//	basic.HeapSortTest(arr)
//	fmt.Println(arr)
//}

// stack
func TestStackOperation(t *testing.T) {
	s := "3+5*8-6"
	s = "31+5*8-6/4*5-11"
	s = "567*3/7*2-888+1/3*999+2-55+9*5/4*6+3" // -386
	operation := stack.StackOperation(s)
	t.Log(operation)
}
