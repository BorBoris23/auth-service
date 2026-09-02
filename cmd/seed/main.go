package main

import (
	"context"
	"log"

	"auth-service/internal/postgres"
	seed "auth-service/internal/repository/seeds"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	conn, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := seed.SeedRoles(ctx, conn); err != nil {
		log.Fatal(err)
	}

	if err := seed.SeedUsers(ctx, conn); err != nil {
		log.Fatal(err)
	}

	log.Println("Seeds completed successfully")
}
