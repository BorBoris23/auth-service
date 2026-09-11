package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	"auth-service/internal/events"
	grpcauth "auth-service/internal/grpc"
	internalhttp "auth-service/internal/http"
	"auth-service/internal/jwt"
	"auth-service/internal/kafka"
	"auth-service/internal/postgres"
	"auth-service/internal/repository/role"
	"auth-service/internal/repository/user"
	"auth-service/internal/services"
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

	expiresIn, err := time.ParseDuration(os.Getenv("JWT_EXPIRES_IN"))
	if err != nil {
		log.Fatal(err)
	}

	jwtService := jwt.NewJWTService(
		os.Getenv("JWT_SECRET"),
		expiresIn,
	)

	producer := kafka.NewProducer(
		os.Getenv("KAFKA_BROKER"),
		os.Getenv("KAFKA_USER_CREATED_TOPIC"),
	)
	defer producer.Close()

	log.Printf(
		"Kafka producer started on %s",
		os.Getenv("KAFKA_BROKER"),
	)

	publisher := events.NewPublisher(producer)

	dispatcher := events.NewDispatcher()

	userCreatedListener := events.NewUserCreatedListener(
		publisher,
	)

	dispatcher.AddListener(
		events.UserCreatedEventName,
		userCreatedListener,
	)

	authService := services.NewAuthService(
		userRepository,
		roleRepository,
		jwtService,
		dispatcher,
	)

	authController := internalhttp.NewAuthController(
		validate,
		authService,
	)

	userService := services.NewUserService(
		userRepository,
		roleRepository,
	)

	router := internalhttp.NewRouter(authController)

	authServer := grpcauth.NewAuthServer(jwtService)
	userServer := grpcauth.NewUserServer(userService)

	go startGRPCAuthServer(authServer)
	go startGRPCUserServer(userServer)

	httpPort := os.Getenv("AUTH_HTTP_PORT")

	log.Printf("Auth HTTP service started on :%s", httpPort)

	err = http.ListenAndServe(
		fmt.Sprintf(":%s", httpPort),
		router,
	)
	if err != nil {
		log.Fatal(err)
	}
}
