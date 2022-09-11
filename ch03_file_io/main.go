package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "cat":
		if len(os.Args) < 3 {
			showHelp()
			os.Exit(1)
		}
		b, err := os.ReadFile(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(string(b))
	default:
		showHelp()
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Printf("USAGE:\n")
	fmt.Printf("%s cat FILE\n", os.Args[0])
	fmt.Printf("   Shows the contents of FILE\n")
}
