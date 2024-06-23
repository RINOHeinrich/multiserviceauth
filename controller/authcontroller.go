package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"log"

	"github.com/RINOHeinrich/multiserviceauth/helper"
	"github.com/RINOHeinrich/multiserviceauth/models"
)

func Login(w *http.ResponseWriter, r *http.Request) {
	userlogin := models.UserLogin{}
	json.NewDecoder(r.Body).Decode(&userlogin)
	tokenmanager := helper.TokenManager{
		D: 1 * time.Hour,
	}
	loginmanager := helper.LoginManager{
		Userlogin:         userlogin,
		HashPassword:      "",
		LoginErrorMessage: errors.New("invalid username or password"),
		Tm:                &tokenmanager,
		Db:                &DB,
	}
	user, err := loginmanager.CheckUser()
	if err != nil {
		http.Error(*w, err.Error(), http.StatusUnauthorized)
		return
	}
	loginmanager.HashPassword = user.Password
	_, err = loginmanager.CheckPassword()
	if err != nil {
		http.Error(*w, err.Error(), http.StatusUnauthorized)
		return
	}
	tokenmanager.User = user
	privatekeypath := "config/keys/private.pem"
	publickeypath := "config/keys/public.pem"
	if !helper.CheckKeys(privatekeypath) {
		privatekey, pubkey := helper.GenerateKeys()
		helper.SavePublicToDisk(publickeypath, pubkey)
		helper.SavePrivateToDisk(privatekeypath, privatekey)
	}
	tokenmanager.PrivateKey, err = helper.LoadPrivateKey(privatekeypath)
	if err != nil {
		log.Fatal(err)
		return
	}

	token, err := tokenmanager.GenerateToken()
	if err != nil {
		log.Fatal(err)
		return
	}
	json.NewEncoder(*w).Encode(token)
}
