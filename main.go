package main

import (
	"fmt"
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/config"
	"github.com/RINOHeinrich/multiserviceauth/controller"
	"github.com/RINOHeinrich/multiserviceauth/middleware"
	routeHandler "github.com/RINOHeinrich/multiserviceauth/routeHandling"
	_ "github.com/lib/pq"
)

func init() {
	config.LoadConfig("config/.env")
	controller.InitDB()
}

func main() {
	var cors middleware.CORSHandler
	cors.LoadConfig(&config.Config.Corsconfig)
	http.HandleFunc("/users", cors.Handle(routeHandler.Userhandler))
	http.HandleFunc("/login", cors.Handle(routeHandler.Loginhandler))
	http.HandleFunc("/register", cors.Handle(routeHandler.Registerhandler))
	if config.Config.PublicKeyAccess.Enable {
		http.HandleFunc("/pubkey", cors.Handle(routeHandler.Pubkeyhandler))
	}
	door := config.Config.Host + ":" + fmt.Sprint(config.Config.Port)
	fmt.Printf("Server started on http://%s\n", door)
	http.ListenAndServe(door, nil)
}
