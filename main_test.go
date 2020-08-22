package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootHandler(t *testing.T) {
	t.Run("Success Case", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()

		rootHandler(resp, req)

		got := resp.Body.String()
		want := "Hello Gocker!"

		assert.Equal(t, got, want)
	})
}
