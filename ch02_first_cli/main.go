package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		showHelp()
		os.Exit(1)
	}
	var kind string
	switch os.Args[1] {
	case "add":
		kind = "addition"
	case "substract":
		kind = "substraction"
	default:
		showHelp()
		os.Exit(1)
	}

	x, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("ERROR: %q is not a number\n", os.Args[2])
		os.Exit(1)
	}
	y, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("ERROR: %q is not a number\n", os.Args[3])
		os.Exit(1)
	}

	var res int
	switch kind {
	case "addition":
		res = x + y
	case "substraction":
		res = x - y
	}
	fmt.Printf("%s: %d\n", kind, res)
}

func showHelp() {
	fmt.Printf("USAGE:\n")
	fmt.Printf("%s (add|substract) X Y\n", os.Args[0])
	fmt.Printf("   Shows addion or substraction with X and Y\n")
	fmt.Printf("   X and Y must be number\n")
}
