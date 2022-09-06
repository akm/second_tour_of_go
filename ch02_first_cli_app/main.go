package main

import (
	"fmt"
	"os"
	"strconv"
)

var calcMap = map[string]func() Calculation{
	"add":      func() Calculation { return &Addition{} },
	"subtract": func() Calculation { return &Subtraction{} },
}

func main() {
	if len(os.Args) < 4 {
		showHelp()
		os.Exit(1)
	}

	calcFactory, ok := calcMap[os.Args[1]]
	if !ok {
		showHelp()
		os.Exit(1)
	}
	calc := calcFactory()

	x, err := parseInt(os.Args[2])
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	y, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s: %d\n", calc.String(), calc.Do(x, y))
}

func showHelp() {
	fmt.Printf("USAGE:\n")
	fmt.Printf("%s (add|subtract) X Y\n", os.Args[0])
	fmt.Printf("   Shows addition or subtraction with X and Y\n")
	fmt.Printf("   X and Y must be number\n")
}

func parseInt(s string) (int, error) {
	r, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("ERROR: %q is not a number", s)
	}
	return r, nil
}
