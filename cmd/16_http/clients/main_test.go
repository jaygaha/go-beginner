package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestGetMethodExample tests the GET method functionality
func TestGetMethodExample(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		// Verify path contains a post ID
		if !strings.Contains(r.URL.Path, "/posts/") {
			t.Errorf("Expected path to contain /posts/, got %s", r.URL.Path)
		}

		// Return a mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := postBody{
			UserID: 1,
			Id:     42,
			Title:  "Test Title",
			Body:   "Test Body",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Override the HTTP client to use our test server
	testClient := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	// Test the function
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL+"/posts/42", nil)
	if err != nil {
		t.Fatalf("Error creating request: %s", err)
	}

	resp, err := testClient.Do(req)
	if err != nil {
		t.Fatalf("Error making request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var resBody postBody
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		t.Fatalf("Error decoding body: %s", err)
	}

	// Verify response
	if resBody.Id != 42 {
		t.Errorf("Expected ID 42, got %d", resBody.Id)
	}
	if resBody.Title != "Test Title" {
		t.Errorf("Expected title 'Test Title', got '%s'", resBody.Title)
	}
}

// TestPostMethodExample tests the POST method functionality
func TestPostMethodExample(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		// Verify path
		if r.URL.Path != "/posts" {
			t.Errorf("Expected path /posts, got %s", r.URL.Path)
		}

		// Verify content type
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", contentType)
		}

		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Error reading request body: %s", err)
		}

		// Verify request body contains expected fields
		var reqBody postBody
		if err := json.Unmarshal(body, &reqBody); err != nil {
			t.Fatalf("Error unmarshalling request body: %s", err)
		}

		if reqBody.UserID == 0 {
			t.Error("Expected UserID to be set")
		}
		if reqBody.Title == "" {
			t.Error("Expected Title to be set")
		}
		if reqBody.Body == "" {
			t.Error("Expected Body to be set")
		}

		// Return a mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response := postBody{
			UserID: reqBody.UserID,
			Id:     101, // New post ID
			Title:  reqBody.Title,
			Body:   reqBody.Body,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Test the function
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	postData := postBody{
		UserID: 5,
		Title:  "Test Post Title",
		Body:   "Test Post Body",
	}

	postBodyJson, err := json.Marshal(postData)
	if err != nil {
		t.Fatalf("Error marshalling post body: %s", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, server.URL+"/posts", bytes.NewBuffer(postBodyJson))
	if err != nil {
		t.Fatalf("Error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error making request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	// Read response body
	var resBody postBody
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		t.Fatalf("Error decoding response body: %s", err)
	}

	// Verify response
	if resBody.Id != 101 {
		t.Errorf("Expected ID 101, got %d", resBody.Id)
	}
	if resBody.UserID != postData.UserID {
		t.Errorf("Expected UserID %d, got %d", postData.UserID, resBody.UserID)
	}
	if resBody.Title != postData.Title {
		t.Errorf("Expected title '%s', got '%s'", postData.Title, resBody.Title)
	}
}

// TestPutMethodExample tests the PUT method functionality
func TestPutMethodExample(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}

		// Verify path contains a post ID
		if !strings.Contains(r.URL.Path, "/posts/") {
			t.Errorf("Expected path to contain /posts/, got %s", r.URL.Path)
		}

		// Verify content type
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", contentType)
		}

		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Error reading request body: %s", err)
		}

		// Verify request body contains expected fields
		var reqBody postBody
		if err := json.Unmarshal(body, &reqBody); err != nil {
			t.Fatalf("Error unmarshalling request body: %s", err)
		}

		if reqBody.UserID == 0 {
			t.Error("Expected UserID to be set")
		}
		if reqBody.Title == "" {
			t.Error("Expected Title to be set")
		}
		if reqBody.Body == "" {
			t.Error("Expected Body to be set")
		}

		// Return a mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := postBody{
			UserID: reqBody.UserID,
			Id:     42, // Updated post ID
			Title:  reqBody.Title,
			Body:   reqBody.Body,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Test the function
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	putData := postBody{
		UserID: 5,
		Title:  "Updated Title",
		Body:   "Updated Body",
	}

	putBodyJson, err := json.Marshal(putData)
	if err != nil {
		t.Fatalf("Error marshalling put body: %s", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		server.URL+"/posts/42",
		bytes.NewBuffer(putBodyJson),
	)
	if err != nil {
		t.Fatalf("Error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error making request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Read response body
	var resBody postBody
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		t.Fatalf("Error decoding response body: %s", err)
	}

	// Verify response
	if resBody.Id != 42 {
		t.Errorf("Expected ID 42, got %d", resBody.Id)
	}
	if resBody.UserID != putData.UserID {
		t.Errorf("Expected UserID %d, got %d", putData.UserID, resBody.UserID)
	}
	if resBody.Title != putData.Title {
		t.Errorf("Expected title '%s', got '%s'", putData.Title, resBody.Title)
	}
}

// TestPatchMethodExample tests the PATCH method functionality
func TestPatchMethodExample(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method
		if r.Method != http.MethodPatch {
			t.Errorf("Expected PATCH request, got %s", r.Method)
		}

		// Verify path contains a post ID
		if !strings.Contains(r.URL.Path, "/posts/") {
			t.Errorf("Expected path to contain /posts/, got %s", r.URL.Path)
		}

		// Verify content type
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", contentType)
		}

		// Read request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Error reading request body: %s", err)
		}

		// Verify request body contains expected fields
		var reqBody postBody
		if err := json.Unmarshal(body, &reqBody); err != nil {
			t.Fatalf("Error unmarshalling request body: %s", err)
		}

		if reqBody.Title == "" {
			t.Error("Expected Title to be set")
		}

		// Return a mock response with original data plus patch
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := postBody{
			UserID: 1, // Original user ID
			Id:     42,
			Title:  reqBody.Title,   // Updated title
			Body:   "Original body", // Original body
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Test the function
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	patchData := postBody{
		Title: "Patched Title Only",
	}

	patchBodyJson, err := json.Marshal(patchData)
	if err != nil {
		t.Fatalf("Error marshalling patch body: %s", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPatch,
		server.URL+"/posts/42",
		bytes.NewBuffer(patchBodyJson),
	)
	if err != nil {
		t.Fatalf("Error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error making request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Read response body
	var resBody postBody
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		t.Fatalf("Error decoding response body: %s", err)
	}

	// Verify response
	if resBody.Id != 42 {
		t.Errorf("Expected ID 42, got %d", resBody.Id)
	}
	if resBody.Title != patchData.Title {
		t.Errorf("Expected title '%s', got '%s'", patchData.Title, resBody.Title)
	}
	if resBody.Body != "Original body" {
		t.Errorf("Expected body 'Original body', got '%s'", resBody.Body)
	}
}

// TestDeleteMethodExample tests the DELETE method functionality
func TestDeleteMethodExample(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		// Verify path contains a post ID
		if !strings.Contains(r.URL.Path, "/posts/") {
			t.Errorf("Expected path to contain /posts/, got %s", r.URL.Path)
		}

		// Return a mock response
		w.WriteHeader(http.StatusOK)
		// DELETE typically returns an empty body
	}))
	defer server.Close()

	// Test the function
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		server.URL+"/posts/42",
		nil,
	)
	if err != nil {
		t.Fatalf("Error creating request: %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error making request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

// TestErrorHandling tests error handling in HTTP requests
func TestErrorHandling(t *testing.T) {
	// Create a mock server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	// Test with GET request
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL+"/posts/42", nil)
	if err != nil {
		t.Fatalf("Error creating request: %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error making request: %s", err)
	}
	defer resp.Body.Close()

	// Verify error status code is detected
	if resp.StatusCode == http.StatusOK {
		t.Errorf("Expected non-OK status code, got %d", resp.StatusCode)
	}

	// Verify status code is 500
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}
}
