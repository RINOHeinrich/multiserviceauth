package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	login1 := "Rino"
	password1 := "123456"
	login2 := "Rino1234"
	password2 := "OSNIA1"
	endpoint := "http://localhost:8080/login"

	// Number of concurrent requests
	numRequests := 100

	// WaitGroup to synchronize the goroutines
	var wg sync.WaitGroup
	wg.Add(numRequests)

	// Channel to collect response statuses
	statusChan := make(chan int, numRequests)

	// Function to perform the request with JSON payload
	doRequest := func(login, password string) {
		defer wg.Done()

		// Create JSON payload
		data := map[string]string{
			"login":    login,
			"password": password,
		}
		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		// Create a POST request with JSON body
		req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error creating request:", err)
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

		//fmt.Printf("Request for login=%s, password=%s | Status: %s\n", login, password, resp.Status)
	}

	// Launch concurrent requests
	for i := 0; i < numRequests; i++ {
		// Alternate between login1 and login2/password2
		if i%2 == 0 {
			go doRequest(login1, password1)
		} else {
			go doRequest(login2, password2)
		}
	}

	// Close channel after all requests are done
	go func() {
		wg.Wait()
		close(statusChan)
	}()

	// Counters for status codes
	count200 := 0
	count401 := 0

	// Collect status codes
	for status := range statusChan {
		switch status {
		case http.StatusOK:
			count200++
		case http.StatusUnauthorized:
			count401++
		}
	}

	// Print the counts
	fmt.Printf("\nNumber of 200 OK responses: %d\n", count200)
	fmt.Printf("Number of 401 Unauthorized responses: %d\n", count401)
}
