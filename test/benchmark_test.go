package test

import (
	"container/heap"
	"testing"
)

//func BenchmarkN(b *testing.B) {
//	n := 100000
//	arr := make([]int, n)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < n; j++ {
//			arr[j] = j
//		}
//	}
//	b.StopTimer()
//}
//
///*
//BenchmarkN-8      					 36399             31917 ns/op
//BenchmarkNMultiplyOne-8         	 36369             32602 ns/op
//*/
//func BenchmarkNMultiplyOne(b *testing.B) {
//	n := 100001
//	arr := make([]int, n-1)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < n-1; j++ {
//			arr[j] = j
//		}
//	}
//	b.StopTimer()
//}

/*
go test -bench="." test/benchmark_test.go -benchmem

不用指针：heap.Pop(&hp)
BenchmarkHeapNormal-8               7628            177797 ns/op           31021 B/op       1243 allocs/op
BenchmarkHeapAllocs-8              10000            175653 ns/op           31586 B/op       1243 allocs/op

用指针：heap.Pop(hp)
BenchmarkHeapNormal-8               8542            181210 ns/op           33492 B/op       1243 allocs/op
BenchmarkHeapAllocs-8               9343            175985 ns/op           33107 B/op       1243 allocs/op

总结：用指针略快，初始化容量略快
	h := hpTest(make([]int, 0, n>>1))
	hp := &h
*/

func BenchmarkHeapNormal(b *testing.B) {
	b.ResetTimer()
	n := 2000
	hp := &hpTest{}
	for i := 0; i < b.N; i++ {
		for j := 0; j < n>>1; j++ {
			heap.Push(hp, j)
			if j&1 > 0 {
				heap.Pop(hp)
			}
		}
	}
	b.StopTimer()
}
func BenchmarkHeapAllocs(b *testing.B) {
	b.ResetTimer()
	n := 2000
	h := hpTest(make([]int, 0, n>>1))
	hp := &h
	for i := 0; i < b.N; i++ {
		for j := 0; j < n>>1; j++ {
			heap.Push(hp, j)
			if j&1 > 0 {
				heap.Pop(hp)
			}
		}
	}
	b.StopTimer()
}

type hpTest []int

func (h hpTest) Len() int {
	return len(h)
}

func (h hpTest) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h hpTest) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hpTest) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *hpTest) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

var _ heap.Interface = (*hpTest)(nil)
