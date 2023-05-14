package main

import (
	"fmt"
	"os"

	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hey there")

	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	fmt.Println("Port: ", port)
}