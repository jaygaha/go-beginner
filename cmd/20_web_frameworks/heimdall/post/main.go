package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gojek/heimdall/httpclient"
)

/*
POST
-> A POST request is used to send data to the server.
For example, it is used to send JSON data to a server.
*/

func main() {
	// create a new client with heimdall
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(10*time.Second), // set the timeout to 10 seconds
		httpclient.WithRetryCount(3),               // set the number of retries to 3
	)

	// JSON data
	jsonData := []byte(`{"body": "Comments by commenter", "postId": 3, "userId": 1}`)
	headers := http.Header{
		"Content-Type": []string{"application/json"},
	}

	//   make a POST request
	resp, err := client.Post("https://dummyjson.com/comments/add", bytes.NewBuffer(jsonData), headers)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	defer resp.Body.Close()

	// read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// print the response
	fmt.Println(string(body))
	// {"id":341,"body":"Comments by commenter","postId":3,"user":{"id":1,"username":"emilys","fullName":"Emily Johnson"}}
}
