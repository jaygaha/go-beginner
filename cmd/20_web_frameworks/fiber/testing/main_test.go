package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	app := NewServer()

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response, err := app.Test(request)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	// read the response body
	body := make([]byte, response.ContentLength)
	_, err = response.Body.Read(body)

	// EOF is expected when the entire body has been read
	assert.Equal(t, io.EOF, err)
	// Verify the response body content
	assert.Equal(t, "Welcome to Fiber!", string(body))
}

func TestInternalServerError(t *testing.T) {
	app := NewServer()
	request := httptest.NewRequest(http.MethodGet, "/500", nil)
	response, _ := app.Test(request)
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

	// read the response body
	body := make([]byte, response.ContentLength)
	_, err := response.Body.Read(body)
	// EOF is expected when the entire body has been read
	assert.Equal(t, io.EOF, err)
	// Verify the response body content
	assert.Equal(t, "{\"message\":\"Internal Server Error\"}", string(body))

	// alternative way to read the response body
	// body := new(strings.Builder)
	// _, _ = io.Copy(body, response.Body)
	// assert.JSONEq(t, `{"message": "Internal Server Error"}`, body.String())
}
