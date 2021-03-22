package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// Errors
var (
	ErrNoSuchRoute = errors.New("no such route")
)

// A StateServer exports the openbgpd state.
// It implements a request handler decoding the request
// and querying the openbgpd.
type StateServer struct {
	// OpenBGPD
}

// ServeHTTP handles the HTTP request and implements
// the http.Handler interface
func (s *StateServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Route request
	handler := s.routeRequest(req)
	if handler == nil {
		s.handleError(res, req, ErrNoSuchRoute)
		return
	}
	handler.ServeHTTP(res, req)
}

// Route request by selecting an appropriate sub handler
func (s *StateServer) routeRequest(req *http.Request) http.Handler {
	if req.URL.Path == "/" {
		return http.HandlerFunc(s.indexHandler)
	} else if req.URL.Path == "/api/v1/status" {
		return nil
	}
	return nil
}

// HTTP Handlers

// Index, Route: /
func (s *StateServer) indexHandler(res http.ResponseWriter, req *http.Request) {
	s.handleSuccess(res, req, []byte("hi there"))
}

// Handle successful response
func (s *StateServer) handleSuccess(res http.ResponseWriter, req *http.Request, body []byte) {
	res.WriteHeader(http.StatusOK)
	res.Write(body)
}

// Handle errors and generate error response
func (s *StateServer) handleError(res http.ResponseWriter, req *http.Request, err error) {
	res.WriteHeader(http.StatusInternalServerError)

	// Encode error es json
	log.Println(err)
	data, encErr := json.Marshal(err.Error())
	if encErr != nil {
		log.Println("error while encoding error in response:", encErr)
	}
	res.Write(data)
}

// StartHTTP starts the HTTP server at a given listen address
func (s *StateServer) StartHTTP(addr string) {
	log.Fatal(http.ListenAndServe(addr, s))
}
