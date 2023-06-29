package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FadyGamilM/Rss-Aggregator/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	//! read env variables
	// ➜  Load the .env file
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port is not found in the environment variables !")
	} else {
		fmt.Printf("Server port is %v \n", port)
	}

	// define a router which handles the http requests
	router := chi.NewRouter()
	testRouter := chi.NewRouter()

	// use middleware to handle cors errors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// define a server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	testRouter.Get("/health", handlers.Test_server_health)
	testRouter.Get("/error-handling", handlers.Test_error_handler)
	router.Mount("/api/v1", testRouter)

	// run the server
	log.Printf("server is running up on port %v \n", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
