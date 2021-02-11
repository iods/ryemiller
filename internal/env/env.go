package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvVar(key string) string {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("error loading the .env file")
	}
	return os.Getenv(key)
}