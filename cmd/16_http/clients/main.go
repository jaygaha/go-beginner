package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"os"
	"time"
)

const serverPort = 8010

func main() {
	// Receive Localhost
	// StartLocalClient()

	// Receive Remote
	StartRemoteClient()
}

func StartLocalClient() {
	fmt.Println("Starting local client...")
	// create a new client
	requestUrl := fmt.Sprintf("http://localhost:%d", serverPort)

	// make a request
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)

	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		os.Exit(1)
	}

	// Get vs DefaultClient
	// Get is a shortcut for DefaultClient.Get
	// Get returns the response body and discards the response.
	// It returns an error if caused by client policy (such as CheckRedirect),
	// or failure to speak HTTP (such as a network connectivity problem).
	// A non-2xx status code doesn't cause an error.
	// The response is always non-nil and must be closed when done reading from it.
	// RawResponse contains the underlying Response so you can inspect it.
	// res, err := http.Get(requestUrl)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		os.Exit(1)
	}

	// read body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading body: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("cleint response: ", res)
	fmt.Println("Status: ", res.Status)
	fmt.Println("Status code: ", res.StatusCode)
	fmt.Println("Body: ", res.Body)
	fmt.Println("Body: ", body)
}

type postBody struct {
	UserID int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func StartRemoteClient() {
	fmt.Println("Starting remote client...")

	fmt.Println("\nGET method example...")
	GetMethodExample()

	fmt.Println("\nPOST method example...")
	PostMethodExample()

	fmt.Println("\nPUT method example...")
	PutMethodExample()

	fmt.Println("\nPATCH method example...")
	PatchMethodExample()

	fmt.Println("\nDELETE method example...")
	DeleteMethodExample()
}

func GetMethodExample() {
	// why context?
	// context is used to pass values to the request as default http does not have a timeout
	ctx := context.Background()
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	postId := randRange(1, 1000)
	// make a request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postId), nil)

	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", resp.Status)
		return
	}

	var resBody postBody
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		fmt.Printf("Error decoding body: %s\n", err)
		return
	}

	fmt.Printf("%+v\n", resBody)
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func PostMethodExample() {
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	postBody := postBody{
		UserID: randRange(1, 10),
		Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
	}

	// marshal postBody to json
	postBodyJson, err := json.Marshal(postBody)
	if err != nil {
		fmt.Printf("Error marshalling post body: %s\n", err)
		return
	}

	// make a request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(postBodyJson))

	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Error: %s\n", resp.Status)
		return
	}

	// read body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body: %s\n", err)
		return
	}

	fmt.Println("Status: ", resp.Status)
	fmt.Println("Status code: ", resp.StatusCode)
	fmt.Println("Body: ", string(bodyBytes))
}

func PutMethodExample() {
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	postBody := postBody{
		UserID: randRange(1, 10),
		Title:  "updated title",
		Body:   "I am updated body",
	}

	// marshal postBody to json
	postBodyJson, err := json.Marshal(postBody)
	if err != nil {
		fmt.Printf("Error marshalling put body: %s\n", err)
		return
	}

	postId := randRange(1, 100)

	// make a request
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postId),
		bytes.NewBuffer(postBodyJson),
	)

	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", resp.Status)
		return
	}

	// read body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body: %s\n", err)
		return
	}

	fmt.Println("Status: ", resp.Status)
	fmt.Println("Status code: ", resp.StatusCode)
	fmt.Println("Body: ", string(bodyBytes))
}

func PatchMethodExample() {
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	postBody := postBody{
		Title: "updated title using patch",
	}

	// marshal postBody to json
	postBodyJson, err := json.Marshal(postBody)
	if err != nil {
		fmt.Printf("Error marshalling patch body: %s\n", err)
		return
	}

	postId := randRange(1, 100)

	// make a request
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPatch,
		fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postId),
		bytes.NewBuffer(postBodyJson),
	)

	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", resp.Status)
		return
	}

	// read body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body: %s\n", err)
		return
	}

	fmt.Println("Status: ", resp.Status)
	fmt.Println("Status code: ", resp.StatusCode)
	fmt.Println("Body: ", string(bodyBytes))
}

func DeleteMethodExample() {
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	postId := randRange(1, 100)
	// make a request
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postId),
		nil,
	)

	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", resp.Status)
		return
	}
	fmt.Println("Status: ", resp.Status)
	fmt.Println("Status code: ", resp.StatusCode)
}
