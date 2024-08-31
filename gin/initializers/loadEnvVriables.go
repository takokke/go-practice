package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// 環境変数GO_ENVは適当につけました
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "development")
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatal("error loading .env file")
	}
}
