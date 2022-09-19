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

type People []*Person

func (p People) AverageAge() int {
	num := len(p)
	if num == 0 {
		return 0
	}
	s := 0
	for _, i := range p {
		s += i.Age
	}
	return s / num
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
	case "summary":
		var people People
		if len(os.Args) < 3 {
			showHelp()
			os.Exit(1)
		}
		if err := readAndUnmarshal(os.Args[2], &people); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
			os.Exit(1)
		}
		fmt.Printf("%d people, average age: %d\n", len(people), people.AverageAge())
	case "estimate":
		if len(os.Args) < 4 {
			showHelp()
			os.Exit(1)
		}
		var productMap ProductMap
		if err := readAndUnmarshal(os.Args[2], &productMap); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
			os.Exit(1)
		}
		var request Request
		if err := readAndUnmarshal(os.Args[3], &request); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
			os.Exit(1)
		}
		res, err := productMap.Calculate(request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
		}
		b, err := json.MarshalIndent(res, "", "  ")
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

func readAndUnmarshal(path string, dest interface{}) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dest)
}

func showHelp() {
	fmt.Printf("Usage:\n")
	fmt.Printf("  %s example\n", os.Args[0])
	fmt.Printf("    Shows an example of JSON data\n")
	fmt.Printf("  %s summary FILE\n", os.Args[0])
	fmt.Printf("    Shows summary of people from FILE\n")
	fmt.Printf("  %s estimate PRODUCT_FILE REQUEST_FILE\n", os.Args[0])
	fmt.Printf("    Shows estimate for REQUEST_FILE with PRODUCT_FILE\n")
}
