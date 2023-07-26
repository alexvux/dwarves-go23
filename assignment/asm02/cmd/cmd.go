package cmd

import (
	"flag"
	"log"

	"github.com/alexvux/dwarves-go23/assignment/asm02/helper"
)

var intFlag, floatFlag, stringFlag *bool

func setUpFlag() {
	intFlag = flag.Bool("int", false, "Sort integer array")
	floatFlag = flag.Bool("float", false, "Sort float array")
	stringFlag = flag.Bool("string", false, "Sort string array")
}

func Execute() {
	setUpFlag()

	flag.Parse()
	args := flag.Args()

	switch {
	case *intFlag:
		arr, err := helper.ParseInts(args)
		if err != nil {
			log.Fatalln("Error on parsing to interger array: ", err)
		}

		sorted := helper.SortInts(arr)
		helper.PrintInts(sorted)

	case *floatFlag:
		arr, err := helper.ParseFloats(args)
		if err != nil {
			log.Fatalln("Error on parsing to float array: ", err)
		}

		sorted := helper.SortFloats(arr)
		helper.PrintFloats(sorted)

	case *stringFlag:
		sorted := helper.SortStrings(args)
		helper.PrintStrings(sorted)

	default:
		log.Fatalln("Unknow data type, please use 1 of these flags: -int -float -string")
	}
}
