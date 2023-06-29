package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FadyGamilM/Rss-Aggregator/routers"
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
	router := routers.Routes()
	// define a server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// run the server
	log.Printf("server is running up on port %v \n", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
