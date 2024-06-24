package main

import (
	"fmt"
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/controller"
	"github.com/RINOHeinrich/multiserviceauth/middleware"
	_ "github.com/lib/pq"
)

func init() {
	controller.InitDB()
}

func main() {
	var cors middleware.CORSHandler
	cors.LoadConfig("config/.env")
	http.HandleFunc("/users", cors.Handle(userhandler))
	http.HandleFunc("/login", cors.Handle(loginhandler))
	http.HandleFunc("/register", cors.Handle(registerhandler))
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
	case "PUT":
		controller.UpdateUser(&w, r)
	case "DELETE":
		controller.DeleteUser(&w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
func loginhandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	switch r.Method {
	case "POST":
		controller.Login(&w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
func registerhandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	switch r.Method {
	case "POST":
		controller.InsertUser(&w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
