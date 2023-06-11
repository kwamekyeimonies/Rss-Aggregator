package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	serverPort := os.Getenv("PORT")

	if serverPort == "" {
		log.Fatal("PORT is not found in the Environment")
	}

	fmt.Println("Port:", serverPort)
}
