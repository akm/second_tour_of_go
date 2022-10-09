package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "--start line--\n")
		fmt.Fprintf(w, "method: %s\n", req.Method)
		fmt.Fprintf(w, "scheme: %s\n", req.URL.Scheme)
		fmt.Fprintf(w, "domain: %s\n", req.URL.Hostname())
		fmt.Fprintf(w, "port: %s\n", req.URL.Port())
		fmt.Fprintf(w, "path: %s\n", req.URL.Path)
		fmt.Fprintf(w, "parameters:\n")
		for k, values := range req.URL.Query() {
			fmt.Fprintf(w, "  %s: %s\n", k, strings.Join(values, ","))
		}
		fmt.Fprintf(w, "anchor: %s\n", req.URL.Fragment)

		fmt.Fprintf(w, "\n--headers--\n")
		for k, values := range req.Header {
			fmt.Fprintf(w, "%s: %s\n", k, strings.Join(values, ","))
		}

		fmt.Fprintf(w, "\n--body--\n")
		w.Write(b)
		fmt.Fprintf(w, "\n")
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
