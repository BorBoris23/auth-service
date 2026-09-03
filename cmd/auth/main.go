package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	"auth-service/internal/auth"
	grpcauth "auth-service/internal/grpc"
	internalhttp "auth-service/internal/http"
	"auth-service/internal/jwt"
	"auth-service/internal/postgres"
	"auth-service/internal/repository/role"
	"auth-service/internal/repository/user"
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

	authServer := grpcauth.NewAuthServer(jwtService)

	go startGRPCServer(authServer)

	log.Println("Auth HTTP service started on :8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
