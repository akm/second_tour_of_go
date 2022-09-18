package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Birthday  string `json:"birthday"`
	Age       int    `json:"age"`
}

func main() {
	if len(os.Args) < 2 {
		showHelp()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "example":
		people := []*Person{
			{
				FirstName: "Blake",
				LastName:  "Serilda",
				Birthday:  "1989-07-10",
				Age:       33,
			},
			{
				FirstName: "Libbie",
				LastName:  "Drisko",
				Birthday:  "1998-06-15",
				Age:       24,
			},
			{
				FirstName: "Anestassia",
				LastName:  "Truc",
				Birthday:  "1973-04-02",
				Age:       48,
			},
		}
		b, err := json.MarshalIndent(people, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
			os.Exit(1)
		}
		fmt.Println(string(b))
	default:
		showHelp()
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Printf("Usage:\n")
	fmt.Printf("  %s example\n", os.Args[0])
	fmt.Printf("    Shows an example of JSON data\n")
}
