package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestWelcomeHandler(t *testing.T) {
	// Create a request to pass to our handler
	req := httptest.NewRequest("GET", "/", nil)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler directly
	welcomeHandler(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Welcome to go-phers world!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name           string
		queryName      string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid name",
			queryName:      "John",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, John!",
		},
		{
			name:           "Empty name",
			queryName:      "",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, !",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create URL with query parameters
			params := url.Values{}
			params.Add("name", tt.queryName)
			req := httptest.NewRequest("GET", "/hello?"+params.Encode(), nil)

			rr := httptest.NewRecorder()
			helloHandler(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

func TestHelloPostHandler(t *testing.T) {
	tests := []struct {
		name           string
		formData       string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid name",
			formData:       "name=John",
			expectedStatus: http.StatusCreated,
			expectedBody:   "Hello, John!",
		},
		{
			name:           "Empty name",
			formData:       "name=",
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   "Name is required\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/hello",
				strings.NewReader(tt.formData))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			rr := httptest.NewRecorder()
			helloPostHandler(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expectedBody)
			}
		})
	}
}

func TestWelcomeMultipleHandler(t *testing.T) {
	tests := []struct {
		name           string
		serverAddr     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Server One",
			serverAddr:     ":8010",
			expectedStatus: http.StatusOK,
			expectedBody:   "Welcome to go-phers world! from server :8010",
		},
		{
			name:           "Server Two",
			serverAddr:     ":8011",
			expectedStatus: http.StatusOK,
			expectedBody:   "Welcome to go-phers world! from server :8011",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			// Set the server address in the context
			ctx := context.WithValue(req.Context(), CusKeyServerAddr, tt.serverAddr)
			req = req.WithContext(ctx)

			rr := httptest.NewRecorder()
			welcomeMultipleHandler(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}
