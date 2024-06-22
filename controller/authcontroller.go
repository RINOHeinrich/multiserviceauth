package controller

import (
	"encoding/json"
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/database"
	"github.com/RINOHeinrich/multiserviceauth/helper"
	"github.com/RINOHeinrich/multiserviceauth/models"
)

func Login(w *http.ResponseWriter, r *http.Request) {
	// Handle POST request
	userlogin := models.UserLogin{}
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	// Check if user exists in the database usin
	database.Find(&db, userlogin.Email)
	iscorrect, err := helper.ComparePassword(user.Password, userlogin.Password)
	if err != nil {
		http.Error(*w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if iscorrect {
		json.NewEncoder(*w).Encode("Login successful")
		return
	}
	json.NewEncoder(*w).Encode("Login failed")

}
