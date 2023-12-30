package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	filename := fmt.Sprintf(".env.%s", env)
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Error loading %s file", filename)
	}
}
