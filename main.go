package main

import (
	"fmt"
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/controller"
)

func main() {

	http.HandleFunc("/", userhandler)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func userhandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	switch r.Method {
	case "GET":
		// Handle GET request
		if len(query["id"]) == 0 {
			controller.GetAllUsers(&w, r)
			return
		}
		controller.GetUser(&w, r)
	case "POST":
		controller.InsertUser(&w, r)
	case "PUT":
		controller.UpdateUser(&w, r)
	case "DELETE":
		controller.DeleteUser(&w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
