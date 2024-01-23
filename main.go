package main

import (
	routes "Portfoliobackend/routers"
	utils "Portfoliobackend/utils"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func init() { // Init function to run before main
	fmt.Println("Server is starting...")
}

func main() {
	utils.ConnectDb() // Connect to database
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/api/tags", routes.TagsRoutes()) // Mount tags router

	server := http.ListenAndServe(":8080", r) // Run server on port 8080

	if server != nil { // If server is not running, panic
		panic(server)
	} else {
		fmt.Println("Server is running on port 8080")
	}
}
