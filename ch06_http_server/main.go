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
	readReqBody := readReqBodyFunc(req)
	if req.URL.Query().Get("debug") != "" {
		body, err := readReqBody()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		echo(req, body)
	}
	pathParts := strings.Split(req.URL.Path, "/")
	if len(pathParts) < 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch pathParts[1] {
	case "add":
		w.WriteHeader(handleAdd(w, req))
	case "subtract":
		w.WriteHeader(handleSubtract(w, req, pathParts))
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
		body, err := readReqBody()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		var params map[string]interface{}
		if err := json.Unmarshal(body, &params); err != nil {
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

func handleAdd(w http.ResponseWriter, req *http.Request) int {
	if req.Method != "GET" {
		return http.StatusMethodNotAllowed
	}
	a, err := strconv.Atoi(req.URL.Query().Get("a"))
	if err != nil {
		return http.StatusBadRequest
	}
	b, err := strconv.Atoi(req.URL.Query().Get("b"))
	if err != nil {
		return http.StatusBadRequest
	}
	fmt.Fprintf(w, "%d\n", a+b)
	return http.StatusOK
}

func handleSubtract(w http.ResponseWriter, req *http.Request, pathParts []string) int {
	if req.Method != "GET" {
		return http.StatusMethodNotAllowed
	}
	if len(pathParts) < 4 {
		return http.StatusBadRequest
	}
	a, err := strconv.Atoi(pathParts[2])
	if err != nil {
		return http.StatusBadRequest
	}
	b, err := strconv.Atoi(pathParts[3])
	if err != nil {
		return http.StatusBadRequest
	}
	fmt.Fprintf(w, "%d\n", a-b)
	return http.StatusOK
}

func readReqBodyFunc(req *http.Request) func() ([]byte, error) {
	// https://ja.wikipedia.org/wiki/メモ化 を参照
	var memo []byte
	return func() ([]byte, error) {
		if memo == nil {
			var err error
			memo, err = ioutil.ReadAll(req.Body)
			if err != nil {
				return nil, err
			}
		}
		return memo, nil
	}
}

func echo(req *http.Request, b []byte) {
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
}
