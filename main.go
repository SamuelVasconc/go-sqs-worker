package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//Load environment variables by .env file
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
