package main

import (
	"fmt"
	"io"
	"os"
	"strings"
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
		f, err := os.Open(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
		b, err := io.ReadAll(f)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(string(b))
	case "write":
		if len(os.Args) < 4 {
			showHelp()
			os.Exit(1)
		}
		if _, err := os.Stat(os.Args[2]); err != nil {
			if !os.IsNotExist(err) {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("%s already exists. Overwrite? (y/n): ", os.Args[2])
			var answer string
			fmt.Scanln(&answer)
			if !strings.HasPrefix(strings.ToLower(answer), "y") {
				fmt.Println("Quit writing")
				os.Exit(1)
			}
		}
		f, err := os.Create(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
		if _, err := fmt.Fprintln(f, os.Args[3]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case "append":
		if len(os.Args) < 4 {
			showHelp()
			os.Exit(1)
		}
		f, err := os.OpenFile(os.Args[2], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
		if _, err = fmt.Fprintln(f, os.Args[3]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		showHelp()
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Printf("USAGE:\n")
	fmt.Printf("%s cat FILE\n", os.Args[0])
	fmt.Printf("   Shows the contents of FILE\n")
	fmt.Printf("%s write FILE CONTENT\n", os.Args[0])
	fmt.Printf("   Writes CONTENT to FILE\n")
	fmt.Printf("%s append FILE CONTENT\n", os.Args[0])
	fmt.Printf("   Appends CONTENT to FILE\n")
}
