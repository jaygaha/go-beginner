package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gojek/heimdall"
	"github.com/gojek/heimdall/httpclient"
)

/*
Retry using index back-off
-> The index back-off is a strategy to gradually increase the retry interval.

For example, we retry after 500 milliseconds at first, then doubling the interval, such as 1 second and 2 seconds.
This strategy is useful if you want to increase the likelihood of success through retry while reducing the load on resources.

The following is an example of code that implements a retry using exponential backoff.
This example sets the retry to repeat the initial interval from 500 milliseconds to up to 5 seconds interval.
*/

func main() {
	// Set index backoff
	backoff := heimdall.NewExponentialBackoff(
		500*time.Millisecond,    // Initial standby time (first interval of retrie)
		5*time.Second,           // Maximum waiting time (maximum interval of retries)
		float64(15*time.Second), // Maximum elapsed time (maximum time for all retries)
		2.0,                     // Exponential increase rate (doubling the interval for each retry)
	)

	// create a new client with heimdall
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(10*time.Second),           // set the timeout to 10 seconds
		httpclient.WithRetrier(heimdall.NewRetrier(backoff)), // set the retrier to backoff
	)

	// make a GET request
	resp, err := client.Get("https://not-existance.com", nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Response code: %d\n", resp.StatusCode)
}
