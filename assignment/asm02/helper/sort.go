package helper

import (
	"sort"
)

func SortInts(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func SortFloats(arr []float64) []float64 {
	sort.Float64s(arr)
	return arr
}

func SortStrings(arr []string) []string {
	sort.Strings(arr)
	return arr
}
