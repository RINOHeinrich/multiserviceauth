package middleware

import (
	"net/http"

	_ "github.com/lib/pq"
)

type CORSHandler struct {
	AccessControlAllowOrigin      string
	AccessControlAllowMethods     string
	AccessControlAllowHeaders     string
	AccessControlAllowCredentials string
	IgnoredOptions                bool
}

func (c *CORSHandler) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set the headers for CORS
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Access-Control-Allow-Origin", c.AccessControlAllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", c.AccessControlAllowMethods)
		w.Header().Set("Access-Control-Allow-Headers", c.AccessControlAllowHeaders)
		w.Header().Set("Access-Control-Allow-Credentials", c.AccessControlAllowCredentials)
		// If this is a preflight request, we don't want to continue
		if r.Method == "OPTIONS" && c.IgnoredOptions {
			return
		}
		next(w, r)
	}
}
