package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/alice-lg/openbgpd-state-server/pkg/bgpctl"
)

// Errors
var (
	ErrNoSuchRoute = errors.New("no such route")
)

// Constants
const (
	BgpdRequestTimeout = 60 * time.Second
)

// A StateServer exports the openbgpd state.
// It implements a request handler decoding the request
// and querying the openbgpd.
type StateServer struct {
	// OpenBGPD
}

// Status returns the current server status
func (s *StateServer) Status() *Status {
	// Get bgpctl status
	return &Status{
		Service: "openbgpd-state-server",
		Version: Version,
		Build:   Build,
	}
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
		return http.HandlerFunc(s.serveStatus)
	} else if req.URL.Path == "/api/v1/status" {
		return http.HandlerFunc(s.serveStatus)
	} else if strings.HasPrefix(req.URL.Path, "/api/v1/bgpd") {
		return http.HandlerFunc(s.serveBGPD)
	}
	return nil
}

// HTTP Handlers

// Status: Show the current server status
// Route: /
//
func (s *StateServer) serveStatus(res http.ResponseWriter, req *http.Request) {
	s.respondJSON(res, http.StatusOK, s.Status())
}

// BGPD API: bgpctl commands are encoded in the path
// of the request URL.
func (s *StateServer) serveBGPD(res http.ResponseWriter, req *http.Request) {
	// Prepare path
	path := strings.TrimPrefix(req.URL.Path, "/api/v1/bgpd")
	path = strings.Trim(path, "/ ")
	prefix := strings.ReplaceAll(path, "/", " ")
	query, _ := url.QueryUnescape(req.URL.RawQuery)
	cmd := strings.TrimSpace(prefix + " " + query)

	// We use the request context. This will be cancelled if the
	// connection is closed. Also we add a timelimit.
	ctx, cancel := context.WithTimeout(req.Context(), BgpdRequestTimeout)
	defer cancel()

	// Decode request
	ctlreq := bgpctl.RequestFromString(cmd).Sanitize()

	// Query bgpctl
	result, err := bgpctl.DefaultBGPCTL.Do(ctx, ctlreq)
	if err != nil {
		s.handleError(res, req, err)
		return
	}

	// Respond with raw json
	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

// Handle errors and generate error response
func (s *StateServer) handleError(
	res http.ResponseWriter,
	req *http.Request,
	err error,
) {
	log.Println("ERROR:", err.Error())
	status := http.StatusInternalServerError
	s.respondJSON(res, status, err.Error())
}

// Create JSON response
func (s *StateServer) respondJSON(
	res http.ResponseWriter, status int, body interface{},
) {
	res.Header().Set("content-type", "application/json")
	res.WriteHeader(status)

	// Encode error es json
	data, encErr := json.Marshal(body)
	if encErr != nil {
		log.Println("error while encoding response:", encErr)
	}
	res.Write(data)
}

// StartHTTP starts the HTTP server at a given listen address
func (s *StateServer) StartHTTP(addr string) {
	log.Fatal(http.ListenAndServe(addr,
		RequestLogger(s)))
}
