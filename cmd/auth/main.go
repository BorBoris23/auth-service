package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"auth-service/internal/database"
	"auth-service/internal/repository"
	"auth-service/internal/server"
	"auth-service/internal/server/handlers"
	"auth-service/internal/server/service"
	"auth-service/internal/server/token"

	"github.com/go-playground/validator/v10"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	validate := validator.New()

	conn, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	userRepository := repository.NewUserRepository(conn)

	jwtService := token.NewJWTService(os.Getenv("JWT_SECRET"))

	authService := service.NewAuthService(
		userRepository,
		jwtService,
	)

	authHandler := handlers.NewAuthHandler(
		validate,
		authService,
	)

	router := server.NewRouter(authHandler)

	log.Println("Auth service started on :8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
