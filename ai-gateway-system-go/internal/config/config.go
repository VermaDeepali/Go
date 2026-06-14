package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system env")
	}

	if os.Getenv("OPENAI_API_KEY") == "" {
		log.Println("warning: OPENAI_API_KEY not set")
	}
}
