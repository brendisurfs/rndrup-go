package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Info struct: for .env return
type Info struct {
	email, password string
}

// Env parses dotenv file from the user and sends it to main
func Env() []string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env")
	}

	userEmail := os.Getenv("EMAIL")
	userPassword := os.Getenv("PASSWORD")

	return []string{userEmail, userPassword}
}
