/*
Test for benchmarking Golang with a simple slice addition
*/
package sliceadd

import "testing"

// func BenchmarkSliceAdd100000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		sliceAdd(100000)
// 	}
// }
// func BenchmarkSliceAdd10000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		sliceAdd(10000)
// 	}
// }

// func BenchmarkSliceAdd10(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		sliceAdd(10)
// 	}
// }

func BenchmarkSliceAddOnly10(b *testing.B) {
	slice := makeRange(0, 10-1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceAddOnly(slice, 10)
	}
}

func BenchmarkSliceAddOnly10000(b *testing.B) {
	slice := makeRange(0, 10000-1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceAddOnly(slice, 10000)
	}
}

func BenchmarkSliceAddOnly100000(b *testing.B) {
	slice := makeRange(0, 100000-1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceAddOnly(slice, 100000)
	}
}
