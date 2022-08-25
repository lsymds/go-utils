package http_api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Server encompasses a router and all dependencies required to make a functioning HTTP service.
type Server struct {
	router *mux.Router
}

// NewServer initializes a new, functioning server.
func NewServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}

	// s.mapStandardRoutes()

	return s
}

// ok writes a 200 OK response with an optional body to the response writer.
func ok(w http.ResponseWriter, body interface{}) {
	writeResponse(w, 200, body)
}

// notAuthorized writes a 401 NOT AUTHORIZED response with an optional body to the response writer.
func notAuthorized(w http.ResponseWriter, body interface{}) {
	writeResponse(w, 401, body)
}

// forbidden writes a 403 FORBIDDEN response with an optional body to the response writer.
func forbidden(w http.ResponseWriter, body interface{}) {
	writeResponse(w, 403, body)
}

// notFound writes a 404 NOT FOUND response with an optional body to the response writer.
func notFound(w http.ResponseWriter, body interface{}) {
	writeResponse(w, 404, body)
}

// internalServerError writes a 500 INTERNAL SERVER ERROR response with an optional body to the
// response writer.
func internalServerError(w http.ResponseWriter, body interface{}) {
	writeResponse(w, 500, body)
}

// writeResponse sets the response code and writes an optional JSON body to the response writer.
func writeResponse(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
}
