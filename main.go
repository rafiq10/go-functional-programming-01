package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type Stream[T Number] struct {
	s []T
}

func (s *Stream[T]) filter(f func(T) bool) *Stream[T] {
	s.s = Filter(s.s, f)
	return s
}

func main() {
	intSlice := []int{1, 4, -3, 5, -9, 3, -1}
	positiveIntSlice := Filter[int](intSlice, IsPositive[int])
	fmt.Println(positiveIntSlice)
	s := Stream[int]{s: intSlice}
	s.filter(IsPositive[int])
	fmt.Printf("Filtered positive numbers: %v \n", s)

	s.filter(IsPositive[int]).filter(IsEven[int])
	fmt.Printf("Filtered positive and even numbers: %v \n", s)

}

func Filter[T Number](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func IsPositive[T Number](v T) bool {
	return v > 0
}

func IsEven[T constraints.Integer](a T) bool {
	return (a % 2) == 0
}
