package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// handleRequest is a generic handler function for all API requests.
func (s *APIServer) handleRequest(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteJSON(w, http.StatusBadRequest, APIError{Error: "Invalid request"})
		return
	}

	// Process the request here. This is just a placeholder.
	response := map[string]interface{}{"message": "Request processed successfully"}

	WriteJSON(w, http.StatusOK, response)
}

// WriteJSON writes a JSON response to the client.
func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// NewAPIServer creates a new instance of APIServer.
func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{listenAddr: listenAddr}
}

// Run starts the API server.
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/", s.handleRequest).Methods("POST")
	log.Println("API server running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}
