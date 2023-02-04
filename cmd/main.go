package main

import (
	"fmt"

	"github.com/henry-l-zhou/restaurant-elo-system/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load("../'.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	database.Connect()
}
