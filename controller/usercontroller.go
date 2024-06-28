package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/config"
	"github.com/RINOHeinrich/multiserviceauth/database"
	"github.com/RINOHeinrich/multiserviceauth/helper"
	"github.com/RINOHeinrich/multiserviceauth/models"
)

var DB database.Postgres
var Bh helper.BcryptHandler

func InitDB() {
	dbconfig := config.Config.Dbconfig
	// Load the database configuration from the .env file
	DB.LoadConfig(&dbconfig)
	err := DB.Connect()
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
	}

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
	id_str := &query["login"][0]
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
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	err := DB.DB.Ping()
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return
	}
	User, _ := Loginmanager.CheckUser(&DB)
	if User.Login != "" {
		err := errors.New("user already exist")
		http.Error(*w, err.Error(), http.StatusUnauthorized)
		return
	}
	user1 := User
	if user1.Login == user.Login {
		http.Error(*w, "User already exists", http.StatusUnauthorized)
		return
	}
	Bh.Config = config.Config.Bcryptconfig
	// Handle POST request
	user.Password = Bh.HashPassword(user.Password)

	err = DB.Insert(&user)

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
	if len(query["login"]) == 0 {
		http.Error(*w, "User not exist", http.StatusMethodNotAllowed)
		return
	}
	id_str := &query["login"][0]
	if *id_str == "" {
		http.Error(*w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	err := database.Update(&DB, *id_str, &user)
	if err != nil {
		fmt.Println("Error updating user: ", err)
		return
	}
	fmt.Fprintf(*w, "PUT request received on login %s", *id_str)
}
func DeleteUser(w *http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if len(query["login"]) == 0 {
		http.Error(*w, "User not exist", http.StatusMethodNotAllowed)
		return
	}
	id_str := &query["login"][0]
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
