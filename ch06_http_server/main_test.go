package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	textResponse := func(method, path string, result string) func(t *testing.T) {
		return func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(method, path, nil)
			req.Host = "localhost:8080"
			handler(w, req)
			resp := w.Result()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			body, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Equal(t, result, strings.TrimSpace(string(body)))
		}
	}

	errorResponse := func(method, path string, status int) func(t *testing.T) {
		return func(t *testing.T) {
			w := httptest.NewRecorder()
			handler(w, httptest.NewRequest(method, path, nil))
			resp := w.Result()
			assert.Equal(t, status, resp.StatusCode)
		}
	}

	t.Run("add", func(t *testing.T) {
		t.Run("valid case1", textResponse("GET", "/add?a=1&b=2", "3"))
		t.Run("valid case1", textResponse("GET", "/add?a=100&b=200", "300"))
		t.Run("with debug", textResponse("GET", "/add?a=100&b=200&debug=1", "300"))
		t.Run("no parameters", errorResponse("GET", "/add", http.StatusBadRequest))
		t.Run("without b", errorResponse("GET", "/add?a=100", http.StatusBadRequest))
		t.Run("without a", errorResponse("GET", "/add?b=100", http.StatusBadRequest))
		t.Run("invalid a", errorResponse("GET", "/add?a=foo&bar=200", http.StatusBadRequest))
		t.Run("invalid b", errorResponse("GET", "/add?a=100&bar=bar", http.StatusBadRequest))
		t.Run("POST", errorResponse("POST", "/add?a=1&b=2", http.StatusMethodNotAllowed))
	})

	t.Run("subtract", func(t *testing.T) {
		t.Run("valid case1", textResponse("GET", "/subtract/1/2", "-1"))
		t.Run("valid case1", textResponse("GET", "/subtract/200/100", "100"))
		t.Run("with debug", textResponse("GET", "/subtract/100/200?debug=1", "-100"))
		t.Run("no parameters", errorResponse("GET", "/subtract", http.StatusBadRequest))
		t.Run("without b", errorResponse("GET", "/subtract/100", http.StatusBadRequest))
		t.Run("invalid a", errorResponse("GET", "/subtract/foo/200", http.StatusBadRequest))
		t.Run("invalid b", errorResponse("GET", "/subtract/100/bar", http.StatusBadRequest))
		t.Run("POST", errorResponse("POST", "/subtract/1/2", http.StatusMethodNotAllowed))
	})
}