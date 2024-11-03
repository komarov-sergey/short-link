package main

import (
	"fmt"
	"net/http"
	"short-link/internal/hello"
)

func main() {
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":7081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 7081")
	server.ListenAndServe()
}
