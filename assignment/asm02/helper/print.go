package helper

import (
	"fmt"
)

var str = "Output:"

func PrintInts(arr []int) {
	for _, val := range arr {
		str += " " + fmt.Sprint(val)
	}
	fmt.Println(str)
}

func PrintFloats(arr []float64) {
	for _, val := range arr {
		str += " " + fmt.Sprint(val)
	}
	fmt.Println(str)
}

func PrintStrings(arr []string) {
	for _, val := range arr {
		str += " " + fmt.Sprint(val)
	}
	fmt.Println(str)
}
