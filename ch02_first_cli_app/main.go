package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("USAGE:\n")
		fmt.Printf("%s X Y\n", os.Args[0])
		fmt.Printf("   Shows sum of X and Y\n")
		fmt.Printf("   X and Y must be number\n")
		os.Exit(1)
	}
	x, err := parseInt(os.Args[1])
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	y, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Result: %d\n", x+y)
}

func parseInt(s string) (int, error) {
	r, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("ERROR: %q is not a number", s)
	}
	return r, nil
}
