package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	var reqBody []byte
	if req.URL.Query().Get("debug") != "" {
		var err error
		reqBody, err = ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		if err := echo(req, reqBody); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	pathParts := strings.Split(req.URL.Path, "/")
	if len(pathParts) < 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch pathParts[1] {
	case "add":
		if req.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		a, err := strconv.Atoi(req.URL.Query().Get("a"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		b, err := strconv.Atoi(req.URL.Query().Get("b"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "%d\n", a+b)
		return
	case "subtract":
		if req.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if len(pathParts) < 4 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		a, err := strconv.Atoi(pathParts[2])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		b, err := strconv.Atoi(pathParts[3])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "%d\n", a-b)
	case "multiply":
		if req.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		a, err := strconv.Atoi(req.Header.Get("X-VALUE-A"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		b, err := strconv.Atoi(req.Header.Get("X-VALUE-B"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		result := map[string]int{"result": a * b}
		resBody, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(resBody)
		fmt.Fprintln(w)
	case "divide":
		if req.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if reqBody == nil {
			var err error
			reqBody, err = ioutil.ReadAll(req.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		var params map[string]interface{}
		if err := json.Unmarshal(reqBody, &params); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		a, ok := params["a"].(float64)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		b, ok := params["b"].(float64)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if b == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		result := map[string]float64{"result": a / b}
		resBody, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(resBody)
		fmt.Fprintln(w)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func echo(req *http.Request, b []byte) error {
	w := os.Stdout
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
