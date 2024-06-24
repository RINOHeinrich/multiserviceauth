package main

//this is a simple client to test the CORS middleware
import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest("GET", "http://localhost:8080/users", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers to test CORS
	req.Header.Set("Origin", "http://localhost:3000")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status code
	fmt.Println("Response Status:", resp.Status)
}
