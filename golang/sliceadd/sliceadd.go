/*
Package for a simple addition using Golang
*/
package sliceadd

const NUMITER = 100

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func sliceAdd(size int) []int {
	//slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	slice := makeRange(0, size-1)
	for i := 0; i < NUMITER; i++ {
		for key, value := range slice {
			slice[key] = value + 1
		}
	}
	return slice
}

func sliceAddOnly(slice []int, size int) []int {
	for i := 0; i < NUMITER; i++ {
		for key, value := range slice {
			slice[key] = value + 1
		}
	}
	return slice
}
