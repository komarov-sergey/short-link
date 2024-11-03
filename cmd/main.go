package main

import (
	"fmt"
	"net/http"
	"short-link/configs"
	"short-link/internal/hello"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":7081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 7081")
	server.ListenAndServe()
}
