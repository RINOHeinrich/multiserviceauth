package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/config"
	"github.com/RINOHeinrich/multiserviceauth/database"
	"github.com/RINOHeinrich/multiserviceauth/helper"
	"github.com/RINOHeinrich/multiserviceauth/models"
)

var DB database.Postgres
var Bh helper.BcryptHandler

func InitDB() {
	// Load the database configuration from the .env file
	err := DB.LoadConfig("config/.env")
	if err != nil {
		log.Default().Println(err)
	}
	err = DB.Connect()
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
	}
	fmt.Println("Connected to database")
}

func GetAllUsers(w *http.ResponseWriter, r *http.Request) {

	users, err := database.FindAll(&DB)
	if err != nil {
		fmt.Println("Error getting users: ", err)
		return
	}
	json.NewEncoder(*w).Encode(users)
}
func GetUser(w *http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id_str := &query["id"][0]
	if *id_str == "" {
		http.Error(*w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	user, err := database.Find(&DB, *id_str)
	if err != nil {
		fmt.Println("Error getting user: ", err)
		return
	}
	json.NewEncoder(*w).Encode(user)
}
func InsertUser(w *http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)
	Bh.Config = config.Config.Bcryptconfig
	// Handle POST request
	user.Password = Bh.HashPassword(user.Password)
	err := database.Insert(&DB, user)
	if err != nil {
		fmt.Println("Error inserting user: ", err)
		return
	}
	fmt.Fprintf(*w, "User inserted: %v\n", user)
}

func UpdateUser(w *http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	Bh.Config = config.Config.Bcryptconfig
	user.Password = Bh.HashPassword(user.Password)
	query := r.URL.Query()
	id_str := &query["id"][0]
	if *id_str == "" {
		http.Error(*w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	err := database.Update(&DB, *id_str, &user)
	if err != nil {
		fmt.Println("Error updating user: ", err)
		return
	}
	fmt.Fprintf(*w, "PUT request received on id %s", *id_str)
}
func DeleteUser(w *http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id_str := &query["id"][0]
	if *id_str == "" {
		http.Error(*w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	err := database.Delete(&DB, *id_str)
	if err != nil {
		fmt.Println("Error deleting user: ", err)
		return
	}
	fmt.Fprintf(*w, "DELETE request received on id %s", *id_str)
}
