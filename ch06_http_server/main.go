package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		if err := echo(w, req); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echo(w io.Writer, req *http.Request) error {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	hostParts := strings.SplitN(req.Host, ":", 2)

	fmt.Fprintf(w, "--start line--\n")
	fmt.Fprintf(w, "method: %s\n", req.Method)
	fmt.Fprintf(w, "scheme: %s\n", req.URL.Scheme)
	fmt.Fprintf(w, "domain: %s\n", hostParts[0])
	fmt.Fprintf(w, "port: %s\n", hostParts[1])
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
	return nil
}
