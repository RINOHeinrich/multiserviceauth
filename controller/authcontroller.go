package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/config"
	"github.com/RINOHeinrich/multiserviceauth/helper"
	"github.com/RINOHeinrich/multiserviceauth/models"
)

var Loginmanager helper.LoginManager
var Tokenmanager helper.TokenManager
var Keymanager helper.KeyManager

func Login(w *http.ResponseWriter, r *http.Request) {
	userlogin := models.UserLogin{}
	err := json.NewDecoder(r.Body).Decode(&userlogin)
	if err != nil {
		log.Println(err)
	}
	//json.Unmarshal(r.Body)
	fmt.Printf("user login: %s", userlogin.Login)
	Loginmanager.Userlogin = userlogin
	Loginmanager.Db = &DB
	Loginmanager.LoginErrorMessage = fmt.Errorf("invalid username or password")
	Loginmanager.Bh = &Bh
	Tokenmanager.LoadConfig(&config.Config)
	user, err := Loginmanager.CheckUser()

	if err != nil {
		http.Error(*w, err.Error(), http.StatusUnauthorized)
		return
	}
	if user.Login == "" {
		http.Error(*w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	Loginmanager.HashPassword = user.Password
	err = Loginmanager.CheckPassword()

	if err != nil {
		http.Error(*w, err.Error(), http.StatusUnauthorized)
		return
	}

	Keymanager.LoadConfig(&config.Config.Keyconfig)
	Tokenmanager.User = &user
	Tokenmanager.Keymanager = Keymanager
	token, err := Tokenmanager.GenerateToken()

	if err != nil {
		log.Default().Println(err)
		return
	}

	json.NewEncoder(*w).Encode(token)

}
