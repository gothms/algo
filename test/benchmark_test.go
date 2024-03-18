package test

import "testing"

func BenchmarkN(b *testing.B) {
	n := 100000
	arr := make([]int, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			arr[j] = j
		}
	}
	b.StopTimer()
}

/*
BenchmarkN-8      					 36399             31917 ns/op
BenchmarkNMultiplyOne-8         	 36369             32602 ns/op
*/
func BenchmarkNMultiplyOne(b *testing.B) {
	n := 100001
	arr := make([]int, n-1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n-1; j++ {
			arr[j] = j
		}
	}
	b.StopTimer()
}
