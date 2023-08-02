package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Sortable interface {
	int | float64 | string
}

func main() {
	Root()
}

func Root() {
	intFlag := flag.Bool("int", false, "Use to sort integer array")
	floatFlag := flag.Bool("float", false, "Use to sort float array")
	stringFlag := flag.Bool("string", false, "Use to sort string array")
	mixFlag := flag.Bool("mix", false, "Use to sort mixed array of interger, float or string")

	if len(os.Args) == 1 {
		fmt.Println("Expected one sub-command, ex: -int")
		os.Exit(1)
	}

	if len(os.Args) == 2 {
		fmt.Println("No input provided")
		os.Exit(1)
	}

	flag.Parse()
	switch {
	case *intFlag:
		arr, err := ParseInts(flag.Args())
		if err != nil {
			fmt.Println("Error when parsing:", err)
			os.Exit(1)
		}
		sorted := SortItems[int](arr)
		PrintItems[int](sorted)

	case *floatFlag:
		arr, err := ParseFloats(flag.Args())
		if err != nil {
			fmt.Println("Error when parsing:", err)
			os.Exit(1)
		}
		sorted := SortItems[float64](arr)
		PrintItems[float64](sorted)

	case *stringFlag:
		arr := flag.Args()
		sorted := SortItems[string](arr)
		PrintItems[string](sorted)

	case *mixFlag:
		numArr, strArr := ParseMixedItems(flag.Args())
		if len(numArr) != 0 {
			sorted := SortItems[float64](numArr)
			PrintItems(sorted)
		}
		if len(strArr) != 0 {
			sorted := SortItems[string](strArr)
			PrintItems(sorted)
		}

	default:
		fmt.Println("Unknow flag:", os.Args[1])
	}
}

func ParseInts(items []string) ([]int, error) {
	intArr := make([]int, len(items))

	for i, item := range items {
		v, err := strconv.Atoi(item)
		if err != nil {
			return []int{}, err
		}
		intArr[i] = v
	}
	return intArr, nil
}

func ParseFloats(items []string) ([]float64, error) {
	floatArr := make([]float64, len(items))

	for i, item := range items {
		v, err := strconv.ParseFloat(item, 64)
		if err != nil {
			return []float64{}, err
		}
		floatArr[i] = v
	}
	return floatArr, nil
}

func ParseMixedItems(items []string) ([]float64, []string) {
	var numArr []float64
	var strArr []string

	for _, item := range items {
		v, err := strconv.ParseFloat(item, 64)
		if err == nil {
			numArr = append(numArr, v)
			continue
		}
		strArr = append(strArr, item)
	}
	return numArr, strArr
}

func SortItems[T Sortable](items []T) []T {
	return quickSort(items, 0, len(items)-1)
}

func quickSort[T Sortable](items []T, low, high int) []T {
	if low < high {
		var p int
		items, p = partition(items, low, high)
		items = quickSort(items, low, p-1)
		items = quickSort(items, p+1, high)
	}
	return items
}

func partition[T Sortable](arr []T, low, high int) ([]T, int) {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func PrintItems[T Sortable](items []T) {
	for _, v := range items {
		fmt.Print(v, " ")
	}
}
