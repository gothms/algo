package test

import (
	"algo/maths"
	"testing"
)

// go test -bench="." test/math_test.go -benchmem
/*
BenchmarkPermuteUniqueMemo-8         680           1806008 ns/op         2688256 B/op      15140 allocs/op
BenchmarkPermuteUnique-8             850           1335197 ns/op         2688239 B/op      15140 allocs/op
BenchmarkSubsetsAndPermute-8         742           1563163 ns/op         2688365 B/op      15143 allocs/op

[]int{1, 2, 3, 3, 1, 5, 3, 2, 4, 9, 5, 9}
BenchmarkPermuteUnique-8               2         518461350 ns/op        1139993604 B/op  4989649 allocs/op
BenchmarkSubsetsAndPermute-8           2         606576850 ns/op        1139994560 B/op  4989650 allocs/op
*/

var arr_math_test = []int{1, 2, 3, 3, 1, 5, 3, 2, 4, 9, 5, 9}

//	func BenchmarkPermuteUniqueMemo(b *testing.B) {
//		b.ResetTimer()
//		for i := 0; i < b.N; i++ {
//			math.PermuteUniqueMemo(arr_math_test)
//		}
//		b.StopTimer()
//	}
func BenchmarkPermuteUnique(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maths.PermuteUnique(arr_math_test)
	}
	b.StopTimer()
}
func BenchmarkSubsetsAndPermute(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maths.SubsetsAndPermute(arr_math_test)
	}
	b.StopTimer()
}
