package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Command line flags
	numRequests := flag.Int("numRequests", 180, "Number of concurrent requests to send")
	sleepDuration := flag.Duration("sleepDuration", 7*time.Millisecond, "Duration to sleep between requests")
	endpointURL := flag.String("endpoint", "http://localhost:8080/login", "URL of the login endpoint")
	jsonFilePath := flag.String("jsonFile", "", "Path to the JSON file containing the payload")
	flag.Parse()

	// Validate numRequests
	if *numRequests <= 0 {
		fmt.Println("Number of requests must be greater than zero")
		return
	}

	// Read JSON payload from file
	fileContent, err := ioutil.ReadFile(*jsonFilePath)
	if err != nil {
		fmt.Printf("Error reading JSON file: %s\n", err)
		return
	}

	// WaitGroup to synchronize the goroutines
	var wg sync.WaitGroup
	wg.Add(*numRequests)

	// Channel to collect response statuses
	statusChan := make(chan int, *numRequests)

	// Function to perform the request with JSON payload from file
	doRequest := func(payload []byte) {
		defer wg.Done()

		// Create a POST request with JSON body
		req, err := http.NewRequest("POST", *endpointURL, bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error creating request:", err)
			statusChan <- http.StatusInternalServerError
			return
		}
		req.Header.Set("Content-Type", "application/json")

		// Send the request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			statusChan <- http.StatusInternalServerError
			return
		}
		defer resp.Body.Close()

		// Send status to channel for counting
		statusChan <- resp.StatusCode
	}

	startTime := time.Now()
	// Launch concurrent requests
	for i := 0; i < *numRequests; i++ {
		go doRequest(fileContent)
		time.Sleep(*sleepDuration)
	}

	// Wait for all requests to finish
	go func() {
		wg.Wait()
		close(statusChan)
	}()

	// Counters for status codes
	count200 := 0
	count401 := 0
	count500 := 0

	// Collect status codes
	for status := range statusChan {
		switch status {
		case http.StatusOK:
			count200++
		case http.StatusUnauthorized:
			count401++
		case http.StatusInternalServerError:
			count500++
		}
	}

	// Calculate and print the duration
	duration := time.Since(startTime)
	fmt.Printf("\nTest duration: %s\n", duration)

	// Print the counts
	fmt.Printf("\nNumber of 200 OK responses: %d\n", count200)
	fmt.Printf("Number of 401 Unauthorized responses: %d\n", count401)
	fmt.Printf("Number of 500 Internal Server Error responses: %d\n", count500)
}
