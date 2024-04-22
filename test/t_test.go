package test

import (
	"algo/sort"
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
