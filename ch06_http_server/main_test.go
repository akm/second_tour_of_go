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
}
