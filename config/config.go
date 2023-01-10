package config

import (
	"fmt"
	"os"
	
    "github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	
	if err != nil {
        fmt.Print("Error loading .env file")
	}
	
	fmt.Println(os.Getenv(key))
	return os.Getenv(key)
}