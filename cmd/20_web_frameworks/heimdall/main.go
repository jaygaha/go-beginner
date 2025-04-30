package main

import (
	"fmt"
	"io"
	"time"

	"github.com/gojek/heimdall/httpclient"
)

/*
Heimdall:
-> Heimdall is an HTTP client that helps your application make a large number of requests, at scale.

With Heimdall, you can:

    - Use a hystrix-like circuit breaker to control failing requests
    - Add synchronous in-memory retries to each request, with the option of setting your own retrier strategy
    - Create clients with different timeouts for every request

All HTTP methods are exposed as a fluent interface.

Installation:
    $ go get -u github.com/gojek/heimdall/v7
*/

// GET
// A GET request is the most basic HTTP request and is used to retrieve data from a server.
// Heimdall makes it easy to set retries and timeouts.

func main() {
	// create a new client with heimdall
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(10*time.Second), // set the timeout to 10 seconds
		httpclient.WithRetryCount(3),               // set the number of retries to 3
	)

	// make a GET request
	resp, err := client.Get("https://dummyjson.com/quotes", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// output the response
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Status Code:", resp.StatusCode)
	fmt.Println("Response Body:", string(body))
	// fmt.Println("Response Headers:", resp.Header)
}
