package server

import (
	"log"
	"net/http"
)

// RequestLogger is a middleware logging the request.
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Println(req.Method, req.URL.Path, "from", req.RemoteAddr)
		next.ServeHTTP(res, req)
	})
}
