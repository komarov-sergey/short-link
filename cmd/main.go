package main

import (
	"fmt"
	"net/http"
	"short-link/configs"
	"short-link/internal/auth"
	"short-link/internal/link"
	"short-link/pkg/db"
	"short-link/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(db)

	// Handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	// Middlewares
	stack := middleware.Chain(middleware.CORS, middleware.Logging)

	server := http.Server{
		Addr:    ":7081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 7081")
	server.ListenAndServe()
}
