package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello golang")

	//! read env variables
	// ➜  Load the .env file
	godotenv.Load(".env")
	 
	port := os.Getenv("PORT")
	if port == ""{
		log.Fatal("port is not found in the environment variables !")
	}else{
		fmt.Printf("Server port is %v \n",port)
	}
}