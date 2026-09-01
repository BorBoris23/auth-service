package server

import (
	"net/http"

	"auth-service/internal/server/handlers"
)

func NewRouter(authHandler *handlers.AuthHandler) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /login", authHandler.Login)

	return router
}
