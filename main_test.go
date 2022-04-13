package main

import "testing"

func BenchmarkFilterWithFP(b *testing.B) {
	intSlice := []int{1, 4, -3, 5, -9, 3, -1}
	for i := 0; i < b.N; i++ {
		intStream := Stream[int]{intSlice}
		intStream.filter(IsPositive[int]).filter(IsEven[int])
	}
}

func BenchmarkWithloops(b *testing.B) {
	intSlice := []int{1, 4, -3, 5, -9, 3, -1}
	for i := 0; i < b.N; i++ {
		var positiveIntSlie []int
		for _, v := range intSlice {
			if v > 0 {
				positiveIntSlie = append(positiveIntSlie, v)
			}
		}
		var results []int
		for _, v := range positiveIntSlie {
			if v%2 == 0 {
				results = append(results, v)
			}
		}
	}
}

func BenchmarkWithOneloop(b *testing.B) {
	intSlice := []int{1, 4, -3, 5, -9, 3, -1}
	for i := 0; i < b.N; i++ {
		var result []int
		for _, v := range intSlice {
			if v > 0 && v%2 == 0 {
				result = append(result, v)
			}
		}

	}
}
