package main

import (
	"context"
	"log"

	"auth-service/internal/database"
	seed "auth-service/internal/seeds"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	conn, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	if err := seed.SeedRoles(ctx, conn); err != nil {
		log.Fatal(err)
	}

	if err := seed.SeedUsers(ctx, conn); err != nil {
		log.Fatal(err)
	}

	log.Println("Seeds completed successfully")
}
