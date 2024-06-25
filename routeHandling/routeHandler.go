package routeHandler

import (
	"crypto/x509"
	"fmt"
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/controller"
)

func Loginhandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	switch r.Method {
	case "POST":
		controller.Login(&w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
func Registerhandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	switch r.Method {
	case "POST":
		controller.InsertUser(&w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
func Userhandler(w http.ResponseWriter, r *http.Request) {
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
func Pubkeyhandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	switch r.Method {
	case "GET":
		pubkey, err := controller.Keymanager.GetPublicKey()
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		pubkeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(pubkeyBytes)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
