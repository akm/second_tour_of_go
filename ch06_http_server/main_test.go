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
	assertTextResponse := func(t *testing.T, method, path string, result string) {
		w := httptest.NewRecorder()
		handler(w, httptest.NewRequest(method, path, nil))
		resp := w.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, result, strings.TrimSpace(string(body)))
	}

	assertError := func(t *testing.T, method, path string, status int) {
		w := httptest.NewRecorder()
		handler(w, httptest.NewRequest(method, path, nil))
		resp := w.Result()
		assert.Equal(t, status, resp.StatusCode)
	}

	t.Run("add", func(t *testing.T) {
		t.Run("GET", func(t *testing.T) {
			assertTextResponse(t, "GET", "/add?a=1&b=2", "3")
		})
		t.Run("GET", func(t *testing.T) {
			assertTextResponse(t, "GET", "/add?a=100&b=200", "300")
		})
		t.Run("GET", func(t *testing.T) {
			assertError(t, "GET", "/add", http.StatusBadRequest)
		})
		t.Run("GET", func(t *testing.T) {
			assertError(t, "GET", "/add?a=100", http.StatusBadRequest)
		})
		t.Run("GET", func(t *testing.T) {
			assertError(t, "GET", "/add?b=100", http.StatusBadRequest)
		})
		t.Run("GET", func(t *testing.T) {
			assertError(t, "GET", "/add?a=foo&bar=200", http.StatusBadRequest)
		})
		t.Run("GET", func(t *testing.T) {
			assertError(t, "GET", "/add?a=100&bar=bar", http.StatusBadRequest)
		})
		t.Run("POST", func(t *testing.T) {
			assertError(t, "POST", "/add?a=1&b=2", http.StatusMethodNotAllowed)
		})
	})
}
