/*
Test for benchmarking Golang with a simple vector addition
*/
package vecadd

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

// func BenchmarkVecAdd100000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		vecAdd(100000)
// 	}
// }

// func BenchmarkVecAdd10000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		vecAdd(10000)
// 	}
// }
// func BenchmarkVecAdd10(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		vecAdd(10)
// 	}
// }

// func BenchmarkMatAdd10000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		matAdd(10000)
// 	}
// }

func BenchmarkVecAddOnly10(b *testing.B) {
	size := 10
	vec := mat.NewVecDense(size, makeRange(0, size-1))
	one := mat.NewVecDense(size, rep(1, size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vecAddOnly(vec, one)
	}
}

func BenchmarkVecAddOnly10000(b *testing.B) {
	size := 10000
	vec := mat.NewVecDense(size, makeRange(0, size-1))
	one := mat.NewVecDense(size, rep(1, size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vecAddOnly(vec, one)
	}
}

func BenchmarkVecAddOnly100000(b *testing.B) {
	size := 100000
	vec := mat.NewVecDense(size, makeRange(0, size-1))
	one := mat.NewVecDense(size, rep(1, size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vecAddOnly(vec, one)
	}
}

func BenchmarkVecAddOnly1000000(b *testing.B) {
	size := 1000000
	vec := mat.NewVecDense(size, makeRange(0, size-1))
	one := mat.NewVecDense(size, rep(1, size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vecAddOnly(vec, one)
	}
}
func BenchmarkVecAddOnly10000000(b *testing.B) {
	size := 10000000
	vec := mat.NewVecDense(size, makeRange(0, size-1))
	one := mat.NewVecDense(size, rep(1, size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vecAddOnly(vec, one)
	}
}
