package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/levigross/grequests"
)

/*
grequests:
-> Inspired by python requests library for making http requests.
-> simple interface for making http requests helps to reduce the boilerplate code.
-> It is a wrapper around the net/http package.
-> useful while making async http requests.

Features:
-> Simple interface for making http requests.
-> Compatible with goroutines.
-> Customizable options like timeout, headers, etc.

*/

type recipe struct {
	Id              int      `json:"id"`
	Name            string   `json:"name"`
	Ingredients     []string `json:"ingredients"`
	Instructions    []string `json:"instructions"`
	PrepTimeMinutes int      `json:"prepTimeMinutes"`
	CookTimeMinutes int      `json:"cookTimeMinutes"`
}

func main() {
	// A structure that specifies the options for the request. You can set headers, cookies, timeouts, etc.
	// Setting the request options
	requestOptions := grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-API-Key":    "1234567890",
			// "Authorization": "Bearer XXX.XXX.XXX",
		},
	}

	// get request
	resp, err := grequests.Get("https://dummyjson.com/recipes/1", &requestOptions)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// print the response
	// Using persistent response
	var per_recipe recipe
	err = json.Unmarshal(resp.Bytes(), &per_recipe)
	if err != nil {
		fmt.Println("JSON parse error:", err)
		return
	}

	fmt.Printf("streamline response: %+v\n", per_recipe)

	// Request Quirks
	// grequests automatically sets the User-Agent header to "grequests". You can override this by setting the UserAgent field in the RequestOptions struct.

	ro := grequests.RequestOptions{
		UserAgent: "My Custom User Agent",
		Params: map[string]string{
			"limit":  "10",
			"skip":   "0",
			"select": "name,image",
			"search": "pizza",
			"order":  "desc",
		},
	}

	resp, err = grequests.Get("https://dummyjson.com/recipes", &ro)
	// Above Params will be converted to query string

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// print the response
	fmt.Println("Request Quirks")
	fmt.Println(resp.String())

	// Request with JSON data for the new receipt
	jsonData := map[string]string{
		"name": "MoMo",
		// other input fields
	}

	ro = grequests.RequestOptions{
		JSON: jsonData, // set the JSON data for the request
	}

	// send the POST request
	resp, err = grequests.Post("https://dummyjson.com/recipes/add", &ro)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// print the response
	fmt.Println("Add Request with JSON data")
	fmt.Println(resp.String())

	// Error handling
	ro = grequests.RequestOptions{
		RequestTimeout: 5 * time.Second, // set the timeout for the request
	}

	// to demonstrate error handling request non existing url

	resp, err = grequests.Get("http://dummyjson.com/recipes/1000", &ro)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// print the response
	fmt.Println(resp.String())

	// Concurrency
	// grequests supports concurrent requests. You can use the grequests.Map function to make concurrent requests.
	// The Map function takes a slice of URLs and a function that returns a grequests.RequestOptions struct.
	// The function is called for each URL and the result is returned in a slice of grequests.Response structs.
	urls := []string{
		"https://dummyjson.com/recipes/1",
		"https://dummyjson.com/recipes/2",
		"https://dummyjson.com/recipes/3",
		"https://dummyjson.com/recipes/4",
		"https://dummyjson.com/recipes/5",
	}

	var wg sync.WaitGroup // it waits until all the goroutines are done

	fmt.Println("Concurrency")

	// make concurrent requests
	// process each request asynchronously in a goroutine
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			resp, err := grequests.Get(url, nil)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(resp.String())
		}(url)
	}

	wg.Wait()
	fmt.Println("All requests completed")

	// Download
	fmt.Println("Download")

	resp, err = grequests.Get("https://dummyjson.com/image/120/120/ffffff?text=Hello+Gopher", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// save the response to a file
	if err := resp.DownloadToFile("image"); err != nil {
		log.Println("Unable to download file: ", err)
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Image downloaded successfully")
}
