package main

import (
	"log"
	"time"

	"github.com/gojek/heimdall"
	"github.com/gojek/heimdall/httpclient"
)

/*
Retry with back-offs
-> Retry with back-offs is a technique used to retry a request after a failure.
-> The back-off is a technique used to wait for a certain amount of time before retrying a request.
-> The back-off is usually exponential.

For example:
- The following is an example of code that implements a retry using fixed backoffs.
- This example retries is performed at 2-second intervals to set a timeout for the HTTP request.
*/

func main() {
	// Set fixed backoff (retry interval: 2 seconds)
	// NewConstantBackoff creates a new backoff with a constant retry interval.
	backoff := heimdall.NewConstantBackoff(2*time.Second, 1*time.Second)

	// http client
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(10*time.Millisecond),      // Set the timeout to ten seconds.
		httpclient.WithRetrier(heimdall.NewRetrier(backoff)), // Set the retrier to use the fixed backoff.
	)

	// Make a request
	res, err := client.Get("https://example.com", nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Response code: %d\n", res.StatusCode)
}
