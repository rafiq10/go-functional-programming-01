package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func main() {
	intSlice := []int{1, 4, -3, 5, -9, 3, -1}
	positiveIntSlice := FilterPositive[int](intSlice)
	fmt.Println(positiveIntSlice)
}

func FilterPositive[T Number](s []T) []T {
	var r []T
	for _, v := range s {
		if v > 0 {
			r = append(r, v)
		}
	}
	return r
}
