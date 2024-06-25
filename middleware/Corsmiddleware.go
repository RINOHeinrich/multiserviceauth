package middleware

import (
	"net/http"

	"github.com/RINOHeinrich/multiserviceauth/models"
	_ "github.com/lib/pq"
)

type CORSHandler struct {
	Corsconfig     models.Corsconfig
	IgnoredOptions bool
}

func (c *CORSHandler) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set the headers for CORS
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Origin", c.Corsconfig.AccessControlAllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", c.Corsconfig.AccessControlAllowMethods)
		w.Header().Set("Access-Control-Allow-Headers", c.Corsconfig.AccessControlAllowHeaders)
		w.Header().Set("Access-Control-Allow-Credentials", c.Corsconfig.AccessControlAllowCredentials)
		// If this is a preflight request, we don't want to continue
		if r.Method == "OPTIONS" && c.IgnoredOptions {
			return
		}
		next(w, r)
	}
}
func (c *CORSHandler) LoadConfig(config *models.Corsconfig) {
	c.Corsconfig = *config
}
