package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func GMin[T constraints.Ordered](arr []T, size int) []T {
	for i := size - 1; i > 0; i-- {

		f := 1
		for e := 0; e < i; e++ {

			if arr[e] > arr[f] {
				g := arr[e]
				arr[e] = arr[f]
				arr[f] = g
			}
			f++
		}
	}

	return arr
}

func main() {
	arr := []int{1, 5, 3}
	fmt.Println(GMin[int](arr, 3))
}
