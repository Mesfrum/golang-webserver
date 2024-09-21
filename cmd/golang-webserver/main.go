package main

import "golang-webserver/internal/api"

func main() {
	server := NewAPIServer(":3000")
	server.Run()
}
