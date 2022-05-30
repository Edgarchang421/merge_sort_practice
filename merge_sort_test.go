package main

import (
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	sortedNums := MergeSortRecursive(nums)

	sortedOrNot := sort.SliceIsSorted(sortedNums, func(p, q int) bool {
		return sortedNums[p] < sortedNums[q]
	})
	if !sortedOrNot {
		t.Error("sort.SliceIsSorted check fail")
	}
	t.Log("MergeSortRecursive pass sort.SliceIsSorted test")
	// write some test
}

func BenchmarkMergeSortRecursiveClean(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSortRecursiveClean([]int{22, 99, 27, 175, 300, 201, 1, 22, 56, 88, 3, 176, 22, 9, 73, 13, 11, 22, 143, 88})
	}
	b.StopTimer()
}

func BenchmarkMergeSortIterateCleans(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSortIterateClean([]int{22, 99, 27, 175, 300, 201, 1, 22, 56, 88, 3, 176, 22, 9, 73, 13, 11, 22, 143, 88})
	}
	b.StopTimer()
}

/*
goos: windows
goarch: amd64
pkg: marge_sort
cpu: AMD Ryzen 7 PRO 4750G with Radeon Graphics
BenchmarkMergeSortRecursiveClean-16       592401              1987 ns/op            1776 B/op         56 allocs/op
BenchmarkMergeSortIterateCleans-16        391989              3120 ns/op            2760 B/op         80 allocs/op
*/
