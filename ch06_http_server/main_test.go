package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	newReq := func(method, path string) *http.Request {
		req := httptest.NewRequest(method, path, nil)
		req.Host = "localhost:8080"
		return req
	}

	textResponse := func(req *http.Request, result string) func(t *testing.T) {
		return func(t *testing.T) {
			w := httptest.NewRecorder()
			handler(w, req)
			resp := w.Result()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			body, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Equal(t, result, strings.TrimSpace(string(body)))
		}
	}

	errorResponse := func(req *http.Request, status int) func(t *testing.T) {
		return func(t *testing.T) {
			w := httptest.NewRecorder()
			handler(w, req)
			resp := w.Result()
			assert.Equal(t, status, resp.StatusCode)
		}
	}

	t.Run("add", func(t *testing.T) {
		t.Run("valid case1", textResponse(newReq("GET", "/add?a=1&b=2"), "3"))
		t.Run("valid case1", textResponse(newReq("GET", "/add?a=100&b=200"), "300"))
		t.Run("with debug", textResponse(newReq("GET", "/add?a=100&b=200&debug=1"), "300"))
		t.Run("no parameters", errorResponse(newReq("GET", "/add"), http.StatusBadRequest))
		t.Run("without b", errorResponse(newReq("GET", "/add?a=100"), http.StatusBadRequest))
		t.Run("without a", errorResponse(newReq("GET", "/add?b=100"), http.StatusBadRequest))
		t.Run("invalid a", errorResponse(newReq("GET", "/add?a=foo&bar=200"), http.StatusBadRequest))
		t.Run("invalid b", errorResponse(newReq("GET", "/add?a=100&bar=bar"), http.StatusBadRequest))
		t.Run("POST", errorResponse(newReq("POST", "/add?a=1&b=2"), http.StatusMethodNotAllowed))
	})

	t.Run("subtract", func(t *testing.T) {
		t.Run("valid case1", textResponse(newReq("GET", "/subtract/1/2"), "-1"))
		t.Run("valid case1", textResponse(newReq("GET", "/subtract/200/100"), "100"))
		t.Run("with debug", textResponse(newReq("GET", "/subtract/100/200?debug=1"), "-100"))
		t.Run("no parameters", errorResponse(newReq("GET", "/subtract"), http.StatusBadRequest))
		t.Run("without b", errorResponse(newReq("GET", "/subtract/100"), http.StatusBadRequest))
		t.Run("invalid a", errorResponse(newReq("GET", "/subtract/foo/200"), http.StatusBadRequest))
		t.Run("invalid b", errorResponse(newReq("GET", "/subtract/100/bar"), http.StatusBadRequest))
		t.Run("POST", errorResponse(newReq("POST", "/subtract/1/2"), http.StatusMethodNotAllowed))
	})

	jsonResponse := func(req *http.Request, expectedResult float64) func(t *testing.T) {
		return func(t *testing.T) {
			w := httptest.NewRecorder()
			handler(w, req)
			resp := w.Result()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			body, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)
			var resData map[string]interface{}
			assert.NoError(t, json.Unmarshal(body, &resData))
			assert.Equal(t, expectedResult, resData["result"])
		}
	}

	t.Run("multiply", func(t *testing.T) {
		req := func(path, a, b string) *http.Request {
			r := httptest.NewRequest("POST", path, nil)
			r.Host = "localhost:8080"
			if a != "" {
				r.Header.Set("X-VALUE-A", a)
			}
			if b != "" {
				r.Header.Set("X-VALUE-B", b)
			}
			return r
		}
		t.Run("valid case1", jsonResponse(req("/multiply", "2", "3"), 6))
		t.Run("valid case2", jsonResponse(req("/multiply", "20", "30"), 600))
		t.Run("with debug", jsonResponse(req("/multiply?debug=1", "20", "-30"), -600))
		t.Run("no parameters", errorResponse(req("/multiply", "", ""), http.StatusBadRequest))
		t.Run("without a", errorResponse(req("/multiply", "", "2"), http.StatusBadRequest))
		t.Run("without b", errorResponse(req("/multiply", "3", ""), http.StatusBadRequest))
		t.Run("invalid a", errorResponse(req("/multiply", "foo", "4"), http.StatusBadRequest))
		t.Run("invalid b", errorResponse(req("/multiply", "5", "bar"), http.StatusBadRequest))
		t.Run("GET", errorResponse(newReq("GET", "/multiply"), http.StatusMethodNotAllowed))
	})

	t.Run("divide", func(t *testing.T) {
		req := func(path, a, b string) *http.Request {
			parts := []string{}
			if a != "" {
				parts = append(parts, fmt.Sprintf(`"a":%s`, a))
			}
			if b != "" {
				parts = append(parts, fmt.Sprintf(`"b":%s`, b))
			}
			jsonData := fmt.Sprintf("{%s}", strings.Join(parts, ","))
			t.Logf("jsonData: %s\n", jsonData)
			buf := bytes.NewBufferString(jsonData)
			r := httptest.NewRequest("POST", path, buf)
			r.Host = "localhost:8080"
			return r
		}
		t.Run("valid case1", jsonResponse(req("/divide", "6", "3"), 2))
		t.Run("valid case2", jsonResponse(req("/divide", "30", "20"), 1.5))
		t.Run("with debug", jsonResponse(req("/divide?debug=1", "90", "-30"), -3))
		t.Run("no parameters", errorResponse(req("/divide", "", ""), http.StatusBadRequest))
		t.Run("without a", errorResponse(req("/divide", "", "2"), http.StatusBadRequest))
		t.Run("without b", errorResponse(req("/divide", "3", ""), http.StatusBadRequest))
		t.Run("invalid a", errorResponse(req("/divide", "foo", "4"), http.StatusBadRequest))
		t.Run("invalid b", errorResponse(req("/divide", "5", "bar"), http.StatusBadRequest))
		t.Run("GET", errorResponse(newReq("GET", "/divide"), http.StatusMethodNotAllowed))
	})
}
