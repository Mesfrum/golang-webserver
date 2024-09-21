package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandleRequest(t *testing.T) {
	// Create a new instance of APIServer
	server := NewAPIServer(":3000")

	// Create a new request
	reqBody := []byte(`{"test": "data"}`)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a new router and register the handleRequest function
	router := mux.NewRouter()
	router.HandleFunc("/", server.handleRequest).Methods("POST")

	// Serve the request
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := map[string]interface{}{"message": "Request processed successfully"}
	var got map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	if got["message"] != expected["message"] {
		t.Errorf("handler returned unexpected body: got %v want %v", got, expected)
	}
}

func TestWriteJSON(t *testing.T) {
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Test data
	testData := map[string]string{"test": "data"}

	// Call WriteJSON
	WriteJSON(rr, http.StatusOK, testData)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("WriteJSON set wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the Content-Type header
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("WriteJSON set wrong Content-Type: got %v want %v", ctype, "application/json")
	}

	// Check the response body
	var got map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	if got["test"] != testData["test"] {
		t.Errorf("WriteJSON wrote unexpected body: got %v want %v", got, testData)
	}
}

func TestNewAPIServer(t *testing.T) {
	listenAddr := ":3000"
	server := NewAPIServer(listenAddr)

	if server == nil {
		t.Error("NewAPIServer returned nil")
	}

	if server.listenAddr != listenAddr {
		t.Errorf("NewAPIServer set wrong listenAddr: got %v want %v", server.listenAddr, listenAddr)
	}
}
