package main

import (
	"fmt"
	"net/http"
	"short-link/configs"
	"short-link/internal/auth"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})

	server := http.Server{
		Addr:    ":7081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 7081")
	server.ListenAndServe()
}
