package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Env parses dotenv file from the user and sends it to main
func Env() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env")
	}

	userEmail := os.Getenv("EMAIL")
	userPassword := os.Getenv("PASSWORD")

	fmt.Printf("%v + %v", userEmail, userPassword)
}
