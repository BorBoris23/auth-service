package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"auth-service/internal/auth"
	internalhttp "auth-service/internal/http"
	"auth-service/internal/jwt"
	"auth-service/internal/postgres"
	"auth-service/internal/role"
	"auth-service/internal/user"

	"github.com/go-playground/validator/v10"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	validate := validator.New()

	conn, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userRepository := user.NewUserRepository(conn)
	roleRepository := role.NewRoleRepository(conn)

	jwtService := jwt.NewJWTService(os.Getenv("JWT_SECRET"))

	authService := auth.NewAuthService(
		userRepository,
		roleRepository,
		jwtService,
	)

	authController := internalhttp.NewAuthController(
		validate,
		authService,
	)

	router := internalhttp.NewRouter(authController)

	log.Println("Auth service started on :8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
