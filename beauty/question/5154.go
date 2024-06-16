package main

import "fmt"

func main() {
	arr := []int{3, 2, 5, 1, 6}
	quickSort(arr, len(arr))
	fmt.Println(arr)
}
func quickSort(arr []int, n int) {
	insertSort := func(l, r int) {
		for i := l + 1; i <= r; i++ {
			v, j := arr[i], i-1
			for ; j >= l && arr[j] > v; j-- {
				arr[j+1] = arr[j]
			}
			arr[j+1] = v
		}
	}
	var qSort func(int, int)
	qSort = func(l, r int) {
		if r-l < 8 {
			insertSort(l, r)
			return
		}
		m := (l + r) >> 1
		if arr[l] > arr[r] { // 三数取中
			arr[l], arr[r] = arr[r], arr[l]
		}
		if arr[m] < arr[r] {
			if arr[l] > arr[m] {
				arr[l], arr[r] = arr[r], arr[l]
			} else {
				arr[m], arr[r] = arr[r], arr[m]
			}
		}
		v, counter := arr[r], l // pivot 取 r
		for i := l; i < r; i++ {
			if arr[i] <= v {
				if counter != i {
					arr[counter], arr[i] = arr[i], arr[counter]
				}
				counter++
			}
		}
		arr[counter], arr[r] = arr[r], arr[counter] // pivot = counter
		qSort(l, counter-1)
		qSort(counter+1, r)
	}
	qSort(0, n-1)
}
