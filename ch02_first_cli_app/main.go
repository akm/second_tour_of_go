package main

import "os"

func main() {
	if len(os.Args) > 1 {
		println("Hello, " + os.Args[1] + "!")
	} else {
		println("Hello, someone!")
	}
}
