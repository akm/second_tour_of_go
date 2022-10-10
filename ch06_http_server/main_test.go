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
	t.Run("add", func(t *testing.T) {
		t.Run("GET", func(t *testing.T) {
			w := httptest.NewRecorder()
			handler(w, httptest.NewRequest("GET", "/add?a=1&b=2", nil))
			resp := w.Result()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			body, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Equal(t, "3", strings.TrimSpace(string(body)))
		})
		t.Run("GET", func(t *testing.T) {
			w := httptest.NewRecorder()
			handler(w, httptest.NewRequest("GET", "/add?a=100&b=200", nil))
			resp := w.Result()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			body, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Equal(t, "300", strings.TrimSpace(string(body)))
		})
		t.Run("GET", func(t *testing.T) {
			w := httptest.NewRecorder()
			handler(w, httptest.NewRequest("GET", "/add", nil))
			resp := w.Result()
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})
		t.Run("GET", func(t *testing.T) {
			w := httptest.NewRecorder()
			handler(w, httptest.NewRequest("GET", "/add?a=100", nil))
			resp := w.Result()
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})
		t.Run("GET", func(t *testing.T) {
			w := httptest.NewRecorder()
			handler(w, httptest.NewRequest("GET", "/add?b=100", nil))
			resp := w.Result()
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})
		t.Run("POST", func(t *testing.T) {
			w := httptest.NewRecorder()
			handler(w, httptest.NewRequest("POST", "/add?a=1&b=2", nil))
			resp := w.Result()
			assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
		})
	})
}
