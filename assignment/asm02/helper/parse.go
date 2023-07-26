package helper

import (
	"strconv"
)

func ParseInts(arr []string) ([]int, error) {
	result := make([]int, len(arr))

	for i, item := range arr {
		num, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		result[i] = num
	}
	return result, nil
}

func ParseFloats(arr []string) ([]float64, error) {
	result := make([]float64, len(arr))

	for i, val := range arr {
		num, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, err
		}
		result[i] = num
	}
	return result, nil
}
