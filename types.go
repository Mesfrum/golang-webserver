package main

type APIServer struct {
	listenAddr string
}

type APIError struct {
	Error string `json:"error"`
}
