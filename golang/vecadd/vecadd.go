/*
Package for a simple vector addition using Golang's GoNum library
*/
package vecadd

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

const NUMOPS = 100000

func makeRange(min int, max int) []float64 {
	a := make([]float64, max-min+1)
	for i := range a {
		a[i] = float64(min + i)
	}
	return a
}

func rep(val float64, num int) []float64 {
	a := make([]float64, num)
	for i := range a {
		a[i] = val
	}
	return a
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func vecAdd(size int) *mat.VecDense {
	res := mat.NewVecDense(size, makeRange(0, size-1))
	one := mat.NewVecDense(size, rep(1, size))
	for i := 0; i < NUMOPS; i++ {
		res.AddVec(res, one)
	}
	return res
}

func vecAddOnly(vec *mat.VecDense, one *mat.VecDense) *mat.VecDense {
	for i := 0; i < NUMOPS; i++ {
		vec.AddVec(vec, one)
	}
	return vec
}

func matAdd(size int) *mat.Dense {
	res := mat.NewDense(1, size, makeRange(0, size-1))
	one := mat.NewDense(1, size, rep(1, size))
	for i := 0; i < NUMOPS; i++ {
		res.Add(res, one)
	}
	return res
}
