package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	switch os.Args[1] {
	case "example":
		if len(os.Args) < 2 {
			showHelp()
			os.Exit(1)
		}
		b, err := json.Marshal(EstimateRequestExample1)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(b))
	case "unmarshal":
		if len(os.Args) < 3 {
			showHelp()
			os.Exit(1)
		}
		b, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var req EstimateRequest
		if err := json.Unmarshal(b, &req); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(req.Text())
	default:
		showHelp()
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Printf("USAGE:\n")
	fmt.Printf("  %s example\n", os.Args[0])
	fmt.Printf("  %s unmarshal path/to/file.json\n", os.Args[0])
}
