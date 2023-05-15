package initializers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("failed to load environment variables")
		os.Exit(-1)
	}
}
